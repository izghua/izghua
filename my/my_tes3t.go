/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-03
 * Time: 22:19
 */
package my

import (
	"fmt"
	"runtime"
	"strings"
)


func Testaa() {
	test2()
}

func test2(){
	fmt.Println("尝试隔断")
	pc,file,line,ok := runtime.Caller(2)
	fmt.Println(pc)
	fmt.Println(file)
	fmt.Println(line)
	fmt.Println(ok)
	f := runtime.FuncForPC(pc)
	fmt.Println(f.Name())
	fmt.Println("尝试隔断2")

	pc,file,line,ok = runtime.Caller(0)
	fmt.Println(pc)
	fmt.Println(file)
	fmt.Println(line)
	fmt.Println(ok)
	f = runtime.FuncForPC(pc)
	fmt.Println(f.Name())
	fmt.Println("尝试隔断3")

	pc,file,line,ok = runtime.Caller(1)
	fmt.Println(pc)
	fmt.Println(file)
	fmt.Println(line)
	fmt.Println(ok)
	f = runtime.FuncForPC(pc)
	fmt.Println(f.Name())
	arrStr := strings.Split(f.Name(),"/")

	fmt.Println("尝试隔断4",arrStr[len(arrStr)-1])

}