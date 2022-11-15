package main

/*
static int div(int a,int b){
	return a/b;
}
*/
import "C"

func main() {
	println(C.div(6, 3))
}
