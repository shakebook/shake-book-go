package auth

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	accountpb "shakebook/service/account/proto/api/v1"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	authorization = "token"
	bearerPrefix  = "BEARER "
)

//Interceptor .
type Interceptor struct {
	token Verifier
}

//Verifier implement this interface
type Verifier interface {
	Verify(token string) (string, error)
}

//NewInterceptor new interceptor
func NewInterceptor(path string) (grpc.UnaryServerInterceptor, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannnot open public key file: %v", err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("cannot read public key: %v", err)
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("canot parse public key: %v", err)
	}
	in := &Interceptor{
		token: &TokenVerifier{
			PublicKey: pubKey,
		},
	}
	return in.UnaryServerInterceptor, nil
}

//UnaryServerInterceptor .
func (i *Interceptor) UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if IsAuth(info.FullMethod) {
		token, err := getTokenFromContext(ctx)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
		accountID, err := i.token.Verify(token)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
		ctx = ContextWithAccountID(ctx, accountID)
	}

	return handler(ctx, req)
}

//ContextWithAccountIDKey context with account id key
type ContextWithAccountIDKey struct{}

// ContextWithAccountID creates a context with given account ID.
func ContextWithAccountID(c context.Context, aid string) context.Context {
	return context.WithValue(c, ContextWithAccountIDKey{}, aid)
}

func getTokenFromContext(ctx context.Context) (string, error) {
	unauthenticated := status.Error(codes.Unauthenticated, "")
	m, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", unauthenticated
	}

	token := ""
	for _, v := range m[authorization] {
		if strings.HasPrefix(v, bearerPrefix) {
			token = v[len(bearerPrefix):]
		}
	}

	if token == "" {
		return "", unauthenticated
	}

	return token, nil
}

//GetAccountIDFromContext get account id from context
func GetAccountIDFromContext(c context.Context) (*accountpb.AccountId, error) {
	accountStr, ok := c.Value(ContextWithAccountIDKey{}).(string)
	var accountID int64
	if !ok {
		return nil, errors.New(" get account id from context failed")
	}
	accid, err := strconv.Atoi(accountStr)
	if err != nil {
		return nil, err
	}
	accountID = int64(accid)
	return &accountpb.AccountId{Id: accountID}, nil
}
