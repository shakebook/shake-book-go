package impl

import (
	"context"
	"net/smtp"
	"shakebook/common/tools"
	accountpb "shakebook/service/account/proto/api/v1"

	"github.com/jordan-wright/email"
)

//ValidEmailer defined valid email interface
type ValidEmailer interface {
	sedCodeToEmail(sendCode, to, bcc, cc string) error
	genRandCode() (string, error)
}

//ValidEmail valid email
func (s *Server) ValidEmail(c context.Context, req *accountpb.ValidEmailRequest) (*accountpb.Response, error) {
	res := &accountpb.Response{}
	randCode, err := s.ValidEmailer.genRandCode()
	if err != nil {
		res.Message = err.Error()
		return res, nil
	}

	if err := s.ValidEmailer.sedCodeToEmail(randCode,
		req.AccountEmail, req.AccountEmail,
		req.AccountEmail); err != nil {
		res.Message = err.Error()
		return res, nil
	}

	if err := s.Dao.WriteEmailCodeToRedis(c, req.AccountEmail, randCode); err != nil {
		res.Message = err.Error()
		return res, nil
	}

	res.Success = true
	return res, nil
}

//ValidEmail defied email params
type ValidEmail struct {
	From          string
	To            string
	Bcc           string
	Cc            string
	Subject       string
	Text          string
	Addr          string
	Identity      string
	Username      string
	Password      string
	Host          string
	OptionCharts  string
	RandCodeLenth int
}

func (v *ValidEmail) sedCodeToEmail(sendCode, to, bcc, cc string) error {
	e := email.NewEmail()
	e.From = v.Subject + " <" + v.From + ">"
	e.To = []string{to}
	e.Bcc = []string{bcc}
	e.Cc = []string{cc}
	e.Subject = v.Subject
	e.Text = []byte(sendCode)
	e.HTML = []byte(`<p>http://localhost:3000/home</p> </br>` + `<strong>` + sendCode + `</strong>`)
	return e.Send(v.Addr, smtp.PlainAuth(v.Identity, v.Username, v.Password, v.Host))
}

func (v *ValidEmail) genRandCode() (string, error) {
	return tools.GenerateRandCode(v.RandCodeLenth, v.OptionCharts)
}
