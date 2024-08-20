package typesys

// 自定义类型
// 自定义结构类型 里面的各个字段会初始化其零值
// 如果是被其他包导入使用 则类型名称和字段名称需要使用大驼峰的格式
// 大驼峰 其他包可见
// 小驼峰 只在该包下可见

type User struct {
	Name       string
	Email      string
	Ext        int
	Privileged bool
	SliceInfo  []string
}

// ModifyUserWithOutReturn 值传递方式 修改副本的值 不会影响原始对象的属性值
func (u User) ModifyUserWithOutReturn() {
	u.Name = "modify"
}

// ModifyUserNameWithReturn 指针方式传递 指针与原始对象是指向同一个对象的
func (u User) ModifyUserNameWithReturn() (u1 User) {
	u.Name = "modify1"
	return u
}

// EnLargeSlice 修改对象切片的取值原始对象的取值保持不变 说明是完全的值传递 切片头部值 底层的数组是一致的
func (u User) EnLargeSlice() []string {
	// 会返回一个新的切片 未扩充容量底层数组是一致的 如果扩充容量会生成新的底层数据 则返回的新切片指向的是新的底层数据 与原始切片底层数据不一样
	u.SliceInfo = append(u.SliceInfo, "e1")
	u.SliceInfo[0] = "e0"
	return u.SliceInfo
}
