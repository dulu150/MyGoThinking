
// testarrayslice project main.go
package main

import (
	"fmt"
)

func main() {
	var myarray [10]int = [10]int{1, 2, 3, 4, 5, 6}
	var myslice []int = myarray[0:10]

	for i, v := range myslice {
		fmt.Println(i, " ", v, "\r\n")
	}

	//测试数组切片的自动增长能力
	myslice2 := make([]int, 5, 10)
	fmt.Println("length", len(myslice2))
	fmt.Println("cap", cap(myslice2))

	myslice3 := []int{1, 3, 4, 6, 9}
	myslice2 = append(myslice2, myslice3...)
	for i, v := range myslice2 {
		fmt.Println(i, "", v, "\r\n")
	}
}
