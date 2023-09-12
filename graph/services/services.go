package services

import "github.com/volatiletech/sqlboiler/v4/boil"

type services struct {
	userService UserService
}

type Services interface {
	UserService() UserService
}

// TODO: ここの処理理解できてない
func (s *services) UserService() UserService {
	return s.userService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService: &userService{exec: exec},
	}
}
