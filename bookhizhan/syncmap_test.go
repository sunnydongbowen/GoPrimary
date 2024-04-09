package bookhizhan

import(
	"testing"
	"fmt"
	"sync"
)
func TestSyncMap(t *testing.T) {
	var scene sync.Map
	//将键值对保存到sync.Map
	scene.Store("greece",97)
	scene.Store("london",100)
	scene.Store("egypt",200)

	fmt.Println(scene.Load("london"))
	scene.Delete("london")

	scene.Range(func(key, value any) bool {
		fmt.Println("iterate：",key,value)
		return true
	})	
}