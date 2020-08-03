package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"reflect"
)

type PrivateKey = ecdsa.PrivateKey

var (
	secp256k1    *CurveParams
	secp256k1OID asn1.ObjectIdentifier = []int{1, 3, 132, 0, 10}
)

const (
	ecPrivateKeyBlockType = "EC PRIVATE KEY"
	ecPrivateKeyVersion   = 1
)

// Initializes parameters for secp256k1 elliptic curve.
func init() {
	secp256k1 = new(CurveParams)
	secp256k1.Name = "secp256k1"
	secp256k1.P, _ = new(big.Int).SetString("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f", 16)  // Prime
	secp256k1.N, _ = new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)  // Order
	secp256k1.B, _ = new(big.Int).SetString("0000000000000000000000000000000000000000000000000000000000000007", 16)  // B
	secp256k1.Gx, _ = new(big.Int).SetString("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798", 16) // Generator X
	secp256k1.Gy, _ = new(big.Int).SetString("483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8", 16) // Generator Y
	secp256k1.BitSize = 256
}

type ecPrivateKey struct {
	Version       int
	PrivateKey    []byte
	NamedCurveOID asn1.ObjectIdentifier `asn1:"optional,explicit,tag:0"`
	PublicKey     asn1.BitString        `asn1:"optional,explicit,tag:1"`
}

// SignTool it utility, which helps to sign message.
type SignTool struct {
	B2A func([]byte) string
	A2B func(string) ([]byte, error)
}

// DefaultSignTool returns new instance of SignTool SignTool with default encoding parameters.
func DefaultSignTool() *SignTool {
	return &SignTool{
		B2A: base64.StdEncoding.EncodeToString,
		A2B: base64.StdEncoding.DecodeString,
	}
}

// DecodePrivateKey decodes privage key into Elliptic Curve Digital Signature Algorithm private key.
func (t *SignTool) DecodePrivateKey(b []byte) (*PrivateKey, error) {
	var privateKeyPemBlock *pem.Block

	for {
		privateKeyPemBlock, b = pem.Decode(b)
		if privateKeyPemBlock == nil {
			break
		}

		if privateKeyPemBlock.Type == ecPrivateKeyBlockType {
			ret, err := ParseCustomECPrivateKey(privateKeyPemBlock.Bytes)
			if err != nil {
				return nil, err
			}
			return ret, err
		}
	}

	return nil, fmt.Errorf("failed to find private key block")
}

// Sign signs string with specified private key.
func (t *SignTool) Sign(key *PrivateKey, str string) (string, error) {
	hash := sha256.Sum256([]byte(str))

	r, s, err := ecdsa.Sign(rand.Reader, key, hash[:])
	if err != nil {
		return "", err
	}

	asn1Data := []*big.Int{r, s}

	bb, err := asn1.Marshal(asn1Data)
	if err != nil {
		return "", err
	}

	ret := t.B2A(bb)

	return ret, nil
}

// VerifyBytes verifies a digital signature. Returns nil if all is well or an error indicating
// what went wrong.  The equivalent command at the command line is:
// echo -n 'Make America Great Again!' | openssl dgst -verify pub1.pem -signature signature.dat
func (t *SignTool) VerifyBytes(pubkey *ecdsa.PublicKey, b []byte, s string) error {
	return t.VerifyBytesN([]*ecdsa.PublicKey{pubkey}, b, s)
}

// VerifyBytesN verifies with multiple public keys possibly matching.
func (t *SignTool) VerifyBytesN(pubkeys []*ecdsa.PublicKey, b []byte, s string) error {

	if len(pubkeys) == 0 {
		return fmt.Errorf("you must provide at least one key")
	}

	sigb, err := t.A2B(s)
	if err != nil {
		return err
	}

	asn1Data := make([]*big.Int, 0)

	_, err = asn1.Unmarshal(sigb, &asn1Data)
	if err != nil {
		return err
	}

	if len(asn1Data) != 2 {
		return fmt.Errorf("hhile decoding ASN.1 data, expected exactly 2 values, instead got %d", len(asn1Data))
	}

	er, es := asn1Data[0], asn1Data[1]

	hash := sha256.Sum256(b)

	for _, pubkey := range pubkeys {
		if ecdsa.Verify(pubkey, hash[:], er, es) {
			return nil
		}
	}

	return fmt.Errorf("verification failed, no keys matched")
}

func namedCurveFromOID(oid asn1.ObjectIdentifier) elliptic.Curve {
	switch {
	case reflect.DeepEqual(oid, secp256k1OID):
		return secp256k1
	}
	return nil
}

// ParseCustomECPrivateKey returns Elliptic Curve Digital Signature Algorithm private key from file content.
func ParseCustomECPrivateKey(der []byte) (key *PrivateKey, err error) {
	var privKey ecPrivateKey
	if _, err := asn1.Unmarshal(der, &privKey); err != nil {
		return nil, errors.New("x509: failed to parse EC private key: " + err.Error())
	}
	if privKey.Version != ecPrivateKeyVersion {
		return nil, fmt.Errorf("x509: unknown EC private key version %d", privKey.Version)
	}

	curve := namedCurveFromOID(privKey.NamedCurveOID)
	if curve == nil {
		return x509.ParseECPrivateKey(der)
	}

	k := new(big.Int).SetBytes(privKey.PrivateKey)
	curveOrder := curve.Params().N
	if k.Cmp(curveOrder) >= 0 {
		return nil, errors.New("x509: invalid elliptic curve private key value")
	}
	priv := new(PrivateKey)
	priv.Curve = curve
	priv.D = k

	privateKey := make([]byte, (curveOrder.BitLen()+7)/8)

	// Some private keys have leading zero padding. This is invalid
	// according to [SEC1], but this code will ignore it.
	for len(privKey.PrivateKey) > len(privateKey) {
		if privKey.PrivateKey[0] != 0 {
			return nil, errors.New("x509: invalid private key length")
		}
		privKey.PrivateKey = privKey.PrivateKey[1:]
	}

	// Some private keys remove all leading zeros, this is also invalid
	// according to [SEC1] but since OpenSSL used to do this, we ignore
	// this too.
	copy(privateKey[len(privateKey)-len(privKey.PrivateKey):], privKey.PrivateKey)
	priv.X, priv.Y = curve.ScalarBaseMult(privateKey)

	return priv, nil
}

// CurveParams contains the parameters of an elliptic curve and also provides
// a generic, non-constant time implementation of Curve.
type CurveParams struct {
	elliptic.CurveParams
}

// IsOnCurve returns boolean if the point (x,y) is on the curve.
// Part of the elliptic.Curve interface.
func (curve *CurveParams) IsOnCurve(x, y *big.Int) bool {
	// y² = x³ + b
	y2 := new(big.Int).Mul(y, y)
	y2.Mod(y2, curve.P)

	x3 := new(big.Int).Mul(x, x)
	x3.Mul(x3, x)

	//x3.Sub(x3, threeX)
	x3.Add(x3, curve.B)
	x3.Mod(x3, curve.P)

	return x3.Cmp(y2) == 0
}

// affineFromJacobian reverses the Jacobian transform. See the comment at the
// top of the file. If the point is ∞ it returns 0, 0.
func (curve *CurveParams) affineFromJacobian(x, y, z *big.Int) (xOut, yOut *big.Int) {
	if z.Sign() == 0 {
		return new(big.Int), new(big.Int)
	}

	zinv := new(big.Int).ModInverse(z, curve.P)
	zinvsq := new(big.Int).Mul(zinv, zinv)

	xOut = new(big.Int).Mul(x, zinvsq)
	xOut.Mod(xOut, curve.P)
	zinvsq.Mul(zinvsq, zinv)
	yOut = new(big.Int).Mul(y, zinvsq)
	yOut.Mod(yOut, curve.P)
	return
}

func (curve *CurveParams) addJacobian(
	x1, y1, z1, x2, y2, z2 *big.Int,
) (*big.Int, *big.Int, *big.Int) {
	// See http://hyperelliptic.org/EFD/g1p/auto-shortw-jacobian-3.html#addition-add-2007-bl
	x3, y3, z3 := new(big.Int), new(big.Int), new(big.Int)
	if z1.Sign() == 0 {
		x3.Set(x2)
		y3.Set(y2)
		z3.Set(z2)
		return x3, y3, z3
	}
	if z2.Sign() == 0 {
		x3.Set(x1)
		y3.Set(y1)
		z3.Set(z1)
		return x3, y3, z3
	}

	z1z1 := new(big.Int).Mul(z1, z1)
	z1z1.Mod(z1z1, curve.P)
	z2z2 := new(big.Int).Mul(z2, z2)
	z2z2.Mod(z2z2, curve.P)

	u1 := new(big.Int).Mul(x1, z2z2)
	u1.Mod(u1, curve.P)
	u2 := new(big.Int).Mul(x2, z1z1)
	u2.Mod(u2, curve.P)
	h := new(big.Int).Sub(u2, u1)
	xEqual := h.Sign() == 0
	if h.Sign() == -1 {
		h.Add(h, curve.P)
	}
	i := new(big.Int).Lsh(h, 1)
	i.Mul(i, i)
	j := new(big.Int).Mul(h, i)

	s1 := new(big.Int).Mul(y1, z2)
	s1.Mul(s1, z2z2)
	s1.Mod(s1, curve.P)
	s2 := new(big.Int).Mul(y2, z1)
	s2.Mul(s2, z1z1)
	s2.Mod(s2, curve.P)
	r := new(big.Int).Sub(s2, s1)
	if r.Sign() == -1 {
		r.Add(r, curve.P)
	}
	yEqual := r.Sign() == 0
	if xEqual && yEqual {
		return curve.doubleJacobian(x1, y1, z1)
	}
	r.Lsh(r, 1)
	v := new(big.Int).Mul(u1, i)

	x3.Set(r)
	x3.Mul(x3, x3)
	x3.Sub(x3, j)
	x3.Sub(x3, v)
	x3.Sub(x3, v)
	x3.Mod(x3, curve.P)

	y3.Set(r)
	v.Sub(v, x3)
	y3.Mul(y3, v)
	s1.Mul(s1, j)
	s1.Lsh(s1, 1)
	y3.Sub(y3, s1)
	y3.Mod(y3, curve.P)

	z3.Add(z1, z2)
	z3.Mul(z3, z3)
	z3.Sub(z3, z1z1)
	z3.Sub(z3, z2z2)
	z3.Mul(z3, h)
	z3.Mod(z3, curve.P)

	return x3, y3, z3
}

func (curve *CurveParams) doubleJacobian(x, y, z *big.Int) (*big.Int, *big.Int, *big.Int) {
	// See http://hyperelliptic.org/EFD/g1p/auto-shortw-jacobian-0.html#doubling-dbl-2009-l
	a := new(big.Int).Mul(x, x) //X1²
	b := new(big.Int).Mul(y, y) //Y1²
	c := new(big.Int).Mul(b, b) //B²

	d := new(big.Int).Add(x, b) //X1+B
	d.Mul(d, d)                 //(X1+B)²
	d.Sub(d, a)                 //(X1+B)²-A
	d.Sub(d, c)                 //(X1+B)²-A-C
	d.Mul(d, big.NewInt(2))     //2*((X1+B)²-A-C)

	e := new(big.Int).Mul(big.NewInt(3), a) //3*A
	f := new(big.Int).Mul(e, e)             //E²

	x3 := new(big.Int).Mul(big.NewInt(2), d) //2*D
	x3.Sub(f, x3)                            //F-2*D
	x3.Mod(x3, curve.P)

	y3 := new(big.Int).Sub(d, x3)                  //D-X3
	y3.Mul(e, y3)                                  //E*(D-X3)
	y3.Sub(y3, new(big.Int).Mul(big.NewInt(8), c)) //E*(D-X3)-8*C
	y3.Mod(y3, curve.P)

	z3 := new(big.Int).Mul(y, z) //Y1*Z1
	z3.Mul(big.NewInt(2), z3)    //3*Y1*Z1
	z3.Mod(z3, curve.P)

	return x3, y3, z3
}

// ScalarMult returns k*(Bx, By) where k is a big endian integer.
// Part of the elliptic.Curve interface.
func (curve *CurveParams) ScalarMult(Bx, By *big.Int, k []byte) (*big.Int, *big.Int) {
	Bz := new(big.Int).SetInt64(1)
	x, y, z := new(big.Int), new(big.Int), new(big.Int)

	for _, b := range k {
		for bitNum := 0; bitNum < 8; bitNum++ {
			x, y, z = curve.doubleJacobian(x, y, z)
			if b&0x80 == 0x80 {
				x, y, z = curve.addJacobian(Bx, By, Bz, x, y, z)
			}
			b <<= 1
		}
	}

	return curve.affineFromJacobian(x, y, z)
}

// ScalarBaseMult returns k*G where G is the base point of the group and k is a
// big endian integer.
// Part of the elliptic.Curve interface.
func (curve *CurveParams) ScalarBaseMult(k []byte) (*big.Int, *big.Int) {
	return curve.ScalarMult(curve.Gx, curve.Gy, k)
}
