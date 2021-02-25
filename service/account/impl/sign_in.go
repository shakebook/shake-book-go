package impl

import (
	"context"
	accountpb "shakebook/service/account/proto/api/v1"
	"strconv"

	authpb "shakebook/service/auth/proto/api/v1"
)

//SignIn implement server
func (s *Server) SignIn(c context.Context, req *accountpb.SignInRequest) (*accountpb.SignInResponse, error) {
	res := &accountpb.SignInResponse{}
	if err := req.Validate(); err != nil {
		res.Message = err.Error()
		return res, nil
	}

	//TODO: 登录
	resp, err := s.Dao.SignIn(c, req)
	if err != nil {
		res.Message = err.Error()
		return res, nil
	}

	//生成TOKEN
	gen, err := s.AuthClient.GenToken(c, &authpb.GenTokenRequest{
		Id: strconv.Itoa(int(resp.Id)),
	})
	if err != nil {
		res.Message = err.Error()
		return res, nil
	}

	res.Success = true
	res.Data = gen.Token
	return res, nil
}
