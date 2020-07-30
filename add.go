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
    C struct{
        D string
    }
}

//export Add
func Add(a C.int, b unsafe.Pointer)C.int{
	data :=  (*Data)(b)
	//data :=  (*C.char)(b)
	//fmt.Printf("a %d b %s \n", int(a), C.GoString(b) )
	fmt.Printf("a %d A %s B %d D %s   \n", int(a), data.A, data.B, data.C.D )
	return a*a
}
