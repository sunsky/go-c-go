#include <stdio.h>
#include <dlfcn.h>


int call_so(char * so_path, char * fn, int a, char* b)
{
	printf("args %s %s %d %s \n", so_path, fn, a, b);
    void* handle = dlopen(so_path, RTLD_LAZY);
    if(handle==NULL){
        printf("dlopen failed %s\n", dlerror());
        return 1;
    }
    printf("dlopen ok %s\n", dlerror());

    int (*add)(int a, char* b);

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
int main(){
	char* b = "i am from 中国 gcc";
	int ret = call_so("./add.so","Add", 2, b);
	printf("ret %d\n", ret);
	return 0;
}
