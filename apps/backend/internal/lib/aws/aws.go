package aws

import (
	"context"
	"fmt"

	"github.com/ApoorvYdv/go-tasker/internal/server"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

type AWS struct {
	S3Client *S3Client
}

func NewAWS(server *server.Server) (*AWS, error) {
	awsConfig := server.Config.AWS

	configOptions := []func(*config.LoadOptions) error{
		config.WithRegion(awsConfig.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			awsConfig.AccessKeyID,
			awsConfig.SecretAccessKey,
			"",
		)),
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), configOptions...)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	return &AWS{
		S3Client: NewS3Client(server, cfg),
	}, nil
}
