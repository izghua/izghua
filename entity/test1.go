package entity

type Test1 struct {
	Id   int    `xorm:"not null pk autoincr INT(11)"`
	Name string `xorm:"CHAR(1)"`
}
