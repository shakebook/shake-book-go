package impl

import (
	"context"
	token "shakebook/common/auth"
	accountpb "shakebook/service/account/proto/api/v1"
)

//GetAccount implement account
func (s *Server) GetAccount(c context.Context, req *accountpb.EmptyRequest) (*accountpb.GetAccountResponse, error) {
	res := &accountpb.GetAccountResponse{}
	a, err := token.GetAccountIDFromContext(c)
	if err != nil {
		res.Message = err.Error()
		return res, nil
	}
	account, err := s.Dao.GetAccount(a)
	if err != nil {
		res.Message = err.Error()
		return res, nil
	}
	res.Success = true
	res.Data = account
	return res, nil
}
