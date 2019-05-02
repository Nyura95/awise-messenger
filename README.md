# Messenger Appointrip

This package is the http/socket server for the chat of https://appointrip.com

---

- [Install](#install)
- [Explanation](#Explanation)
- [Run](#Run)
- [Dependances](#Dependances)
- [Build](#Build)

---

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go install
```

## Run

For run the project, open a cmd with Go installed and run

```sh
go run main.go
```

## Explanation

See the package for read the doc

## Build

Build for ubuntu run

```
GOOS=linux GOARCH=amd64 go build -o messenger main.go
```

Build for macos run

```
GOOS=darwin GOARCH=amd64 go build -o messenger main.go
```

Build for window run

```
GOOS=windows GOARCH=amd64 go build -o messenger.exe main.go
```

This command create a binary file named 'messenger' or 'messenger.exe'
