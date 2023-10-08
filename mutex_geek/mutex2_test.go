package mutexgeek

import (
	"fmt"
	"sync"
	"testing"
)

type SafeMap[K comparable,V any] struct{
	values map[K]V
	lock sync.RWMutex
}

func (s *SafeMap[K, V]) LoadOrstoreV1(key K,newValue V) (V,bool){
	//加读锁
	s.lock.RLock()
	oldVal,ok:=s.values[key]
	s.lock.RUnlock()

	if ok {
		return oldVal,true
	}
	fmt.Println("add WMutex")

	//写锁
	s.lock.Lock()
	fmt.Println("add WMutex done!")
	defer s.lock.Unlock()

	//double check
	oldVal,ok=s.values[key]
	if ok{
		return oldVal,true
	}
	s.values[key]=newValue
	return newValue,false
}

func TestDeferRlock(t *testing.T) {
	sm:=SafeMap[string,string]{
		values: make(map[string]string, 4),
	}
	sm.LoadOrstoreV1("a","b")
	fmt.Println(sm.values)
}