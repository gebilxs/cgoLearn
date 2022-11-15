package main

/*
#include <stdio.h>

static void SayHello(const char* s){
	puts(s);
}
*/
import "C"

func main() {
	C.SayHello(C.CString("this is a C_self demo!\n"))
}
