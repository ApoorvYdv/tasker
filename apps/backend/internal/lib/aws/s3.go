package aws

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ApoorvYdv/go-tasker/internal/server"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	client *s3.Client
	server *server.Server
}

func NewS3Client(server *server.Server, cfg aws.Config) *S3Client {
	return &S3Client{
		client: s3.NewFromConfig(cfg),
		server: server,
	}
}

func (s *S3Client) UploadFile(ctx context.Context, bucket string, fileKey string, file io.Reader) (string, error) {
	var buffer bytes.Buffer
	_, err := io.Copy(&buffer, file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	_, err = s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(fileKey),
		Body:        bytes.NewReader(buffer.Bytes()),
		ContentType: aws.String(http.DetectContentType(buffer.Bytes())),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %w", err)
	}

	return fileKey, nil
}

func (s *S3Client) GetPresignedUrl(ctx context.Context, bucket string, objectKey string) (string, error) {
	presignedClient := s3.NewPresignClient(s.client)

	expiration := time.Minute * 15

	presignedUrl, err := presignedClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	}, s3.WithPresignExpires(expiration))
	if err != nil {
		return "", fmt.Errorf("Failed to generate presigned URL: %w", err)
	}

	return presignedUrl.URL, nil
}

func (s *S3Client) DeleteFile(ctx context.Context, bucket string, objectKey string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return fmt.Errorf("Failed to delete file: %w", err)
	}

	return nil
}
