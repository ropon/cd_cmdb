package hwyun

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
	bss "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2"
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	ecsregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
	rds "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rds/v3"
	rdsregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rds/v3/region"
)

type HwClient struct {
	Region    string
	AccessKey string
	SecretKey string

	ecsConn *ecs.EcsClient
	rdsConn *rds.RdsClient
	bssConn *bss.BssClient
}

func NewHwClient(region, accessKey, secretKey string) *HwClient {
	return &HwClient{
		Region:    region,
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
}

func (c *HwClient) Credentials() basic.Credentials {
	auth := basic.NewCredentialsBuilder().
		WithAk(c.AccessKey).
		WithSk(c.SecretKey).
		Build()
	return auth
}

func (c *HwClient) GlobalCredentials() global.Credentials {
	auth := global.NewCredentialsBuilder().
		WithAk(c.AccessKey).
		WithSk(c.SecretKey).
		Build()
	return auth
}

// EcsClient 客户端
func (c *HwClient) EcsClient() (*ecs.EcsClient, error) {
	if c.ecsConn != nil {
		return c.ecsConn, nil
	}

	client := ecs.EcsClientBuilder().
		WithRegion(ecsregion.ValueOf(c.Region)).
		WithCredential(c.Credentials()).
		Build()
	c.ecsConn = ecs.NewEcsClient(client)
	return c.ecsConn, nil
}

// RdsClient 客户端
func (c *HwClient) RdsClient() (*rds.RdsClient, error) {
	if c.rdsConn != nil {
		return c.rdsConn, nil
	}

	client := rds.RdsClientBuilder().
		WithRegion(rdsregion.ValueOf(c.Region)).
		WithCredential(c.Credentials()).
		Build()
	c.rdsConn = rds.NewRdsClient(client)
	return c.rdsConn, nil
}

// BssClient 客户端
func (c *HwClient) BssClient() (*bss.BssClient, error) {
	if c.bssConn != nil {
		return c.bssConn, nil
	}

	client := bss.BssClientBuilder().
		WithEndpoint("bss.myhuaweicloud.com").
		WithCredential(c.GlobalCredentials()).
		Build()
	c.bssConn = bss.NewBssClient(client)
	return c.bssConn, nil
}
