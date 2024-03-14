package main

import (
	"fmt"
	"github.com/ropon/cd_cmdb/conf"
	"github.com/ropon/cd_cmdb/routers"
)

// @title cd_cmdb
// @version 1.0
// @description 资产管理系统

// @contact.name Ropon
// @contact.url https://www.ropon.top
// @contact.email luopeng@ropon.top

// @license.name Apache 2.0
// @license.url https://www.apache.org/licenses/LICENSE-2.0.html

// @host cd-cmdb.in.ropon.top:8867
// @BasePath /
func main() {
	err := conf.Init()
	if err != nil {
		fmt.Printf("init failed, err: %v\n", err)
		return
	}

	//go logics.TestSendMail()

	routers.Run(conf.Cfg.Listen)
}
