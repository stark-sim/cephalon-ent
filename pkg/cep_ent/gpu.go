// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/gpu"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// 显卡信息，在表里有显卡型号的，才能在系统中选择使用等
type Gpu struct {
	config `json:"-"`
	// ID of the ent.
	// 19 位雪花 ID
	ID int64 `json:"id,string"`
	// 创建者 ID
	CreatedBy int64 `json:"created_by,string"`
	// 更新者 ID
	UpdatedBy int64 `json:"updated_by,string"`
	// 创建时刻，带时区
	CreatedAt time.Time `json:"created_at"`
	// 更新时刻，带时区
	UpdatedAt time.Time `json:"updated_at"`
	// 软删除时刻，带时区
	DeletedAt time.Time `json:"deleted_at"`
	// 显卡型号
	Version enums.GpuVersion `json:"version"`
	// 显卡能力值
	Power int `json:"power"`
	// 显存
	VideoMemory int `json:"video_memory"`
	// CPU
	CPU int `json:"cpu"`
	// 内存
	Memory int `json:"memory"`
	// 保底最低月收益
	LowestEarnMonth int64 `json:"lowest_earn_month"`
	// 保底最高月收益
	HighestEarnMonth int64 `json:"highest_earn_month"`
	// 故障扣费金额，单位：厘/小时
	TroubleDeductAmount int64 `json:"trouble_deduct_amount"`
	// 提现保留最低金额（押金），单位：厘
	WithdrawRetainAmount int64 `json:"withdraw_retain_amount"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GpuQuery when eager-loading is set.
	Edges        GpuEdges `json:"edges"`
	selectValues sql.SelectValues
}

// GpuEdges holds the relations/edges for other nodes in the graph.
type GpuEdges struct {
	// DeviceGpuMissions holds the value of the device_gpu_missions edge.
	DeviceGpuMissions []*DeviceGpuMission `json:"device_gpu_missions,omitempty"`
	// Prices holds the value of the prices edge.
	Prices []*Price `json:"prices,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// DeviceGpuMissionsOrErr returns the DeviceGpuMissions value or an error if the edge
// was not loaded in eager-loading.
func (e GpuEdges) DeviceGpuMissionsOrErr() ([]*DeviceGpuMission, error) {
	if e.loadedTypes[0] {
		return e.DeviceGpuMissions, nil
	}
	return nil, &NotLoadedError{edge: "device_gpu_missions"}
}

// PricesOrErr returns the Prices value or an error if the edge
// was not loaded in eager-loading.
func (e GpuEdges) PricesOrErr() ([]*Price, error) {
	if e.loadedTypes[1] {
		return e.Prices, nil
	}
	return nil, &NotLoadedError{edge: "prices"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Gpu) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case gpu.FieldID, gpu.FieldCreatedBy, gpu.FieldUpdatedBy, gpu.FieldPower, gpu.FieldVideoMemory, gpu.FieldCPU, gpu.FieldMemory, gpu.FieldLowestEarnMonth, gpu.FieldHighestEarnMonth, gpu.FieldTroubleDeductAmount, gpu.FieldWithdrawRetainAmount:
			values[i] = new(sql.NullInt64)
		case gpu.FieldVersion:
			values[i] = new(sql.NullString)
		case gpu.FieldCreatedAt, gpu.FieldUpdatedAt, gpu.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Gpu fields.
func (gp *Gpu) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case gpu.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gp.ID = int64(value.Int64)
		case gpu.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				gp.CreatedBy = value.Int64
			}
		case gpu.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				gp.UpdatedBy = value.Int64
			}
		case gpu.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gp.CreatedAt = value.Time
			}
		case gpu.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gp.UpdatedAt = value.Time
			}
		case gpu.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				gp.DeletedAt = value.Time
			}
		case gpu.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				gp.Version = enums.GpuVersion(value.String)
			}
		case gpu.FieldPower:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field power", values[i])
			} else if value.Valid {
				gp.Power = int(value.Int64)
			}
		case gpu.FieldVideoMemory:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field video_memory", values[i])
			} else if value.Valid {
				gp.VideoMemory = int(value.Int64)
			}
		case gpu.FieldCPU:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cpu", values[i])
			} else if value.Valid {
				gp.CPU = int(value.Int64)
			}
		case gpu.FieldMemory:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field memory", values[i])
			} else if value.Valid {
				gp.Memory = int(value.Int64)
			}
		case gpu.FieldLowestEarnMonth:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field lowest_earn_month", values[i])
			} else if value.Valid {
				gp.LowestEarnMonth = value.Int64
			}
		case gpu.FieldHighestEarnMonth:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field highest_earn_month", values[i])
			} else if value.Valid {
				gp.HighestEarnMonth = value.Int64
			}
		case gpu.FieldTroubleDeductAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field trouble_deduct_amount", values[i])
			} else if value.Valid {
				gp.TroubleDeductAmount = value.Int64
			}
		case gpu.FieldWithdrawRetainAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field withdraw_retain_amount", values[i])
			} else if value.Valid {
				gp.WithdrawRetainAmount = value.Int64
			}
		default:
			gp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Gpu.
// This includes values selected through modifiers, order, etc.
func (gp *Gpu) Value(name string) (ent.Value, error) {
	return gp.selectValues.Get(name)
}

// QueryDeviceGpuMissions queries the "device_gpu_missions" edge of the Gpu entity.
func (gp *Gpu) QueryDeviceGpuMissions() *DeviceGpuMissionQuery {
	return NewGpuClient(gp.config).QueryDeviceGpuMissions(gp)
}

// QueryPrices queries the "prices" edge of the Gpu entity.
func (gp *Gpu) QueryPrices() *PriceQuery {
	return NewGpuClient(gp.config).QueryPrices(gp)
}

// Update returns a builder for updating this Gpu.
// Note that you need to call Gpu.Unwrap() before calling this method if this Gpu
// was returned from a transaction, and the transaction was committed or rolled back.
func (gp *Gpu) Update() *GpuUpdateOne {
	return NewGpuClient(gp.config).UpdateOne(gp)
}

// Unwrap unwraps the Gpu entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gp *Gpu) Unwrap() *Gpu {
	_tx, ok := gp.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: Gpu is not a transactional entity")
	}
	gp.config.driver = _tx.drv
	return gp
}

// String implements the fmt.Stringer.
func (gp *Gpu) String() string {
	var builder strings.Builder
	builder.WriteString("Gpu(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gp.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", gp.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", gp.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(gp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(gp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(gp.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", gp.Version))
	builder.WriteString(", ")
	builder.WriteString("power=")
	builder.WriteString(fmt.Sprintf("%v", gp.Power))
	builder.WriteString(", ")
	builder.WriteString("video_memory=")
	builder.WriteString(fmt.Sprintf("%v", gp.VideoMemory))
	builder.WriteString(", ")
	builder.WriteString("cpu=")
	builder.WriteString(fmt.Sprintf("%v", gp.CPU))
	builder.WriteString(", ")
	builder.WriteString("memory=")
	builder.WriteString(fmt.Sprintf("%v", gp.Memory))
	builder.WriteString(", ")
	builder.WriteString("lowest_earn_month=")
	builder.WriteString(fmt.Sprintf("%v", gp.LowestEarnMonth))
	builder.WriteString(", ")
	builder.WriteString("highest_earn_month=")
	builder.WriteString(fmt.Sprintf("%v", gp.HighestEarnMonth))
	builder.WriteString(", ")
	builder.WriteString("trouble_deduct_amount=")
	builder.WriteString(fmt.Sprintf("%v", gp.TroubleDeductAmount))
	builder.WriteString(", ")
	builder.WriteString("withdraw_retain_amount=")
	builder.WriteString(fmt.Sprintf("%v", gp.WithdrawRetainAmount))
	builder.WriteByte(')')
	return builder.String()
}

// Gpus is a parsable slice of Gpu.
type Gpus []*Gpu
