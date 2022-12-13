package gokadena

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"golang.org/x/crypto/blake2b"
)

const (
	// DefaultTTL is default time to live for tx in mempool.
	DefaultTTL int64 = 28000 // 8 hours.

	// DefaultGasPrice is default gas price in Kadena network.
	DefaultGasPrice float64 = 1e-5
)

// Request defines command request aka payload.
type Request struct {
	Payload   Payload  `json:"payload"`
	NetworkID string   `json:"networkId"`
	Meta      Meta     `json:"meta"`
	Nonce     string   `json:"nonce"`
	Signers   []Signer `json:"signers"`
	Type      Type     `json:"type"`
}

// Network defines list of all possible Kadena networks.
type Network string

const (
	// NetworkTestnet defines testnet network.
	NetworkTestnet Network = "testnet"
	// NetworkMainnet defines mainnet network.
	NetworkMainnet Network = "mainnet"
)

// Type defines different types of command execution.
type Type string

const (
	// TypeExec defines immediate command execution.
	TypeExec Type = "exec"
	// TypeCont defines continuous command execution, used for listener etc.
	TypeCont Type = "cont"
)

// Payload is wrapper above exec.
type Payload struct {
	Exec Exec `json:"exec"`
}

// Exec contains Pact code where defines what to do with tx and which data should be passed into.
type Exec struct {
	Data interface{} `json:"data"`
	Code string      `json:"code"`
}

// Meta contains public/private metadata for txs.
type Meta struct {
	ChainID      string  `json:"chainId"`
	Sender       string  `json:"sender"`
	GasLimit     int     `json:"gasLimit"`
	GasPrice     float64 `json:"gasPrice"`
	TTL          int64   `json:"ttl"`
	CreationTime int64   `json:"creationTime"`
}

// Signer is list of signers, corresponding with list of signatures in outer command.
type Signer struct {
	PubKey  string         `json:"pubKey"`
	Address string         `json:"address"`
	Schema  Schema         `json:"schema"`
	CList   CapabilityItem `json:"clist"`
}

// Schema defines approach to sign data.
type Schema string

const (
	// SchemaED25519 defines ED25519 way.
	SchemaED25519 Schema = "ED25519"
	// SchemaETH defines ETH way.
	SchemaETH Schema = "ETH"
)

// CapabilityItem defines scope what the signing keys are allowed to sign.
type CapabilityItem struct {
	Name string      `json:"name"`
	Args interface{} `json:"args"`
}

// CreationTime is helper that returns creation unix time of tx.
func CreationTime() int64 {
	return time.Now().UTC().Unix()
}

// Nonce is helper that returns nonce of current tx.
func Nonce() string {
	return time.Now().UTC().String()
}

// ToJSON converts request to json.
func (request Request) ToJSON() (string, error) {
	marshalled, err := json.Marshal(request)
	return string(marshalled), err
}

// ToHash returns hash of payload.
// Hash in Kadena could be can be calculated using this algorithm base64(blake2b(command_json)).
func (request Request) ToHash() (string, error) {
	requestJSON, err := request.ToJSON()
	if err != nil {
		return "", err
	}

	hashedReq := blake2b.Sum256([]byte(requestJSON))
	var hashBytes = make([]byte, 0, len(hashedReq))
	for _, char := range hashedReq {
		hashBytes = append(hashBytes, char)
	}

	hash := base64.StdEncoding.EncodeToString(hashBytes)
	numOfPadding, isPresent := IsPaddingPresent(hash)
	if isPresent {
		hash = hash[:len(hash)-int(numOfPadding)]
	}

	return hash, nil
}

// IsPaddingPresent checks is base64 encoded string has padding in the end
// and returns num of paddings.
func IsPaddingPresent(hash string) (uint8, bool) {
	const paddingSymbol = "="
	var (
		numOfPadding uint8
		present      bool
	)
	if len(hash) < 2 {
		return numOfPadding, present
	}

	// padding can only be in the last two characters of hash.
	chars := hash[len(hash)-2:]
	for _, ch := range chars {
		if string(ch) == paddingSymbol {
			numOfPadding++
			present = true
		}
	}

	return numOfPadding, present
}

// ToCmd returns cmd hash with decoded to json command payload.
func (request Request) ToCmd() (Command, error) {
	hash, err := request.ToHash()
	if err != nil {
		return Command{}, err
	}

	cmd, err := request.ToJSON()
	if err != nil {
		return Command{}, err
	}

	return Command{
		Hash: hash,
		Cmd:  cmd,
	}, nil
}
