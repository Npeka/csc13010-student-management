package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/notification"
	"gorm.io/gorm"
)

func TestNewNotificationRepository(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want notification.INotificationRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNotificationRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNotificationRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_notificationRepository_GetStatuses(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx      context.Context
		statuses []uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Status
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &notificationRepository{
				db: tt.fields.db,
			}
			got, err := r.GetStatuses(tt.args.ctx, tt.args.statuses)
			if (err != nil) != tt.wantErr {
				t.Errorf("notificationRepository.GetStatuses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("notificationRepository.GetStatuses() = %v, want %v", got, tt.want)
			}
		})
	}
}
