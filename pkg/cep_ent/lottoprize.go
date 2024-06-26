// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/lotto"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/lottoprize"
)

// 抽奖活动奖品表
type LottoPrize struct {
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
	// 奖品等级名称
	LevelName string `json:"level_name"`
	// 奖品等级权重
	Weight int64 `json:"weight"`
	// 奖品名称
	Name string `json:"name"`
	// 状态
	Status lottoprize.Status `json:"status"`
	// 类型
	Type lottoprize.Type `json:"type"`
	// 类型为 get_cep 时，cep 的数量
	CepAmount int64 `json:"cep_amount"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LottoPrizeQuery when eager-loading is set.
	Edges        LottoPrizeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LottoPrizeEdges holds the relations/edges for other nodes in the graph.
type LottoPrizeEdges struct {
	// Lotto holds the value of the lotto edge.
	Lotto *Lotto `json:"lotto,omitempty"`
	// LottoRecords holds the value of the lotto_records edge.
	LottoRecords []*LottoRecord `json:"lotto_records,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// LottoOrErr returns the Lotto value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LottoPrizeEdges) LottoOrErr() (*Lotto, error) {
	if e.loadedTypes[0] {
		if e.Lotto == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: lotto.Label}
		}
		return e.Lotto, nil
	}
	return nil, &NotLoadedError{edge: "lotto"}
}

// LottoRecordsOrErr returns the LottoRecords value or an error if the edge
// was not loaded in eager-loading.
func (e LottoPrizeEdges) LottoRecordsOrErr() ([]*LottoRecord, error) {
	if e.loadedTypes[1] {
		return e.LottoRecords, nil
	}
	return nil, &NotLoadedError{edge: "lotto_records"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LottoPrize) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case lottoprize.FieldID, lottoprize.FieldCreatedBy, lottoprize.FieldUpdatedBy, lottoprize.FieldLottoID, lottoprize.FieldWeight, lottoprize.FieldCepAmount:
			values[i] = new(sql.NullInt64)
		case lottoprize.FieldLevelName, lottoprize.FieldName, lottoprize.FieldStatus, lottoprize.FieldType:
			values[i] = new(sql.NullString)
		case lottoprize.FieldCreatedAt, lottoprize.FieldUpdatedAt, lottoprize.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LottoPrize fields.
func (lp *LottoPrize) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case lottoprize.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			lp.ID = int64(value.Int64)
		case lottoprize.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				lp.CreatedBy = value.Int64
			}
		case lottoprize.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				lp.UpdatedBy = value.Int64
			}
		case lottoprize.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				lp.CreatedAt = value.Time
			}
		case lottoprize.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				lp.UpdatedAt = value.Time
			}
		case lottoprize.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				lp.DeletedAt = value.Time
			}
		case lottoprize.FieldLottoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field lotto_id", values[i])
			} else if value.Valid {
				lp.LottoID = value.Int64
			}
		case lottoprize.FieldLevelName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field level_name", values[i])
			} else if value.Valid {
				lp.LevelName = value.String
			}
		case lottoprize.FieldWeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				lp.Weight = value.Int64
			}
		case lottoprize.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				lp.Name = value.String
			}
		case lottoprize.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				lp.Status = lottoprize.Status(value.String)
			}
		case lottoprize.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				lp.Type = lottoprize.Type(value.String)
			}
		case lottoprize.FieldCepAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cep_amount", values[i])
			} else if value.Valid {
				lp.CepAmount = value.Int64
			}
		default:
			lp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the LottoPrize.
// This includes values selected through modifiers, order, etc.
func (lp *LottoPrize) Value(name string) (ent.Value, error) {
	return lp.selectValues.Get(name)
}

// QueryLotto queries the "lotto" edge of the LottoPrize entity.
func (lp *LottoPrize) QueryLotto() *LottoQuery {
	return NewLottoPrizeClient(lp.config).QueryLotto(lp)
}

// QueryLottoRecords queries the "lotto_records" edge of the LottoPrize entity.
func (lp *LottoPrize) QueryLottoRecords() *LottoRecordQuery {
	return NewLottoPrizeClient(lp.config).QueryLottoRecords(lp)
}

// Update returns a builder for updating this LottoPrize.
// Note that you need to call LottoPrize.Unwrap() before calling this method if this LottoPrize
// was returned from a transaction, and the transaction was committed or rolled back.
func (lp *LottoPrize) Update() *LottoPrizeUpdateOne {
	return NewLottoPrizeClient(lp.config).UpdateOne(lp)
}

// Unwrap unwraps the LottoPrize entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (lp *LottoPrize) Unwrap() *LottoPrize {
	_tx, ok := lp.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: LottoPrize is not a transactional entity")
	}
	lp.config.driver = _tx.drv
	return lp
}

// String implements the fmt.Stringer.
func (lp *LottoPrize) String() string {
	var builder strings.Builder
	builder.WriteString("LottoPrize(")
	builder.WriteString(fmt.Sprintf("id=%v, ", lp.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", lp.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", lp.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(lp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(lp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(lp.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("lotto_id=")
	builder.WriteString(fmt.Sprintf("%v", lp.LottoID))
	builder.WriteString(", ")
	builder.WriteString("level_name=")
	builder.WriteString(lp.LevelName)
	builder.WriteString(", ")
	builder.WriteString("weight=")
	builder.WriteString(fmt.Sprintf("%v", lp.Weight))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(lp.Name)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", lp.Status))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", lp.Type))
	builder.WriteString(", ")
	builder.WriteString("cep_amount=")
	builder.WriteString(fmt.Sprintf("%v", lp.CepAmount))
	builder.WriteByte(')')
	return builder.String()
}

// LottoPrizes is a parsable slice of LottoPrize.
type LottoPrizes []*LottoPrize
