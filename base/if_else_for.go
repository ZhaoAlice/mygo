package base

import (
	"fmt"
	"mygo/errorinfo"
	"mygo/typesys"
)

// 多条件判断
func testMultiCondition() {
	a := 2
	b := 4
	if a < b && b > 3 {
		fmt.Println("both conditions are true a < b && b > 3")
	}

	if a > 5 || b < 10 {
		fmt.Println("at least one condition is true: a > 5 || b < 10")
	}
}

// 使用if语句进行错误处理
func testIfError(refType *typesys.RefType) (string, error) {
	fmt.Printf("refType: %v\n", *refType)
	if refType.Name == "correct" {
		return "success", nil
	}
	return "", &errorinfo.MyError1{Code: "500", Message: "error msg"}
}

func executeTestErrorIf() {
	obj, err := testIfError(&typesys.RefType{Name: "no"})
	if err != nil {
		fmt.Printf("error:%v", err.Error())
	} else {
		fmt.Printf("obj is: %v", obj)
	}
}

// 使用if语句的短变量声明
func testShortVar(x int) {
	if y := x * 2; y > 10 {
		fmt.Println("y is greater than 10", y)
	} else {
		fmt.Println("y is less than or equal to 10", y)
	}
}

// switch的用法
func testSwitch(value int) {
	switch value {
	case 1:
		fmt.Println("value1:", value)
	case 2:
		fmt.Println("value2:", value)
	default:
		fmt.Println("default value is:", value)
	}

	// 声明与使用
	switch day := "Monday"; day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("End of the week")
	default:
		fmt.Println("Midweek")
	}
}

// for
func testFor() {
	// 第一种方式
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 第二种方式
	slice := []int{1, 2, 3}
	for index, value := range slice {
		fmt.Printf("index is: %d, value is: %v", index, value)
	}
}

func TestAll() {
	testMultiCondition()
	executeTestErrorIf()
	testShortVar(2)
	testShortVar(10)
	testSwitch(1)
	testSwitch(2)
	testSwitch(10000)
	testFor()

}
