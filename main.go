/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-02
 * Time: 01:34
 */
package main

import (
	"archive/zip"
	"fmt"
	"github.com/izghua/zgh/utils"
	"io"

	"github.com/izghua/zghua/conf"
	"github.com/izghua/zghua/entity"
	"github.com/izghua/zghua/my"
	"github.com/izghua/zghua/router"
	"os"
	"time"
)


func main() {

	//csrf
	//建表


	my.Testaa()

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
	//fmt.Println("没有错")
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

	//res,err := conf.ZHashId.Encode([]int{1})
	//re2,err := conf.ZHashId.Encode([]int{2})
	//re3,err := conf.ZHashId.Encode([]int{23})
	//re4,err := conf.ZHashId.Encode([]int{2322})
	//re5,err := conf.ZHashId.Encode([]int{2234324})
	//fmt.Println(res,re2,re3,re4,re5,err)
	//req,err := conf.ZHashId.DecodeWithError(res)
	//fmt.Println(req,err)
	//err = conf.CacheClient.Set("test1","11234",1*time.Minute).Err()
	//fmt.Println(err,"看缓存是否错了")
	//cache,err := conf.CacheClient.Get("test1").Result()
	//fmt.Println(cache,err,"解雇")
	fp := conf.SqlServer.Dialect().URI().DbName + ".sql"
	_ = os.Remove(fp)
	err = conf.SqlServer.DumpAllToFile("./backup/"+fp)
	fmt.Println(err)
	//StartTimer()
	//BoottimeTimingSettlement(ttt)
	//f3, err := os.Open("./backup")
	//if err != nil {
	//	utils.ZLog().Error("error",err.Error())
	//}
	//defer f3.Close()
	//var files = []*os.File{ f3}
	//dest := "./test.zip"
	//err = Compress(files, dest)
	//if err != nil {
	//	utils.ZLog().Error("error",err.Error())
	//}
	r := router.RouterInit()
	utils.ZLog().Warn(r.Run(":8081"))

}




// compress the files
// it is not designed by my
// i copy it what from https://studygolang.com/articles/7471
//files file arrays，it can be the diff file or directory
//dest be used to the addr
func Compress(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := compress(file, "", w)
		if err != nil {
			utils.ZLog().Error("message","compress the files","error",err.Error())
			return err
		}
	}
	return nil
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		utils.ZLog().Error("message","compress the files","error",err.Error())
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				utils.ZLog().Error("message","compress the files","error",err.Error())
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				utils.ZLog().Error("message","compress the files","error",err.Error())
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			utils.ZLog().Error("message","compress the files","error",err.Error())
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			utils.ZLog().Error("message","compress the files","error",err.Error())
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			utils.ZLog().Error("message","compress the files","error",err.Error())
			return err
		}
	}
	return nil
}


func StartTimer() {
	ticker := time.NewTicker(time.Second * 10)
	i := 0
	go func() {
		for _ = range ticker.C {
			i++
			utils.ZLog().Info("循环",i)
			fmt.Printf("ticked at %v", time.Now())
		}
	}()
}
