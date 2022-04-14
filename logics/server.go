package logics

import (
	"context"
	"github.com/ropon/cd_cmdb/client/awsyun"
)

type ServerReq struct {
	Region   string `json:"region" form:"region" binding:"required"`
	PageSize int    `json:"page_size" form:"page_size"`
	PageNum  int    `json:"page_num" form:"page_num"`
}

type ServerRes struct {
	TotalCount int         `json:"total_count"`
	ServerList interface{} `json:"server_list"`
}

func GetServerS(ctx context.Context, req *ServerReq) (*ServerRes, error) {
	//ecsRes, err := aliyun.ListEcs(req.Region, req.PageSize, req.PageNum)
	//if err != nil {
	//	return nil, err
	//}
	//res := new(ServerRes)
	//res.TotalCount = ecsRes.TotalCount
	//res.ServerList = ecsRes.Instances.Instance

	//cvmRes, err := txyun.ListCvm(req.Region, req.PageSize, req.PageNum)
	//if err != nil {
	//	return nil, err
	//}
	//res := new(ServerRes)
	//res.TotalCount = int(*cvmRes.Response.TotalCount)
	//res.ServerList = cvmRes.Response.InstanceSet

	ec2Res, err := awsyun.ListEc2(req.Region, req.PageSize, req.PageNum)
	if err != nil {
		return nil, err
	}
	res := new(ServerRes)
	res.TotalCount = 0
	res.ServerList = ec2Res.Reservations

	//ecsRes, err := hwyun.ListEcs(req.Region, req.PageSize, req.PageNum)
	//if err != nil {
	//	return nil, err
	//}
	//res := new(ServerRes)
	//res.TotalCount = int(*ecsRes.Count)
	//res.ServerList = ecsRes.Servers

	//bssRes, err := hwyun.ListBss(req.Region, req.PageSize, req.PageNum)
	//if err != nil {
	//	return nil, err
	//}
	//res := new(ServerRes)
	//res.TotalCount = int(*bssRes.TotalCount)
	//res.ServerList = bssRes.FeeRecords

	return res, nil
}
