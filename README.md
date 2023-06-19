# sdk-core-go
**Vottun Core API SDK For Golang**

This repository contains an SDK with the necessary functions to make calls to the Vottun Core Blockchain API.

Requires Go version 1.18 or higher.

## Installation

```shell
go get github.com/vottunio/coreapi
```

## Usage

To use the Core API, first import the dependency into your `.go` file:

```go
import "github.com/vottunio/coreapi"
```

Build a CoreApi client to use all the API calls. For that, you'll need to obtain a token and an application identifier. You can obtain them at [https://apis.vottun.tech/trial](https://apis.vottun.tech/trial). You'll also need the root URL of the application, which will be published very soon.

An example of creating a client would be:

```go
coreApiClient := coreapi.New(<TOKEN_AUTH>, <APP_ID>, <ENV_ROOT_URL>)
```

Where `coreApiClient` is the new client that allows you to call different functions of the Core API.

The following example allows us to make a call to create a new custodied wallet for a user:

```go
newWalletRequest := &coreapi.NewWalletRequestDTO{User: "nameATmail.example", Pin: "6666666"}
if res, err := coa.CreateNewCustodiedWallet(newWalletRequest); err != nil {
    fmt.Printf("err: %v\n", err)
} else {
    fmt.Printf("res: %+v\n", res)
}
```

Please note that the example assumes you have already created an instance of the `coa` client using the `coreApiClient` from the previous step.