package errorinfo

import (
	"errors"
	"fmt"
)

type MyError struct {
	code string
	msg  string
}

// 实现error接口
func (err *MyError) Error() string {
	return err.msg
}

func TestErrorInfo() {
	// 实现了接口error
	err1 := errors.New("this is an error,test")
	fmt.Println(err1.Error())

	err2 := fmt.Errorf("this is an error wiht a value: %d", 42)
	fmt.Println(err2)
}

func (err *MyError) testError1(flag bool) (int, error) {
	if flag {
		return 0, &MyError{code: "123", msg: "test error info"}
	}
	return 0, nil
}

func (err *MyError) TestError2() {
	result, error1 := err.testError1(false)
	if error1 != nil {
		// 类习惯转换
	}
	fmt.Printf("result is %d", result)
}

func testError3(flag bool) (int, error) {
	if flag {
		return 0, &MyError{code: "123", msg: "test error info"}
	}
	return 0, nil
}
