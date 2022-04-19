package txyun

import (
	"github.com/ropon/cd_cmdb/conf"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

var (
	txClient *TxClient
)

func ListCvm(region string, pageSize, pageNum int) (*cvm.DescribeInstancesResponse, error) {
	if txClient == nil {
		txClient = NewTxClient(region, conf.GetCfgByExternalName("txKey"),
			conf.GetCfgByExternalName("txSecret"))
	}
	client, err := txClient.CvmClient()
	if err != nil {
		return nil, err
	}

	request := cvm.NewDescribeInstancesRequest()
	response, err := client.DescribeInstances(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
