package impl

import (
	"context"
	managerpb "shakebook/service/manager/proto/api/v1"
)

//GetMenuList implement Server
func (s *Server) GetMenuList(c context.Context, req *managerpb.EmptyRequest) (*managerpb.GetMenuListResponse, error) {
	res := &managerpb.GetMenuListResponse{}
	list, err := s.Dao.GetMenuList()
	if err != nil {
		res.Message = err.Error()
		return res, nil
	}
	res.Data = list
	res.Success = true
	return res, nil
}
