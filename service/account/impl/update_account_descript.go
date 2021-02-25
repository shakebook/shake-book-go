package impl

import (
	"context"
	token "shakebook/common/auth"
	accountpb "shakebook/service/account/proto/api/v1"
)

//UpdateAccountDescript account information
func (s *Server) UpdateAccountDescript(c context.Context, req *accountpb.UpdateAccountDescriptRequest) (*accountpb.Response, error) {
	res := &accountpb.Response{}
	if err := req.Validate(); err != nil {
		res.Message = err.Error()
		return res, nil
	}
	a, err := token.GetAccountIDFromContext(c)
	if err != nil {
		res.Message = err.Error()
		return res, nil
	}
	err = s.Dao.UpdateAccountDescript(a.Id, req)
	if err != nil {
		res.Message = err.Error()
		return res, nil
	}
	res.Success = true
	return res, nil
}
