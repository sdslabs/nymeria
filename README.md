# nymeria (accounts-v2)

Complete reimplementation of `Login` in Go using ory.sh in our applications.

# Getting Started

Clone the repository

```sh
$ git clone git@github.com:sdslabs/nymeria.git
```

Enter the directory and download the vendor files

```sh
$ cd nymeria/
$ go mod vendor
```

Build the binary using the command.

```sh
$ go build -o nymeria ./cmd/server/main.go
```

Run the binary using the command

```sh
$ ./nymeria
```

Add new packages to the repository using the command

```sh
$ go get -u <package_path>
```
