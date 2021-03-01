/*
SPDX-License-Identifier: Apache-2.0
PROGRAMMERS: kutenglaoshu, carrie
*/

package common

import (
	"net"
	"time"
)

type TXARGS struct {
	Txid    string   `json:"txid"`
	Err     string   `json:"err,omitempty"`
	Valid   int32    `json:"valid,omitempty"`
	Elapsed int64    `json:"elapsed,omitempty"`
	Name    string   `json:"name,omitempty"`
	Verson  string   `json:"verson,omitempty"`
	Path    string   `json:"path,omitempty"`
	Args    []string `json:"args"`
	// The payload of response
	Payload []byte `json:"payload,omitempty"`
}

type TXMOCK struct {
	From    int    `json:"from"`
	To      int    `json:"to"`
	Name    string `json:"name"`
	Valid   int    `json:"valid,omitempty"`
	Invalid int    `json:"invalid,omitempty"`
	Elapsed int64  `json:"elapsed,omitempty"`
}

func TestAddress(addr string) error {
	c, err := net.DialTimeout("tcp", addr, time.Second*2) //同dail 多加超时参数
	if err != nil {
		return err
	}
	c.Close()
	return nil
}
