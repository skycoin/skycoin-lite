package liteclient

import (
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/cipher/secp256k1-go"
)

/*
Functions used mainly during test procedures.
*/

// VerifySignature verifies that hash was signed by PubKey
func VerifySignature(pubkey string, sig string, hash string) {
	p := cipher.MustPubKeyFromHex(pubkey)
	s := cipher.MustSigFromHex(sig)
	h := cipher.MustSHA256FromHex(hash)

	err := cipher.VerifySignature(p, s, h)
	if err != nil {
		panic(err)
	}
}

// ChkSig checks whether PubKey corresponding to address hash signed hash
func ChkSig(address string, hash string, sig string) {
	a := cipher.MustDecodeBase58Address(address)
	h := cipher.MustSHA256FromHex(hash)
	s := cipher.MustSigFromHex(sig)

	err := cipher.ChkSig(a, h, s)
	if err != nil {
		panic(err)
	}
}

// VerifySignedHash this only checks that the signature can be converted to a public key
// Since there is no pubkey or address argument, it cannot check that the
// signature is valid in that context.
func VerifySignedHash(sig string, hash string) {
	s := cipher.MustSigFromHex(sig)
	h := cipher.MustSHA256FromHex(hash)

	err := cipher.VerifySignedHash(s, h)
	if err != nil {
		panic(err)
	}
}

// VerifySeckey validate a private key
func VerifySeckey(seckey string) int {
	s := cipher.MustSecKeyFromHex(seckey)
	return secp256k1.VerifySeckey(s[:])
}

// VerifyPubkey validate a public key
func VerifyPubkey(pubkey string) int {
	p := cipher.MustPubKeyFromHex(pubkey)
	return secp256k1.VerifyPubkey(p[:])
}

// AddressFromPubKey creates Address from PubKey as ripemd160(sha256(sha256(pubkey)))
func AddressFromPubKey(pubkey string) string {
	p := cipher.MustPubKeyFromHex(pubkey)
	return cipher.AddressFromPubKey(p).String()
}

// AddressFromSecKey generates address from secret key
func AddressFromSecKey(seckey string) string {
	s := cipher.MustSecKeyFromHex(seckey)
	return cipher.AddressFromSecKey(s).String()
}

// PubKeyFromSig recovers the public key from a signed hash
func PubKeyFromSig(sig string, hash string) string {
	s := cipher.MustSigFromHex(sig)
	h := cipher.MustSHA256FromHex(hash)
	
	pubKey, err := cipher.PubKeyFromSig(s, h)
	if err != nil {
		panic(err)
	}

	return pubKey.Hex()
}

// SignHash sign hash
func SignHash(hash string, seckey string) string {
	h := cipher.MustSHA256FromHex(hash)
	s := cipher.MustSecKeyFromHex(seckey)

	sig := cipher.SignHash(h, s)
	return sig.Hex()
}