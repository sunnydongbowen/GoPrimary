package shizhan
import (
	"fmt"
	"testing"
)


func TestChangeString(t *testing.T) {
	
	angle:="I Love you"
	angleByte:=[]byte(angle)

	for i := 3; i < 8; i++ {
		angleByte[i]=' '
	}
	fmt.Println(string(angleByte))
}