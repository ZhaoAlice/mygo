package main

import (
	"fmt"
	"mygo/base"
)

func descriptMain() {
	fmt.Println("程序入口：每个go程序必须有一个main包和main函数作为程序运行的入口以及打包的入口")
	fmt.Println("windows打包生成exe文件：go build -o mygo.exe")
}

func main() {
	descriptMain()
	//sliceinfo.TestSlice()
	//sliceinfo.TestRefType()
	//
	//errorinfo.TestErrorInfo()

	//channel.ReadData()
	base.TestVar()
}
