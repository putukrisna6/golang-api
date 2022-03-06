package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/putukrisna6/golang-api/dto"
	"github.com/putukrisna6/golang-api/entity"
	"github.com/putukrisna6/golang-api/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredentials(email, password)
	if v, ok := res.(entity.User); ok {
		matchPassword := matchPassword(v.Password, []byte(password))
		if v.Email == email && matchPassword {
			return res
		}
	}
	return false
}

func (service *authService) CreateUser(user dto.RegisterDTO) entity.User {
	newUser := entity.User{}
	err := smapping.FillStruct(&newUser, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("map failed %v", err)
	}
	res := service.userRepository.InsertUser(newUser)
	return res
}

func (service *authService) FindByEmail(email string) entity.User {
	return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func matchPassword(hashedPassword string, plainPassword []byte) bool {
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
