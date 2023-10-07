// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/mission"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionbatch"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/symbol"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// 任务订单，记录由任务产生的一些金额变动情况，取代任务消耗订单和任务生产订单
type MissionOrder struct {
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
	// 任务 id，外键关联任务
	MissionID int64 `json:"mission_id"`
	// 任务订单的状态，注意不强关联任务的状态
	Status enums.MissionOrderStatus `json:"status"`
	// 币种 id
	SymbolID int64 `json:"symbol_id"`
	// 外键关联发起任务的用户 id
	ConsumeUserID int64 `json:"consume_user_id"`
	// 订单的货币消耗量
	ConsumeAmount int64 `json:"consume_amount"`
	// 外键关联完成任务的用户 id
	ProduceUserID int64 `json:"produce_user_id"`
	// 订单的货币分润量
	ProduceAmount int64 `json:"produce_amount"`
	// 订单的平台收量，不记录用户 id，因为都是记载到 genesis 用户
	GasAmount int64 `json:"gas_amount"`
	// 任务类型，等于任务表的类型字段
	MissionType enums.MissionType `json:"mission_type"`
	// 是否为计时类型任务
	MissionBillingType enums.MissionBillingType `json:"mission_billing_type"`
	// 调用方式，API 调用或者微信小程序调用
	CallWay enums.MissionCallWay `json:"call_way"`
	// 订单序列号
	SerialNumber string `json:"serial_number"`
	// 任务开始执行时刻
	StartedAt time.Time `json:"started_at"`
	// 任务结束执行时刻
	FinishedAt time.Time `json:"finished_at"`
	// 任务计划开始时间（包时）
	PlanStartedAt *time.Time `json:"plan_started_at"`
	// 任务计划结束时间（包时）
	PlanFinishedAt *time.Time `json:"plan_finished_at"`
	// 任务批次外键
	MissionBatchID int64 `json:"mission_batch_id"`
	// 任务批次号，用于方便检索
	MissionBatchNumber string `json:"mission_batch_number"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MissionOrderQuery when eager-loading is set.
	Edges        MissionOrderEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MissionOrderEdges holds the relations/edges for other nodes in the graph.
type MissionOrderEdges struct {
	// ConsumeUser holds the value of the consume_user edge.
	ConsumeUser *User `json:"consume_user,omitempty"`
	// ProduceUser holds the value of the produce_user edge.
	ProduceUser *User `json:"produce_user,omitempty"`
	// Symbol holds the value of the symbol edge.
	Symbol *Symbol `json:"symbol,omitempty"`
	// Bills holds the value of the bills edge.
	Bills []*Bill `json:"bills,omitempty"`
	// MissionBatch holds the value of the mission_batch edge.
	MissionBatch *MissionBatch `json:"mission_batch,omitempty"`
	// Mission holds the value of the mission edge.
	Mission *Mission `json:"mission,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// ConsumeUserOrErr returns the ConsumeUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionOrderEdges) ConsumeUserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.ConsumeUser == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.ConsumeUser, nil
	}
	return nil, &NotLoadedError{edge: "consume_user"}
}

// ProduceUserOrErr returns the ProduceUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionOrderEdges) ProduceUserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.ProduceUser == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.ProduceUser, nil
	}
	return nil, &NotLoadedError{edge: "produce_user"}
}

// SymbolOrErr returns the Symbol value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionOrderEdges) SymbolOrErr() (*Symbol, error) {
	if e.loadedTypes[2] {
		if e.Symbol == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: symbol.Label}
		}
		return e.Symbol, nil
	}
	return nil, &NotLoadedError{edge: "symbol"}
}

// BillsOrErr returns the Bills value or an error if the edge
// was not loaded in eager-loading.
func (e MissionOrderEdges) BillsOrErr() ([]*Bill, error) {
	if e.loadedTypes[3] {
		return e.Bills, nil
	}
	return nil, &NotLoadedError{edge: "bills"}
}

// MissionBatchOrErr returns the MissionBatch value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionOrderEdges) MissionBatchOrErr() (*MissionBatch, error) {
	if e.loadedTypes[4] {
		if e.MissionBatch == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: missionbatch.Label}
		}
		return e.MissionBatch, nil
	}
	return nil, &NotLoadedError{edge: "mission_batch"}
}

// MissionOrErr returns the Mission value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionOrderEdges) MissionOrErr() (*Mission, error) {
	if e.loadedTypes[5] {
		if e.Mission == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: mission.Label}
		}
		return e.Mission, nil
	}
	return nil, &NotLoadedError{edge: "mission"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MissionOrder) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case missionorder.FieldID, missionorder.FieldCreatedBy, missionorder.FieldUpdatedBy, missionorder.FieldMissionID, missionorder.FieldSymbolID, missionorder.FieldConsumeUserID, missionorder.FieldConsumeAmount, missionorder.FieldProduceUserID, missionorder.FieldProduceAmount, missionorder.FieldGasAmount, missionorder.FieldMissionBatchID:
			values[i] = new(sql.NullInt64)
		case missionorder.FieldStatus, missionorder.FieldMissionType, missionorder.FieldMissionBillingType, missionorder.FieldCallWay, missionorder.FieldSerialNumber, missionorder.FieldMissionBatchNumber:
			values[i] = new(sql.NullString)
		case missionorder.FieldCreatedAt, missionorder.FieldUpdatedAt, missionorder.FieldDeletedAt, missionorder.FieldStartedAt, missionorder.FieldFinishedAt, missionorder.FieldPlanStartedAt, missionorder.FieldPlanFinishedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MissionOrder fields.
func (mo *MissionOrder) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case missionorder.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mo.ID = int64(value.Int64)
		case missionorder.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				mo.CreatedBy = value.Int64
			}
		case missionorder.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				mo.UpdatedBy = value.Int64
			}
		case missionorder.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mo.CreatedAt = value.Time
			}
		case missionorder.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mo.UpdatedAt = value.Time
			}
		case missionorder.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				mo.DeletedAt = value.Time
			}
		case missionorder.FieldMissionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mission_id", values[i])
			} else if value.Valid {
				mo.MissionID = value.Int64
			}
		case missionorder.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				mo.Status = enums.MissionOrderStatus(value.String)
			}
		case missionorder.FieldSymbolID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field symbol_id", values[i])
			} else if value.Valid {
				mo.SymbolID = value.Int64
			}
		case missionorder.FieldConsumeUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field consume_user_id", values[i])
			} else if value.Valid {
				mo.ConsumeUserID = value.Int64
			}
		case missionorder.FieldConsumeAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field consume_amount", values[i])
			} else if value.Valid {
				mo.ConsumeAmount = value.Int64
			}
		case missionorder.FieldProduceUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field produce_user_id", values[i])
			} else if value.Valid {
				mo.ProduceUserID = value.Int64
			}
		case missionorder.FieldProduceAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field produce_amount", values[i])
			} else if value.Valid {
				mo.ProduceAmount = value.Int64
			}
		case missionorder.FieldGasAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field gas_amount", values[i])
			} else if value.Valid {
				mo.GasAmount = value.Int64
			}
		case missionorder.FieldMissionType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mission_type", values[i])
			} else if value.Valid {
				mo.MissionType = enums.MissionType(value.String)
			}
		case missionorder.FieldMissionBillingType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mission_billing_type", values[i])
			} else if value.Valid {
				mo.MissionBillingType = enums.MissionBillingType(value.String)
			}
		case missionorder.FieldCallWay:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field call_way", values[i])
			} else if value.Valid {
				mo.CallWay = enums.MissionCallWay(value.String)
			}
		case missionorder.FieldSerialNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial_number", values[i])
			} else if value.Valid {
				mo.SerialNumber = value.String
			}
		case missionorder.FieldStartedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field started_at", values[i])
			} else if value.Valid {
				mo.StartedAt = value.Time
			}
		case missionorder.FieldFinishedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field finished_at", values[i])
			} else if value.Valid {
				mo.FinishedAt = value.Time
			}
		case missionorder.FieldPlanStartedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field plan_started_at", values[i])
			} else if value.Valid {
				mo.PlanStartedAt = new(time.Time)
				*mo.PlanStartedAt = value.Time
			}
		case missionorder.FieldPlanFinishedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field plan_finished_at", values[i])
			} else if value.Valid {
				mo.PlanFinishedAt = new(time.Time)
				*mo.PlanFinishedAt = value.Time
			}
		case missionorder.FieldMissionBatchID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mission_batch_id", values[i])
			} else if value.Valid {
				mo.MissionBatchID = value.Int64
			}
		case missionorder.FieldMissionBatchNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mission_batch_number", values[i])
			} else if value.Valid {
				mo.MissionBatchNumber = value.String
			}
		default:
			mo.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MissionOrder.
// This includes values selected through modifiers, order, etc.
func (mo *MissionOrder) Value(name string) (ent.Value, error) {
	return mo.selectValues.Get(name)
}

// QueryConsumeUser queries the "consume_user" edge of the MissionOrder entity.
func (mo *MissionOrder) QueryConsumeUser() *UserQuery {
	return NewMissionOrderClient(mo.config).QueryConsumeUser(mo)
}

// QueryProduceUser queries the "produce_user" edge of the MissionOrder entity.
func (mo *MissionOrder) QueryProduceUser() *UserQuery {
	return NewMissionOrderClient(mo.config).QueryProduceUser(mo)
}

// QuerySymbol queries the "symbol" edge of the MissionOrder entity.
func (mo *MissionOrder) QuerySymbol() *SymbolQuery {
	return NewMissionOrderClient(mo.config).QuerySymbol(mo)
}

// QueryBills queries the "bills" edge of the MissionOrder entity.
func (mo *MissionOrder) QueryBills() *BillQuery {
	return NewMissionOrderClient(mo.config).QueryBills(mo)
}

// QueryMissionBatch queries the "mission_batch" edge of the MissionOrder entity.
func (mo *MissionOrder) QueryMissionBatch() *MissionBatchQuery {
	return NewMissionOrderClient(mo.config).QueryMissionBatch(mo)
}

// QueryMission queries the "mission" edge of the MissionOrder entity.
func (mo *MissionOrder) QueryMission() *MissionQuery {
	return NewMissionOrderClient(mo.config).QueryMission(mo)
}

// Update returns a builder for updating this MissionOrder.
// Note that you need to call MissionOrder.Unwrap() before calling this method if this MissionOrder
// was returned from a transaction, and the transaction was committed or rolled back.
func (mo *MissionOrder) Update() *MissionOrderUpdateOne {
	return NewMissionOrderClient(mo.config).UpdateOne(mo)
}

// Unwrap unwraps the MissionOrder entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mo *MissionOrder) Unwrap() *MissionOrder {
	_tx, ok := mo.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: MissionOrder is not a transactional entity")
	}
	mo.config.driver = _tx.drv
	return mo
}

// String implements the fmt.Stringer.
func (mo *MissionOrder) String() string {
	var builder strings.Builder
	builder.WriteString("MissionOrder(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mo.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", mo.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", mo.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(mo.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(mo.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(mo.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("mission_id=")
	builder.WriteString(fmt.Sprintf("%v", mo.MissionID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", mo.Status))
	builder.WriteString(", ")
	builder.WriteString("symbol_id=")
	builder.WriteString(fmt.Sprintf("%v", mo.SymbolID))
	builder.WriteString(", ")
	builder.WriteString("consume_user_id=")
	builder.WriteString(fmt.Sprintf("%v", mo.ConsumeUserID))
	builder.WriteString(", ")
	builder.WriteString("consume_amount=")
	builder.WriteString(fmt.Sprintf("%v", mo.ConsumeAmount))
	builder.WriteString(", ")
	builder.WriteString("produce_user_id=")
	builder.WriteString(fmt.Sprintf("%v", mo.ProduceUserID))
	builder.WriteString(", ")
	builder.WriteString("produce_amount=")
	builder.WriteString(fmt.Sprintf("%v", mo.ProduceAmount))
	builder.WriteString(", ")
	builder.WriteString("gas_amount=")
	builder.WriteString(fmt.Sprintf("%v", mo.GasAmount))
	builder.WriteString(", ")
	builder.WriteString("mission_type=")
	builder.WriteString(fmt.Sprintf("%v", mo.MissionType))
	builder.WriteString(", ")
	builder.WriteString("mission_billing_type=")
	builder.WriteString(fmt.Sprintf("%v", mo.MissionBillingType))
	builder.WriteString(", ")
	builder.WriteString("call_way=")
	builder.WriteString(fmt.Sprintf("%v", mo.CallWay))
	builder.WriteString(", ")
	builder.WriteString("serial_number=")
	builder.WriteString(mo.SerialNumber)
	builder.WriteString(", ")
	builder.WriteString("started_at=")
	builder.WriteString(mo.StartedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("finished_at=")
	builder.WriteString(mo.FinishedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := mo.PlanStartedAt; v != nil {
		builder.WriteString("plan_started_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := mo.PlanFinishedAt; v != nil {
		builder.WriteString("plan_finished_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("mission_batch_id=")
	builder.WriteString(fmt.Sprintf("%v", mo.MissionBatchID))
	builder.WriteString(", ")
	builder.WriteString("mission_batch_number=")
	builder.WriteString(mo.MissionBatchNumber)
	builder.WriteByte(')')
	return builder.String()
}

// MissionOrders is a parsable slice of MissionOrder.
type MissionOrders []*MissionOrder
