package main

// #include <stdio.h>
// #include <errno.h>
// #include <stdlib.h>
// #include <string.h>
//extern int go_qsort_compare(void* a, void* b);
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {}

//export handler
func handler(a, b C.int) C.int {
	return a + b
}

//export printChar
func printChar(str *C.char) {
	str1 := C.GoString(str)
	//defer C.free(unsafe.Pointer(str1))
	fmt.Println(str1)
}

////export numbers
//func numbers(num []C.int, len int) {
//	var temp int
//	var value []int
//	for i := 0; i < len; i++ {
//		temp = C.GoInt64(num[i])
//		value = append(value, temp)
//	}
//	// 通过 reflect.SliceHeader 转换
//	fmt.Println(value)
//}

//export numbers
func numbers(addr unsafe.Pointer, lens C.int) {
	var arr []int
	//直接访问内存空间
	arrHeader := (*reflect.SliceHeader)(unsafe.Pointer(&arr))
	//arrHeader 为arr空数组的指针
	arrHeader.Data = uintptr(addr)
	//数据是 找到addr的指针然后访问内存取到数据
	arrHeader.Len = int(lens)
	//分配长度
	arrHeader.Cap = int(lens)
	//分配cap
	//arr = arrHeader.Data

	quickSort(arr, 0, 9)
	fmt.Println(arr)

}

//export quickSort
func quickSort(nums []int, l, r int) { //[l,r]
	if l < r {
		m := partition(nums, l, r)
		quickSort(nums, l, m-1)
		quickSort(nums, m+1, r)
	}
}

//export partition
func partition(nums []int, l int, r int) int {
	key := nums[r]
	//all in [l,i) < key
	//all in [i,j] > key
	i := l
	j := l
	for j < r {
		if nums[j] < key {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
	nums[i], nums[r] = nums[r], nums[i]
	return i
}
