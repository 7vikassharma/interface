package main

import "fmt"

type WTest interface {
	Age(int)
	Name(string)
	Gender()
}
type MTest interface {
	WTest
	Bhaii()
}
type Woman struct{}
type Man struct {
	Woman
	single string
}

func (w Woman) Age(age int) {
	fmt.Println(age)
}
func (w Woman) Name(name string) {
	fmt.Println(name)
}
func (m Woman) Gender() {
	fmt.Println("f")
}
func (w Man) Age(age int) {
	fmt.Println(3 * age)
}
func (m Man) Bhaii() {
	m.single = "yessss"
	fmt.Println(m)
}

func main() {
	var in MTest
	in = Man{}
	in.Age(25)
	in.Gender()
	in.Bhaii()
	var ini WTest
	ini = Woman{}
	ini.Age(35)
	ini.Gender()
}
