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

Or everything else:
```
$ ln -s ditto cat
$ ln -s ditto echo
$ ./echo Hello from ditto | ./cat -
Hello from ditto
```

## Build

```shell
go build cmd/ditto.go
```
