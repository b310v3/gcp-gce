# gcp-gce

用來獲得host的記憶體，CPU，運作環境的數據。
先下載需要的package
```
go get golang.org/x/sys/unix
go get github.com/shirou/gopsutil
go get github.com/StackExchange/wmi
```
然後編譯這個go程式
```
go build sysinfo.go
```
接下來執行便會現實出目前系統的資訊
```
./sysinfo
```
