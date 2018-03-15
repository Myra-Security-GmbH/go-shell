package authentication

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"strings"
)

//
// MyraSignature ...
//
type MyraSignature struct {
	Secret string
}

//
// Build ...
//
func (s *MyraSignature) Build(
	content string,
	method string,
	url string,
	requestType string,
	date string,
) (string, error) {
	signingString := fmt.Sprintf(
		"%x#%s#%s#%s#%s",
		md5.Sum([]byte(content)),
		strings.ToUpper(method),
		url,
		requestType,
		date,
	)

	h1 := hmac.New(sha256.New, []byte("MYRA"+s.Secret))
	_, err := h1.Write([]byte(date))

	if err != nil {
		return "", err
	}

	dateKey := fmt.Sprintf("%x", h1.Sum(nil))

	h2 := hmac.New(sha256.New, []byte(dateKey))
	_, err = h2.Write([]byte("myra-api-request"))
	if err != nil {
		return "", err
	}

	signingKey := fmt.Sprintf("%x", h2.Sum(nil))
	h3 := hmac.New(sha512.New, []byte(signingKey))
	_, err = h3.Write([]byte(signingString))

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(h3.Sum(nil)), nil
}
