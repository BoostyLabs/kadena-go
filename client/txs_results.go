package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zeebo/errs"
)

// CommandsResultRequest holds unique IDs of a pact transaction consisting of its hash.
type CommandsResultRequest struct {
	RequestKeys []string `json:"requestKeys"`
}

// CommandsResultResponse holds response for single tx which have been pulled.
type CommandsResultResponse struct {
	Gas      int      `json:"gas"`
	ReqKey   string   `json:"reqKey"`
	Result   Result   `json:"result"`
	TxID     int64    `json:"txId"`
	Logs     string   `json:"logs"`
	Metadata Metadata `json:"metaData"`
	Events
	Continuation Continuation `json:"continuation"`
}

// Result consists tx status with error message/data if it's required.
type Result struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}

// Status is list of all possible tx status.
type Status string

const (
	// StatusSuccess is status for successful tx.
	StatusSuccess Status = "success"
	// StatusFailure is status for failed tx.
	StatusFailure Status = "failure"
)

// Metadata contains tx metadata.
type Metadata struct {
	BlockTime     int64  `json:"blockTime" help:"UNIX number"`
	PrevBlockHash string `json:"prevBlockHash"`
	BlockHash     string `json:"blockHash"`
	BlockHeight   int64  `json:"blockHeight"`
}

// Events holds all event specific characteristic.
type Events struct {
	Name       string      `json:"name"`
	Module     Module      `json:"module"`
	ModuleHash string      `json:"moduleHash"`
	Params     interface{} `json:"params"`
}

// Module defines module name where event occurred.
type Module struct {
	Namespace interface{} `json:"namespace"`
	Name      string      `json:"name"`
}

// Continuation describes result of a defpact execution.
type Continuation struct {
	PactID          string `json:"pactId"`
	Step            int    `json:"step"`
	StepCount       int    `json:"stepCount"`
	Executed        bool   `json:"executed"`
	StepHasRollback bool   `json:"stepHasRollback"`
	// TODO: add remaining fields.
}

// TxsResult pulls for one or more command results by request key.
func (client *Client) TxsResult(ctx context.Context, commandsResult CommandsResultRequest) (_ map[string]CommandsResultResponse, err error) {
	url := fmt.Sprintf(
		"%v/%v/%v/chain/%v/pact/api/v1/poll",
		client.config.NodeAddress,
		client.config.APIVersion,
		client.config.ChainName,
		client.config.ChainID,
	)

	body, err := json.Marshal(commandsResult)
	if err != nil {
		return map[string]CommandsResultResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return map[string]CommandsResultResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.http.Do(req.WithContext(ctx))
	if err != nil {
		return map[string]CommandsResultResponse{}, err
	}

	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	result := make(map[string]CommandsResultResponse, len(commandsResult.RequestKeys))
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return map[string]CommandsResultResponse{}, err
	}

	return result, nil
}
