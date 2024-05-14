// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/device"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/troublededuct"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// 故障扣费记录，节点故障时，需要扣费，记录到这个表里
type TroubleDeduct struct {
	config `json:"-"`
	// ID of the ent.
	// 19 位雪花 ID
	ID int64 `json:"id,string"`
	// 创建者 ID
	CreatedBy int64 `json:"created_by"`
	// 更新者 ID
	UpdatedBy int64 `json:"updated_by"`
	// 创建时刻，带时区
	CreatedAt time.Time `json:"created_at"`
	// 更新时刻，带时区
	UpdatedAt time.Time `json:"updated_at"`
	// 软删除时刻，带时区
	DeletedAt time.Time `json:"deleted_at"`
	// 设备 id
	DeviceID int64 `json:"device_id,string"`
	// 故障开始时刻
	StartedAt time.Time `json:"started_at"`
	// 故障结束时刻
	FinishedAt time.Time `json:"finished_at"`
	// 持续时长，单位：小时
	TimeOfDuration float64 `json:"time_of_duration,string"`
	// 扣费金额，单位：分
	Amount int64 `json:"amount"`
	// 状态
	Status enums.TroubleDeductStatus `json:"status"`
	// 扣费原因
	Reason string `json:"reason"`
	// 拒绝扣费原因
	RejectReason string `json:"reject_reason"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TroubleDeductQuery when eager-loading is set.
	Edges        TroubleDeductEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TroubleDeductEdges holds the relations/edges for other nodes in the graph.
type TroubleDeductEdges struct {
	// Device holds the value of the device edge.
	Device *Device `json:"device,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DeviceOrErr returns the Device value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TroubleDeductEdges) DeviceOrErr() (*Device, error) {
	if e.loadedTypes[0] {
		if e.Device == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: device.Label}
		}
		return e.Device, nil
	}
	return nil, &NotLoadedError{edge: "device"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TroubleDeduct) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case troublededuct.FieldTimeOfDuration:
			values[i] = new(sql.NullFloat64)
		case troublededuct.FieldID, troublededuct.FieldCreatedBy, troublededuct.FieldUpdatedBy, troublededuct.FieldDeviceID, troublededuct.FieldAmount:
			values[i] = new(sql.NullInt64)
		case troublededuct.FieldStatus, troublededuct.FieldReason, troublededuct.FieldRejectReason:
			values[i] = new(sql.NullString)
		case troublededuct.FieldCreatedAt, troublededuct.FieldUpdatedAt, troublededuct.FieldDeletedAt, troublededuct.FieldStartedAt, troublededuct.FieldFinishedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TroubleDeduct fields.
func (td *TroubleDeduct) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case troublededuct.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			td.ID = int64(value.Int64)
		case troublededuct.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				td.CreatedBy = value.Int64
			}
		case troublededuct.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				td.UpdatedBy = value.Int64
			}
		case troublededuct.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				td.CreatedAt = value.Time
			}
		case troublededuct.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				td.UpdatedAt = value.Time
			}
		case troublededuct.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				td.DeletedAt = value.Time
			}
		case troublededuct.FieldDeviceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field device_id", values[i])
			} else if value.Valid {
				td.DeviceID = value.Int64
			}
		case troublededuct.FieldStartedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field started_at", values[i])
			} else if value.Valid {
				td.StartedAt = value.Time
			}
		case troublededuct.FieldFinishedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field finished_at", values[i])
			} else if value.Valid {
				td.FinishedAt = value.Time
			}
		case troublededuct.FieldTimeOfDuration:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field time_of_duration", values[i])
			} else if value.Valid {
				td.TimeOfDuration = value.Float64
			}
		case troublededuct.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				td.Amount = value.Int64
			}
		case troublededuct.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				td.Status = enums.TroubleDeductStatus(value.String)
			}
		case troublededuct.FieldReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reason", values[i])
			} else if value.Valid {
				td.Reason = value.String
			}
		case troublededuct.FieldRejectReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reject_reason", values[i])
			} else if value.Valid {
				td.RejectReason = value.String
			}
		default:
			td.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TroubleDeduct.
// This includes values selected through modifiers, order, etc.
func (td *TroubleDeduct) Value(name string) (ent.Value, error) {
	return td.selectValues.Get(name)
}

// QueryDevice queries the "device" edge of the TroubleDeduct entity.
func (td *TroubleDeduct) QueryDevice() *DeviceQuery {
	return NewTroubleDeductClient(td.config).QueryDevice(td)
}

// Update returns a builder for updating this TroubleDeduct.
// Note that you need to call TroubleDeduct.Unwrap() before calling this method if this TroubleDeduct
// was returned from a transaction, and the transaction was committed or rolled back.
func (td *TroubleDeduct) Update() *TroubleDeductUpdateOne {
	return NewTroubleDeductClient(td.config).UpdateOne(td)
}

// Unwrap unwraps the TroubleDeduct entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (td *TroubleDeduct) Unwrap() *TroubleDeduct {
	_tx, ok := td.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: TroubleDeduct is not a transactional entity")
	}
	td.config.driver = _tx.drv
	return td
}

// String implements the fmt.Stringer.
func (td *TroubleDeduct) String() string {
	var builder strings.Builder
	builder.WriteString("TroubleDeduct(")
	builder.WriteString(fmt.Sprintf("id=%v, ", td.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", td.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", td.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(td.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(td.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(td.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("device_id=")
	builder.WriteString(fmt.Sprintf("%v", td.DeviceID))
	builder.WriteString(", ")
	builder.WriteString("started_at=")
	builder.WriteString(td.StartedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("finished_at=")
	builder.WriteString(td.FinishedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("time_of_duration=")
	builder.WriteString(fmt.Sprintf("%v", td.TimeOfDuration))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", td.Amount))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", td.Status))
	builder.WriteString(", ")
	builder.WriteString("reason=")
	builder.WriteString(td.Reason)
	builder.WriteString(", ")
	builder.WriteString("reject_reason=")
	builder.WriteString(td.RejectReason)
	builder.WriteByte(')')
	return builder.String()
}

// TroubleDeducts is a parsable slice of TroubleDeduct.
type TroubleDeducts []*TroubleDeduct