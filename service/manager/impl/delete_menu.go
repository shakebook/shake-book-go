package impl

import (
	"context"
	managerpb "shakebook/service/manager/proto/api/v1"
)

//DeleteMenu implement Server
func (s *Server) DeleteMenu(c context.Context, req *managerpb.MenuId) (*managerpb.Response, error) {
	res := &managerpb.Response{}
	if err := req.Validate(); err != nil {
		res.Message = err.Error()
		return res, nil
	}
	err := s.Dao.DeleteMenu(req)
	if err != nil {
		res.Message = err.Error()
		return res, nil
	}
	res.Success = true
	return res, nil
}
