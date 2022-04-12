package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ropon/cd_cmdb/logics"
	"github.com/ropon/cd_cmdb/utils"
)

// GetServers 获取服务器列表接口
// @Summary 获取服务器列表接口
// @Description 获取服务器列表接口
// @Tags 服务器相关接口
// @Produce application/json
// @Param data query logics.ServiceReq true "请求参数"
// @Success 200 {object} logics.ServiceRes "服务器列表返回结果"
// @Router /cd_cmdb/api/v1/server [get]
func GetServers(c *gin.Context) {
	req := new(logics.ServiceReq)
	if !checkData(c, req) {
		return
	}

	ctx := utils.ExtractStdContext(nil, c.Request.Header)
	resList, err := logics.GetServices(ctx, req)
	if err != nil {
		utils.GinErrRsp(c, utils.ErrCodeGeneralFail, err.Error())
		return
	}
	utils.GinOKRsp(c, resList, "获取列表成功")
}

// GetServer 获取单个服务器接口
// @Summary 获取单个服务器接口
// @Description 获取单个服务器接口
// @Tags 服务器相关接口
// @Produce application/json
// @Param id path uint true "id"
// @Success 200 {object} models.Service "服务器返回结果"
// @Router /cd_cmdb/api/v1/server/{id} [get]
func GetServer(c *gin.Context) {
	id, flag := checkParamsId(c)
	if !flag {
		return
	}

	res, err := logics.GetService(id)
	if err != nil {
		utils.GinErrRsp(c, utils.ErrCodeGeneralFail, err.Error())
		return
	}
	utils.GinOKRsp(c, res, "获取成功")
}
