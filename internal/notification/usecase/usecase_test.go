package usecase

import (
	"reflect"
	"testing"

	"github.com/csc13010-student-management/internal/notification"
)

func TestNewNotificationUsecase(t *testing.T) {
	type args struct {
		fr notification.INotificationRepository
	}
	tests := []struct {
		name string
		args args
		want notification.INotificationUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNotificationUsecase(tt.args.fr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNotificationUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
