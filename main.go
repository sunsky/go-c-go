package main 


/*
#include <stdio.h>
#include <dlfcn.h>

int call_so(char * so_path, char * fn, int a, void *b)
{

	int* bp= (int*)(b);
	printf("args %s %s %d %d \n", so_path, fn, a, *bp);
    void* handle = dlopen(so_path, RTLD_LAZY);
    int (*add)(int a, void* b);
    add = dlsym(handle, fn);
    int sum = add(a, b );
    printf("add  %d\n", sum);

    dlclose(handle);
    return sum;
}
*/
import "C"
//import "fmt"
import "unsafe"


func main(){
	
	//_ :=  struct{
	//	A string
	//	B int
	//}{"i am go", 99}
	 var num int= 2
	//ret := C.call_so(C.CString("/Users/qiangjian/go-c-go/add.so"), C.CString("add"), C.int(8) ,unsafe.Pointer(&info) )
	//var ret C.int = i
	C.call_so(C.CString("/Users/qiangjian/go-c-go/add.so"), C.CString("Add"), C.int(8) ,unsafe.Pointer(&num) )
	//fmt.Println("ret C.int ", ret)
}

