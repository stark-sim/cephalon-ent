// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missioncategory"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// 任务大类，任务类型的最高抽象层，记录了所有任务大类
type MissionCategory struct {
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
	// 任务大类
	Category string `json:"category"`
	// 任务大类类型，比如热门类型等
	Type enums.CategoryType `json:"type"`
	// 权重（目前排序可以用到）
	Weight       int `json:"weight"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MissionCategory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case missioncategory.FieldID, missioncategory.FieldCreatedBy, missioncategory.FieldUpdatedBy, missioncategory.FieldWeight:
			values[i] = new(sql.NullInt64)
		case missioncategory.FieldCategory, missioncategory.FieldType:
			values[i] = new(sql.NullString)
		case missioncategory.FieldCreatedAt, missioncategory.FieldUpdatedAt, missioncategory.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MissionCategory fields.
func (mc *MissionCategory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case missioncategory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mc.ID = int64(value.Int64)
		case missioncategory.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				mc.CreatedBy = value.Int64
			}
		case missioncategory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				mc.UpdatedBy = value.Int64
			}
		case missioncategory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mc.CreatedAt = value.Time
			}
		case missioncategory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mc.UpdatedAt = value.Time
			}
		case missioncategory.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				mc.DeletedAt = value.Time
			}
		case missioncategory.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				mc.Category = value.String
			}
		case missioncategory.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				mc.Type = enums.CategoryType(value.String)
			}
		case missioncategory.FieldWeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				mc.Weight = int(value.Int64)
			}
		default:
			mc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MissionCategory.
// This includes values selected through modifiers, order, etc.
func (mc *MissionCategory) Value(name string) (ent.Value, error) {
	return mc.selectValues.Get(name)
}

// Update returns a builder for updating this MissionCategory.
// Note that you need to call MissionCategory.Unwrap() before calling this method if this MissionCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (mc *MissionCategory) Update() *MissionCategoryUpdateOne {
	return NewMissionCategoryClient(mc.config).UpdateOne(mc)
}

// Unwrap unwraps the MissionCategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mc *MissionCategory) Unwrap() *MissionCategory {
	_tx, ok := mc.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: MissionCategory is not a transactional entity")
	}
	mc.config.driver = _tx.drv
	return mc
}

// String implements the fmt.Stringer.
func (mc *MissionCategory) String() string {
	var builder strings.Builder
	builder.WriteString("MissionCategory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mc.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", mc.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", mc.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(mc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(mc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(mc.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(mc.Category)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", mc.Type))
	builder.WriteString(", ")
	builder.WriteString("weight=")
	builder.WriteString(fmt.Sprintf("%v", mc.Weight))
	builder.WriteByte(')')
	return builder.String()
}

// MissionCategories is a parsable slice of MissionCategory.
type MissionCategories []*MissionCategory
