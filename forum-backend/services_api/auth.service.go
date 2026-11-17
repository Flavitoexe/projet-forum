package services_api

import (
	"bluedit/forum-backend/auth"
	"bluedit/forum-backend/dto"
	"bluedit/forum-backend/models"
	"bluedit/forum-backend/repositories"
	"errors"
	"fmt"
	"strconv"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func InitAuthService(repo *repositories.UserRepository) *AuthService {
	return &AuthService{
		userRepo: repo,
	}
}

func (s *AuthService) Login(data dto.LoginRequestDto) (*dto.LoginResponseDto, error) {
	user, err := s.userRepo.FindByNameOrEmail(data.Username)
	if err != nil {
		fmt.Println("err auth.service, identifants pas trouvés/invalides")
		return nil, errors.New("Erreur identifiant")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		fmt.Println("err auth.service, mdp", err)
		return nil, errors.New("Erreur mdp")
	}

	userIDStr := strconv.Itoa(user.IdUser)

	token, err := auth.GenerateToken(userIDStr, user.Role)
	if err != nil {
		return nil, fmt.Errorf("Erreur génération token - Erreur\n\t%s", err.Error())
	}

	return &dto.LoginResponseDto{
		IdUser:      user.IdUser,
		Type:        "Bearer",
		AccessToken: token,
		ExpiresIn:   300,
	}, nil
}

func (s *AuthService) Register(data dto.RegisterRequestDto) (int, error) {
	if len(data.Password) < 12 {
		return -1, errors.New("Le mot de passe doit faire au moins 12 caractères")
	}

	var hasUpper, hasSpecial bool
	for _, char := range data.Password {
		if unicode.IsUpper(char) {
			hasUpper = true
		}
		if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			hasSpecial = true
		}
	}

	if !hasUpper || !hasSpecial {
		return -1, errors.New("Le mot de passe nécessite au moins une majuscule et un caractère spécial")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, errors.New("Erreur du mot de passe")
	}

	newUser := models.User{
		Name:     data.Username,
		Email:    data.Email,
		Password: string(hashedPassword),
	}

	id, err := s.userRepo.Create(newUser)
	if err != nil {
		return -1, fmt.Errorf("Erreur lors de la création du compte : %v", err)
	}

	return id, nil
}
