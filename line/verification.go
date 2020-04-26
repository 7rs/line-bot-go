package line

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"

	"github.com/labstack/echo"
)

func (p *Client) getDigest(key []byte, body []byte) ([]byte, error) {
	hash := hmac.New(sha256.New, []byte(p.Secret))
	_, err := hash.Write(body)
	if err != nil {
		return []byte{}, err
	}
	return hash.Sum(nil), nil
}

func (p *Client) verifySignature(c echo.Context, req *http.Request, body []byte) error {
	a, err := base64.StdEncoding.DecodeString(req.Header.Get("X-Line-Signature"))
	if err != nil {
		return newLineAPIError(c, http.StatusBadRequest, "X-Line-Signature is invalid.")
	}

	b, err := p.getDigest([]byte(p.Secret), body)
	if err != nil {
		return newLineAPIError(c, http.StatusInternalServerError, "Failed getting degest.")
	}

	if !hmac.Equal(a, b) {
		return newLineAPIError(c, http.StatusForbidden, "X-Line-Signature is invalid.")
	}
	return nil
}
