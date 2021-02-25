package impl

import (
	"context"
	accountpb "shakebook/service/account/proto/api/v1"

	"golang.org/x/crypto/bcrypt"
)

//SignUp implement server
func (s *Server) SignUp(c context.Context, req *accountpb.SignUpRequest) (*accountpb.Response, error) {
	res := &accountpb.Response{}
	s.Logger.Sugar().Infof("signup request:%v", req)
	if err := req.Validate(); err != nil {
		res.Message = err.Error()
		return res, nil
	}

	if err := s.Dao.ValidEmail(c, req.AccountEmail,
		req.EmailValidCode); err != nil {
		res.Message = err.Error()
		return res, nil
	}

	enpassword, _ := bcrypt.GenerateFromPassword([]byte(req.AccountPassword), bcrypt.DefaultCost)
	req.AccountPassword = string(enpassword)

	if err := s.Dao.SignUp(req); err != nil {
		res.Message = err.Error()
		return res, nil
	}
	res.Success = true
	return res, nil

}
