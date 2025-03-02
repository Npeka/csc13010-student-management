package repository

import (
	"context"

	"github.com/csc13010-student-management/internal/auth"
	"github.com/csc13010-student-management/internal/models"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auth.IAuthRepository {
	return &authRepository{db: db}
}

// CreateUser implements auth.IAuthRepository.
func (ar *authRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepository.CreateUser")
	defer span.Finish()

	if err := ar.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, errors.Wrap(err, "authRepository.CreateUser.Create")
	}

	return user, nil
}

// FindByUsername implements auth.IAuthRepository.
func (ar *authRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepository.FindByUsername")
	defer span.Finish()

	var user models.User
	err := ar.db.WithContext(ctx).Where("username = ?", username).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Không tìm thấy user, trả về nil mà không có lỗi
		}
		return nil, errors.Wrap(err, "authRepository.FindByUsername.First")
	}
	return &user, nil
}

// FindByEmail implements auth.IAuthRepository.
func (ar *authRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepository.FindByEmail")
	defer span.Finish()

	var user models.User
	err := ar.db.WithContext(ctx).Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Không tìm thấy user, trả về nil mà không có lỗi
		}
		return nil, errors.Wrap(err, "authRepository.FindByEmail.First")
	}
	return &user, nil
}

// FindRoleByName implements auth.IAuthRepository.
func (ar *authRepository) FindRoleByName(ctx context.Context, name string) (*models.Role, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepository.FindRoleByName")
	defer span.Finish()

	var role models.Role
	err := ar.db.WithContext(ctx).Where("name = ?", name).First(&role).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Không tìm thấy role, trả về nil mà không có lỗi
		}
		return nil, errors.Wrap(err, "authRepository.FindRoleByName.First")
	}
	return &role, nil
}
