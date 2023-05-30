package service

type AuthService struct {
}

func (as *AuthService) LoginByUsername(username, password string) (string, string, error) {
	//TODO 查询数据库是否存在这个用户和密码
	//db_account := &model.Account{}
	return "s", "s", nil
}
