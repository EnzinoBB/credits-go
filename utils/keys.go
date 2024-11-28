package utils

import (
	"crypto/ed25519"
	"fmt"

	"github.com/akamensky/base58"
)

func GetPublicKey(publicKey ed25519.PublicKey) string {
	return base58.Encode(publicKey)
}

func GetPrivateKey(privateKey ed25519.PrivateKey) (string, error) {

	pubKeyAny := privateKey.Public()
	pubKey, result := pubKeyAny.(ed25519.PublicKey)
	if !result {
		return "", fmt.Errorf("Error converting public key")
	}

	slice := append(privateKey.Seed(), pubKey...)
	return base58.Encode(slice), nil
}

func GetPublicKeyFromPrivate(privateKey string) (string, error) {

	slice, err := base58.Decode(privateKey)
	if err != nil {
		return "", err
	}

	publicKey := slice[32:]
	return GetPublicKey(publicKey), nil
}

func GenerateKeys() (string, string, error) {

	pub, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		return "", "", err
	}
	privateKey, err := GetPrivateKey(priv)
	if err != nil {
		return "", "", err
	}
	publicKey := GetPublicKey(pub)

	return publicKey, privateKey, nil
}
