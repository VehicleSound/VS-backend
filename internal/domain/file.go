package domain

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"io"
)

type File struct {
	Bytes []byte
	Id    string
	Ext   string
	Hash  [16]byte
	Size  int64
}

func NewFile(ext string, reader io.Reader, size int64) (*File, error) {
	bytes := make([]byte, size)
	if _, err := reader.Read(bytes); err != nil {
		return nil, fmt.Errorf("err create file: %w", err)
	}

	return &File{
		Bytes: bytes,
		Id:    uuid.NewString(),
		Ext:   ext,
		Hash:  md5.Sum(bytes),
		Size:  size,
	}, nil
}

// Name returns file name in format [id].[ext]
func (f *File) Name() string {
	return fmt.Sprintf("%s.%s", f.Id, f.Ext)
}

func (f *File) HashString() string {
	return hex.EncodeToString(f.Bytes[:])
}

func (f *File) SetHashAsString(h string) error {
	md5Bytes, err := hex.DecodeString(h)
	if err != nil {
		return err
	}

	copy(f.Hash[:], md5Bytes)
	return nil
}
