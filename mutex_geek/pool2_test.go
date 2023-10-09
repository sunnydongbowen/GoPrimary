package mutexgeek

import (
	"fmt"
	"sync"
	"testing"
)

//reset pool

type User struct{
	ID int
	Name string
}

func (u *User) Reset(){
	u.ID=0
	u.Name=""
}

func TestPool2(t *testing.T) {
	pool:=sync.Pool{
		New: func() any {
			return &User{}
		},
	}

	// 拿出来改了
	u1:=pool.Get().(*User)
	u1.ID=12
	u1.Name="TOM"

	//重置完放进去
	u1.Reset()
	pool.Put(u1)
	u2:=pool.Get().(*User)
	fmt.Println(u2)


	
	
}