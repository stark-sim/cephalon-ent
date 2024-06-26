// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/lotto"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/lottousercount"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// 用户可用抽奖次数表
type LottoUserCount struct {
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
	// 外键：用户 ID
	UserID int64 `json:"user_id"`
	// 外键：抽奖活动 ID
	LottoID int64 `json:"lotto_id"`
	// 剩余抽奖次数
	RemainLottoCount int64 `json:"remain_lotto_count"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LottoUserCountQuery when eager-loading is set.
	Edges        LottoUserCountEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LottoUserCountEdges holds the relations/edges for other nodes in the graph.
type LottoUserCountEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Lotto holds the value of the lotto edge.
	Lotto *Lotto `json:"lotto,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LottoUserCountEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// LottoOrErr returns the Lotto value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LottoUserCountEdges) LottoOrErr() (*Lotto, error) {
	if e.loadedTypes[1] {
		if e.Lotto == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: lotto.Label}
		}
		return e.Lotto, nil
	}
	return nil, &NotLoadedError{edge: "lotto"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LottoUserCount) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case lottousercount.FieldID, lottousercount.FieldCreatedBy, lottousercount.FieldUpdatedBy, lottousercount.FieldUserID, lottousercount.FieldLottoID, lottousercount.FieldRemainLottoCount:
			values[i] = new(sql.NullInt64)
		case lottousercount.FieldCreatedAt, lottousercount.FieldUpdatedAt, lottousercount.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LottoUserCount fields.
func (luc *LottoUserCount) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case lottousercount.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			luc.ID = int64(value.Int64)
		case lottousercount.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				luc.CreatedBy = value.Int64
			}
		case lottousercount.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				luc.UpdatedBy = value.Int64
			}
		case lottousercount.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				luc.CreatedAt = value.Time
			}
		case lottousercount.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				luc.UpdatedAt = value.Time
			}
		case lottousercount.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				luc.DeletedAt = value.Time
			}
		case lottousercount.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				luc.UserID = value.Int64
			}
		case lottousercount.FieldLottoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field lotto_id", values[i])
			} else if value.Valid {
				luc.LottoID = value.Int64
			}
		case lottousercount.FieldRemainLottoCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field remain_lotto_count", values[i])
			} else if value.Valid {
				luc.RemainLottoCount = value.Int64
			}
		default:
			luc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the LottoUserCount.
// This includes values selected through modifiers, order, etc.
func (luc *LottoUserCount) Value(name string) (ent.Value, error) {
	return luc.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the LottoUserCount entity.
func (luc *LottoUserCount) QueryUser() *UserQuery {
	return NewLottoUserCountClient(luc.config).QueryUser(luc)
}

// QueryLotto queries the "lotto" edge of the LottoUserCount entity.
func (luc *LottoUserCount) QueryLotto() *LottoQuery {
	return NewLottoUserCountClient(luc.config).QueryLotto(luc)
}

// Update returns a builder for updating this LottoUserCount.
// Note that you need to call LottoUserCount.Unwrap() before calling this method if this LottoUserCount
// was returned from a transaction, and the transaction was committed or rolled back.
func (luc *LottoUserCount) Update() *LottoUserCountUpdateOne {
	return NewLottoUserCountClient(luc.config).UpdateOne(luc)
}

// Unwrap unwraps the LottoUserCount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (luc *LottoUserCount) Unwrap() *LottoUserCount {
	_tx, ok := luc.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: LottoUserCount is not a transactional entity")
	}
	luc.config.driver = _tx.drv
	return luc
}

// String implements the fmt.Stringer.
func (luc *LottoUserCount) String() string {
	var builder strings.Builder
	builder.WriteString("LottoUserCount(")
	builder.WriteString(fmt.Sprintf("id=%v, ", luc.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", luc.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", luc.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(luc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(luc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(luc.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", luc.UserID))
	builder.WriteString(", ")
	builder.WriteString("lotto_id=")
	builder.WriteString(fmt.Sprintf("%v", luc.LottoID))
	builder.WriteString(", ")
	builder.WriteString("remain_lotto_count=")
	builder.WriteString(fmt.Sprintf("%v", luc.RemainLottoCount))
	builder.WriteByte(')')
	return builder.String()
}

// LottoUserCounts is a parsable slice of LottoUserCount.
type LottoUserCounts []*LottoUserCount
