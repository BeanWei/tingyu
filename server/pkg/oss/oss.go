package oss

import (
	"context"
	"mime/multipart"

	"github.com/BeanWei/tingyu/g"
)

const (
	OSSTypeLocal string = "local"
	OSSTypeAWS   string = "aws"
)

type OSS interface {
	PutObject(ctx context.Context, file *multipart.FileHeader) (output *PutObjectOutput, err error)
}

type PutObjectOutput struct {
	URL       string `json:"url"`
	Filename  string `json:"filename"`
	OSSType   string `json:"oss_type"`
	OSSDomain string `json:"oss_domain"`
	OSSBucket string `json:"oss_bucket"`
}

func New() OSS {
	switch g.Cfg().OSS.Type {
	default:
		return &local{
			Domain: g.Cfg().OSS.Local.Domain,
			Path:   g.Cfg().OSS.Local.Path,
		}
	}
}
