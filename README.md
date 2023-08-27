# nymeria (accounts-v2)

Complete reimplementation of `Login` in Go using ory.sh in our applications.

# Getting Started
## To setup Nymeria
Clone the repository

```sh
 git clone git@github.com:sdslabs/nymeria.git
```

Enter the directory and download the vendor files

```sh
 cd nymeria/
 make vendor
```

Copy the contents of `config.sample.yaml` to `config.yaml`

```sh
cp config.sample.yaml config.yaml
```

Build the binary using the command.

```sh
 make build
```

Run the binary using the command

```sh
 make run
```

To perform lint and formatting of the code, install golangci-lint using the command

```sh
 make install-golangci-lint
```

To lint the code, run the command

```sh
 make lint
```

To format the code, run the command

```sh
 make format
```

Add new packages to the repository using the command

```sh
 go get -u <package_path>
```

Hot reloading support
- run the following command to install `air` (hot reload support)

```sh
 make install-air
```

- run the following command to run nymeria with `air` (hot reload support)

```sh
 make dev
```

## To setup Ory Kratos

Clone the ory/kratos repository 

```sh
 git clone https://github.com/ory/kratos.git
```

Enter the Kratos directory and  Change the Kratos Version to v0.10.0

```sh
 cd kratos
 git checkout v0.10.0
```

Download the dependencies

```sh
 go mod download
 go install -tags sqlite,json1,hsm .
```

Add the kratos binary to your Path

```sh
 $(go env GOPATH)/bin/kratos help
```

Copy the Kratos config file and identity schema from nymeria

```sh
 cp ../nymeria/config/kratos_config.yaml ./contrib/quickstart/kratos/email-password/kratos.yml
 cp ../nymeria/config/identity.schema.json ./contrib/quickstart/kratos/email-password/identity.schema.json
```

Run the following command to use Kratos in containerized form

```sh
 docker-compose -f quickstart.yml -f quickstart-standalone.yml -f quickstart-postgres.yml up --build --force-recreate
```