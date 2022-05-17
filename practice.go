package main

import "fmt"

func todo() {
	arr:=[]int{1,2,3,4}
	arr2:=[]string{"hi","where","you","are"}
	arr2 =append(arr2,"in ","this","city")
	fmt.Println(arr, arr2)

}
func main() {
	todo()
}
