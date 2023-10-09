package mutexgeek


import (
	"fmt"
	"time"
	"testing"
)


func (s *SafeMap[K, V]) LoadOrstoreV3(key K,newvalue V) (V,bool){
	s.lock.RLock()
	oldVal,ok:=s.values[key]
	s.lock.RUnlock()
	if ok{
		fmt.Printf("first check ready exists%v\n",s.values)
		return oldVal,true
	}
	time.Sleep(100*time.Millisecond)

	s.lock.Lock()
	defer s.lock.Unlock()
	oldVal,ok=s.values[key]
	if ok{
		fmt.Printf("second check ready exist %v\n",s.values)
		return oldVal,true
	}

	s.values[key]=newvalue
	fmt.Printf("new key added %v\n",s.values)
	return newvalue,false
}

func TestDoubleCheck(t *testing.T) {
	sm:=SafeMap[string,int]{
		values: make(map[string]int,4),
	}
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func (i int)  {
			defer wg.Done()
			time.Sleep(time.Second)
			sm.LoadOrstoreV3("a",i)
		}(i)
	}
	time.Sleep(3*time.Second)
	wg.Wait()
	fmt.Println(sm)
}