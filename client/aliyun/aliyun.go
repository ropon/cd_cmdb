package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/green"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
)

type AliClient struct {
	Region       string
	AccessKey    string
	AccessSecret string

	ecsConn   *ecs.Client
	rdsConn   *rds.Client
	bssConn   *bssopenapi.Client
	greenConn *green.Client
}

func NewAliClient(region, accessKey, accessSecret string) *AliClient {
	return &AliClient{
		Region:       region,
		AccessKey:    accessKey,
		AccessSecret: accessSecret,
	}
}

// EcsClient 客户端
func (c *AliClient) EcsClient() (*ecs.Client, error) {
	if c.ecsConn != nil {
		return c.ecsConn, nil
	}

	client, err := ecs.NewClientWithAccessKey(c.Region, c.AccessKey, c.AccessSecret)
	if err != nil {
		return nil, err
	}
	c.ecsConn = client
	return client, nil
}

// RdsClient 客户端
func (c *AliClient) RdsClient() (*rds.Client, error) {
	if c.rdsConn != nil {
		return c.rdsConn, nil
	}

	client, err := rds.NewClientWithAccessKey(c.Region, c.AccessKey, c.AccessSecret)
	if err != nil {
		return nil, err
	}
	c.rdsConn = client
	return client, nil
}

// BssClient 客户端
func (c *AliClient) BssClient() (*bssopenapi.Client, error) {
	if c.bssConn != nil {
		return c.bssConn, nil
	}

	client, err := bssopenapi.NewClientWithAccessKey(c.Region, c.AccessKey, c.AccessSecret)
	if err != nil {
		return nil, err
	}
	c.bssConn = client
	return client, nil
}
