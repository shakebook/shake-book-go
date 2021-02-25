package impl

import (
	"context"
	managerpb "shakebook/service/manager/proto/api/v1"
)

//CreateMenu implement server
func (s *Server) CreateMenu(c context.Context, req *managerpb.CreateMenuRequest) (*managerpb.Response, error) {
	res := &managerpb.Response{}
	if err := req.Validate(); err != nil {
		res.Message = err.Error()
		return res, nil
	}

	if err := s.Dao.CreateMenu(req); err != nil {
		res.Message = err.Error()
		return res, nil
	}

	res.Success = true
	return res, nil
}
