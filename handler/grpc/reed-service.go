package grpchandle

import (
	"context"

	"github.com/yasseraboudkhil/cloud-hsm/mykms"
	model "github.com/yasseraboudkhil/reed/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/yasseraboudkhil/reed/usecase"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

type ReedServiceServer struct {
	//redis       map[string]string // maps userId|fieldName -> encryptedDataKey
	kms         *mykms.KMS
	reedUseCase usecase.ReedUseCase
}

func (*ReedServiceServer) CheckAPIVersion(api string) error {

	if len(api) > 0 && api == apiVersion {
		return nil
	}

	return status.Errorf(codes.Unimplemented,
		"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
}
func NewReedServiceServer(kms *mykms.KMS) model.ReedServiceServer {
	redis := make(map[string]string)
	reedUseCase := usecase.NewReedUseCase(redis)
	return &ReedServiceServer{
		kms:         kms,
		reedUseCase: reedUseCase,
	}
}

func (s *ReedServiceServer) Encrypt(ctx context.Context, req *model.EncryptRequest) (*model.EncryptResponse, error) {

	if err := s.CheckAPIVersion(req.Api); err != nil {
		return nil, err
	}

	result, err := s.reedUseCase.Encrypt(req.EncryptPayload, s.kms)
	if err != nil {
		return nil, err
	}

	return &model.EncryptResponse{
		Api:          req.Api,
		CipherAndKey: result,
	}, nil

}

func (s *ReedServiceServer) Decrypt(ctx context.Context, req *model.DecryptRequest) (*model.DecryptResponse, error) {

	if err := s.CheckAPIVersion(req.Api); err != nil {
		return nil, err
	}

	result, err := s.reedUseCase.Decrypt(req.DecryptPayload, s.kms)
	if err != nil {
		return nil, err
	}

	return &model.DecryptResponse{
		Api:             req.Api,
		DecryptedCipher: result,
	}, nil

}
