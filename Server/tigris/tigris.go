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
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	cfg "github.com/elc49/copod/config"
	"github.com/elc49/copod/logger"
	"github.com/sirupsen/logrus"
)

var T tigris

type tigris interface {
	Upload(context.Context, multipart.File, *multipart.FileHeader) (*string, error)
	DeleteObjects(context.Context) error
}

type tigrisClient struct {
	log      *logrus.Logger
	uploader *manager.Uploader
	s3Client *s3.Client
}

func New() {
	log := logger.GetLogger()
	uploaderSdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(
		credentials.NewStaticCredentialsProvider(
			cfg.C.Tigris.AccessKeyId,
			cfg.C.Tigris.SecretAccessKey,
			"",
		),
	))
	if err != nil {
		log.WithError(err).Fatalln("tigris: LoadDefaultConfig")
	}

	c := s3.NewFromConfig(uploaderSdkConfig, func(o *s3.Options) {
		o.Region = cfg.C.Tigris.Region
		o.BaseEndpoint = aws.String(cfg.C.Tigris.S3Endpoint)
	})

	T = &tigrisClient{log, manager.NewUploader(c), c}
}

func (tc *tigrisClient) Upload(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (*string, error) {
	buf := make([]byte, fileHeader.Size)
	file.Read(buf)
	params := &s3.PutObjectInput{
		Bucket: aws.String(cfg.C.Tigris.BucketName),
		Key:    aws.String(fileHeader.Filename),
		Body:   bytes.NewReader(buf),
	}

	res, err := tc.uploader.Upload(ctx, params)
	if err != nil {
		tc.log.WithError(err).WithFields(logrus.Fields{"file_size": fileHeader.Size, "file_name": fileHeader.Filename}).Errorf("tigris: Upload")
		return nil, err
	}

	return &res.Location, nil
}

func (tc *tigrisClient) DeleteObjects(ctx context.Context) error {
	// Get bucket content
	content, err := tc.s3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(cfg.C.Tigris.BucketName),
	})
	if err != nil {
		tc.log.WithError(err).Errorf("tigris: ListObjectsV2")
		return err
	}

	// Prepare object
	var objectsToDelete []types.ObjectIdentifier
	for _, obj := range content.Contents {
		objectsToDelete = append(objectsToDelete, types.ObjectIdentifier{
			Key: obj.Key,
		})
	}

	// Proceed to delete
	if len(objectsToDelete) > 0 {
		_, err := tc.s3Client.DeleteObjects(ctx, &s3.DeleteObjectsInput{
			Bucket: aws.String(cfg.C.Tigris.BucketName),
			Delete: &types.Delete{
				Objects: objectsToDelete,
				Quiet:   aws.Bool(true),
			},
		})
		if err != nil {
			tc.log.WithError(err).Error("tigris: s3.DeleteObjects")
			return err
		}
		tc.log.Printf("Deleted %d objects from bucket %s\n", len(objectsToDelete), cfg.C.Tigris.BucketName)
	} else {
		tc.log.Infoln("No objects to delete")
	}

	return nil
}
