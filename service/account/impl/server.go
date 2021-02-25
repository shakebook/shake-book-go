package impl

import (
	"context"
	accountpb "shakebook/service/account/proto/api/v1"
	authpb "shakebook/service/auth/proto/api/v1"

	"go.uber.org/zap"
)

// Server defined
type Server struct {
	Logger       *zap.Logger
	AuthClient   authpb.AuthServiceClient
	Dao          Dao
	ValidEmailer ValidEmailer
}

//Dao db
type Dao interface {
	SignUp(*accountpb.SignUpRequest) error
	SignIn(context.Context, *accountpb.SignInRequest) (*accountpb.AccountId, error)
	UpdateAccountDescript(int64, *accountpb.UpdateAccountDescriptRequest) error
	GetAccount(*accountpb.AccountId) (*accountpb.AccountInfo, error)
	ValidEmail(context.Context, string, string) error
	WriteEmailCodeToRedis(context.Context, string, string) error
	GetAccountList() ([]*accountpb.AccountInfo, error)
}
