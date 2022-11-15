package main

//
///*
//struct A {
//    int i;
//    float f;
//};
//*/
//import "C"
//import "fmt"
//
//func main() {
//	var a C.struct_A
//	a.i = 3
//	a.f = 3.5
//	fmt.Println(a.i)
//	fmt.Println(a.f)
//}

/*

 */
//struct A {
//    int type; // type 是 Go 语言的关键字
//};
//*/
//import "C"
//import "fmt"
//
//func main() {
//	var a C.struct_A
//	a._type = 10
//	fmt.Println(a._type) // _type 对应 type
//}

///*
//struct A {
//    int   type;  // type 是 Go 语言的关键字
//    float _type; // 将屏蔽CGO对 type 成员的访问
//};
//*/
//import "C"
//import "fmt"
//
//func main() {
//	var a C.struct_A
//	a._type = 3.5
//
//	fmt.Println(a._type) // _type 对应 _type
//}

///*
//struct A {
//    int   size: 10; // 位字段无法访问
//    float arr[];    // 零长的数组也无法访问
//};
//*/
//import "C"
//import "fmt"
//
//func main() {
//	var a C.struct_A
//	fmt.Println(a.size) // 错误: 位字段无法访问
//	fmt.Println(a.arr)  // 错误: 零长的数组也无法访问
//}

///*
//#include <stdint.h>
//
//union B1 {
//    int i;
//    float f;
//};
//
//union B2 {
//    int8_t i8;
//    int64_t i64;
//};
//*/
//import "C"
//import "fmt"
//
//func main() {
//	var b1 C.union_B1
//	fmt.Printf("%T\n", b1) // [4]uint8
//
//	var b2 C.union_B2
//	fmt.Printf("%T\n", b2) // [8]uint8
//}

///*
//#include <stdint.h>
//
//union B {
//    int i;
//    float f;
//};
//*/
//import "C"
//import (
//	"fmt"
//	"unsafe"
//)
//
//func main() {
//	var b C.union_B
//	//b.i = 3
//	fmt.Println("b.i:", *(*C.int)(unsafe.Pointer(&b)))
//	fmt.Println("b.f:", *(*C.float)(unsafe.Pointer(&b)))
//}

/*
enum C {
    ONE,
    TWO,
};
*/
import "C"
import "fmt"

func main() {
	var c C.enum_C = C.TWO
	fmt.Println(c)
	fmt.Println(C.ONE)
	fmt.Println(C.TWO)
}
