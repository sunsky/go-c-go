
#include <stdio.h>
#include <dlfcn.h>
//#include "add.h"

int call_so(char * so_path, char * fn, int a, void *b)
{

	int* bp= (int*)(b);
	printf("args %s %s %d %d \n", so_path, fn, a, *bp);
	//return 0;
    void* handle = dlopen(so_path, RTLD_LAZY);
    int (*add)(int a, void* b);

    add = dlsym(handle, fn);
    int sum = add(a, b );
    printf("add  %d\n", sum);

    dlclose(handle);
    return sum;
}

int main(){
	int b = 3;
	int ret = call_so("/Users/qiangjian/go-c-go/add.so","Add", 2, &b);
	printf("ret %d\n", ret);
	return 0;
}
