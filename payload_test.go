package gokadena_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	gokadena "github.com/BoostyLabs/kadena-go"
)

func TestToHash(t *testing.T) {
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

	actualHash, err := req.ToHash()
	require.NoError(t, err)
	assert.Len(t, actualHash, 43)
}

func TestIsPaddingPresent(t *testing.T) {
	tests := []struct {
		hash             string
		numOfPaddings    uint8
		isPaddingPresent bool
	}{
		{hash: "", numOfPaddings: 0, isPaddingPresent: false},
		{hash: "abs=", numOfPaddings: 1, isPaddingPresent: true},
		{hash: "abs==", numOfPaddings: 2, isPaddingPresent: true},
	}

	for _, test := range tests {
		num, isPresent := gokadena.IsPaddingPresent(test.hash)
		assert.Equal(t, num, test.numOfPaddings)
		assert.Equal(t, isPresent, test.isPaddingPresent)
	}
}

func TestToCmd(t *testing.T) {
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

	cmd, err := req.ToCmd()
	require.NoError(t, err)
	assert.Len(t, cmd.Hash, 43)
	assert.NotEmpty(t, cmd.Cmd)
}
