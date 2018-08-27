package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

func NewSigningKey() (*ecdsa.PrivateKey, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return key, err
}

func Sign(data []byte, priv *ecdsa.PrivateKey) ([]byte, error) {
	digest := sha256.Sum256(data)
	r, s, err := ecdsa.Sign(rand.Reader, priv, digest[:])
	if err != nil {
		return nil, err
	}
	// encode the signature {R, S}
	params := priv.Curve.Params()
	curveByteSize := params.P.BitLen() / 8
	rBytes, sBytes := r.Bytes(), s.Bytes()
	signature := make([]byte, curveByteSize*2)
	copy(signature[curveByteSize-len(rBytes):], rBytes)
	copy(signature[curveByteSize*2-len(sBytes):], sBytes)
	return signature, nil
}

// Verify Returns true if it's valid and false if not.
func Verify(data, sig []byte, pub *ecdsa.PublicKey) bool {
	digest := sha256.Sum256(data)
	curveByteSize := pub.Curve.Params().P.BitLen() / 8
	r, s := new(big.Int), new(big.Int)
	r.SetBytes(sig[:curveByteSize])
	s.SetBytes(sig[curveByteSize:])
	return ecdsa.Verify(pub, digest[:], r, s)
}

func main() {
	priv, err := NewSigningKey()
	if err != nil {
		fmt.Println("priv:", priv)
		return
	}
	// 接口类型转换
	pub := priv.Public().(*ecdsa.PublicKey)
	data := []byte("hello")
	sig, err := Sign(data, priv)
	if err != nil {
		return
	}
	// 验证需要原文，密文，公钥三个进行验证
	fmt.Println(Verify(data, sig, pub))

}
