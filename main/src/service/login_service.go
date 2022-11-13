package service

type LoginService interface {
	IsUserValid(email string, password string) bool
}

type loginInformation struct {
	email    string
	password string
}

func StaticLoginService() LoginService {
	return &loginInformation{
		email:    "ali.akber@brainstation-23.com",
		password: "Pass.1234#",
	}
}

func (info *loginInformation) IsUserValid(email string, password string) bool {
	return info.email == email && info.password == password
}
