package base

import (
	"fmt"
	"os"
)

// go语言中，defer语句用户延迟函数的执行，直到包含它的函数执行结束返回时，适合于资源清理、解锁、关闭文件等操作
// 语法 defer FuncCall() 当包含defer语句的函数执行结束后，才会执行funcCall

// 执行顺序 与栈类似 后进先出
func illustrateMultiDefer() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	fmt.Println("This is the main function")
}

// 参数求值
func calculateVarValue() {
	x := 10
	// x的值在defer语句被调用时进行计算 为10
	defer fmt.Println("value is: ", x)
	// 不影响defer语句中x的取值
	x = 20
}

// 资源管理

func OpenFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(file)
	fmt.Println("File open successfully")
	return file
}

func TestAll1() {
	/**
	This is the main function
	defer 3
	defer 2
	defer 1
	*/
	illustrateMultiDefer()
	calculateVarValue()
}
