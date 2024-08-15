package typesys

type RefType struct {
	Name string
	age  int8
	Ref  bool
}

// TestRefType 指针接收修改的是原始对象的值
func (ref *RefType) TestRefType() {
	ref.Name = "modifiedRefName"
}
