package impl

import (
	"context"
	accountpb "shakebook/service/account/proto/api/v1"
)

//GetAccountList implement Server
func (s *Server) GetAccountList(context.Context, *accountpb.EmptyRequest) (*accountpb.GetAccountListResponse, error) {
	res := &accountpb.GetAccountListResponse{}
	data, err := s.Dao.GetAccountList()
	if err != nil {
		res.Message = err.Error()
		return res, nil
	}
	res.Data = data
	res.Success = true
	return res, nil
}
