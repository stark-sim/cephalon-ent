// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/symbol"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/transferorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/vxsocial"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/withdrawrecord"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// 转账订单，谁给谁转了多少什么货币
type TransferOrder struct {
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
	// 转账来源的用户 id
	SourceUserID int64 `json:"source_user_id,string"`
	// 转账目标的用户 id
	TargetUserID int64 `json:"target_user_id,string"`
	// 转账订单的状态，比如微信发起支付后可能没完成支付
	Status transferorder.Status `json:"status"`
	// 币种 id
	SymbolID int64 `json:"symbol_id,string"`
	// 充值多少货币量
	Amount int64 `json:"amount"`
	// 充值订单的类型
	Type enums.TransferOrderType `json:"type"`
	// 充值订单的序列号
	SerialNumber string `json:"serial_number"`
	// 关联充值来源的身份源 id
	SocialID int64 `json:"social_id,string"`
	// 第三方平台的返回，给到前端才能发起支付
	ThirdAPIResp string `json:"third_api_resp"`
	// 平台方订单号
	OutTransactionID string `json:"out_transaction_id"`
	// 操作的用户 id，管理后台手动充值才有数据，默认为 0
	OperateUserID int64 `json:"operate_user_id,string"`
	// 充值订单活动赠送的类型
	GiftType enums.TransferOrderGiftType `json:"gift_type"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TransferOrderQuery when eager-loading is set.
	Edges        TransferOrderEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TransferOrderEdges holds the relations/edges for other nodes in the graph.
type TransferOrderEdges struct {
	// SourceUser holds the value of the source_user edge.
	SourceUser *User `json:"source_user,omitempty"`
	// TargetUser holds the value of the target_user edge.
	TargetUser *User `json:"target_user,omitempty"`
	// Bills holds the value of the bills edge.
	Bills []*Bill `json:"bills,omitempty"`
	// VxSocial holds the value of the vx_social edge.
	VxSocial *VXSocial `json:"vx_social,omitempty"`
	// Symbol holds the value of the symbol edge.
	Symbol *Symbol `json:"symbol,omitempty"`
	// WithdrawRecord holds the value of the withdraw_record edge.
	WithdrawRecord *WithdrawRecord `json:"withdraw_record,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// SourceUserOrErr returns the SourceUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransferOrderEdges) SourceUserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.SourceUser == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.SourceUser, nil
	}
	return nil, &NotLoadedError{edge: "source_user"}
}

// TargetUserOrErr returns the TargetUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransferOrderEdges) TargetUserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.TargetUser == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.TargetUser, nil
	}
	return nil, &NotLoadedError{edge: "target_user"}
}

// BillsOrErr returns the Bills value or an error if the edge
// was not loaded in eager-loading.
func (e TransferOrderEdges) BillsOrErr() ([]*Bill, error) {
	if e.loadedTypes[2] {
		return e.Bills, nil
	}
	return nil, &NotLoadedError{edge: "bills"}
}

// VxSocialOrErr returns the VxSocial value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransferOrderEdges) VxSocialOrErr() (*VXSocial, error) {
	if e.loadedTypes[3] {
		if e.VxSocial == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: vxsocial.Label}
		}
		return e.VxSocial, nil
	}
	return nil, &NotLoadedError{edge: "vx_social"}
}

// SymbolOrErr returns the Symbol value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransferOrderEdges) SymbolOrErr() (*Symbol, error) {
	if e.loadedTypes[4] {
		if e.Symbol == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: symbol.Label}
		}
		return e.Symbol, nil
	}
	return nil, &NotLoadedError{edge: "symbol"}
}

// WithdrawRecordOrErr returns the WithdrawRecord value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransferOrderEdges) WithdrawRecordOrErr() (*WithdrawRecord, error) {
	if e.loadedTypes[5] {
		if e.WithdrawRecord == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: withdrawrecord.Label}
		}
		return e.WithdrawRecord, nil
	}
	return nil, &NotLoadedError{edge: "withdraw_record"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TransferOrder) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case transferorder.FieldID, transferorder.FieldCreatedBy, transferorder.FieldUpdatedBy, transferorder.FieldSourceUserID, transferorder.FieldTargetUserID, transferorder.FieldSymbolID, transferorder.FieldAmount, transferorder.FieldSocialID, transferorder.FieldOperateUserID:
			values[i] = new(sql.NullInt64)
		case transferorder.FieldStatus, transferorder.FieldType, transferorder.FieldSerialNumber, transferorder.FieldThirdAPIResp, transferorder.FieldOutTransactionID, transferorder.FieldGiftType:
			values[i] = new(sql.NullString)
		case transferorder.FieldCreatedAt, transferorder.FieldUpdatedAt, transferorder.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TransferOrder fields.
func (to *TransferOrder) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transferorder.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			to.ID = int64(value.Int64)
		case transferorder.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				to.CreatedBy = value.Int64
			}
		case transferorder.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				to.UpdatedBy = value.Int64
			}
		case transferorder.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				to.CreatedAt = value.Time
			}
		case transferorder.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				to.UpdatedAt = value.Time
			}
		case transferorder.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				to.DeletedAt = value.Time
			}
		case transferorder.FieldSourceUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field source_user_id", values[i])
			} else if value.Valid {
				to.SourceUserID = value.Int64
			}
		case transferorder.FieldTargetUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field target_user_id", values[i])
			} else if value.Valid {
				to.TargetUserID = value.Int64
			}
		case transferorder.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				to.Status = transferorder.Status(value.String)
			}
		case transferorder.FieldSymbolID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field symbol_id", values[i])
			} else if value.Valid {
				to.SymbolID = value.Int64
			}
		case transferorder.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				to.Amount = value.Int64
			}
		case transferorder.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				to.Type = enums.TransferOrderType(value.String)
			}
		case transferorder.FieldSerialNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial_number", values[i])
			} else if value.Valid {
				to.SerialNumber = value.String
			}
		case transferorder.FieldSocialID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field social_id", values[i])
			} else if value.Valid {
				to.SocialID = value.Int64
			}
		case transferorder.FieldThirdAPIResp:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field third_api_resp", values[i])
			} else if value.Valid {
				to.ThirdAPIResp = value.String
			}
		case transferorder.FieldOutTransactionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field out_transaction_id", values[i])
			} else if value.Valid {
				to.OutTransactionID = value.String
			}
		case transferorder.FieldOperateUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field operate_user_id", values[i])
			} else if value.Valid {
				to.OperateUserID = value.Int64
			}
		case transferorder.FieldGiftType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gift_type", values[i])
			} else if value.Valid {
				to.GiftType = enums.TransferOrderGiftType(value.String)
			}
		default:
			to.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TransferOrder.
// This includes values selected through modifiers, order, etc.
func (to *TransferOrder) Value(name string) (ent.Value, error) {
	return to.selectValues.Get(name)
}

// QuerySourceUser queries the "source_user" edge of the TransferOrder entity.
func (to *TransferOrder) QuerySourceUser() *UserQuery {
	return NewTransferOrderClient(to.config).QuerySourceUser(to)
}

// QueryTargetUser queries the "target_user" edge of the TransferOrder entity.
func (to *TransferOrder) QueryTargetUser() *UserQuery {
	return NewTransferOrderClient(to.config).QueryTargetUser(to)
}

// QueryBills queries the "bills" edge of the TransferOrder entity.
func (to *TransferOrder) QueryBills() *BillQuery {
	return NewTransferOrderClient(to.config).QueryBills(to)
}

// QueryVxSocial queries the "vx_social" edge of the TransferOrder entity.
func (to *TransferOrder) QueryVxSocial() *VXSocialQuery {
	return NewTransferOrderClient(to.config).QueryVxSocial(to)
}

// QuerySymbol queries the "symbol" edge of the TransferOrder entity.
func (to *TransferOrder) QuerySymbol() *SymbolQuery {
	return NewTransferOrderClient(to.config).QuerySymbol(to)
}

// QueryWithdrawRecord queries the "withdraw_record" edge of the TransferOrder entity.
func (to *TransferOrder) QueryWithdrawRecord() *WithdrawRecordQuery {
	return NewTransferOrderClient(to.config).QueryWithdrawRecord(to)
}

// Update returns a builder for updating this TransferOrder.
// Note that you need to call TransferOrder.Unwrap() before calling this method if this TransferOrder
// was returned from a transaction, and the transaction was committed or rolled back.
func (to *TransferOrder) Update() *TransferOrderUpdateOne {
	return NewTransferOrderClient(to.config).UpdateOne(to)
}

// Unwrap unwraps the TransferOrder entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (to *TransferOrder) Unwrap() *TransferOrder {
	_tx, ok := to.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: TransferOrder is not a transactional entity")
	}
	to.config.driver = _tx.drv
	return to
}

// String implements the fmt.Stringer.
func (to *TransferOrder) String() string {
	var builder strings.Builder
	builder.WriteString("TransferOrder(")
	builder.WriteString(fmt.Sprintf("id=%v, ", to.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", to.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", to.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(to.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(to.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(to.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("source_user_id=")
	builder.WriteString(fmt.Sprintf("%v", to.SourceUserID))
	builder.WriteString(", ")
	builder.WriteString("target_user_id=")
	builder.WriteString(fmt.Sprintf("%v", to.TargetUserID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", to.Status))
	builder.WriteString(", ")
	builder.WriteString("symbol_id=")
	builder.WriteString(fmt.Sprintf("%v", to.SymbolID))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", to.Amount))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", to.Type))
	builder.WriteString(", ")
	builder.WriteString("serial_number=")
	builder.WriteString(to.SerialNumber)
	builder.WriteString(", ")
	builder.WriteString("social_id=")
	builder.WriteString(fmt.Sprintf("%v", to.SocialID))
	builder.WriteString(", ")
	builder.WriteString("third_api_resp=")
	builder.WriteString(to.ThirdAPIResp)
	builder.WriteString(", ")
	builder.WriteString("out_transaction_id=")
	builder.WriteString(to.OutTransactionID)
	builder.WriteString(", ")
	builder.WriteString("operate_user_id=")
	builder.WriteString(fmt.Sprintf("%v", to.OperateUserID))
	builder.WriteString(", ")
	builder.WriteString("gift_type=")
	builder.WriteString(fmt.Sprintf("%v", to.GiftType))
	builder.WriteByte(')')
	return builder.String()
}

// TransferOrders is a parsable slice of TransferOrder.
type TransferOrders []*TransferOrder
