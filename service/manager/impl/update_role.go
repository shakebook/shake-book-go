package impl

import (
	"context"
	managerpb "shakebook/service/manager/proto/api/v1"
)

//UpdateRole implement Server
func (s *Server) UpdateRole(c context.Context, req *managerpb.Role) (*managerpb.Response, error) {
	res := &managerpb.Response{}
	if err := req.Validate(); err != nil {
		res.Message = err.Error()
		return res, nil
	}

	err := s.Dao.UpdateRole(req)
	if err != nil {
		res.Message = err.Error()
		return res, nil
	}
	res.Success = true
	return res, nil
}
