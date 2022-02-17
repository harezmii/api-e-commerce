package minioUpload

import (
	"api/pkg/config"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"mime/multipart"
)

type Config struct {
	Endpoint    string
	AccessKey   string
	SecretKey   string
	Context     context.Context
	BucketName  string
	ContentType string
}

func ConfigDefault(bucketName, contentType string) Config {
	cfg := Config{}

	if cfg.Endpoint == "" {
		cfg.Endpoint = config.GetEnvironment("ENDPOINT", config.STRING).(string)
	}
	if cfg.AccessKey == "" {
		cfg.AccessKey = config.GetEnvironment("ACCESS_KEY", config.STRING).(string)
	}
	if cfg.SecretKey == "" {
		cfg.SecretKey = config.GetEnvironment("SECRET_KEY", config.STRING).(string)
	}
	if cfg.Context == nil {
		cfg.Context = context.Background()
	}
	cfg.ContentType = contentType
	cfg.BucketName = bucketName
	return cfg
}
func (c Config) PutImage(objectName string, buffer multipart.File, objectSize int64) minio.UploadInfo {
	cli, _ := minio.New(config.GetEnvironment("ENDPOINT", config.STRING).(string), &minio.Options{Secure: false, Creds: credentials.NewStaticV4(config.GetEnvironment("ACCESS_KEY", config.STRING).(string), config.GetEnvironment("SECRET_KEY", config.STRING).(string), "")})
	object, err := cli.PutObject(c.Context, c.BucketName, objectName, buffer, objectSize, minio.PutObjectOptions{ContentType: c.ContentType})
	if err != nil {
		return minio.UploadInfo{}
	}
	return object
}
