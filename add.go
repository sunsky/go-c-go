package main

import (
	"C"
	"fmt"
	"unsafe"
)
func main(){}

type Data struct{
	A string
	B int
}

//export Add
func Add(a C.int, b unsafe.Pointer)C.int{
	//data :=  (*Data)(b)
	data :=  (*C.int)(b)
	fmt.Println("b ", *data )
	return a*a;
}
