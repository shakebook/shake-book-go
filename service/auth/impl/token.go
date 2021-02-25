package impl

import (
	"context"
	"crypto/rsa"
	"time"

	authpb "shakebook/service/auth/proto/api/v1"

	"github.com/dgrijalva/jwt-go"
)

//Server implement
type Server struct {
	privateKey *rsa.PrivateKey
	issuer     string
	nowFunc    func() time.Time
	expire     time.Duration
}

//GenToken generate token
func (s *Server) GenToken(c context.Context, req *authpb.GenTokenRequest) (*authpb.GenTokenResponse, error) {
	nowSec := s.nowFunc().Unix()
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		Issuer:    s.issuer,
		IssuedAt:  nowSec,
		ExpiresAt: nowSec + int64(s.expire.Seconds()),
		Subject:   req.Id,
	})

	tokenstring, err := tkn.SignedString(s.privateKey)
	return &authpb.GenTokenResponse{
		Token: tokenstring,
	}, err

}

//NewServer new a token server
func NewServer(privateKey *rsa.PrivateKey, issuer string, expire time.Duration) *Server {
	return &Server{
		privateKey: privateKey,
		issuer:     issuer,
		expire:     expire,
		nowFunc:    time.Now,
	}
}
