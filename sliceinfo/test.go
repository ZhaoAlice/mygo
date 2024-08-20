package sliceinfo

import (
	"fmt"
	"mygo/base"
	"mygo/typesys"
)

func TestSlice() {
	var bill typesys.User
	bill.ModifyUserWithOutReturn()
	fmt.Printf("user name= %s,email= %s\n", bill.Name, bill.Email)
	bill1 := typesys.User{
		Name:      "zlg",
		Email:     "test@qq.com",
		SliceInfo: make([]string, 3, 5), // 指定切片类型 切片长度 容量
	}
	newUser := bill1.ModifyUserNameWithReturn()
	fmt.Printf("ap:%p, fp:%p, bill1:%v, newUser:%v\n", &bill1, &newUser, bill1, newUser)
	// 值引用时会创建新的切片 但是两个切片引用的底层数组是一个
	fmt.Printf("oldSlice:%p, newSlice:%p, oldSliceData:%p, newSliceData:%p\n", &bill1.SliceInfo, &newUser.SliceInfo,
		base.GetStrSliceAddr(bill1.SliceInfo), base.GetStrSliceAddr(newUser.SliceInfo))

	// 这个方法里面 会对切片扩容 底层数组发生变更 旧的底层数组依旧存在 直到没有被引用 go的垃圾回收机制会回收
	sliceInfo := bill1.EnLargeSlice()

	// 指向的地址发生变更
	fmt.Printf("new addr is %p, old addr is %p\n", &sliceInfo[0], &bill1.SliceInfo[0])
	fmt.Printf("value:%s  len: %d  capacity:%d\n", sliceInfo, len(sliceInfo), cap(sliceInfo))
	fmt.Printf("user name= %s,email= %s,sliceinfo=%s\n", bill1.Name, bill1.Email, bill1.SliceInfo)
}

func TestRefType() {
	lili := typesys.RefType{
		Name: "lili",
		Ref:  true,
	}
	(&lili).TestRefType()
	fmt.Printf("ref type name is %s, ref is %t\n", lili.Name, lili.Ref)
}

func BasicInfo() {
	// make初始化
	// 创建并初始化一个长度和容量均为5的切片 元素的初始值均为零值
	makeSlice := make([]int, 5)
	for key, value := range makeSlice {
		fmt.Printf("key is %d, value is %d\n", key, value)
	}
	/**
	key is 0, value is 0
	key is 1, value is 0
	key is 2, value is 0
	key is 3, value is 0
	key is 4, value is 0
	*/
	// 创建并初始化一个长度为3，容量为5的切片 元素初始化为零值
	makeSlice1 := make([]string, 3, 5)
	for key, value := range makeSlice1 {
		fmt.Printf("key1 is %d, value1 is %s+v\n", key, value)
	}
	/**
	key1 is 0, value1 is +v
	key1 is 1, value1 is +v
	key1 is 2, value1 is +v
	*/
	// 注释的区别 切片 与数组的区别是不设置长度 字面量初始化
	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		fmt.Printf("index: %d, Value: %d\n", i, num)
	}
	fmt.Println("sum is", sum)

	for _, val := range nums {
		fmt.Println("value is ", val)
	}

	// 映射
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Printf("key is: %s, value is %d\n", k, v)
	}
	fmt.Printf("hello my go program")

	// 根据nums1 创建切片1 新的切片与原始切片共享底层数组 类似与java list的subList方法
	m1 := nums[1:3]

	for _, v1 := range m1 {
		fmt.Printf("value is %d\n", v1)
	}
}
