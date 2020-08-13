package bgn

import (
	"math/big"
	"testing"
)

const KEYBITS = 512
const POLYBASE = 3
const MSGSPACE = 1021 // message space for polynomial coefficients
const FPSCALEBASE = 3
const FPPREC = 0.0001
const DET = true // deterministic ops

func TestMarshalUnmarshalPublicKey(t *testing.T) {
	pk, _, err := NewKeyGen(KEYBITS, big.NewInt(MSGSPACE), POLYBASE, FPSCALEBASE, FPPREC, DET)
	if err != nil {
		panic(err)
	}

	bytes, _ := pk.MarshalBinary()
	pk.UnmarshalBinary(bytes)

}

func TestMarshalUnmarshalCiphertext(t *testing.T) {
	pk, _, err := NewKeyGen(KEYBITS, big.NewInt(MSGSPACE), POLYBASE, FPSCALEBASE, FPPREC, DET)
	if err != nil {
		panic(err)
	}

	ct := pk.encryptZero()

	bytes, _ := ct.MarshalBinary()
	ct.UnmarshalBinary(bytes)
}

func TestMarshalUnmarshalPublicKeyNil(t *testing.T) {
	pk := &PublicKey{}
	bytes, _ := pk.MarshalBinary()
	pk.UnmarshalBinary(bytes)

}

func TestMarshalUnmarshalCiphertextNil(t *testing.T) {

	ct := &Ciphertext{}
	bytes, _ := ct.MarshalBinary()
	ct.UnmarshalBinary(bytes)
}

func BenchmarkKeyGen(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_, _, err := NewKeyGen(KEYBITS, big.NewInt(MSGSPACE), POLYBASE, FPSCALEBASE, FPPREC, DET)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	pk, _, err := NewKeyGen(KEYBITS, big.NewInt(MSGSPACE), POLYBASE, FPSCALEBASE, FPPREC, DET)
	if err != nil {
		panic(err)
	}

	c := pk.Encrypt(big.NewInt(1))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pk.Add(c, c)
	}
}

func BenchmarkMultConstant(b *testing.B) {
	pk, _, err := NewKeyGen(KEYBITS, big.NewInt(MSGSPACE), POLYBASE, FPSCALEBASE, FPPREC, DET)
	if err != nil {
		panic(err)
	}

	c := pk.Encrypt(big.NewInt(1))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pk.MultConst(c, big.NewInt(1))
	}
}

func BenchmarkMult(b *testing.B) {
	pk, _, err := NewKeyGen(KEYBITS, big.NewInt(MSGSPACE), POLYBASE, FPSCALEBASE, FPPREC, DET)
	if err != nil {
		panic(err)
	}

	c := pk.Encrypt(big.NewInt(1))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pk.Mult(c, c)
	}
}
