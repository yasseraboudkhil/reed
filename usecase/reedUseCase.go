package usecase

import (
	"errors"

	"github.com/yasseraboudkhil/cloud-hsm/mykms"
	model "github.com/yasseraboudkhil/reed/proto/v1"
	"github.com/yasseraboudkhil/ya-lib/cryptutil"
	"github.com/yasseraboudkhil/ya-lib/strutil"
)

type ReedUseCase interface {
	Encrypt(encryptPayload *model.EncryptPayload, kmsClient *mykms.KMS) (*model.CipherAndKey, error)
	Decrypt(decryptPayload *model.DecryptPayload, kmsClient *mykms.KMS) (string, error)
}

type reedUseCase struct {
	redis map[string]string
}

type ReedEntryKey struct {
	Key string
}

func NewReedUseCase(redis map[string]string) *reedUseCase {
	return &reedUseCase{
		redis: redis,
	}
}

/*
 * Ensure Encrypt caches (userId + "|" + fieldName + "|" + ciphertextInBase64) -> encryptedDataKeyInBase64
 */
func (useCase *reedUseCase) Encrypt(encryptPayload *model.EncryptPayload, kmsClient *mykms.KMS) (*model.CipherAndKey, error) {

	var additionalData []byte = nil

	plaintextBinary := []byte(encryptPayload.Plaintext)

	dataKey, err := useCase.generateDataKey(encryptPayload, kmsClient)
	if err != nil {
		return nil, err
	}

	ciphertext, err := cryptutil.EncryptWithGCM(dataKey.Plaintext, plaintextBinary, additionalData)
	if err != nil {
		return nil, err
	}

	cipherAndKey, err := encodeToBase64(ciphertext, dataKey.Ciphertext)
	if err != nil {
		return nil, err
	}

	err = useCase.storeEncryptedDataKey(encryptPayload, cipherAndKey)
	if err != nil {
		return nil, err
	}

	return cipherAndKey, nil

}

/*
 * Ensure Decrypt uses the cache (userId + "|" + fieldName + "|" + ciphertextInBase64) -> encryptedDataKeyInBase64
 * The cache is to be used if and only if decryptPayload.CipherAndKey.EncryptedDataKey == "".
 */
func (useCase *reedUseCase) Decrypt(decryptPayload *model.DecryptPayload, kmsClient *mykms.KMS) (string, error) {

	var additionalData []byte = nil

	dataKeyInBinary, err := useCase.getDecryptedDataKey(decryptPayload, kmsClient)
	if err != nil {
		return "", err
	}

	ciphertextInBinary, err := strutil.DecodeFromBase64([]byte(decryptPayload.CipherAndKey.Ciphertext))
	if err != nil {
		return "", err
	}

	plaintext, err := cryptutil.DecryptWithGCM(dataKeyInBinary, ciphertextInBinary, additionalData)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func (useCase *reedUseCase) generateDataKey(encryptPayload *model.EncryptPayload, kmsClient *mykms.KMS) (*mykms.KMSDatakey, error) {
	encryptionContext := generateEncryptionContext(encryptPayload)
	return kmsClient.GenerateDataKey(encryptionContext)
}

/*
 * Stores in the cache using the following map:
 * (userId + "|" + fieldName + "|" + ciphertextInBase64) -> encryptedDataKeyInBase64
 */
func (useCase *reedUseCase) storeEncryptedDataKey(encryptPayload *model.EncryptPayload, cipherAndKey *model.CipherAndKey) error {
	reedEntrykey := GetReedEntryKey(encryptPayload.UserId, encryptPayload.FieldName, cipherAndKey.Ciphertext)
	return StoreReedEntryInCacheNX(reedEntrykey, cipherAndKey, useCase.redis)
}

/*
 * Ensure Decrypt uses the cache (userId + "|" + fieldName + "|" + ciphertextInBase64) -> encryptedDataKeyInBase64
 * The cache is to be used if and only if decryptPayload.CipherAndKey.EncryptedDataKey == "".
 * Returns encryptedKey in binary form (i.e as originally generated) as a byte[]
 */
func (useCase *reedUseCase) getEncryptedKeyInBinary(decryptPayload *model.DecryptPayload) ([]byte, error) {
	var encryptedDataKeyEncoded64 string

	if decryptPayload.CipherAndKey.EncryptedDataKey != "" {
		encryptedDataKeyEncoded64 = decryptPayload.CipherAndKey.EncryptedDataKey
	} else {
		reedEntryKey := GetReedEntryKey(decryptPayload.UserId, decryptPayload.FieldName, decryptPayload.CipherAndKey.Ciphertext)
		var err error
		encryptedDataKeyEncoded64, err = GetReedEntryFromCache(reedEntryKey, useCase.redis)
		if err != nil {
			return nil, err
		}
	}

	encryptedDataKey, err := strutil.DecodeFromBase64([]byte(encryptedDataKeyEncoded64))
	if err != nil {
		return nil, err
	}

	return encryptedDataKey, nil
}

/*
 * Returns decryptedDataKey in binary form (i.e as originally generated) as a byte[]
 */
func (useCase *reedUseCase) getDecryptedDataKey(decryptPayload *model.DecryptPayload, kmsClient *mykms.KMS) ([]byte, error) {
	decryptionContext := generateDecryptionContext(decryptPayload)

	encryptedDataKeyInBinary, err := useCase.getEncryptedKeyInBinary(decryptPayload)
	if err != nil {
		return nil, err
	}

	dataKeyInBinary, err := kmsClient.DecryptDataKey(encryptedDataKeyInBinary, decryptionContext)
	if err != nil {
		return nil, err
	}

	return dataKeyInBinary, nil
}

func GetReedEntryFromCache(reedKey ReedEntryKey, redis map[string]string) (string, error) {
	value, ok := redis[reedKey.Key]
	if !ok {
		return "", errors.New("DataKey not found in store")
	}

	return value, nil
}

func StoreReedEntryInCacheNX(reedEntryKey ReedEntryKey, cipherAndKey *model.CipherAndKey, redis map[string]string) error {
	if _, ok := redis[reedEntryKey.Key]; ok {
		return errors.New("This userId,FieldName,cipherText combination already has an encrypted dataKey stored")
	}

	redis[reedEntryKey.Key] = cipherAndKey.EncryptedDataKey
	return nil
}

func GetReedEntryKey(userId, fieldName, ciphertextInBase64 string) ReedEntryKey {
	return ReedEntryKey{
		Key: (userId + ":" + fieldName + ":" + ciphertextInBase64),
	}
}

func encodeToBase64(ciphertextInBinary, encryptedDataKeyInBinary []byte) (*model.CipherAndKey, error) {
	ciphertextInBase64, err := strutil.EncodeToBase64(ciphertextInBinary)
	if err != nil {
		return nil, err
	}

	encryptedDataKeyInBase64, err := strutil.EncodeToBase64(encryptedDataKeyInBinary)
	if err != nil {
		return nil, err
	}

	return &model.CipherAndKey{
		Ciphertext:       string(ciphertextInBase64),
		EncryptedDataKey: string(encryptedDataKeyInBase64),
	}, nil
}

func makeEncryptionContext(userId, fieldName string) map[string]*string {
	encryptionContext := make(map[string]*string)
	str := fieldName
	encryptionContext[userId] = &str

	return encryptionContext
}
func generateEncryptionContext(encryptPayload *model.EncryptPayload) map[string]*string {
	return makeEncryptionContext(encryptPayload.UserId, encryptPayload.FieldName)
}

func generateDecryptionContext(decryptPayload *model.DecryptPayload) map[string]*string {
	return makeEncryptionContext(decryptPayload.UserId, decryptPayload.FieldName)
}
