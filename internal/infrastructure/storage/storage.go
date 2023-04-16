package storage

import (
	"bytes"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/minio/minio-go"
	"github.com/timickb/transport-sound/internal/domain"
	"github.com/timickb/transport-sound/internal/interfaces"
	"io"
	"strings"
)

type Params struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	// RequiredBuckets contains buckets names which will be created if they don't exist.
	RequiredBuckets []string
}

type Storage struct {
	client *minio.Client
	log    interfaces.Logger
}

func New(log interfaces.Logger, p Params) (*Storage, error) {
	mc, err := minio.New(p.Endpoint, p.AccessKey, p.SecretKey, false)
	if err != nil {
		return nil, fmt.Errorf("err create minio client: %w", err)
	}

	for _, bName := range p.RequiredBuckets {
		ok, err := mc.BucketExists(bName)
		if err != nil {
			return nil, fmt.Errorf("err create minio client: %w", err)
		}
		if !ok {
			if err := mc.MakeBucket(bName, ""); err != nil {
				return nil, fmt.Errorf("err create minio client: %w", err)
			}
		}
		log.Info("created minio bucket: ", bName)
	}

	return &Storage{
		client: mc,
		log:    log,
	}, nil
}

func (s *Storage) CreateFile(bucket string, file *domain.File) error {
	s.log.Info("create minio object in bucket ", bucket, " with name ", file.Name())

	_, err := s.client.PutObject(bucket, file.Name(), bytes.NewReader(file.Bytes), file.Size, minio.PutObjectOptions{})
	if err != nil {
		return fmt.Errorf("err put minio object: %w", err)
	}

	s.log.Info("object created")
	return nil
}

func (s *Storage) DeleteFile(bucket, filename string) error {
	s.log.Info("remove minio object in bucket ", bucket, " with name ", filename)

	if err := s.client.RemoveObject(bucket, filename); err != nil {
		return fmt.Errorf("err remove minio object: %w", err)
	}
	s.log.Info("object removed")
	return nil
}

func (s *Storage) GetFile(bucket, filename string) (*domain.File, error) {
	s.log.Info("get minio object in bucket ", bucket, " with name ", filename)

	obj, err := s.client.GetObject(bucket, filename, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer obj.Close()

	info, err := obj.Stat()
	if err != nil {
		return nil, err
	}

	bytes := make([]byte, info.Size)
	if _, err := obj.Read(bytes); err != nil && err != io.EOF {
		return nil, fmt.Errorf("err read minio object: %w", err)
	}

	split := strings.Split(info.Key, ".")
	if len(split) != 2 {
		return nil, errors.New("err read minio object: invalid key")
	}

	return &domain.File{
		Bytes: bytes,
		Id:    split[0],
		Ext:   split[1],
		Hash:  md5.Sum(bytes),
		Size:  info.Size,
	}, nil
}
