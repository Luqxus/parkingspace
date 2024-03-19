package service


type AuthService interface {}

type authService struct {}


func NewAuthService() *authService {
	return &authService{}
}


func (s *authService) CreateUser() {
//	TODO: create a new user
	
}

func (s *authService) LoginUser () {
//	TODO: login user logic
}


func (s *authService) DeleteUser() {
//	TODO: delete user account logic
}

func (s *authService) UpdateUser() {
//	TODO: update user logic
}

