package oss

import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"path"
	"time"

	"github.com/duke-git/lancet/v2/random"
)

var _ OSS = (*local)(nil)

type local struct {
	Domain string
	Path   string
}

func (s *local) PutObject(ctx context.Context, file *multipart.FileHeader) (*PutObjectOutput, error) {
	bucket := time.Now().Format("20060102")
	dir := path.Join(s.Path, bucket)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return nil, err
	}

	reader, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	uuid, _ := random.UUIdV4()
	filename := uuid + path.Ext(file.Filename)
	writer, err := os.Create(path.Join(dir, filename))
	if err != nil {
		return nil, err
	}
	defer writer.Close()

	_, err = io.Copy(writer, reader)
	if err != nil {
		return nil, err
	}
	return &PutObjectOutput{
		URL:       s.Domain + "/upload/" + bucket + "/" + filename,
		Filename:  filename,
		OSSType:   OSSTypeLocal,
		OSSDomain: s.Domain,
		OSSBucket: bucket,
	}, nil
}
