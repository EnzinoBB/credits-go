package model

import (
	"crypto/ed25519"
	"fmt"
	"strconv"

	"github.com/EnzinoBB/credits-go/api"
	"github.com/EnzinoBB/credits-go/utils"
)

func GetWalletData_Out(data *api.WalletData) *WalletData {

	if data == nil {
		return nil
	}

	balance, _ := strconv.ParseFloat(fmt.Sprintf("%d.%d", data.Balance.Integral, data.Balance.Fraction), 64)

	return &WalletData{
		WalletId:          int32(data.WalletId),
		Balance:           balance,
		LastTransactionId: int64(data.LastTransactionId),
		Delegated:         GetDelegated_Out(data.Delegated),
	}

}

func GetWalletData_In(data *WalletData) *api.WalletData {
	return nil
}

func GetDelegated_Out(data *api.Delegated) *Delegated {

	if data == nil {
		return nil
	}

	incoming, _ := strconv.ParseFloat(fmt.Sprintf("%d.%d", data.Incoming.Integral, data.Incoming.Fraction), 64)
	outgoing, _ := strconv.ParseFloat(fmt.Sprintf("%d.%d", data.Outgoing.Integral, data.Outgoing.Fraction), 64)

	donors := []*DelegatedItem{}
	for _, donor := range data.Donors {

		donor_out := GetDelegatedItem_Out(donor)
		donors = append(donors, donor_out)

	}

	recipients := []*DelegatedItem{}
	for _, recipient := range data.Recipients {

		recipient_out := GetDelegatedItem_Out(recipient)
		recipients = append(recipients, recipient_out)

	}

	return &Delegated{
		Incoming:   incoming,
		Outgoing:   outgoing,
		Donors:     donors,
		Recipients: recipients,
	}
}

func GetDelegated_In(data *Delegated) *api.Delegated {
	return nil
}

func GetDelegatedItem_Out(data *api.DelegatedItem) *DelegatedItem {

	if data == nil {
		return nil
	}

	sum, _ := strconv.ParseFloat(fmt.Sprintf("%d.%d", data.Sum.Integral, data.Sum.Fraction), 64)

	return &DelegatedItem{
		Wallet:     utils.GetPublicKey(ed25519.PublicKey(data.Wallet)),
		Sum:        sum,
		ValidUntil: *data.ValidUntil,
		FromTime:   *data.FromTime,
		Coeff:      *data.Coeff,
	}
}

func GetDelegatedItem_In(data *DelegatedItem) *api.DelegatedItem {
	return nil
}
