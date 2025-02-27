package repository

import (
	"reflect"
	"testing"

	"github.com/csc13010-student-management/internal/notification"
)

func TestNewNotificationRepository(t *testing.T) {
	tests := []struct {
		name string
		want notification.INotificationRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNotificationRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNotificationRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
