package mutexgeek

import (
	"errors"
	"sync"
	"sync/atomic"
	"unsafe"
)

//封装缓冲池
type Mypool struct{
	sync.Pool
	cnt int32
	maxCnt int32
}


func (p *Mypool) Get() any{
	//sync 有这个方法，它就有这个方法，要熟悉go的这种特点
	return p.Get()
}


func (p *Mypool) Put(val any) any{
	// 大对象不放回去
	if unsafe.Sizeof(val) >1024{
		return errors.New("超过1024，大对象不可以放")
	}

	//超过数量限制
	cnt:=atomic.AddInt32(&p.cnt,1)
	if cnt>p.maxCnt{
		atomic.AddInt32(&p.cnt,-1)
		return errors.New("too many object,exit!")
	}
	p.Pool.Put(val)
	return nil
}


