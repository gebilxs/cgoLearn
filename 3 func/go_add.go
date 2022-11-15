package main

import "C"

/*
static int add (int a,int b){
return a+b;
}
*/
import "C"

func main() {
	println(C.add(3, 5))
}
