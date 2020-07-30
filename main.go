package main 


/*
#cgo linux LDFLAGS: -ldl
#include <stdio.h>
#include <dlfcn.h>

int call_so(char * so_path, char * fn, int a, void* b)
{
	printf("args %s %s %d %s \n", so_path, fn, a, b);
    void* handle = dlopen(so_path, RTLD_LAZY);
    if(handle==NULL){
        printf("dlopen failed %s\n", dlerror());
        return 1;
    }
    printf("dlopen ok %s\n", dlerror());

    int (*add)(int a, void* b);

    add = dlsym(handle, fn);
    if(add==NULL){
        printf("dlsym failed %s\n", dlerror());
        return 1;
    }
    printf("dlsym ok %s  %lu \n", dlerror(), (unsigned long)(&add));

    int sum = add(a, b);
    printf("add  %d\n", sum);
    int flg=(int)dlclose(handle);
    if(flg==-1){
        printf("dlclose failed %s\n", dlerror());
        return 1;
    }
    printf("dlclose ok %s   \n", dlerror());
    return sum;
}
*/
import "C"
import "fmt"
import "unsafe"

func main(){
    type C_OBJ struct{
    C string
}
        innerObj:=C_OBJ{"inner  obj"}
	obj :=  struct{
		A string
		B int
                C C_OBJ 
	}{"i am  中国 object 2", 99, innerObj}
	 //str := "i am 中国 2"
	//ret := C.call_so(C.CString("/Users/qiangjian/go-c-go/add.so"), C.CString("add"), C.int(8) ,unsafe.Pointer(&info) )
	//var ret C.int = i
	//fmt.Printf("%#v %#v \n", &str, unsafe.Pointer(&str))
	//C.call_so(C.CString("./add.so"), C.CString("Add"), C.int(8) , C.CString(str))
	ret := C.call_so(C.CString("./add.so"), C.CString("Add"), C.int(8) , unsafe.Pointer(&obj))
	fmt.Println("ret C.int ", ret)
}

