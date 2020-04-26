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

func (p *Client) verifySignature(c echo.Context, req *http.Request, body []byte) (*MessagingAPIErrorJSON, error) {
	a, err := base64.StdEncoding.DecodeString(req.Header.Get("X-Line-Signature"))
	if err != nil {
		return DoMessagingAPIError(c, "X-Line-Signature is invalid.", http.StatusBadRequest)
	}

	b, err := p.getDigest([]byte(p.Secret), body)
	if err != nil {
		return DoMessagingAPIError(c, "Failed getting degest.", http.StatusInternalServerError)
	}

	if !hmac.Equal(a, b) {
		return DoMessagingAPIError(c, "X-Line-Signature is invalid.", http.StatusForbidden)
	}
	return nil, nil
}
