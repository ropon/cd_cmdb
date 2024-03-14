package logics

import (
	"fmt"
	"github.com/ropon/cd_cmdb/client/awsyun"
)

func TestSendMail() {
	err := awsyun.SendMail("u-west-1")
	if err != nil {
		fmt.Println(err.Error())
	}
}
