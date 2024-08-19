package base

import (
	"fmt"
	"unsafe"
)

func TestVar() {
	// 通过var关键字进行声明并初始化 也可以不声明类型 go会根据值进行类型推断
	var a int = 2
	var b = 3
	var c string = "hello"
	fmt.Println(a, b, c)

	// 短变量声明 在函数内部使用
	x := 8
	y := "test111"
	fmt.Println(x, y)

	// 默认值 声明变量但未初始化 则使用默认值
	var d int      // 0
	var s string   // 空字符串
	var e bool     // false
	var s1 *string // nil
	fmt.Println(d, s, e, s1)

	// 数组/切片(类似java中的list)/映射的初始化
	arr := [3]int{1, 2, 3}
	// 切片 使用make
	slice1 := make([]int, 10)
	slice := []int{4, 5, 6}  // 一个指向底层数据的引用 切片大小与容量的区别
	var arr1 [5]int          // 每个元素都为零值
	sliceFromArr := arr[1:2] // 与数组共享底层数据结构
	fmt.Println(arr, arr1, slice, slice1, sliceFromArr)
	arr[1] = 10
	fmt.Println(arr, slice, sliceFromArr)
	sliceFromArr[0] = 20
	// result arr [1 20 3], arrAddr: 0xc00011e060, slice:[20], sliceAddr:0xc00011e068
	printArrAndSlice(&arr, &sliceFromArr)
	sliceFromArr = append(sliceFromArr, 30, 40, 50, 60)
	// result arr [1 20 3], arrAddr: 0xc00011e060, slice:[20 30 40 50 60], sliceAddr:0xc000128090
	// 切片扩容后 底层数组会发生变化 此时数组指向的是旧的数组地址
	printArrAndSlice(&arr, &sliceFromArr)

	// 结构体的初始化
	type person struct {
		name string
		age  int
	}
	p1 := person{name: "111", age: 5}
	fmt.Println(p1, &p1)

	// 使用new关键字
	p := new(int) // 指针
	*p = 42       // 给指针指向的地址赋值 *在变量之前时 意思是引用变量
	fmt.Println(*p, p)

	// map
	var map1 map[string]int
	//map1["nil"] = 5 仅声明未进行初始化则map为nil, 直接使用则会报错
	// 最后一个元素后面必须要有逗号 语法
	map2 := map[string]int{
		"key1": 2,
		"key2": 3,
	}
	map3 := make(map[string]int)
	map3["key3"] = 5
	fmt.Println(map1, map2, map3)

	// 结构体 接口 方法
}

// getSliceAddr 获取slice底层数组的地址
func getSliceAddr(slice []int) *int {
	return unsafe.SliceData(slice)
}

// printArrAndSlice 打印数组和切片的地址 进行对比
func printArrAndSlice(arr *[3]int, sliceFromArr *[]int) {
	fmt.Printf("arr %v, arrAddr: %p, slice:%v, sliceAddr:%p\n", *arr, arr, *sliceFromArr, getSliceAddr(*sliceFromArr))
}

type Reader interface {
}

// NamingConventions 命名约束说明
func NamingConventions() {
	// 大小写规则
	fmt.Println("如果标识符号首字母大写，则未public可导出，其他包可以引用，如果首字母为小写，则为private只能在该包中引用")
	// 命名风格
	fmt.Println("简洁性、描述性、驼峰命名法，不使用下划线")
	// 特殊命名
	fmt.Println("布尔类型：通常使用is,has,can等前缀标识状态\n" +
		"接口命名：以er结尾？，以表示接口的行为")
	// 包命名
	fmt.Println("小写字符，避免使用下划线以及大写字母，简短且明确")
	// 其他
	fmt.Println("命名时注意作用域，避免出现命名冲突")
}
