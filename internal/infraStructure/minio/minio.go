package minioUpload

import (
	"api/pkg/config"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"mime/multipart"
	"net/url"
	"time"
)

type Minio interface {
	PutImage(objectName string, buffer multipart.File, objectSize int64) error
	RemoveImage(objectName string) bool
	GetImage(objectName string) (*url.URL, error)
}

type Config struct {
	Endpoint    string
	AccessKey   string
	SecretKey   string
	Context     context.Context
	BucketName  string
	ContentType string
}

func ConfigDefault(params ...string) Config {
	c := config.GetConf()

	cfg := Config{}

	if cfg.Endpoint == "" {
		cfg.Endpoint = c.Minio.EndPoint
	}
	if cfg.AccessKey == "" {
		cfg.AccessKey = c.Minio.AccessKey
	}
	if cfg.SecretKey == "" {
		cfg.SecretKey = c.Minio.SecretKey
	}
	if cfg.Context == nil {
		cfg.Context = context.Background()
	}

	cfg.ContentType = params[1]

	cfg.BucketName = params[0]
	return cfg
}
func (c Config) PutImage(objectName string, buffer multipart.File, objectSize int64) error {
	cli, cliError := minio.New(c.Endpoint, &minio.Options{Secure: false, Creds: credentials.NewStaticV4(c.AccessKey, c.SecretKey, "")})

	if cliError != nil {
		return cliError
	}
	exist, _ := cli.BucketExists(c.Context, c.BucketName)

	if !exist {
		bucketError := cli.MakeBucket(c.Context, c.BucketName, minio.MakeBucketOptions{})
		if bucketError != nil {

		} else {
			_, err := cli.PutObject(c.Context, c.BucketName, objectName, buffer, objectSize, minio.PutObjectOptions{ContentType: c.ContentType})
			if err != nil {
				return err
			}
		}
	}
	_, err := cli.PutObject(c.Context, c.BucketName, objectName, buffer, objectSize, minio.PutObjectOptions{ContentType: c.ContentType})
	if err != nil {
		return err
	}
	return nil
}
func (c Config) RemoveImage(objectName string) bool {
	cli, cliError := minio.New(c.Endpoint, &minio.Options{Secure: false, Creds: credentials.NewStaticV4(c.AccessKey, c.SecretKey, "")})
	if cliError != nil {
		return false
	}
	err := cli.RemoveObject(c.Context, c.BucketName, objectName, minio.RemoveObjectOptions{GovernanceBypass: true})
	if err != nil {
		return false
	}
	return true
}
func (c Config) GetImage(objectName string) (*url.URL, error) {
	cli, cliError := minio.New(c.Endpoint, &minio.Options{Secure: false, Creds: credentials.NewStaticV4(c.AccessKey, c.SecretKey, "")})
	if cliError != nil {
		return nil, cliError
	}

	getObject, err := cli.PresignedGetObject(c.Context, c.BucketName, objectName, time.Second*24*60*60, nil)
	if err != nil {
		return nil, err
	}
	return getObject, nil
}
