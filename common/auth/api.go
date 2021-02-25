package auth

var apiMap = map[string]bool{
	"/shakebook.AccountService/SignIn":     true,
	"/shakebook.AccountService/SignUp":     true,
	"/shakebook.AccountService/ValidEmail": true,
	// "/shakebook.AccountService/SignOut":       true,
	// "/shakebook.AccountService/UpdateAccount": true,
	// "/shakebook.AccountService/GetAccount":    true,
	// "/shakebook.ManagerService/CreateRole":    true,
	// "/shakebook.ManagerService/CreateRole":    true,
	// "/shakebook.ManagerService/CreateRole":    true,
}

//IsAuth is auth
func IsAuth(path string) bool {
	return !apiMap[path]
}
