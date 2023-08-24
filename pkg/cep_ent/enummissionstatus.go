// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/enummissionstatus"
)

// EnumMissionStatus is the model entity for the EnumMissionStatus schema.
type EnumMissionStatus struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy int64 `json:"created_by"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy int64 `json:"updated_by"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at"`
	// 给到前端的任务状态
	FrontStatus string `json:"front_status"`
	// 任务类型
	MissionType string `json:"mission_type"`
	// 任务状态
	MissionStatus string `json:"mission_status"`
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EnumMissionStatus) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case enummissionstatus.FieldID, enummissionstatus.FieldCreatedBy, enummissionstatus.FieldUpdatedBy:
			values[i] = new(sql.NullInt64)
		case enummissionstatus.FieldFrontStatus, enummissionstatus.FieldMissionType, enummissionstatus.FieldMissionStatus:
			values[i] = new(sql.NullString)
		case enummissionstatus.FieldCreatedAt, enummissionstatus.FieldUpdatedAt, enummissionstatus.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EnumMissionStatus fields.
func (ems *EnumMissionStatus) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case enummissionstatus.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ems.ID = int64(value.Int64)
		case enummissionstatus.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ems.CreatedBy = value.Int64
			}
		case enummissionstatus.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ems.UpdatedBy = value.Int64
			}
		case enummissionstatus.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ems.CreatedAt = value.Time
			}
		case enummissionstatus.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ems.UpdatedAt = value.Time
			}
		case enummissionstatus.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ems.DeletedAt = value.Time
			}
		case enummissionstatus.FieldFrontStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field front_status", values[i])
			} else if value.Valid {
				ems.FrontStatus = value.String
			}
		case enummissionstatus.FieldMissionType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mission_type", values[i])
			} else if value.Valid {
				ems.MissionType = value.String
			}
		case enummissionstatus.FieldMissionStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mission_status", values[i])
			} else if value.Valid {
				ems.MissionStatus = value.String
			}
		default:
			ems.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EnumMissionStatus.
// This includes values selected through modifiers, order, etc.
func (ems *EnumMissionStatus) Value(name string) (ent.Value, error) {
	return ems.selectValues.Get(name)
}

// Update returns a builder for updating this EnumMissionStatus.
// Note that you need to call EnumMissionStatus.Unwrap() before calling this method if this EnumMissionStatus
// was returned from a transaction, and the transaction was committed or rolled back.
func (ems *EnumMissionStatus) Update() *EnumMissionStatusUpdateOne {
	return NewEnumMissionStatusClient(ems.config).UpdateOne(ems)
}

// Unwrap unwraps the EnumMissionStatus entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ems *EnumMissionStatus) Unwrap() *EnumMissionStatus {
	_tx, ok := ems.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: EnumMissionStatus is not a transactional entity")
	}
	ems.config.driver = _tx.drv
	return ems
}

// String implements the fmt.Stringer.
func (ems *EnumMissionStatus) String() string {
	var builder strings.Builder
	builder.WriteString("EnumMissionStatus(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ems.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", ems.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", ems.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ems.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ems.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ems.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("front_status=")
	builder.WriteString(ems.FrontStatus)
	builder.WriteString(", ")
	builder.WriteString("mission_type=")
	builder.WriteString(ems.MissionType)
	builder.WriteString(", ")
	builder.WriteString("mission_status=")
	builder.WriteString(ems.MissionStatus)
	builder.WriteByte(')')
	return builder.String()
}

// EnumMissionStatusSlice is a parsable slice of EnumMissionStatus.
type EnumMissionStatusSlice []*EnumMissionStatus
