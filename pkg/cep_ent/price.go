// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/price"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// 任务定价表，表里有数据，任务才有单价，才可以被创建
type Price struct {
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
	// 显卡型号
	GpuVersion enums.GpuVersion `json:"gpu_version"`
	// 任务类型
	MissionType enums.MissionType `json:"mission_type"`
	// 任务大类
	MissionCategory enums.MissionCategory `json:"mission_category"`
	// 任务计费类型
	MissionBillingType enums.MissionBillingType `json:"mission_billing_type"`
	// 包时类型，只有包时任务才有
	RenewalType enums.RenewalType `json:"renewal_type,omitempty" renewal_type`
	// 任务单价
	Cep int64 `json:"cep"`
	// 价格有效时间开始，为空表示永久有效
	StartedAt *time.Time `json:"started_at"`
	// 价格有效时间结束，为空表示永久有效
	FinishedAt *time.Time `json:"finished_at"`
	// 价格是否屏蔽，前端置灰，硬选也可以
	IsDeprecated bool `json:"is_deprecated"`
	// 价格是否敏感，用于特殊类型任务，不能让外部看到选项
	IsSensitive  bool `json:"is_sensitive"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Price) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case price.FieldIsDeprecated, price.FieldIsSensitive:
			values[i] = new(sql.NullBool)
		case price.FieldID, price.FieldCreatedBy, price.FieldUpdatedBy, price.FieldCep:
			values[i] = new(sql.NullInt64)
		case price.FieldGpuVersion, price.FieldMissionType, price.FieldMissionCategory, price.FieldMissionBillingType, price.FieldRenewalType:
			values[i] = new(sql.NullString)
		case price.FieldCreatedAt, price.FieldUpdatedAt, price.FieldDeletedAt, price.FieldStartedAt, price.FieldFinishedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Price fields.
func (pr *Price) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case price.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int64(value.Int64)
		case price.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pr.CreatedBy = value.Int64
			}
		case price.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pr.UpdatedBy = value.Int64
			}
		case price.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case price.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case price.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pr.DeletedAt = value.Time
			}
		case price.FieldGpuVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gpu_version", values[i])
			} else if value.Valid {
				pr.GpuVersion = enums.GpuVersion(value.String)
			}
		case price.FieldMissionType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mission_type", values[i])
			} else if value.Valid {
				pr.MissionType = enums.MissionType(value.String)
			}
		case price.FieldMissionCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mission_category", values[i])
			} else if value.Valid {
				pr.MissionCategory = enums.MissionCategory(value.String)
			}
		case price.FieldMissionBillingType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mission_billing_type", values[i])
			} else if value.Valid {
				pr.MissionBillingType = enums.MissionBillingType(value.String)
			}
		case price.FieldRenewalType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field renewal_type", values[i])
			} else if value.Valid {
				pr.RenewalType = enums.RenewalType(value.String)
			}
		case price.FieldCep:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cep", values[i])
			} else if value.Valid {
				pr.Cep = value.Int64
			}
		case price.FieldStartedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field started_at", values[i])
			} else if value.Valid {
				pr.StartedAt = new(time.Time)
				*pr.StartedAt = value.Time
			}
		case price.FieldFinishedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field finished_at", values[i])
			} else if value.Valid {
				pr.FinishedAt = new(time.Time)
				*pr.FinishedAt = value.Time
			}
		case price.FieldIsDeprecated:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_deprecated", values[i])
			} else if value.Valid {
				pr.IsDeprecated = value.Bool
			}
		case price.FieldIsSensitive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_sensitive", values[i])
			} else if value.Valid {
				pr.IsSensitive = value.Bool
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Price.
// This includes values selected through modifiers, order, etc.
func (pr *Price) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// Update returns a builder for updating this Price.
// Note that you need to call Price.Unwrap() before calling this method if this Price
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Price) Update() *PriceUpdateOne {
	return NewPriceClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Price entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Price) Unwrap() *Price {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: Price is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Price) String() string {
	var builder strings.Builder
	builder.WriteString("Price(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", pr.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", pr.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(pr.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("gpu_version=")
	builder.WriteString(fmt.Sprintf("%v", pr.GpuVersion))
	builder.WriteString(", ")
	builder.WriteString("mission_type=")
	builder.WriteString(fmt.Sprintf("%v", pr.MissionType))
	builder.WriteString(", ")
	builder.WriteString("mission_category=")
	builder.WriteString(fmt.Sprintf("%v", pr.MissionCategory))
	builder.WriteString(", ")
	builder.WriteString("mission_billing_type=")
	builder.WriteString(fmt.Sprintf("%v", pr.MissionBillingType))
	builder.WriteString(", ")
	builder.WriteString("renewal_type=")
	builder.WriteString(fmt.Sprintf("%v", pr.RenewalType))
	builder.WriteString(", ")
	builder.WriteString("cep=")
	builder.WriteString(fmt.Sprintf("%v", pr.Cep))
	builder.WriteString(", ")
	if v := pr.StartedAt; v != nil {
		builder.WriteString("started_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := pr.FinishedAt; v != nil {
		builder.WriteString("finished_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("is_deprecated=")
	builder.WriteString(fmt.Sprintf("%v", pr.IsDeprecated))
	builder.WriteString(", ")
	builder.WriteString("is_sensitive=")
	builder.WriteString(fmt.Sprintf("%v", pr.IsSensitive))
	builder.WriteByte(')')
	return builder.String()
}

// Prices is a parsable slice of Price.
type Prices []*Price
