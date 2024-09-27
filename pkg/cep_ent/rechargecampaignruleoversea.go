// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/rechargecampaignruleoversea"
)

// 海外充值活动的规则，死表，与国内区分，国内是赠送充值比例，海外是固定
type RechargeCampaignRuleOversea struct {
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
	// 美元价格，为了大数计算，都用字符串存
	DollarPrice string `json:"dollar_price"`
	// RMB 价格
	RmbPrice string `json:"rmb_price"`
	// RMB 原价
	OriginalRmbPrice string `json:"original_rmb_price"`
	// 总共能获得的脑力值
	TotalCep int64 `json:"total_cep"`
	// 优惠前能获得的脑力值
	BeforeDiscountCep int64 `json:"before_discount_cep"`
	// 优惠力度（现价基于原价的比例）
	DiscountRatio int64 `json:"discount_ratio"`
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RechargeCampaignRuleOversea) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case rechargecampaignruleoversea.FieldID, rechargecampaignruleoversea.FieldCreatedBy, rechargecampaignruleoversea.FieldUpdatedBy, rechargecampaignruleoversea.FieldTotalCep, rechargecampaignruleoversea.FieldBeforeDiscountCep, rechargecampaignruleoversea.FieldDiscountRatio:
			values[i] = new(sql.NullInt64)
		case rechargecampaignruleoversea.FieldDollarPrice, rechargecampaignruleoversea.FieldRmbPrice, rechargecampaignruleoversea.FieldOriginalRmbPrice:
			values[i] = new(sql.NullString)
		case rechargecampaignruleoversea.FieldCreatedAt, rechargecampaignruleoversea.FieldUpdatedAt, rechargecampaignruleoversea.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RechargeCampaignRuleOversea fields.
func (rcro *RechargeCampaignRuleOversea) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case rechargecampaignruleoversea.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rcro.ID = int64(value.Int64)
		case rechargecampaignruleoversea.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				rcro.CreatedBy = value.Int64
			}
		case rechargecampaignruleoversea.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				rcro.UpdatedBy = value.Int64
			}
		case rechargecampaignruleoversea.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				rcro.CreatedAt = value.Time
			}
		case rechargecampaignruleoversea.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				rcro.UpdatedAt = value.Time
			}
		case rechargecampaignruleoversea.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				rcro.DeletedAt = value.Time
			}
		case rechargecampaignruleoversea.FieldDollarPrice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dollar_price", values[i])
			} else if value.Valid {
				rcro.DollarPrice = value.String
			}
		case rechargecampaignruleoversea.FieldRmbPrice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field rmb_price", values[i])
			} else if value.Valid {
				rcro.RmbPrice = value.String
			}
		case rechargecampaignruleoversea.FieldOriginalRmbPrice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field original_rmb_price", values[i])
			} else if value.Valid {
				rcro.OriginalRmbPrice = value.String
			}
		case rechargecampaignruleoversea.FieldTotalCep:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_cep", values[i])
			} else if value.Valid {
				rcro.TotalCep = value.Int64
			}
		case rechargecampaignruleoversea.FieldBeforeDiscountCep:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field before_discount_cep", values[i])
			} else if value.Valid {
				rcro.BeforeDiscountCep = value.Int64
			}
		case rechargecampaignruleoversea.FieldDiscountRatio:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field discount_ratio", values[i])
			} else if value.Valid {
				rcro.DiscountRatio = value.Int64
			}
		default:
			rcro.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RechargeCampaignRuleOversea.
// This includes values selected through modifiers, order, etc.
func (rcro *RechargeCampaignRuleOversea) Value(name string) (ent.Value, error) {
	return rcro.selectValues.Get(name)
}

// Update returns a builder for updating this RechargeCampaignRuleOversea.
// Note that you need to call RechargeCampaignRuleOversea.Unwrap() before calling this method if this RechargeCampaignRuleOversea
// was returned from a transaction, and the transaction was committed or rolled back.
func (rcro *RechargeCampaignRuleOversea) Update() *RechargeCampaignRuleOverseaUpdateOne {
	return NewRechargeCampaignRuleOverseaClient(rcro.config).UpdateOne(rcro)
}

// Unwrap unwraps the RechargeCampaignRuleOversea entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rcro *RechargeCampaignRuleOversea) Unwrap() *RechargeCampaignRuleOversea {
	_tx, ok := rcro.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: RechargeCampaignRuleOversea is not a transactional entity")
	}
	rcro.config.driver = _tx.drv
	return rcro
}

// String implements the fmt.Stringer.
func (rcro *RechargeCampaignRuleOversea) String() string {
	var builder strings.Builder
	builder.WriteString("RechargeCampaignRuleOversea(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rcro.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", rcro.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", rcro.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(rcro.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(rcro.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(rcro.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("dollar_price=")
	builder.WriteString(rcro.DollarPrice)
	builder.WriteString(", ")
	builder.WriteString("rmb_price=")
	builder.WriteString(rcro.RmbPrice)
	builder.WriteString(", ")
	builder.WriteString("original_rmb_price=")
	builder.WriteString(rcro.OriginalRmbPrice)
	builder.WriteString(", ")
	builder.WriteString("total_cep=")
	builder.WriteString(fmt.Sprintf("%v", rcro.TotalCep))
	builder.WriteString(", ")
	builder.WriteString("before_discount_cep=")
	builder.WriteString(fmt.Sprintf("%v", rcro.BeforeDiscountCep))
	builder.WriteString(", ")
	builder.WriteString("discount_ratio=")
	builder.WriteString(fmt.Sprintf("%v", rcro.DiscountRatio))
	builder.WriteByte(')')
	return builder.String()
}

// RechargeCampaignRuleOverseas is a parsable slice of RechargeCampaignRuleOversea.
type RechargeCampaignRuleOverseas []*RechargeCampaignRuleOversea
