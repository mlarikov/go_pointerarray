package main

import (
	"fmt"
	"reflect"
)

func modifyArray(arr *[3][3]string) {
	arr[2][2] = "hello"
	fmt.Println("Arr", reflect.TypeOf(arr), arr)
}

func main() {
	myArr := [3][3]string{{"qwe", "asd", "zxc"}, {"rty", "fgh", "vbn"}, {"uio", "jkl", "m,."}}
	fmt.Println("myArr", reflect.TypeOf(myArr), myArr)
	modifyArray(&myArr)
	fmt.Println("myArr", reflect.TypeOf(myArr), myArr)
}
