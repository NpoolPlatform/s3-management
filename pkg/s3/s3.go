package s3

import (
	"bytes"
	"context"
	"encoding/base64"
	"io/ioutil"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/s3-management/message/npool"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"golang.org/x/xerrors"
)

var ServiceName string

type Config struct {
	AccessKey string
	SecretKey string
	EndPoint  string
	Region    string
	Bucket    string
}

func getConfig() *Config {
	ServiceName = config.GetStringValueWithNameSpace("", config.KeyHostname)
	return &Config{
		AccessKey: config.GetStringValueWithNameSpace(ServiceName, "s3_access_key"),
		SecretKey: config.GetStringValueWithNameSpace(ServiceName, "s3_secret_key"),
		EndPoint:  config.GetStringValueWithNameSpace(ServiceName, "s3_endpoint"),
		Region:    config.GetStringValueWithNameSpace(ServiceName, "s3_region"),
		Bucket:    config.GetStringValueWithNameSpace(ServiceName, "s3_bucket"),
	}
}

func newS3(s3Config *Config) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(s3Config.AccessKey, s3Config.SecretKey, ""),
		Endpoint:         aws.String(s3Config.EndPoint),
		Region:           aws.String(s3Config.Region),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(false),
	})
	if err != nil {
		return nil, xerrors.Errorf("new session error: %v", err)
	}
	return sess, nil
}

func UploadImgToS3(ctx context.Context, in *npool.UploadImgToS3Request) (*npool.UploadImgToS3Response, error) {
	s3Config := getConfig()
	sess, err := newS3(s3Config)
	if err != nil {
		return nil, err
	}
	encodeImg := base64.StdEncoding.EncodeToString([]byte(in.ImgBase64))
	s3Key := in.ImgType + in.UserID
	s3Body := bytes.NewReader([]byte(encodeImg))

	service := s3.New(sess)
	_, err = service.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s3Config.Bucket),
		Key:    aws.String("kyc/" + s3Key),
		Body:   s3Body,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail to upload img to s3: %v", err)
	}
	return &npool.UploadImgToS3Response{
		Info: s3Key,
	}, nil
}

func GetImgFromS3(ctx context.Context, in *npool.GetImgFromS3Request) (*npool.GetImgFromS3Response, error) {
	s3Config := getConfig()
	sess, err := newS3(s3Config)
	if err != nil {
		return nil, err
	}

	resp, err := s3.New(sess).GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s3Config.Bucket),
		Key:    aws.String("kyc/" + in.ImgID),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail to get img from s3: %v", err)
	}
	if resp == nil || resp.Body == nil {
		return nil, xerrors.Errorf("empty response")
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, xerrors.Errorf("fail to get resp data: %v", err)
	}

	decodeImg, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return nil, xerrors.Errorf("fail to decode img: %v", err)
	}
	return &npool.GetImgFromS3Response{
		Info: string(decodeImg),
	}, nil
}
