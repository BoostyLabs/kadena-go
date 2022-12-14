# Kadena and Pact SDK

[![](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://pkg.go.dev/github.com/BoostyLabs/kadena-go)

A simple library to interact with Pact API and easy-to-build command structure. 

## How to install 
```sh
go get github.com/BoostyLabs/kadena-go
```

----
## Usage

> **Note**
>
> The library is still in progress and some endpoints that are provided by Pact API **have not** been implemented yet.

#### Example how to construct command payload which named as request.

```go 
    req := gokadena.Request{
		Payload: gokadena.Payload{
			Exec: gokadena.Exec{
				Data: "some_data",
				Code: "(+a b)",
			},
		},
		NetworkID: "testnet04",
		Meta: gokadena.Meta{
			ChainID:      "testnet04",
			Sender:       "ba54b224d1924dd98403f5c751abdd10de6cd81b0121800bf7bdbdcfaec7388d",
			GasLimit:     1000,
			GasPrice:     gokadena.DefaultGasPrice,
			TTL:          gokadena.DefaultTTL,
			CreationTime: gokadena.CreationTime(),
		},
		Nonce: gokadena.Nonce(),
		Signers: []gokadena.Signer{
			{
				PubKey:  "ba54b224d1924dd98403f5c751abdd10de6cd81b0121800bf7bdbdcfaec7388d",
				Address: "ba54b224d1924dd98403f5c751abdd10de6cd81b0121800bf7bdbdcfaec7388d",
				Schema:  gokadena.SchemaED25519,
				CList: gokadena.CapabilityItem{
					Name: "coin.TRANSFER",
					Args: []interface{}{"bob", "alice", 0.1},
				},
			},
		},
		Type: gokadena.TypeExec,
    }
```

#### Example of API client creation and interaction with endpoint supported by Pact 

```go
	kadena := client.New(client.Config{
		NodeAddress: "https://api.chainweb.com/chainweb",
		ChainName:   "mainnet01",
		APIVersion:  "0.0",
		ChainID:     2,
	})

	cmdsResult := client.CommandsResultRequest{
		RequestKeys: []string{"vpIr7t79FRROh-uzZ12VLvXgXdNJlhXIlYYOwOLH6qI"},
	}
```

## Tests
You can also interact with this library using tests. Repository has two types of tests:
- The first type is unit tests, mostly used for CI, to run use the command from the makefile 
```sh 
make tests
```
- The second type is integration tests that are used during development and interact with 3rd party services/nodes.
To use run this command
```sh 
make dev_tests
```
