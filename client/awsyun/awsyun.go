package awsyun

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type AwsClient struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string

	ec2Conn *ec2.Client
}

func NewAwsClient(region, accessKeyID, secretAccessKey string) *AwsClient {
	return &AwsClient{
		Region:          region,
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
	}
}

// Ec2Client 客户端
func (c *AwsClient) Ec2Client() (*ec2.Client, error) {
	if c.ec2Conn != nil {
		return c.ec2Conn, nil
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(c.AccessKeyID, c.SecretAccessKey, "")),
		config.WithRegion(c.Region),
	)
	if err != nil {
		return nil, err
	}
	c.ec2Conn = ec2.NewFromConfig(cfg)
	return c.ec2Conn, nil
}
