package repository

import (
	model "github.com/yasseraboudkhil/reed/proto/v1"
)


New(cache ReedCache)
StoreReedEntry(reedEntryKey ReedEntryKey, cipherAndKey *model.CipherAndKey) error
GetReedEntry(reedKey ReedEntryKey) (string, error)