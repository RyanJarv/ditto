# ditto

Mimic any binary

## Why?

Just sprinkle on some code for logging stdin, let sit in prod for 2-3 weeks, and you have a really dumb keylogger.

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
