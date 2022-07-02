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
$ make vendor
```

Build the binary using the command.

```sh
$ make build
```

Run the binary using the command

```sh
$ make run
```

To perform lint and formatting of the code, install golangci-lint using the command

```sh
$ make install-golangci-lint
```

To lint the code, run the command

```sh
$ make lint
```

To format the code, run the command

```sh
$ make format
```

Add new packages to the repository using the command

```sh
$ go get -u <package_path>
```
