/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-04
 * Time: 22:53
 */
package conf

import (
	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
	"github.com/izghua/zgh/conn"
	"github.com/izghua/zgh/utils"
	"github.com/speps/go-hashids"
)

var (
	SqlServer *xorm.Engine
	ZHashId *hashids.HashID
	CacheClient *redis.Client
)


func init() {
	DbInit()
	AlarmInit()
	MailInit()
	ZLogInit()
	ZHashIdInit()
	RedisInit()
	JwtInit()
	QCaptchaInit()

	//err := utils.SendMail("2067930913@qq.com")
	//fmt.Println(err,"看发送邮件")
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

func ZLogInit() {
	zog := new(utils.ZLogParam)
	fileName := zog.SetFileName("zghua")
	err := zog.ZLogInit(fileName)
	if err != nil {
		utils.ZLog().Error(err.Error())
	}
}

func ZHashIdInit() {
	hd := new(utils.HashIdParams)
	salt := hd.SetHashIdSalt(HashIdSalt)
	hdLength := hd.SetHashIdLength(HashIdLength)
	zHashId,err := hd.HashIdInit(hdLength,salt)
	if err != nil {
		utils.ZLog().Error(err.Error())
	}
	ZHashId = zHashId

}

func RedisInit() {
	rc := new(conn.RedisClient)
	addr := rc.SetRedisAddr(RedisAddr)
	pwd := rc.SetRedisPwd(RedisPwd)
	db := rc.SetRedisDb(RedisDb)
	client,err := rc.RedisInit(addr,db,pwd)
	if err != nil {
		utils.ZLog().Error(err.Error())
	}
	CacheClient = client
}

func JwtInit() {
	jt := new(utils.JwtParam)
 	ad := jt.SetDefaultAudience("zgh")
 	jti := jt.SetDefaultJti("izghua")
 	iss := jt.SetDefaultIss("izghua")
 	sk := jt.SetDefaultSecretKey("izghua")
 	rc := jt.SetRedisCache(CacheClient)
 	_ = jt.JwtInit(ad,jti,iss,sk,rc)

}

func QCaptchaInit() {
	qc := new(utils.QQCaptcha)
	aid := qc.SetAid(QCaptchaAid)
	sk := qc.SetSecretKey(QCaptchaSecreptKey)
	_ = qc.QQCaptchaInit(aid,sk)
}