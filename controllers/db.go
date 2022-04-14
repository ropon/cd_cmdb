package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ropon/cd_cmdb/logics"
	"github.com/ropon/cd_cmdb/utils"
)

// GetDbs 获取数据库列表接口
// @Summary 获取数据库列表接口
// @Description 获取数据库列表接口
// @Tags 数据库相关接口
// @Produce application/json
// @Param data query logics.ServiceReq true "请求参数"
// @Success 200 {object} logics.ServiceRes "数据库列表返回结果"
// @Router /cd_cmdb/api/v1/db [get]
func GetDbs(c *gin.Context) {
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

// GetDb 获取单个数据库接口
// @Summary 获取单个数据库接口
// @Description 获取单个数据库接口
// @Tags 数据库相关接口
// @Produce application/json
// @Param id path uint true "id"
// @Success 200 {object} models.Service "数据库返回结果"
// @Router /cd_cmdb/api/v1/db/{id} [get]
func GetDb(c *gin.Context) {
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
