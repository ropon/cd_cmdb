package txyun

import (
	billing "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"
	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

type TxClient struct {
	Region    string
	SecretId  string
	SecretKey string

	cvmConn  *cvm.Client
	cdbConn  *cdb.Client
	billConn *billing.Client
}

func NewTxClient(region, secretId, secretKey string) *TxClient {
	return &TxClient{
		Region:    region,
		SecretId:  secretId,
		SecretKey: secretKey,
	}
}

// CvmClient 客户端
func (c *TxClient) CvmClient() (*cvm.Client, error) {
	if c.cvmConn != nil {
		return c.cvmConn, nil
	}

	credential := common.NewCredential(
		c.SecretId,
		c.SecretKey,
	)
	cpf := profile.NewClientProfile()
	client, err := cvm.NewClient(credential, c.Region, cpf)
	if err != nil {
		return nil, err
	}
	c.cvmConn = client
	return client, nil
}

// CdbClient 客户端
func (c *TxClient) CdbClient() (*cdb.Client, error) {
	if c.cvmConn != nil {
		return c.cdbConn, nil
	}

	credential := common.NewCredential(
		c.SecretId,
		c.SecretKey,
	)
	cpf := profile.NewClientProfile()
	client, err := cdb.NewClient(credential, c.Region, cpf)
	if err != nil {
		return nil, err
	}
	c.cdbConn = client
	return client, nil
}

// BillClient 客户端
func (c *TxClient) BillClient() (*billing.Client, error) {
	if c.cvmConn != nil {
		return c.billConn, nil
	}

	credential := common.NewCredential(
		c.SecretId,
		c.SecretKey,
	)
	cpf := profile.NewClientProfile()
	client, err := billing.NewClient(credential, c.Region, cpf)
	if err != nil {
		return nil, err
	}
	c.billConn = client
	return client, nil
}
