package http

import (
	"reflect"
	"testing"

	"github.com/csc13010-student-management/internal/notification"
)

func TestNewNotificationHandlers(t *testing.T) {
	tests := []struct {
		name string
		want notification.INotificationHandlers
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNotificationHandlers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNotificationHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}
