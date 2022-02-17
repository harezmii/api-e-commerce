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
func (c Config) PutImage(objectName string, buffer multipart.File, objectSize int64) minio.UploadInfo {
	cli, _ := minio.New(c.Endpoint, &minio.Options{Secure: false, Creds: credentials.NewStaticV4(c.AccessKey, c.SecretKey, "")})
	object, err := cli.PutObject(c.Context, c.BucketName, objectName, buffer, objectSize, minio.PutObjectOptions{ContentType: c.ContentType})
	if err != nil {
		return minio.UploadInfo{}
	}
	return object
}
func (c Config) RemoveImage(objectName string) bool {
	cli, _ := minio.New(c.Endpoint, &minio.Options{Secure: false, Creds: credentials.NewStaticV4(c.AccessKey, c.SecretKey, "")})
	err := cli.RemoveObject(c.Context, c.BucketName, objectName, minio.RemoveObjectOptions{GovernanceBypass: true})
	if err != nil {
		return false
	}
	return true
}
func (c Config) GetImage(objectName string) *url.URL {
	cli, _ := minio.New(c.Endpoint, &minio.Options{Secure: false, Creds: credentials.NewStaticV4(c.AccessKey, c.SecretKey, "")})
	//object, err := cli.GetObject(c.Context, c.BucketName, objectName, minio.GetObjectOptions{})
	//if err != nil {
	//	return nil
	//}
	getObject, err := cli.PresignedGetObject(c.Context, c.BucketName, objectName, time.Second*24*60*60, nil)
	if err != nil {
		return nil
	}
	return getObject
}
