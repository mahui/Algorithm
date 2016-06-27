package main

import (
	"fmt"
)


func isLess (a,b int) bool{
	return a >= b
}

func replace (arr []int, idx1 int, idx2 int){
	temp := arr[idx1]
	arr[idx1] = arr[idx2]
	arr[idx2] = temp
}

func Asc(arr []int){
	for i := 0; i < len(arr); i ++ {
		min := arr[i]
		for j := i; j < len(arr); j++ {
			if(isLess(min,arr[j])){
				replace(arr,i,j)
			}
		}
	}
}

func Desc(arr []int){
	for i := 0; i < len(arr); i ++ {
		max := arr[i]
		for j := i; j < len(arr); j++ {
			if(!isLess(max,arr[j])){
				replace(arr,i,j)
			}
		}
	}
}

func main(){
	fmt.Printf("%t\n",isLess(1,2))
	arr1 := []int {1,2,3,4,8,6,5}
	fmt.Printf("%v\n",arr1)
	replace(arr1,1,2)
	fmt.Printf("%v\n",arr1)
	Asc(arr1)
	fmt.Printf("%v\n",arr1)
	Desc(arr1)
	fmt.Printf("%v\n",arr1)
}