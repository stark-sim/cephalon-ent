// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/outputlog"
)

// OutputLog is the model entity for the OutputLog schema.
type OutputLog struct {
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
	// 请求追踪 id
	TraceID int64 `json:"trace_id"`
	// 请求头
	Headers string `json:"headers"`
	// 请求体
	Body string `json:"body"`
	// 请求地址
	URL string `json:"url"`
	// 客户端 IP
	IP string `json:"ip"`
	// 调用方
	Caller string `json:"caller"`
	// 状态码
	Status int16 `json:"status"`
	// 记录调用者的密钥对
	HmacKey      string `json:"hmac_key"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OutputLog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case outputlog.FieldID, outputlog.FieldCreatedBy, outputlog.FieldUpdatedBy, outputlog.FieldTraceID, outputlog.FieldStatus:
			values[i] = new(sql.NullInt64)
		case outputlog.FieldHeaders, outputlog.FieldBody, outputlog.FieldURL, outputlog.FieldIP, outputlog.FieldCaller, outputlog.FieldHmacKey:
			values[i] = new(sql.NullString)
		case outputlog.FieldCreatedAt, outputlog.FieldUpdatedAt, outputlog.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OutputLog fields.
func (ol *OutputLog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case outputlog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ol.ID = int64(value.Int64)
		case outputlog.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ol.CreatedBy = value.Int64
			}
		case outputlog.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ol.UpdatedBy = value.Int64
			}
		case outputlog.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ol.CreatedAt = value.Time
			}
		case outputlog.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ol.UpdatedAt = value.Time
			}
		case outputlog.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ol.DeletedAt = value.Time
			}
		case outputlog.FieldTraceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field trace_id", values[i])
			} else if value.Valid {
				ol.TraceID = value.Int64
			}
		case outputlog.FieldHeaders:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field headers", values[i])
			} else if value.Valid {
				ol.Headers = value.String
			}
		case outputlog.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				ol.Body = value.String
			}
		case outputlog.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				ol.URL = value.String
			}
		case outputlog.FieldIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip", values[i])
			} else if value.Valid {
				ol.IP = value.String
			}
		case outputlog.FieldCaller:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field caller", values[i])
			} else if value.Valid {
				ol.Caller = value.String
			}
		case outputlog.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ol.Status = int16(value.Int64)
			}
		case outputlog.FieldHmacKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hmac_key", values[i])
			} else if value.Valid {
				ol.HmacKey = value.String
			}
		default:
			ol.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OutputLog.
// This includes values selected through modifiers, order, etc.
func (ol *OutputLog) Value(name string) (ent.Value, error) {
	return ol.selectValues.Get(name)
}

// Update returns a builder for updating this OutputLog.
// Note that you need to call OutputLog.Unwrap() before calling this method if this OutputLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (ol *OutputLog) Update() *OutputLogUpdateOne {
	return NewOutputLogClient(ol.config).UpdateOne(ol)
}

// Unwrap unwraps the OutputLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ol *OutputLog) Unwrap() *OutputLog {
	_tx, ok := ol.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: OutputLog is not a transactional entity")
	}
	ol.config.driver = _tx.drv
	return ol
}

// String implements the fmt.Stringer.
func (ol *OutputLog) String() string {
	var builder strings.Builder
	builder.WriteString("OutputLog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ol.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", ol.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", ol.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ol.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ol.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ol.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("trace_id=")
	builder.WriteString(fmt.Sprintf("%v", ol.TraceID))
	builder.WriteString(", ")
	builder.WriteString("headers=")
	builder.WriteString(ol.Headers)
	builder.WriteString(", ")
	builder.WriteString("body=")
	builder.WriteString(ol.Body)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(ol.URL)
	builder.WriteString(", ")
	builder.WriteString("ip=")
	builder.WriteString(ol.IP)
	builder.WriteString(", ")
	builder.WriteString("caller=")
	builder.WriteString(ol.Caller)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", ol.Status))
	builder.WriteString(", ")
	builder.WriteString("hmac_key=")
	builder.WriteString(ol.HmacKey)
	builder.WriteByte(')')
	return builder.String()
}

// OutputLogs is a parsable slice of OutputLog.
type OutputLogs []*OutputLog
