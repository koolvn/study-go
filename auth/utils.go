package auth

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func loadPrivateKey(path string) (ed25519.PrivateKey, error) {
	// Read the PEM file
	pemData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Decode the PEM data
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, err
	}

	// Parse the private key
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// Assert the type to ed25519.PrivateKey and return
	edPrivateKey, ok := privateKey.(ed25519.PrivateKey)
	if !ok {
		return nil, err
	}

	return edPrivateKey, nil
}

func loadPublicKey(path string) (ed25519.PublicKey, error) {
	// Read the PEM file
	pemData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Decode the PEM data
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, err
	}

	// Parse the public key
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// Assert the type to ed25519.PublicKey and return
	edPublicKey, ok := publicKey.(ed25519.PublicKey)
	if !ok {
		return nil, err
	}

	return edPublicKey, nil
}
