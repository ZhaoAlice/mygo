package main

import (
	"mygo/errorinfo"
	"mygo/sliceinfo"
)

func main() {
	sliceinfo.TestSlice()
	sliceinfo.TestRefType()

	errorinfo.TestErrorInfo()
}
