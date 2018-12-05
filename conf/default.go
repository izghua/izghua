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
	sp := new(conn.Sp)
	dbUser := sp.SetDbUserName(DbUser)
	dbPwd := sp.SetDbPassword(DbPassword)
	dbPort := sp.SetDbPort(DbPort)
	dbHost := sp.SetDbHost(DbHost)
	dbdb := sp.SetDbDataBase(DbDataBase)
	SqlServer = conn.InitMysql(dbUser,dbPwd,dbPort,dbHost,dbdb)

	alarm := new(utils.AlarmParam)
	alarmT := alarm.SetType("mail")
	alarm.Options(alarmT)
}

