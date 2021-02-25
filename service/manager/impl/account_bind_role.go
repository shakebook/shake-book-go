package impl

import (
	"context"
	managerpb "shakebook/service/manager/proto/api/v1"
)

//AccountBindRole implement Server
func (s *Server) AccountBindRole(c context.Context, req *managerpb.AccountBindRoleRequest) (*managerpb.Response, error) {
	res := &managerpb.Response{}
	err := req.Validate()
	if err != nil {
		res.Message = err.Error()
		return res, nil
	}

	if err := s.Dao.AccountBindRole(req); err != nil {
		res.Message = err.Error()
		return res, nil
	}

	res.Success = true
	return res, nil
}
