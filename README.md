# ditto

Mimic any binary

## Examples

???
```shell
$ ./ditto
```

Anything else:
```
$ mv ./ditto ./ls
$ ./ls
cmd	go.mod	ls	pkg
```

Everything else:
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
