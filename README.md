# go-c-go

```
#go invoke c , then c invoke go
#this can be useful for load/unload dynamic library

go build -buildmode=c-shared -o add.so add.go
go run main.go
```
