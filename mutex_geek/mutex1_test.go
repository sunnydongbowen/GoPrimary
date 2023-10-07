package mutexgeek

import (
	"fmt"
	"sync"
	"testing"
)

type safeResource struct{
	resource map[string]string
	lock sync.Mutex
}

func (s *safeResource) Add(key,value string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.resource[key]=value
}

func TestMu(t *testing.T){
	var s safeResource
	//一定要初始化
	s.resource=map[string]string{}
	s.Add("key","value")
	fmt.Println(s.resource)
}

func TestMu1(t *testing.T) {
	s:=safeResource{
		resource: map[string]string{},
		lock: sync.Mutex{},
	}
	fmt.Println(s.resource)	
	s.Add("name","bowen")
	fmt.Println(s.resource)
}

