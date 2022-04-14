package hwyun

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	"github.com/ropon/cd_cmdb/conf"
)

var (
	hwClient *HwClient
)

func ListEcs(region string, pageSize, pageNum int) (*model.ListServersDetailsResponse, error) {
	if hwClient == nil {
		hwClient = NewHwClient(region, conf.GetCfgByExternalName("hwKey"),
			conf.GetCfgByExternalName("hwSecret"))
	}
	client, err := hwClient.EcsClient()
	if err != nil {
		return nil, err
	}
	request := new(model.ListServersDetailsRequest)
	response, err := client.ListServersDetails(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
