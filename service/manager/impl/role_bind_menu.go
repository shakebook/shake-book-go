package impl

import (
	"context"
	managerpb "shakebook/service/manager/proto/api/v1"
)

//RoleBindMenu implement Server
func (s *Server) RoleBindMenu(ctx context.Context, req *managerpb.RoleBindMenuRequest) (*managerpb.Response, error) {
	res := &managerpb.Response{}
	if err := req.Validate(); err != nil {
		res.Message = err.Error()
		return res, nil
	}

	if err := s.Dao.RoleBindMenu(req); err != nil {
		res.Message = err.Error()
		return res, nil
	}

	res.Success = true
	return res, nil
}
