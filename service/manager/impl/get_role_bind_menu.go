package impl

import (
	"context"
	managerpb "shakebook/service/manager/proto/api/v1"
)

//GetRoleBindMenu implement Server
func (s *Server) GetRoleBindMenu(c context.Context, req *managerpb.RoleId) (*managerpb.GetRoleBindMenuResponse, error) {
	res := &managerpb.GetRoleBindMenuResponse{}
	if err := req.Validate(); err != nil {
		res.Message = err.Error()
		return res, nil
	}

	data, err := s.Dao.GetRoleBindMenu(req)
	if err != nil {
		res.Message = err.Error()
		return res, nil
	}
	res.Data = data
	res.Success = true
	return res, nil
}
