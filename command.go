package gokadena

// Command describes Kadena command to interact with blockchain.
type Command struct {
	Hash string   `json:"hash"`
	Sigs []Signer `json:"sigs"`
	Cmd  string   `json:"cmd"`
}
