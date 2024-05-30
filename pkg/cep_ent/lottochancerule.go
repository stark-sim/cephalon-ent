// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/lotto"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/lottochancerule"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// 获取抽奖次数规则表
type LottoChanceRule struct {
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
	// 外键：抽奖活动 ID
	LottoID int64 `json:"lotto_id"`
	// 条件
	Condition enums.LottoCondition `json:"condition"`
	// 奖励抽奖次数
	AwardCount int64 `json:"award_count"`
	// 充值金额，类型为充值时才有数据
	RechargeAmount int64 `json:"recharge_amount"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LottoChanceRuleQuery when eager-loading is set.
	Edges        LottoChanceRuleEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LottoChanceRuleEdges holds the relations/edges for other nodes in the graph.
type LottoChanceRuleEdges struct {
	// Lotto holds the value of the lotto edge.
	Lotto *Lotto `json:"lotto,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// LottoOrErr returns the Lotto value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LottoChanceRuleEdges) LottoOrErr() (*Lotto, error) {
	if e.loadedTypes[0] {
		if e.Lotto == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: lotto.Label}
		}
		return e.Lotto, nil
	}
	return nil, &NotLoadedError{edge: "lotto"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LottoChanceRule) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case lottochancerule.FieldID, lottochancerule.FieldCreatedBy, lottochancerule.FieldUpdatedBy, lottochancerule.FieldLottoID, lottochancerule.FieldAwardCount, lottochancerule.FieldRechargeAmount:
			values[i] = new(sql.NullInt64)
		case lottochancerule.FieldCondition:
			values[i] = new(sql.NullString)
		case lottochancerule.FieldCreatedAt, lottochancerule.FieldUpdatedAt, lottochancerule.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LottoChanceRule fields.
func (lcr *LottoChanceRule) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case lottochancerule.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			lcr.ID = int64(value.Int64)
		case lottochancerule.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				lcr.CreatedBy = value.Int64
			}
		case lottochancerule.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				lcr.UpdatedBy = value.Int64
			}
		case lottochancerule.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				lcr.CreatedAt = value.Time
			}
		case lottochancerule.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				lcr.UpdatedAt = value.Time
			}
		case lottochancerule.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				lcr.DeletedAt = value.Time
			}
		case lottochancerule.FieldLottoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field lotto_id", values[i])
			} else if value.Valid {
				lcr.LottoID = value.Int64
			}
		case lottochancerule.FieldCondition:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field condition", values[i])
			} else if value.Valid {
				lcr.Condition = enums.LottoCondition(value.String)
			}
		case lottochancerule.FieldAwardCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field award_count", values[i])
			} else if value.Valid {
				lcr.AwardCount = value.Int64
			}
		case lottochancerule.FieldRechargeAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field recharge_amount", values[i])
			} else if value.Valid {
				lcr.RechargeAmount = value.Int64
			}
		default:
			lcr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the LottoChanceRule.
// This includes values selected through modifiers, order, etc.
func (lcr *LottoChanceRule) Value(name string) (ent.Value, error) {
	return lcr.selectValues.Get(name)
}

// QueryLotto queries the "lotto" edge of the LottoChanceRule entity.
func (lcr *LottoChanceRule) QueryLotto() *LottoQuery {
	return NewLottoChanceRuleClient(lcr.config).QueryLotto(lcr)
}

// Update returns a builder for updating this LottoChanceRule.
// Note that you need to call LottoChanceRule.Unwrap() before calling this method if this LottoChanceRule
// was returned from a transaction, and the transaction was committed or rolled back.
func (lcr *LottoChanceRule) Update() *LottoChanceRuleUpdateOne {
	return NewLottoChanceRuleClient(lcr.config).UpdateOne(lcr)
}

// Unwrap unwraps the LottoChanceRule entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (lcr *LottoChanceRule) Unwrap() *LottoChanceRule {
	_tx, ok := lcr.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: LottoChanceRule is not a transactional entity")
	}
	lcr.config.driver = _tx.drv
	return lcr
}

// String implements the fmt.Stringer.
func (lcr *LottoChanceRule) String() string {
	var builder strings.Builder
	builder.WriteString("LottoChanceRule(")
	builder.WriteString(fmt.Sprintf("id=%v, ", lcr.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", lcr.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", lcr.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(lcr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(lcr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(lcr.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("lotto_id=")
	builder.WriteString(fmt.Sprintf("%v", lcr.LottoID))
	builder.WriteString(", ")
	builder.WriteString("condition=")
	builder.WriteString(fmt.Sprintf("%v", lcr.Condition))
	builder.WriteString(", ")
	builder.WriteString("award_count=")
	builder.WriteString(fmt.Sprintf("%v", lcr.AwardCount))
	builder.WriteString(", ")
	builder.WriteString("recharge_amount=")
	builder.WriteString(fmt.Sprintf("%v", lcr.RechargeAmount))
	builder.WriteByte(')')
	return builder.String()
}

// LottoChanceRules is a parsable slice of LottoChanceRule.
type LottoChanceRules []*LottoChanceRule
