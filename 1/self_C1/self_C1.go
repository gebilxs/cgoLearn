package main

//void SayHello(const char* s);
import "C"

func main() {
	C.SayHello(C.CString("another self_C demo\n"))
}
