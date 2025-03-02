package usecase

import (
	"context"
	"encoding/json"

	"github.com/casbin/casbin/v2"
	"github.com/csc13010-student-management/internal/auth"
	"github.com/csc13010-student-management/internal/auth/dtos"
	"github.com/csc13010-student-management/internal/events"
	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/csc13010-student-management/pkg/utils/crypto"
	"github.com/csc13010-student-management/pkg/utils/jwt"
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

// Register implements auth.IAuthUsecase.
func (au *authUsecase) Register(ctx context.Context, registerReq *dtos.UserRegisterRequestDTO) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authUsecase.Register")
	defer span.Finish()

	// check if user already exists
	userExists, err := au.ar.FindByUsername(ctx, registerReq.Username)
	if err != nil {
		return err
	}
	if userExists != nil {
		return errors.New("authUsecase.Register: user already exists")
	}

	// check if role exists
	roleExists, err := au.ar.FindRoleByName(ctx, registerReq.Role)
	if err != nil {
		return errors.Wrap(err, "authUsecase.Register.FindRoleByName")
	}
	if roleExists == nil {
		return errors.New("authUsecase.Register: role not found")
	}

	// create user
	hashedPassword := crypto.GetHash(registerReq.Password)
	user := &models.User{
		Username: registerReq.Username,
		Password: hashedPassword,
		RoleId:   roleExists.ID,
	}
	userCreated, err := au.ar.CreateUser(ctx, user)
	if err != nil {
		return errors.New("authUsecase.Register: error creating user")
	}
	if userCreated == nil {
		return errors.New("authUsecase.Register: user not created")
	}

	// add role for user
	_, err = au.e.AddRoleForUser(user.Username, models.RoleUser)
	if err != nil {
		return errors.Wrap(err, "authUsecase.Register.AddRoleForUser")
	}

	// publish user created event
	userCreatedEventJSON, err := json.Marshal(events.UserCreatedEvent{
		UserID:   userCreated.ID,
		Username: userCreated.Username,
	})
	if err != nil {
		return errors.Wrap(err, "authUsecase.Register.MarshalUserCreatedEvent")
	}
	if err := au.kw.WriteMessages(ctx, kafka.Message{Value: userCreatedEventJSON}); err != nil {
		return errors.Wrap(err, "authUsecase.Register.WriteMessages")
	}

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
