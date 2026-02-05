package auth

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	// userRepo repository.UserRepository
	ts TokenService
}

func NewLoginService(ts TokenService) *LoginService {
	return &LoginService{
		// userRepo: userRepo,
		ts: ts,
	}
}

func (service *LoginService) Login(ctx context.Context, email, password string) (*TokenPairs, error) {
	// user, err := service.userRepo.FindByEmail(ctx, email)

	// if err != nil {
	// return nil, err
	// }

	// if !checkPasswordHash(password, user.Password) {
	// return nil, errors.New("invalid credentials")
	// }

	// todo fix it and put user info
	pairs, pairErr := service.ts.GeneratePairs(123)

	if pairErr != nil {
		return nil, pairErr
	}

	return pairs, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
