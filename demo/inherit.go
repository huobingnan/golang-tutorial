package main

import (
	"fmt"
)

// golang只能使用接口来实现多态的性质
// golang的接口是一种鸭子类型的接口

type Animal interface {
	Type() string
}

type Setter interface {
	SetProperty(string, string)
}

type Dog struct {
	properties map[string]string
}

type Cat struct {
}

type HabaDog struct {
	Dog
}

func (d *Dog) SetProperty(key, value string) {
	if d.properties == nil {
		d.properties = make(map[string]string)
	}
	d.properties[key] = value
}

func (d *Dog) Type() string {
	return "Dog"
}

func (c *Cat) Type() string {
	return "Cat"
}

func main() {
	// 接口实质上是是一个指针，所以接口作为参数进行传递时，必须传递数据的地址参照
	func(animal Animal) {
		fmt.Printf("animal type is %s\n", animal.Type())
	}(&Dog{})

	func(animal Animal) {
		fmt.Printf("animal type is %s\n", animal.Type())
	}(&Cat{})

	// interface本质上是一个指针
	func(setter Setter) {
		setter.SetProperty("s", "s")
	}(&Dog{})

	// struct本身没有多态的属性，即使是struct类型的指针，也没有多态的属性
	//var dogPtr *Dog
	//dogPtr = &HabaDog{}
}
