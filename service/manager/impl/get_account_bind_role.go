package impl

import (
	"context"
	managerpb "shakebook/service/manager/proto/api/v1"
)

//GetAccountBindRole implement Server
func (s *Server) GetAccountBindRole(c context.Context, req *managerpb.AccountId) (*managerpb.GetAccountBindRoleResponse, error) {
	res := &managerpb.GetAccountBindRoleResponse{}
	if err := req.Validate(); err != nil {
		res.Message = err.Error()
		return res, nil
	}

	data, err := s.Dao.GetAccountBindRole(req)
	if err != nil {
		res.Message = err.Error()
		return res, nil
	}

	res.Success = true
	res.Data = data
	return res, nil

}
