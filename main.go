package main

import "fmt"

func modifyArray(arr *[3][3]string) {
	arr[2][2] = "hello"
}

func main() {
	myArr := [3][3]string{{"qwe", "asd", "zxc"}, {"rty", "fgh", "vbn"}, {"uio", "jkl", "m,."}}
	modifyArray(&myArr)
	fmt.Println(myArr)
}
