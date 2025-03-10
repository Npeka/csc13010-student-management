package repository

import (
	"context"
	"fmt"

	"github.com/csc13010-student-management/internal/config"
	"github.com/csc13010-student-management/internal/models"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type configRepository struct {
	db *gorm.DB
}

func NewConfigRepository(
	db *gorm.DB,
) config.IConfigRepository {
	return &configRepository{
		db: db,
	}
}

func (c *configRepository) GetConfig(ctx context.Context) (*models.Config, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "configRepository.GetConfig")
	defer span.Finish()

	var cf models.Config
	err := c.db.WithContext(ctx).First(&cf).Error
	if err != nil {
		return nil, errors.Wrap(err, "configRepository.GetConfig")
	}

	return &cf, nil
}

func (c *configRepository) UpdateConfig(ctx context.Context, cf *models.Config) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "configRepository.UpdateConfig")
	defer span.Finish()

	fmt.Println("configRepository.UpdateConfig", cf)

	if err := c.db.WithContext(ctx).Model(&models.Config{}).Where("id = ?", 1).
		Updates(map[string]interface{}{
			"email_domain":   cf.EmailDomain,
			"validate_phone": cf.ValidatePhone,
			"status_rules":   cf.StatusRules,
			"delete_limit":   cf.DeleteLimit,
		}).Error; err != nil {
		return errors.Wrap(err, "configRepository.UpdateConfig")
	}

	return nil
}
