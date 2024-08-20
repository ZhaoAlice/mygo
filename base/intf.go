package base

import "fmt"

type Reader1 interface {
	ReadBook() string
}

type Reader2 interface {
	ReadBook1(name string) string
}

// ComponentReader 嵌套接口 组合Reader1与Reader2 并且可以继续新增接口
type ComponentReader interface {
	Reader1
	Reader2
	ReadBook3() string
}

// ReaderImpl 实现接口的结构体
type ReaderImpl struct {
	name string
}

func (r ReaderImpl) ReadBook() string {
	return "intf ComponentReader Reader1"
}

func (r ReaderImpl) ReadBook1(name string) string {
	return name
}

func (r ReaderImpl) ReadBook3() string {
	return "test read intf impl"
}

// 类型断言 用于检查接口变量的具体类型 并在需要时将其转换为该类型

type Animal interface {
	Speak() string
}

type Dog struct {
}

type Cat struct {
}

func (d Dog) Speak() string {
	return "woof!"
}

func (c Cat) Speak() string {
	return "meow!"
}

func AssertTest() {
	var a Animal = Cat{}
	// 类型断言 这种写法 尝试将a转换为Dog类型 如果成功ok的值为true 转换失败否则为false
	dog, ok := a.(Dog)
	// 注意if else的缩进
	if ok {
		fmt.Println("dog says：", dog.Speak())
	} else {
		fmt.Println("not a dog")
	}
}

// MakeSound 类型切换用于检查接口变量的具体类型，并在需要时将其转换为该类型 多态
func MakeSound(a Animal) {
	// 实现多态
	fmt.Println(a.Speak())
	// 这个是类型切换的概念 获取a的类型 type switch是一种特殊的switch语句
	switch v := a.(type) {
	case Dog:
		fmt.Println("dog says", v.Speak())
	case Cat:
		fmt.Println("cat says", v.Speak())
	default:
		fmt.Println("unKnown animal")
	}
}

// PrintType 空接口
func printType(value interface{}) {
	fmt.Printf("value: %v, Type: %T\n\n", value, value)
}

func TestEmptyInterface() {
	printType(2)
	printType("string")
	printType(3.14)
	printType(true)
	printType(Dog{})
}
