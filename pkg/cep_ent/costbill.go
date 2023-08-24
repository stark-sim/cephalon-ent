// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/costaccount"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/costbill"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionconsumeorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/rechargeorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// 记录额度账户的变动的主流水账单记录
type CostBill struct {
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
	// 额度账户流水的类型，充值或者任务消耗
	Type costbill.Type `json:"type"`
	// 是否增加余额，布尔值默认为否
	IsAdd bool `json:"is_add"`
	// 外键用户 id
	UserID int64 `json:"user_id"`
	// 账单序列号
	SerialNumber string `json:"serial_number"`
	// 外键消耗账户 id
	CostAccountID int64 `json:"cost_account_id"`
	// 消耗多少本金余额
	PureCep int64 `json:"pure_cep"`
	// 消耗多少赠送余额
	GiftCep int64 `json:"gift_cep"`
	// 关联消耗产生的来源外键 id，比如 type 为 mission 时关联任务订单
	ReasonID int64 `json:"reason_id"`
	// 消耗流水一开始不能直接生效，确定关联的消耗时间成功后才可以扣费
	Status enums.BillStatus `json:"status"`
	// 营销流水 id
	MarketBillID int64 `json:"market_bill_id"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CostBillQuery when eager-loading is set.
	Edges        CostBillEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CostBillEdges holds the relations/edges for other nodes in the graph.
type CostBillEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// CostAccount holds the value of the cost_account edge.
	CostAccount *CostAccount `json:"cost_account,omitempty"`
	// RechargeOrder holds the value of the recharge_order edge.
	RechargeOrder *RechargeOrder `json:"recharge_order,omitempty"`
	// MissionConsumeOrder holds the value of the mission_consume_order edge.
	MissionConsumeOrder *MissionConsumeOrder `json:"mission_consume_order,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CostBillEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// CostAccountOrErr returns the CostAccount value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CostBillEdges) CostAccountOrErr() (*CostAccount, error) {
	if e.loadedTypes[1] {
		if e.CostAccount == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: costaccount.Label}
		}
		return e.CostAccount, nil
	}
	return nil, &NotLoadedError{edge: "cost_account"}
}

// RechargeOrderOrErr returns the RechargeOrder value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CostBillEdges) RechargeOrderOrErr() (*RechargeOrder, error) {
	if e.loadedTypes[2] {
		if e.RechargeOrder == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: rechargeorder.Label}
		}
		return e.RechargeOrder, nil
	}
	return nil, &NotLoadedError{edge: "recharge_order"}
}

// MissionConsumeOrderOrErr returns the MissionConsumeOrder value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CostBillEdges) MissionConsumeOrderOrErr() (*MissionConsumeOrder, error) {
	if e.loadedTypes[3] {
		if e.MissionConsumeOrder == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: missionconsumeorder.Label}
		}
		return e.MissionConsumeOrder, nil
	}
	return nil, &NotLoadedError{edge: "mission_consume_order"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CostBill) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case costbill.FieldIsAdd:
			values[i] = new(sql.NullBool)
		case costbill.FieldID, costbill.FieldCreatedBy, costbill.FieldUpdatedBy, costbill.FieldUserID, costbill.FieldCostAccountID, costbill.FieldPureCep, costbill.FieldGiftCep, costbill.FieldReasonID, costbill.FieldMarketBillID:
			values[i] = new(sql.NullInt64)
		case costbill.FieldType, costbill.FieldSerialNumber, costbill.FieldStatus:
			values[i] = new(sql.NullString)
		case costbill.FieldCreatedAt, costbill.FieldUpdatedAt, costbill.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CostBill fields.
func (cb *CostBill) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case costbill.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cb.ID = int64(value.Int64)
		case costbill.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				cb.CreatedBy = value.Int64
			}
		case costbill.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				cb.UpdatedBy = value.Int64
			}
		case costbill.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cb.CreatedAt = value.Time
			}
		case costbill.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cb.UpdatedAt = value.Time
			}
		case costbill.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				cb.DeletedAt = value.Time
			}
		case costbill.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				cb.Type = costbill.Type(value.String)
			}
		case costbill.FieldIsAdd:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_add", values[i])
			} else if value.Valid {
				cb.IsAdd = value.Bool
			}
		case costbill.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				cb.UserID = value.Int64
			}
		case costbill.FieldSerialNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial_number", values[i])
			} else if value.Valid {
				cb.SerialNumber = value.String
			}
		case costbill.FieldCostAccountID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cost_account_id", values[i])
			} else if value.Valid {
				cb.CostAccountID = value.Int64
			}
		case costbill.FieldPureCep:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pure_cep", values[i])
			} else if value.Valid {
				cb.PureCep = value.Int64
			}
		case costbill.FieldGiftCep:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field gift_cep", values[i])
			} else if value.Valid {
				cb.GiftCep = value.Int64
			}
		case costbill.FieldReasonID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field reason_id", values[i])
			} else if value.Valid {
				cb.ReasonID = value.Int64
			}
		case costbill.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				cb.Status = enums.BillStatus(value.String)
			}
		case costbill.FieldMarketBillID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field market_bill_id", values[i])
			} else if value.Valid {
				cb.MarketBillID = value.Int64
			}
		default:
			cb.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CostBill.
// This includes values selected through modifiers, order, etc.
func (cb *CostBill) Value(name string) (ent.Value, error) {
	return cb.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the CostBill entity.
func (cb *CostBill) QueryUser() *UserQuery {
	return NewCostBillClient(cb.config).QueryUser(cb)
}

// QueryCostAccount queries the "cost_account" edge of the CostBill entity.
func (cb *CostBill) QueryCostAccount() *CostAccountQuery {
	return NewCostBillClient(cb.config).QueryCostAccount(cb)
}

// QueryRechargeOrder queries the "recharge_order" edge of the CostBill entity.
func (cb *CostBill) QueryRechargeOrder() *RechargeOrderQuery {
	return NewCostBillClient(cb.config).QueryRechargeOrder(cb)
}

// QueryMissionConsumeOrder queries the "mission_consume_order" edge of the CostBill entity.
func (cb *CostBill) QueryMissionConsumeOrder() *MissionConsumeOrderQuery {
	return NewCostBillClient(cb.config).QueryMissionConsumeOrder(cb)
}

// Update returns a builder for updating this CostBill.
// Note that you need to call CostBill.Unwrap() before calling this method if this CostBill
// was returned from a transaction, and the transaction was committed or rolled back.
func (cb *CostBill) Update() *CostBillUpdateOne {
	return NewCostBillClient(cb.config).UpdateOne(cb)
}

// Unwrap unwraps the CostBill entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cb *CostBill) Unwrap() *CostBill {
	_tx, ok := cb.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: CostBill is not a transactional entity")
	}
	cb.config.driver = _tx.drv
	return cb
}

// String implements the fmt.Stringer.
func (cb *CostBill) String() string {
	var builder strings.Builder
	builder.WriteString("CostBill(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cb.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", cb.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", cb.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(cb.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(cb.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(cb.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", cb.Type))
	builder.WriteString(", ")
	builder.WriteString("is_add=")
	builder.WriteString(fmt.Sprintf("%v", cb.IsAdd))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", cb.UserID))
	builder.WriteString(", ")
	builder.WriteString("serial_number=")
	builder.WriteString(cb.SerialNumber)
	builder.WriteString(", ")
	builder.WriteString("cost_account_id=")
	builder.WriteString(fmt.Sprintf("%v", cb.CostAccountID))
	builder.WriteString(", ")
	builder.WriteString("pure_cep=")
	builder.WriteString(fmt.Sprintf("%v", cb.PureCep))
	builder.WriteString(", ")
	builder.WriteString("gift_cep=")
	builder.WriteString(fmt.Sprintf("%v", cb.GiftCep))
	builder.WriteString(", ")
	builder.WriteString("reason_id=")
	builder.WriteString(fmt.Sprintf("%v", cb.ReasonID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", cb.Status))
	builder.WriteString(", ")
	builder.WriteString("market_bill_id=")
	builder.WriteString(fmt.Sprintf("%v", cb.MarketBillID))
	builder.WriteByte(')')
	return builder.String()
}

// CostBills is a parsable slice of CostBill.
type CostBills []*CostBill
