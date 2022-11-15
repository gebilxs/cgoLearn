package main

//you can not put /n between

//#include <stdio.h>
import "C"

func main() {
	C.puts(C.CString("this is a string from C, demo show!\n"))
}

//package main
//
////#include <stdio.h>
//import "C"
//
//func main() {
//	C.puts(C.CString("Hello, World\n"))
//}
