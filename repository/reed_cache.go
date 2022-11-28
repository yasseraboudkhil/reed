package repository

import (
	model "github.com/yasseraboudkhil/reed/proto/v1"
)

type ReedCache interface {
	StoreReedEntry(reedEntryKey ReedEntryKey, cipherAndKey *model.CipherAndKey) error
	GetReedEntry(reedKey ReedEntryKey) (string, error)
}
