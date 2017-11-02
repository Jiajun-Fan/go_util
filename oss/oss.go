package main

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

type Oss interface {
	io.Reader
	io.Writer
	io.Closer
}

type OssConfig struct {
	Type     string
	Key      string
	Secret   string
	Bucket   string
	FileName string
	EndPoint string
	Write    bool
}

func MakeOss(config OssConfig) Oss {
	if config.Type == "aliyun" {
		return NewAliyunOss(config)
	}
	Fatal(fmt.Sprintf("oss type '%s' is not implemented", config.Type))
	return nil
}

type AliyunOss struct {
	config OssConfig
	client *oss.Client
	bucket *oss.Bucket
	reader io.ReadCloser
	buffer bytes.Buffer
}

func (a *AliyunOss) Read(data []byte) (int, error) {
	if a.config.Write == true {
		Fatal("file is opened in write mode")
	}
	if a.reader == nil {
		if reader, err := a.bucket.GetObject(a.config.FileName); err == nil {
			a.reader = reader
		} else {
			Fatal(fmt.Sprintf("can't open file %s", a.config.FileName))
		}
	}
	return a.reader.Read(data)
}

func (a *AliyunOss) Write(input []byte) (int, error) {
	if a.config.Write == false {
		Fatal("file is opened in read mode")
	}
	a.buffer.Write(input)
	return len(input), nil
}

func (a *AliyunOss) Close() error {
	if a.config.Write {
		return a.bucket.PutObject(a.config.FileName, &a.buffer)
	} else {
		return a.reader.Close()
	}
}

func NewAliyunOss(config OssConfig) Oss {
	aliyun := &AliyunOss{}
	aliyun.config = config
	if client, err := oss.New(aliyun.config.EndPoint, aliyun.config.Key, aliyun.config.Secret); err != nil {
		Fatal(err.Error())
	} else {
		aliyun.client = client
	}
	if bucket, err := aliyun.client.Bucket(aliyun.config.Bucket); err != nil {
		Fatal(err.Error())
	} else {
		aliyun.bucket = bucket
	}
	return aliyun
}
