go build -buildmode=c-shared -o add.so add.go
go run main.go
rm -rf add.so add.h 
