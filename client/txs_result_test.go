package client_test

import (
	"context"
	"flag"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/BoostyLabs/kadena-go/client"
)

var dev = flag.Bool("dev", getenv("DEV_MODE"), "flag 'dev' used ONLY to run tests during development")

func getenv(key string) bool {
	val := os.Getenv(key)
	isDev, err := strconv.ParseBool(val)
	if err != nil {
		return false
	}
	return isDev
}

func TestTxsResult(t *testing.T) {
	if !*dev {
		t.Skip("skip testing because 'dev' flag is absent")
	}

	kadena := client.New(client.Config{
		NodeAddress: "https://api.chainweb.com/chainweb",
		ChainName:   "mainnet01",
		APIVersion:  "0.0",
		ChainID:     2,
	})

	cmdsResult := client.CommandsResultRequest{
		RequestKeys: []string{"vpIr7t79FRROh-uzZ12VLvXgXdNJlhXIlYYOwOLH6qI"},
	}

	_, err := kadena.TxsResult(context.Background(), cmdsResult)
	require.NoError(t, err)
}
