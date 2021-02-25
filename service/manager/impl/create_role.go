package impl

import (
	"context"
	managerpb "shakebook/service/manager/proto/api/v1"
)

//CreateRole implement server
func (s *Server) CreateRole(c context.Context, req *managerpb.Role) (*managerpb.Response, error) {
	if err := req.Validate(); err != nil {
		res := &managerpb.Response{}
		res.Message = err.Error()
		return res, nil
	}

	return s.Dao.CreateRole(req)
}
