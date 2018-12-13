/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-02
 * Time: 01:34
 */
package main

import (
	"fmt"
	"github.com/izghua/zgh/utils"
	"izghua/zghua/conf"
	"izghua/zghua/entity"
	"time"
)


func main() {

	//conf.InitDefault()
	test := new(entity.Test1)
	_,err := conf.SqlServer.Where("id = ?",1).Get(test)
	if err != nil {
		utils.ZLog().Error("error",err.Error())
	}

	utils.Alarm("!")


	subject := "您好"
	text := "你好！"
	body := `
    <html>
    <body>
    <h3>
    "Kubernetes is an open source system for managing containerized applications across multiple hosts; providing basic mechanisms for deployment, maintenance, and scaling of applications.

Kubernetes builds upon a decade and a half of experience at Google running production workloads at scale using a system called Borg, combined with best-of-breed ideas and practices from the community.

Kubernetes is hosted by the Cloud Native Computing Foundation (CNCF). If you are a company that wants to help shape the evolution of technologies that are container-packaged, dynamically-scheduled and microservices-oriented, consider joining the CNCF. For details about who's involved and how Kubernetes plays a role, read the CNCF announcement."` + text + `
    </h3>
    </body>
    </html>
    `
	err = utils.SendMail("xzghua@gmail.com",subject,body)
	if err != nil {
		fmt.Println(err.Error(),"有错")
	}
	fmt.Println("没有错")
	//fmt.Println(err.Error())
	//
	//// 邮箱账号
	//user := "test@g9zz.com"
	////注意，此处为授权码、不是密码
	//password := "1234abcd#"
	////smtp地址及端口
	//host := "smtp.mxhichina.com:25"
	////接收者，内容可重复，邮箱之间用；隔开
	//to := "xzghua@gmail.com"
	////邮件主题
	////邮件内容

	//fmt.Println("send email")
	//执行逻辑函数
	//err = SendMail(user, password, host, to, subject, body, "html2")
	//if err != nil {
	//	fmt.Println("发送邮件失败!")
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("发送邮件成功!")
	//}

	//c,err := smtp.Dial(host + "1")
	//fmt.Println(c.Verify(host),c.Close(),err)

	res,err := conf.ZHashId.Encode([]int{1})
	re2,err := conf.ZHashId.Encode([]int{2})
	re3,err := conf.ZHashId.Encode([]int{23})
	re4,err := conf.ZHashId.Encode([]int{2322})
	re5,err := conf.ZHashId.Encode([]int{2234324})
	fmt.Println(res,re2,re3,re4,re5,err)
	req,err := conf.ZHashId.DecodeWithError(res)
	fmt.Println(req,err)
	err = conf.CacheClient.Set("test1","11234",1*time.Minute).Err()
	fmt.Println(err,"看缓存是否错了")
	cache,err := conf.CacheClient.Get("test1").Result()
	fmt.Println(cache,err,"解雇")
	}

