package impl

import (
	"context"
	accountpb "shakebook/service/account/proto/api/v1"
)

//SignOut implement server
func (s *Server) SignOut(c context.Context, req *accountpb.AccountId) (*accountpb.Response, error) {
	res := &accountpb.Response{}
	res.Success = true
	return res, nil
}
