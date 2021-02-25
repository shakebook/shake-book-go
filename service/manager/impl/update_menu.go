package impl

import (
	"context"
	managerpb "shakebook/service/manager/proto/api/v1"
)

//UpdateMenu implement Server
func (s *Server) UpdateMenu(c context.Context, req *managerpb.UpdateMenuRequest) (*managerpb.Response, error) {
	res := &managerpb.Response{}
	if err := req.Validate(); err != nil {
		res.Message = err.Error()
		return res, nil
	}
	err := s.Dao.UpdateMenu(req)
	if err != nil {
		res.Message = err.Error()
		return res, nil
	}
	res.Success = true
	return res, nil
}
