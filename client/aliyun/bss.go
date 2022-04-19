package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/ropon/cd_cmdb/conf"
)

func ListBssOverview(region, bssDate string, pageSize, pageNum int) (*bssopenapi.QueryBillOverviewResponse, error) {
	if aliClient == nil {
		aliClient = NewAliClient(region, conf.GetCfgByExternalName("aliKey"),
			conf.GetCfgByExternalName("aliSecret"))
	}
	client, err := aliClient.BssClient()
	if err != nil {
		return nil, err
	}
	request := bssopenapi.CreateQueryBillOverviewRequest()
	request.RegionId = region
	request.BillOwnerId = requests.NewInteger(1137658205106772)
	request.BillingCycle = bssDate
	response, err := client.QueryBillOverview(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
