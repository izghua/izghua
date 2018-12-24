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
	"github.com/izghua/zgh"
	"github.com/izghua/zgh/conn"
	"github.com/izghua/zgh/jwt"
	"github.com/izghua/zgh/utils/alarm"
	"github.com/izghua/zgh/utils/hashid"
	"github.com/izghua/zgh/utils/mail"
	"github.com/izghua/zgh/utils/qq_captcha"
	"github.com/speps/go-hashids"
)

var (
	SqlServer *xorm.Engine
	ZHashId *hashids.HashID
	CacheClient *redis.Client
)


func init() {
	AlarmInit()
	MailInit()
	ZHashIdInit()
	RedisInit()
	JwtInit()
	QCaptchaInit()

	//err := utils.SendMail("2067930913@qq.com")
	//fmt.Println(err,"看发送邮件")
}

func BackUpInit() {

}



func AlarmInit() {
	a := new(alarm.AlarmParam)
	alarmT := a.SetType(AlarmType)
	mailTo := a.SetMailTo("xzghua@gmail.com")
	err := a.AlarmInit(alarmT,mailTo)
	if err != nil {
		zgh.ZLog().Error(err.Error())
	}
}

func MailInit() {
	m := new(mail.EmailParam)
	mailUser := m.SetMailUser(MailUser)
	mailPwd := m.SetMailPwd(MailPwd)
	mailHost :=  m.SetMailHost(MailHost)
	err := m.MailInit(mailPwd,mailHost,mailUser)
	if err != nil {
		zgh.ZLog().Error(err.Error())
	}
}



func ZHashIdInit() {
	hd := new(hashid.HashIdParams)
	salt := hd.SetHashIdSalt(HashIdSalt)
	hdLength := hd.SetHashIdLength(HashIdLength)
	zHashId,err := hd.HashIdInit(hdLength,salt)
	if err != nil {
		zgh.ZLog().Error(err.Error())
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
		zgh.ZLog().Error(err.Error())
	}
	CacheClient = client
}

func JwtInit() {
	jt := new(jwt.JwtParam)
	ad := jt.SetDefaultAudience("zgh")
	jti := jt.SetDefaultJti("izghua")
	iss := jt.SetDefaultIss("izghua")
	sk := jt.SetDefaultSecretKey("izghua")
	rc := jt.SetRedisCache(CacheClient)
	_ = jt.JwtInit(ad,jti,iss,sk,rc)

}

func QCaptchaInit() {
	qc := new(qq_captcha.QQCaptcha)
	aid := qc.SetAid(QCaptchaAid)
	sk := qc.SetSecretKey(QCaptchaSecreptKey)
	_ = qc.QQCaptchaInit(aid,sk)
}