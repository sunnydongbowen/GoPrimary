package mutexgeek

import (
	"fmt"
	"sync"
	"time"
	"testing"
)

var wg sync.WaitGroup

func (s *SafeMap[K, V]) LoadOrstoreV2(key K,newValue V) (V,bool){
	s.lock.RLock()
	oldVal,ok:=s.values[key]
	s.lock.RUnlock()
	
	if ok{
		fmt.Printf("all ready have%v\n",s.values)
		return oldVal,true
	}
	time.Sleep(100*time.Millisecond)

	s.lock.Lock()
	defer s.lock.Unlock()
	s.values[key]=newValue
	fmt.Printf("new k-v added %v\n",s.values)
	return newValue,false

}

func TestWithoutDoubleCheck(t *testing.T) {
	sm:=SafeMap[string,int]{
		values: make(map[string]int,4),
	}
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func (i int)  {
			defer wg.Done()
			time.Sleep(time.Second)
			sm.LoadOrstoreV2("a",i)
		}(i)
	}
	time.Sleep(3*time.Second)
	wg.Wait()
	fmt.Println(sm)

}