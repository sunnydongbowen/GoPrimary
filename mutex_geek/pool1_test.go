package mutexgeek

import (
	"fmt"
	"sync"
	"testing"
)
//sync.Pool的使用
func  TestPoolDemo(t *testing.T) {
	var strPool = sync.Pool{
		New: func() any {
			return "test str"
		},
	}

	//取出来
	str:=strPool.Get()
	fmt.Println(str)

	//再放进去
	strPool.Put(str)
}



func TestPool(t *testing.T){
	pool:=sync.Pool{
		New: func() any {
			fmt.Println("jjj,new")
			//返回空字节切片
			return []byte{}
		},
	}

	for i := 0; i < 5; i++ {
		val:=pool.Get()
		fmt.Println(val)
		//pool.Put(val)
		// 想一想，拿出来不放进去会怎么样
		pool.Put(i)
		
	}
}
