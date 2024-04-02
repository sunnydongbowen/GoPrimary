package mutexgeek

import (
	"fmt"
	"sync"
	"testing"
)

//sync.Once的使用
// 大明喜欢这么用
type OnceClose struct {
	close sync.Once
}

func (o *OnceClose) Close() error {
	o.close.Do(func() {
		fmt.Println("close")
	})
	return nil
}

func (no OnceClose) Close1() error{
	no.close.Do(func ()  {
		fmt.Println("close")
	})
return nil
}


func TestOnceClose(t *testing.T) {
	oc:=&OnceClose{}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		oc.Close()
	}
}

func TestOnceClose1(t *testing.T) {
	oc:=OnceClose{}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		oc.Close1()
	}	
}
