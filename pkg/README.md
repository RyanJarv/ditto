# ditto

Mimic any binary

## Examples

Crash your computer:
```shell
$ ./ditto
```

Anything else:
```
$ mv ./ditto ./ls
$ ./ls
cmd	go.mod	ls	pkg
$ mv ./ls ./cat
$ ./cat go.mod 
module github.com/ryanjarv/ditto

go 1.16
```


## Build

```shell
go build cmd/ditto.go
```