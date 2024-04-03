package bookhizhan

import (
	"fmt"
	"testing"
)

type ParseError struct {
	Filename string
	Line  int
}

// 实现error接口，返回错误描述
func (e *ParseError) Error() string{
	return  fmt.Sprintf("%s:%d",e.Filename,e.Line)
}
// 创建一些解析错误
func newParseError(filename string,line int) error{
	return &ParseError{filename,line}
}

func TestParseError(t *testing.T) {
	var e error
	// 创建一个错误实例，包含文件名和行号
	e = newParseError("main.go",1)
	// 根据错误接口的具体类型，获取详细信息
	switch detail :=e.(type){
	case *ParseError:
		fmt.Printf("Filename: %s Line: %d\n",detail.Filename,detail.Line)
	default:
		fmt.Println("other error")
	}
}