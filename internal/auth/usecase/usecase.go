package usecase

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/csc13010-student-management/internal/auth"
	"github.com/csc13010-student-management/internal/auth/dtos"
	"github.com/csc13010-student-management/internal/events"
	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/csc13010-student-management/pkg/utils/crypto"
	"github.com/csc13010-student-management/pkg/utils/jwt"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)

type authUsecase struct {
	ar auth.IAuthRepository
	lg *logger.LoggerZap
	kw *kafka.Writer
	e  *casbin.Enforcer
}

func NewAuthUsecase(
	ar auth.IAuthRepository,
	lg *logger.LoggerZap,
	kw *kafka.Writer,
	e *casbin.Enforcer,
) auth.IAuthUsecase {
	return &authUsecase{
		ar: ar,
		lg: lg,
		kw: kw,
		e:  e,
	}
}

func (au *authUsecase) Register(ctx context.Context, registerReq *dtos.UserRegisterRequestDTO) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authUsecase.Register")
	defer span.Finish()

	var wg sync.WaitGroup
	// var mu sync.Mutex // Để tránh race condition khi ghi dữ liệu
	// var userExists *models.User
	var roleExists *models.Role
	var errRole error

	// Tìm user
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	userExists, errUser = au.ar.FindByUsername(ctx, registerReq.Username)
	// }()

	roleExists, errRole = au.ar.FindRoleByName(ctx, registerReq.Role)

	// wg.Wait() // Chờ cả 2 truy vấn hoàn thành

	// if errUser != nil {
	// 	return errors.Wrap(errUser, "authUsecase.Register.FindByUsername")
	// }
	// if userExists != nil {
	// 	return errors.New("authUsecase.Register: user already exists")
	// }

	if errRole != nil {
		return errors.Wrap(errRole, "authUsecase.Register.FindRoleByName")
	}
	if roleExists == nil {
		return errors.New("authUsecase.Register: role not found")
	}

	// Tạo user
	user := &models.User{
		ID:       uuid.New(),
		Username: registerReq.Username,
		Password: crypto.GetHash(registerReq.Password),
		RoleId:   roleExists.ID,
	}

	userCreated, err := au.ar.CreateUser(ctx, user)
	if err != nil {
		return errors.Wrap(err, "authUsecase.Register.CreateUser")
	}
	if userCreated == nil {
		return errors.New("authUsecase.Register: user not created")
	}

	// Thêm role cho user (có thể chạy song song với publish event)
	wg.Add(1)
	go func() {
		defer wg.Done()
		if _, err := au.e.AddRoleForUser(user.Username, models.RoleUser); err != nil {
			log.Printf("❌ Lỗi AddRoleForUser: %v", err)
		}
	}()

	// Publish event
	wg.Add(1)
	go func() {
		defer wg.Done()
		userCreatedEventJSON, err := json.Marshal(events.UserCreatedEvent{
			UserID:   userCreated.ID,
			Username: userCreated.Username,
		})
		if err != nil {
			log.Printf("❌ Lỗi MarshalUserCreatedEvent: %v", err)
			return
		}
		if err := au.kw.WriteMessages(ctx, kafka.Message{Value: userCreatedEventJSON}); err != nil {
			log.Printf("❌ Lỗi WriteMessages: %v", err)
		}
	}()

	wg.Wait()

	return nil
}

func (au *authUsecase) Login(ctx context.Context, loginReq *dtos.UserLoginRequestDTO) (*dtos.UserLoginResponseDTO, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authUsecase.Login")
	defer span.Finish()

	user, err := au.ar.FindByUsername(ctx, loginReq.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	hashedPassword := crypto.GetHash(loginReq.Password)
	if user.Password != hashedPassword {
		return nil, errors.New("invalid password")
	}

	token, err := jwt.CreateJWTToken(user.Username, jwt.JwtExpAT)
	if err != nil {
		return nil, errors.New("error creating token")
	}

	return &dtos.UserLoginResponseDTO{
		Token: token,
	}, nil
}

// Logout implements auth.IAuthUsecase.
func (au *authUsecase) Logout(ctx context.Context) {
	panic("unimplemented")
}

// Refresh implements auth.IAuthUsecase.
func (au *authUsecase) Refresh(ctx context.Context) {
	panic("unimplemented")
}
