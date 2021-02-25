package impl

import (
	"context"
	managerpb "shakebook/service/manager/proto/api/v1"
)

//GetRoleList implement Server
func (s *Server) GetRoleList(c context.Context, req *managerpb.EmptyRequest) (*managerpb.GetRoleListResponse, error) {
	res := &managerpb.GetRoleListResponse{}
	list, err := s.Dao.GetRoleList()
	if err != nil {
		res.Message = err.Error()
		return res, nil
	}
	res.Data = list
	res.Success = true
	return res, nil
}
