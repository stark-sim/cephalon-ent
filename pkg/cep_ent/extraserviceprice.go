// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/extraservice"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/extraserviceprice"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// 附加服务价格表
type ExtraServicePrice struct {
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
	// 附加服务类型
	ExtraServiceType enums.ExtraServiceType `json:"extra_service_type"`
	// 附加服务订单类型
	ExtraServiceBillingType enums.ExtraServiceBillingType `json:"extra_service_billing_type"`
	// 附加服务 id，外键关联附加服务
	ExtraServiceID int64 `json:"extra_service_id,omitempty" json:extra_service_id,string`
	// 附加服务单价
	Cep int64 `json:"cep"`
	// 价格有效时间开始，为空表示永久有效
	StartedAt *time.Time `json:"started_at"`
	// 价格有效时间结束，为空表示永久有效
	FinishedAt *time.Time `json:"finished_at"`
	// 价格是否屏蔽，前端置灰，硬选也可以
	IsDeprecated bool `json:"is_deprecated"`
	// 价格是否敏感，用于特殊类型任务，不能让外部看到选项
	IsSensitive bool `json:"is_sensitive"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ExtraServicePriceQuery when eager-loading is set.
	Edges        ExtraServicePriceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ExtraServicePriceEdges holds the relations/edges for other nodes in the graph.
type ExtraServicePriceEdges struct {
	// ExtraService holds the value of the extra_service edge.
	ExtraService *ExtraService `json:"extra_service,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ExtraServiceOrErr returns the ExtraService value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ExtraServicePriceEdges) ExtraServiceOrErr() (*ExtraService, error) {
	if e.loadedTypes[0] {
		if e.ExtraService == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: extraservice.Label}
		}
		return e.ExtraService, nil
	}
	return nil, &NotLoadedError{edge: "extra_service"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ExtraServicePrice) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case extraserviceprice.FieldIsDeprecated, extraserviceprice.FieldIsSensitive:
			values[i] = new(sql.NullBool)
		case extraserviceprice.FieldID, extraserviceprice.FieldCreatedBy, extraserviceprice.FieldUpdatedBy, extraserviceprice.FieldExtraServiceID, extraserviceprice.FieldCep:
			values[i] = new(sql.NullInt64)
		case extraserviceprice.FieldExtraServiceType, extraserviceprice.FieldExtraServiceBillingType:
			values[i] = new(sql.NullString)
		case extraserviceprice.FieldCreatedAt, extraserviceprice.FieldUpdatedAt, extraserviceprice.FieldDeletedAt, extraserviceprice.FieldStartedAt, extraserviceprice.FieldFinishedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ExtraServicePrice fields.
func (esp *ExtraServicePrice) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case extraserviceprice.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			esp.ID = int64(value.Int64)
		case extraserviceprice.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				esp.CreatedBy = value.Int64
			}
		case extraserviceprice.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				esp.UpdatedBy = value.Int64
			}
		case extraserviceprice.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				esp.CreatedAt = value.Time
			}
		case extraserviceprice.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				esp.UpdatedAt = value.Time
			}
		case extraserviceprice.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				esp.DeletedAt = value.Time
			}
		case extraserviceprice.FieldExtraServiceType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field extra_service_type", values[i])
			} else if value.Valid {
				esp.ExtraServiceType = enums.ExtraServiceType(value.String)
			}
		case extraserviceprice.FieldExtraServiceBillingType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field extra_service_billing_type", values[i])
			} else if value.Valid {
				esp.ExtraServiceBillingType = enums.ExtraServiceBillingType(value.String)
			}
		case extraserviceprice.FieldExtraServiceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field extra_service_id", values[i])
			} else if value.Valid {
				esp.ExtraServiceID = value.Int64
			}
		case extraserviceprice.FieldCep:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cep", values[i])
			} else if value.Valid {
				esp.Cep = value.Int64
			}
		case extraserviceprice.FieldStartedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field started_at", values[i])
			} else if value.Valid {
				esp.StartedAt = new(time.Time)
				*esp.StartedAt = value.Time
			}
		case extraserviceprice.FieldFinishedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field finished_at", values[i])
			} else if value.Valid {
				esp.FinishedAt = new(time.Time)
				*esp.FinishedAt = value.Time
			}
		case extraserviceprice.FieldIsDeprecated:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_deprecated", values[i])
			} else if value.Valid {
				esp.IsDeprecated = value.Bool
			}
		case extraserviceprice.FieldIsSensitive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_sensitive", values[i])
			} else if value.Valid {
				esp.IsSensitive = value.Bool
			}
		default:
			esp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ExtraServicePrice.
// This includes values selected through modifiers, order, etc.
func (esp *ExtraServicePrice) Value(name string) (ent.Value, error) {
	return esp.selectValues.Get(name)
}

// QueryExtraService queries the "extra_service" edge of the ExtraServicePrice entity.
func (esp *ExtraServicePrice) QueryExtraService() *ExtraServiceQuery {
	return NewExtraServicePriceClient(esp.config).QueryExtraService(esp)
}

// Update returns a builder for updating this ExtraServicePrice.
// Note that you need to call ExtraServicePrice.Unwrap() before calling this method if this ExtraServicePrice
// was returned from a transaction, and the transaction was committed or rolled back.
func (esp *ExtraServicePrice) Update() *ExtraServicePriceUpdateOne {
	return NewExtraServicePriceClient(esp.config).UpdateOne(esp)
}

// Unwrap unwraps the ExtraServicePrice entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (esp *ExtraServicePrice) Unwrap() *ExtraServicePrice {
	_tx, ok := esp.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: ExtraServicePrice is not a transactional entity")
	}
	esp.config.driver = _tx.drv
	return esp
}

// String implements the fmt.Stringer.
func (esp *ExtraServicePrice) String() string {
	var builder strings.Builder
	builder.WriteString("ExtraServicePrice(")
	builder.WriteString(fmt.Sprintf("id=%v, ", esp.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", esp.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", esp.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(esp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(esp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(esp.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("extra_service_type=")
	builder.WriteString(fmt.Sprintf("%v", esp.ExtraServiceType))
	builder.WriteString(", ")
	builder.WriteString("extra_service_billing_type=")
	builder.WriteString(fmt.Sprintf("%v", esp.ExtraServiceBillingType))
	builder.WriteString(", ")
	builder.WriteString("extra_service_id=")
	builder.WriteString(fmt.Sprintf("%v", esp.ExtraServiceID))
	builder.WriteString(", ")
	builder.WriteString("cep=")
	builder.WriteString(fmt.Sprintf("%v", esp.Cep))
	builder.WriteString(", ")
	if v := esp.StartedAt; v != nil {
		builder.WriteString("started_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := esp.FinishedAt; v != nil {
		builder.WriteString("finished_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("is_deprecated=")
	builder.WriteString(fmt.Sprintf("%v", esp.IsDeprecated))
	builder.WriteString(", ")
	builder.WriteString("is_sensitive=")
	builder.WriteString(fmt.Sprintf("%v", esp.IsSensitive))
	builder.WriteByte(')')
	return builder.String()
}

// ExtraServicePrices is a parsable slice of ExtraServicePrice.
type ExtraServicePrices []*ExtraServicePrice
