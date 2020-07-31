package main

import (
	"C"
	"fmt"
	"unsafe"
	"sync"
)
func main(){
	fmt.Println("main in plugin")
}

type Data struct{
	A string
	B int
	C struct{
		D string
	}
}

//export Add
func Add(a C.int, b unsafe.Pointer)C.int{
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(){
		defer wg.Done()
		defer func(){
			fmt.Println("in plugin defer ")
			if err := recover(); err!=nil {
				fmt.Println(err)
			}
		}()
		data :=  (*Data)(b)

		//data :=  (*C.char)(b)
		//fmt.Printf("a %d b %s \n", int(a), C.GoString(b) )
		fmt.Printf(" so changed2  a %d A %s B %d D %s   \n", int(a), data.A, data.B, data.C.D )
		if a == 8 {
			panic(" plugin error occur for a is 8")
		}

	}()
	wg.Wait()
	return a*a
}
