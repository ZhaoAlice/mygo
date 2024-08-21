package errorinfo

import (
	"errors"
	"fmt"
	"net/http"
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

	err2 := fmt.Errorf("this is an error with a value: %d", 42)
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

// panic和recover的使用 go的错误处理更倾向于显式地返回错误 而不是使用异常处理机制 全局异常捕获 对于不可恢复的错误可以使用该方式
// 类似Java中的try catch语句
func handler(w http.ResponseWriter, r *http.Request) {
	defer recoverFromPanic()
	// 抛出错误
	panic("error occur")
}

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("recovered from panic:", r)
		http.Error(http.ResponseWriter(nil), "Internal Server Error", http.StatusInternalServerError)
	}
}

// MyError1 自定义错误类型
type MyError1 struct {
	Code    string
	Message string
}

func (e *MyError1) Error() string {
	return e.Message
}

func ExecuteFunc() error {
	// 方法是指针接收者
	return &MyError1{Message: "test custom error"}
}

func TestError() {
	if err := ExecuteFunc(); err != nil {
		fmt.Println("error:", err)
	}
}
