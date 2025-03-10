package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/csc13010-student-management/internal/auditlog"
	"github.com/csc13010-student-management/internal/events"
	"github.com/csc13010-student-management/internal/models"
	kafkas "github.com/csc13010-student-management/pkg/kafka"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"
)

type auditlogWorker struct {
	au auditlog.IAuditLogUsecase
	lg *logger.LoggerZap
}

func NewAuditLogWorker(
	au auditlog.IAuditLogUsecase,
	lg *logger.LoggerZap,
) auditlog.IAuditLogWorker {
	return &auditlogWorker{
		au: au,
		lg: lg,
	}
}

func (aw *auditlogWorker) Start(kurl string) {
	krs := []kafkas.KafkaReader{
		{
			Topic:   "dbserver1.public.students",
			GroupID: "auditlog-service",
			Handler: aw.HandleTableChangedEvent,
			MaxIns:  1,
		},
	}
	kafkas.StartKafkaConsumers(kurl, krs)
}
func (aw *auditlogWorker) HandleTableChangedEvent(ctx context.Context, msg kafka.Message) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "auditlogWorker.HandleTableChangedEvent")
	defer span.Finish()

	var data events.DebeziumEvent
	data.Payload.Before = make(map[string]interface{})
	data.Payload.After = make(map[string]interface{})
	if err := json.Unmarshal(msg.Value, &data); err != nil {
		return err
	}

	changedFields := getChangedFields(data.Payload.Before, data.Payload.After)

	action := getAction(data.Payload.Op)

	auditlog := createAuditLog(data, changedFields, action)

	if err := aw.au.CreateAuditLog(ctx, auditlog); err != nil {
		return err
	}

	return nil
}

func getChangedFields(before, after map[string]interface{}) map[string]interface{} {
	changedFields := make(map[string]interface{})
	for field, beforeVal := range before {
		afterVal, exists := after[field]
		if !exists || !reflect.DeepEqual(beforeVal, afterVal) {
			changedFields[field] = map[string]interface{}{
				"before": beforeVal,
				"after":  afterVal,
			}
		}
	}

	for field, afterVal := range after {
		if _, exists := before[field]; !exists {
			changedFields[field] = map[string]interface{}{
				"before": nil,
				"after":  afterVal,
			}
		}
	}

	return changedFields
}

func getAction(op interface{}) models.Action {
	switch op {
	case "c", "r":
		return models.ActionCreate
	case "u":
		return models.ActionUpdate
	case "d":
		return models.ActionDelete
	default:
		return models.ActionCreate
	}
}

func createAuditLog(data events.DebeziumEvent, changedFields map[string]interface{}, action models.Action) *models.AuditLog {
	changedFieldsJSON, _ := json.Marshal(changedFields)
	oldRecordJSON, _ := json.Marshal(data.Payload.Before)
	newRecordJSON, _ := json.Marshal(data.Payload.After)

	var recordID uint
	if action == models.ActionDelete {
		if id, ok := data.Payload.Before["id"].(float64); ok {
			recordID = uint(id)
		} else {
			fmt.Println("ERROR: Missing ID in Before for DELETE operation")
		}
	} else {
		if id, ok := data.Payload.After["id"].(float64); ok {
			recordID = uint(id)
		} else {
			fmt.Println("ERROR: Missing ID in After for CREATE/UPDATE operation")
		}
	}

	return &models.AuditLog{
		TableName:    data.Payload.Source.Table,
		RecordID:     recordID,
		Action:       action,
		OldRecord:    string(oldRecordJSON),
		NewRecord:    string(newRecordJSON),
		FieldChanges: string(changedFieldsJSON),
		LSN:          data.Payload.Source.Lsn,
		Transaction:  data.Payload.Transaction,
	}
}
