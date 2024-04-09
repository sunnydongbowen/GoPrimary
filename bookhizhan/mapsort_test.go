package bookhizhan

import (
	"fmt"
	"sort"
	"testing"
)
func TestMapSort(t *testing.T) {
	scene:=make(map[string]int)
	scene["route"]=99
	scene["brazil"]=4
	scene["china"]=78

	var sceneList []string
	for k:=range scene{
		sceneList = append(sceneList, k)
		}
	// 对切片排序
	sort.Strings(sceneList)
	fmt.Println(sceneList)
}