package impl

import (
	accountpb "shakebook/service/account/proto/api/v1"
	authpb "shakebook/service/auth/proto/api/v1"
	managerpb "shakebook/service/manager/proto/api/v1"

	"go.uber.org/zap"
)

// Server defined
type Server struct {
	Logger     *zap.Logger
	AuthClient authpb.AuthServiceClient
	Dao        Dao
}

//Dao db
type Dao interface {
	//角色
	CreateRole(*managerpb.Role) (*managerpb.Response, error)
	GetRoleList() ([]*managerpb.Role, error)
	UpdateRole(*managerpb.Role) error
	DeleteRole(*managerpb.RoleId) error
	RoleBindMenu(*managerpb.RoleBindMenuRequest) error
	AccountBindRole(*managerpb.AccountBindRoleRequest) error
	GetRoleBindMenu(*managerpb.RoleId) ([]int32, error)
	GetAccountBindRole(*managerpb.AccountId) ([]int32, error)
	GetAccountMenu(*accountpb.AccountId) ([]*managerpb.MenuInfo, error)

	//菜单
	CreateMenu(*managerpb.CreateMenuRequest) error
	GetMenuList() ([]*managerpb.MenuInfo, error)
	DeleteMenu(*managerpb.MenuId) error
	UpdateMenu(*managerpb.UpdateMenuRequest) error
}
