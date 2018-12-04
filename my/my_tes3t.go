/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-03
 * Time: 22:19
 */
package my

import "fmt"

type options struct {
	a int64
	b string
	c map[int]string
}

func (o *options) writeA(a int64) *options {
	o.a = a
	return o
}
func (o *options) writeB(b string) *options {
	o.b = b
	return o
}

func (o *options) WriteC(c map[int]string) *options {
	o.c = c
	return o
}

type Foo struct {
	verbosity int
}


func main() {
	op := new(options)
	op.writeA(int64(1)).writeB("test").WriteC(make(map[int]string, 0))

	fmt.Println(op.a, op.b, op.c)
}