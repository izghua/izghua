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
)

var (
	SqlServer *xorm.Engine
)


func InitDefault() {

	sp := new(conn.Sp)
	sql := sp.SetDbHost(DbHost).SetDbPort(DbPort).SetDbDataBase(DbDataBase).SetDbUserName(DbUser).SetDbPassword(DbPassword)

	SqlServer = conn.InitMysql(sql)
}

