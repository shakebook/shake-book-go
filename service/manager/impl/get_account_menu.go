package impl

import (
	"context"
	token "shakebook/common/auth"
	managerpb "shakebook/service/manager/proto/api/v1"
)

//GetAccountMenu implement Server
func (s *Server) GetAccountMenu(c context.Context, req *managerpb.EmptyRequest) (*managerpb.GetAccountMenuResponse, error) {
	res := &managerpb.GetAccountMenuResponse{}
	acc, err := token.GetAccountIDFromContext(c)

	if err != nil {
		res.Message = err.Error()
		return res, nil
	}
	data, err := s.Dao.GetAccountMenu(acc)
	if err != nil {
		res.Message = err.Error()
		return res, nil
	}
	res.Success = true
	res.Data = data
	return res, nil
}
