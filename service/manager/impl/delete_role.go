package impl

import (
	"context"
	managerpb "shakebook/service/manager/proto/api/v1"
)

//DeleteRole implement Server
func (s *Server) DeleteRole(c context.Context, req *managerpb.RoleId) (*managerpb.Response, error) {
	res := &managerpb.Response{}
	if err := req.Validate(); err != nil {
		res.Message = err.Error()
		return res, nil
	}
	err := s.Dao.DeleteRole(req)
	if err != nil {
		res.Message = err.Error()
		return res, nil
	}
	res.Success = true
	return res, nil
}
