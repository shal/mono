package auth

import (
	"crypto/elliptic"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"shal.dev/mono/ecdsa"
)

type CorporateAuth struct {
	signer     *ecdsa.SignTool
	privateKey *ecdsa.PrivateKey
	KeyID      string
}

func NewCorporateAuth(keyData []byte) (*CorporateAuth, error) {
	sign := ecdsa.DefaultSignTool()

	privateKey, err := sign.DecodePrivateKey(keyData)
	if err != nil {
		return nil, errors.New("failed to decode private key")
	}

	publicKey := privateKey.PublicKey
	data := elliptic.Marshal(publicKey, publicKey.X, publicKey.Y)
	hash := sha1.New()
	if _, err := hash.Write(data); err != nil {
		return nil, errors.New("failed to encode public key with sha1")
	}
	keyID := hex.EncodeToString(hash.Sum(nil))

	return &CorporateAuth{
		signer:     sign,
		privateKey: privateKey,
		KeyID:      keyID,
	}, nil
}

func (auth *CorporateAuth) sign(params ...string) (string, error) {
	return auth.signer.Sign(auth.privateKey, strings.Join(params, ""))
}

func (auth *CorporateAuth) Sign(r *http.Request) error {
	switch r.URL.Path {
	case "/personal/auth/request":
		switch r.Method {
		case http.MethodPost:
			timestamp := strconv.Itoa(int(time.Now().Unix()))
			pp := r.Header.Get("X-Permissions")

			sign, err := auth.sign(timestamp, pp, r.URL.Path)
			if err != nil {
				return err
			}

			r.Header.Set("X-Sign", sign)
		case http.MethodGet:
			timestamp := strconv.Itoa(int(time.Now().Unix()))
			reqID := r.Header.Get("X-Request-Id")

			sign, err := auth.sign(timestamp, reqID, r.URL.Path)
			if err != nil {
				return err
			}

			r.Header.Set("X-Sign", sign)
		}
	}

	return nil
}

func (auth *CorporateAuth) Auth(r *http.Request) error {
	//timestamp := strconv.Itoa(int(time.Now().Unix()))

	//r.Header.Set("X-Key-Id", auth.KeyID)
	//r.Header.Set("X-Time", timestamp)

	return nil
}
