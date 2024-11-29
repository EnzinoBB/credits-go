package model

import (
	"encoding/json"
)

type DelegatedItem struct {
	Wallet     string
	Sum        float64
	ValidUntil int64
	FromTime   int64
	Coeff      int8
}

func (o *DelegatedItem) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}

type Delegated struct {
	Incoming   float64
	Outgoing   float64
	Donors     []*DelegatedItem
	Recipients []*DelegatedItem
}

func (o *Delegated) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}

type WalletData struct {
	WalletId          int32
	Balance           float64
	LastTransactionId int64
	Delegated         *Delegated
}

func (o *WalletData) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}
