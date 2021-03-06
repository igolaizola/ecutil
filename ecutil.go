package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	rawHex := flag.String("hex", "", "ecdsa public key in raw hexadecimal format")
	pem := flag.String("pem", "", "ecdsa public key in pem format")
	curve := flag.String("curve", "P256", "elliptic curve")
	flag.Parse()
	switch {
	case *rawHex != "":
		data, err := toPEM(*curve, *rawHex)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))
	case *pem != "":
		data, err := toHex(*curve, *pem)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%x\n", data)
	default:
		log.Fatal("invalid input")
	}
}

func toHex(curve, rawPem string) ([]byte, error) {
	var c elliptic.Curve
	switch strings.ToUpper(curve) {
	case "P224":
		c = elliptic.P224()
	case "P256":
		c = elliptic.P256()
	case "P384":
		c = elliptic.P384()
	default:
		return nil, fmt.Errorf("unsupported curve %s", curve)
	}
	block, _ := pem.Decode([]byte(rawPem))
	output, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pubKey, ok := output.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("invalid type %T", output)
	}
	data := elliptic.Marshal(c, pubKey.X, pubKey.Y)
	return data, nil
}

func toPEM(curve, rawHex string) ([]byte, error) {
	var c elliptic.Curve
	switch strings.ToUpper(curve) {
	case "P224":
		c = elliptic.P224()
	case "P256":
		c = elliptic.P256()
	case "P384":
		c = elliptic.P384()
	default:
		return nil, fmt.Errorf("unsupported curve %s", curve)
	}
	raw, err := hex.DecodeString(rawHex)
	if err != nil {
		return nil, err
	}
	x, y := elliptic.Unmarshal(c, raw)
	if x == nil {
		return nil, err
	}
	var pubKey ecdsa.PublicKey = ecdsa.PublicKey{
		Curve: c,
		X:     x,
		Y:     y,
	}
	derKey, err := x509.MarshalPKIXPublicKey(&pubKey)
	if err != nil {
		return nil, err
	}

	keyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derKey,
	}
	return pem.EncodeToMemory(keyBlock), nil
}
