package main

//#include<demo.h>
import "C"

func main() {
	C.SayHello(C.CString("this is go make c function demo\n"))
}
