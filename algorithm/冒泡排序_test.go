package algorithm

import("testing"
		"fmt"
	)

func BubbleSort(a []int) []int{
	// 外层循环控制比较轮数
	for i := 0; i <len(a)-1; i++ {
		for j:=0;j<len(a)-1-i;j++{
			// 注意这里
			if a[j]>a[j+1]{
				a[j+1],a[j]=a[j],a[j+1]
			}
		}
	}
	return a
}

func TestBulleSort(t *testing.T){
	a:=[]int{13,12,9,18,20,100,122}
	fmt.Print(BubbleSort(a))
}

