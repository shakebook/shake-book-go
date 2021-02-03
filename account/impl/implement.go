package impl

import (
	"context"
	accountpb "shakebook/account/proto/api/v1"
)

// Server defined
type Server struct{}

//GetAccount implement account
func (*Server) GetAccount(context.Context, *accountpb.GetAccountRequest) (*accountpb.GetAccountResponse, error) {
	return &accountpb.GetAccountResponse{
		Id:        1,
		FirstName: "yang",
		LastName:  "jiafeng",
	}, nil
}

//CreateAccount implement server
func (*Server) CreateAccount(c context.Context, req *accountpb.CreateAccountRequest) (*accountpb.GetAccountResponse, error) {
	return &accountpb.GetAccountResponse{
		Id:        1,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}, nil
}
