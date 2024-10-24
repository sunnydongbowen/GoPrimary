package slice

import (
	"fmt"
	"testing"
) 

func changeMap(m map[int]string) {

	m[1]="map2"
}

func TestMap(t *testing.T) {
	m:=map[int]string{
		1:"map1",
	}
	fmt.Println(m)
	changeMap(m)
	fmt.Println(m)
}