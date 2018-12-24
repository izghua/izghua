/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-02
 * Time: 01:34
 */
package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/izghua/zgh"
	"github.com/izghua/zgh/conn"
	"github.com/izghua/zgh/utils/backup"
	"github.com/izghua/zghua/conf"
	"github.com/izghua/zghua/router"
	"os"
	"time"
)

var SqlServer *xorm.Engine

func main() {
	//日志一定要最先初始化
	ZLogInit()
	DbInit()
	//csrf
	//建表

	bp := new(backup.BackUpParam)
	f3, err := os.Open("./backup")
	if err != nil {
		zgh.ZLog().Error("error",err.Error())
	}
	defer f3.Close()
	var files = []*os.File{ f3}
	dest := "./backup"
	backu := bp.SetFilePath("./backup/").SetFiles(files).SetDest(dest).SetDuration(time.Second * 20)
	err = backu.Backup()
	if err != nil {
		fmt.Println("备份出了问题",err.Error())
		//zgh.ZLog().Error("备份出了问题",err.Error())
	} else {
		fmt.Println("嘿嘿,备份没有问题","空空")
		//zgh.ZLog().Info("嘿嘿,备份没有问题","空空")
	}

	//fp := conf.SqlServer.Dialect().URI().DbName + ".sql"
	//fmt.Println(fp,"这hi什么")
	//_ = os.Remove(fp)
	//err = conf.SqlServer.DumpAllToFile("./backup/"+fp)
	//fmt.Println(err,"备份")

	r := router.RoutersInit()
	_ = r.Run(":8081")

	//my.Testaa()

	//conf.InitDefault()
	//test := new(entity.Test1)
	//_,err := conf.SqlServer.Where("id = ?",1).Get(test)
	//if err != nil {
	//	utils.ZLog().Error("error",err.Error())
	//}
	//
	//utils.Alarm("!")



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

	//utils.ZLog().Warn(r.Run(":8081"))
}




// compress the files
// it is not designed by my
// i copy it what from https://studygolang.com/articles/7471
//files file arrays，it can be the diff file or directory
//dest be used to the addr
//func Compress(files []*os.File, dest string) error {
//	d, _ := os.Create(dest)
//	defer d.Close()
//	w := zip.NewWriter(d)
//	defer w.Close()
//	for _, file := range files {
//		err := compress(file, "", w)
//		if err != nil {
//			utils.ZLog().Error("message","compress the files","error",err.Error())
//			return err
//		}
//	}
//	return nil
//}
//
//func compress(file *os.File, prefix string, zw *zip.Writer) error {
//	info, err := file.Stat()
//	if err != nil {
//		utils.ZLog().Error("message","compress the files","error",err.Error())
//		return err
//	}
//	if info.IsDir() {
//		prefix = prefix + "/" + info.Name()
//		fileInfos, err := file.Readdir(-1)
//		if err != nil {
//			return err
//		}
//		for _, fi := range fileInfos {
//			f, err := os.Open(file.Name() + "/" + fi.Name())
//			if err != nil {
//				utils.ZLog().Error("message","compress the files","error",err.Error())
//				return err
//			}
//			err = compress(f, prefix, zw)
//			if err != nil {
//				utils.ZLog().Error("message","compress the files","error",err.Error())
//				return err
//			}
//		}
//	} else {
//		header, err := zip.FileInfoHeader(info)
//		header.Name = prefix + "/" + header.Name
//		if err != nil {
//			utils.ZLog().Error("message","compress the files","error",err.Error())
//			return err
//		}
//		writer, err := zw.CreateHeader(header)
//		if err != nil {
//			utils.ZLog().Error("message","compress the files","error",err.Error())
//			return err
//		}
//		_, err = io.Copy(writer, file)
//		file.Close()
//		if err != nil {
//			utils.ZLog().Error("message","compress the files","error",err.Error())
//			return err
//		}
//	}
//	return nil
//}
//
//
//func StartTimer() {
//	ticker := time.NewTicker(time.Second * 10)
//	i := 0
//	go func() {
//		for _ = range ticker.C {
//			i++
//			utils.ZLog().Info("循环",i)
//			fmt.Printf("ticked at %v", time.Now())
//		}
//	}()
//}


func ZLogInit() {
	zog := new(zgh.ZLogParam)
	fileName := zog.SetFileName("zghua")
	err := zog.ZLogInit(fileName)
	if err != nil {
		zgh.ZLog().Error(err.Error())
	}
}

func DbInit () {
	sp := new(conn.Sp)
	dbUser := sp.SetDbUserName(conf.DbUser)
	dbPwd := sp.SetDbPassword(conf.DbPassword)
	dbPort := sp.SetDbPort(conf.DbPort)
	dbHost := sp.SetDbHost(conf.DbHost)
	dbdb := sp.SetDbDataBase(conf.DbDataBase)
	sqlServer,err := conn.InitMysql(dbUser,dbPwd,dbPort,dbHost,dbdb)
	SqlServer = sqlServer
	if err != nil {
		zgh.ZLog().Error("有错误",err.Error())
	}
}