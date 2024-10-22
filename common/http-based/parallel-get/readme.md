```
// takes about 1 sec for compilation time so that you don't get confused with the result
go build -o parallel-get common/parallel-get/main.go

time ./parallel-get
https://www.google.com:135ms
https://www.apple.com:135ms
https://www.microsoft.com:841ms
All processing done
./parallel-get  0.85s user 0.01s system 75% cpu 1.141 total
```