// Code generated by cmd/cgo; DO NOT EDIT.

//line D:\go\data\huiyan\learning_demo\5 interal\c_use_go.go:1:1
package main

//int sum(int a, int b);
import _ "unsafe"

//export sum
func sum(a, b  /*line :7:15*/_Ctype_int /*line :7:20*/)  /*line :7:22*/_Ctype_int /*line :7:27*/ {
	return a + b
}

func main() {
	println(sum(32, 3))
}
