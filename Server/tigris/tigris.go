package tigris

import (
	"bytes"
	"context"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	cfg "github.com/elc49/copod/config"
	"github.com/elc49/copod/logger"
	"github.com/sirupsen/logrus"
)

var T tigris

type tigris interface {
	Upload(context.Context, multipart.File, *multipart.FileHeader) (*string, error)
}

type tigrisClient struct {
	log *logrus.Logger
	s3  *manager.Uploader
}

func New() {
	log := logger.GetLogger()
	s3SdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(
		credentials.NewStaticCredentialsProvider(
			cfg.C.Tigris.AccessKeyId,
			cfg.C.Tigris.SecretAccessKey,
			"",
		),
	))
	if err != nil {
		log.WithError(err).Fatalln("tigris: LoadDefaultConfig")
	}

	c := s3.NewFromConfig(s3SdkConfig, func(o *s3.Options) {
		o.Region = cfg.C.Tigris.Region
		o.BaseEndpoint = aws.String(cfg.C.Tigris.S3Endpoint)
	})

	T = &tigrisClient{log, manager.NewUploader(c)}
}

func (tc *tigrisClient) Upload(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (*string, error) {
	buf := make([]byte, fileHeader.Size)
	file.Read(buf)
	params := &s3.PutObjectInput{
		Bucket: aws.String(cfg.C.Tigris.BucketName),
		Key:    aws.String(fileHeader.Filename),
		Body:   bytes.NewReader(buf),
	}

	res, err := tc.s3.Upload(ctx, params)
	if err != nil {
		tc.log.WithError(err).WithFields(logrus.Fields{"file_size": fileHeader.Size, "file_name": fileHeader.Filename}).Errorf("tigris: Upload")
		return nil, err
	}

	return &res.Location, nil
}
