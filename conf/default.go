/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-04
 * Time: 22:53
 */
package conf

import (
	"github.com/go-xorm/xorm"
	"izghua/pkg/zgh/conn"
	"izghua/pkg/zgh/utils"
)

var (
	SqlServer *xorm.Engine
)


func InitDefault() {
	DbInit()
	AlarmInit()
	MailInit()
	ZlogInit()

	utils.ZLog().Info("kaiwanxiaone","叶落山城","有东西","还有东西")
}

func DbInit () {
	sp := new(conn.Sp)
	dbUser := sp.SetDbUserName(DbUser)
	dbPwd := sp.SetDbPassword(DbPassword)
	dbPort := sp.SetDbPort(DbPort)
	dbHost := sp.SetDbHost(DbHost)
	dbdb := sp.SetDbDataBase(DbDataBase)
	sqlServer,err := conn.InitMysql(dbUser,dbPwd,dbPort,dbHost,dbdb)
	SqlServer = sqlServer
	if err != nil {
		utils.ZLog().Error("有错误",err.Error())
	}
}

func AlarmInit() {
	alarm := new(utils.AlarmParam)
	alarmT := alarm.SetType(AlarmType)
	mailTo := alarm.SetMailTo("xzghua@gmail.com")
	err := alarm.AlarmInit(alarmT,mailTo)
	if err != nil {
		utils.ZLog().Error(err.Error())
	}
}

func MailInit() {
	mail := new(utils.EmailParam)
	mailUser := mail.SetMailUser(MailUser)
	mailPwd := mail.SetMailPwd(MailPwd)
	mailHost :=  mail.SetMailHost(MailHost)
	err := mail.MailInit(mailPwd,mailHost,mailUser)
	if err != nil {
		utils.ZLog().Error(err.Error())
	}
}

func ZlogInit() {
	zog := new(utils.ZLogParam)
	err := zog.ZLogInit()
	if err != nil {
		utils.ZLog().Error(err.Error())
	}
}