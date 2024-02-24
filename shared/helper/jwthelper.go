package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtHelper struct {
	_token_  string
	_secret_ string
	_jwttk_  *jwt.Token
	jwt.RegisteredClaims
	Payload map[string]interface{} `json:"payload"`
}

func NewJwtHelper(secret string) *JwtHelper {
	return &JwtHelper{
		_token_:  "",
		_secret_: secret,
		Payload:  map[string]interface{}{},
		RegisteredClaims: jwt.RegisteredClaims{
			ID:       "",
			Issuer:   "",
			Subject:  "",
			Audience: []string{},
		},
	}
}

func (o *JwtHelper) Token() string {
	return o._token_
}

func (o *JwtHelper) SetIssuer(value string) *JwtHelper {
	o.RegisteredClaims.Issuer = value
	return o
}

func (o *JwtHelper) SetSubject(value string) *JwtHelper {
	o.RegisteredClaims.Subject = value
	return o
}

func (o *JwtHelper) SetAudience(value string) *JwtHelper {
	o.RegisteredClaims.Audience = append(o.RegisteredClaims.Audience, value)
	return o
}

func (o *JwtHelper) SetExpires(value time.Time) *JwtHelper {
	o.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(value)
	return o
}

func (o *JwtHelper) SetNotBefore(value time.Time) *JwtHelper {
	o.RegisteredClaims.NotBefore = jwt.NewNumericDate(value)
	return o
}

func (o *JwtHelper) SetIssuedAt(value time.Time) *JwtHelper {
	o.RegisteredClaims.IssuedAt = jwt.NewNumericDate(value)
	return o
}

func (o *JwtHelper) SetPayload(key string, value interface{}) *JwtHelper {
	o.Payload[key] = value
	return o
}

func (o *JwtHelper) GetPayload(key string) interface{} {
	return o.Payload[key]
}

func (o *JwtHelper) Sign() error {

	var err error
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, o)
	o._token_, err = token.SignedString([]byte(o._secret_))
	if err != nil {
		return err
	}
	return nil
}

func (o *JwtHelper) Load(tk string) error {

	var err error
	o._token_ = tk
	o._jwttk_, err = jwt.ParseWithClaims(tk, o, func(token *jwt.Token) (interface{}, error) {
		return []byte(o._secret_), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (o *JwtHelper) TokenIsValid() bool {

	if o._jwttk_ == nil {
		return false
	}
	return o._jwttk_.Valid
}
