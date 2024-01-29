// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/lotto"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// 抽奖活动表
type Lotto struct {
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
	// 抽奖活动名称
	Name string `json:"name"`
	// 活动总权重
	TotalWeight int64 `json:"total_weight"`
	// 活动开始时间
	StartedAt time.Time `json:"started_at"`
	// 活动结束时间
	EndedAt time.Time `json:"ended_at"`
	// 状态
	Status enums.LottoStatus `json:"status"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LottoQuery when eager-loading is set.
	Edges        LottoEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LottoEdges holds the relations/edges for other nodes in the graph.
type LottoEdges struct {
	// LottoPrizes holds the value of the lotto_prizes edge.
	LottoPrizes []*LottoPrize `json:"lotto_prizes,omitempty"`
	// LottoRecords holds the value of the lotto_records edge.
	LottoRecords []*LottoRecord `json:"lotto_records,omitempty"`
	// LottoUserCounts holds the value of the lotto_user_counts edge.
	LottoUserCounts []*LottoUserCount `json:"lotto_user_counts,omitempty"`
	// LottoGetCountRecords holds the value of the lotto_get_count_records edge.
	LottoGetCountRecords []*LottoGetCountRecord `json:"lotto_get_count_records,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// LottoPrizesOrErr returns the LottoPrizes value or an error if the edge
// was not loaded in eager-loading.
func (e LottoEdges) LottoPrizesOrErr() ([]*LottoPrize, error) {
	if e.loadedTypes[0] {
		return e.LottoPrizes, nil
	}
	return nil, &NotLoadedError{edge: "lotto_prizes"}
}

// LottoRecordsOrErr returns the LottoRecords value or an error if the edge
// was not loaded in eager-loading.
func (e LottoEdges) LottoRecordsOrErr() ([]*LottoRecord, error) {
	if e.loadedTypes[1] {
		return e.LottoRecords, nil
	}
	return nil, &NotLoadedError{edge: "lotto_records"}
}

// LottoUserCountsOrErr returns the LottoUserCounts value or an error if the edge
// was not loaded in eager-loading.
func (e LottoEdges) LottoUserCountsOrErr() ([]*LottoUserCount, error) {
	if e.loadedTypes[2] {
		return e.LottoUserCounts, nil
	}
	return nil, &NotLoadedError{edge: "lotto_user_counts"}
}

// LottoGetCountRecordsOrErr returns the LottoGetCountRecords value or an error if the edge
// was not loaded in eager-loading.
func (e LottoEdges) LottoGetCountRecordsOrErr() ([]*LottoGetCountRecord, error) {
	if e.loadedTypes[3] {
		return e.LottoGetCountRecords, nil
	}
	return nil, &NotLoadedError{edge: "lotto_get_count_records"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Lotto) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case lotto.FieldID, lotto.FieldCreatedBy, lotto.FieldUpdatedBy, lotto.FieldTotalWeight:
			values[i] = new(sql.NullInt64)
		case lotto.FieldName, lotto.FieldStatus:
			values[i] = new(sql.NullString)
		case lotto.FieldCreatedAt, lotto.FieldUpdatedAt, lotto.FieldDeletedAt, lotto.FieldStartedAt, lotto.FieldEndedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Lotto fields.
func (l *Lotto) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case lotto.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int64(value.Int64)
		case lotto.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				l.CreatedBy = value.Int64
			}
		case lotto.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				l.UpdatedBy = value.Int64
			}
		case lotto.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				l.CreatedAt = value.Time
			}
		case lotto.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				l.UpdatedAt = value.Time
			}
		case lotto.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				l.DeletedAt = value.Time
			}
		case lotto.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				l.Name = value.String
			}
		case lotto.FieldTotalWeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_weight", values[i])
			} else if value.Valid {
				l.TotalWeight = value.Int64
			}
		case lotto.FieldStartedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field started_at", values[i])
			} else if value.Valid {
				l.StartedAt = value.Time
			}
		case lotto.FieldEndedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ended_at", values[i])
			} else if value.Valid {
				l.EndedAt = value.Time
			}
		case lotto.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				l.Status = enums.LottoStatus(value.String)
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Lotto.
// This includes values selected through modifiers, order, etc.
func (l *Lotto) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// QueryLottoPrizes queries the "lotto_prizes" edge of the Lotto entity.
func (l *Lotto) QueryLottoPrizes() *LottoPrizeQuery {
	return NewLottoClient(l.config).QueryLottoPrizes(l)
}

// QueryLottoRecords queries the "lotto_records" edge of the Lotto entity.
func (l *Lotto) QueryLottoRecords() *LottoRecordQuery {
	return NewLottoClient(l.config).QueryLottoRecords(l)
}

// QueryLottoUserCounts queries the "lotto_user_counts" edge of the Lotto entity.
func (l *Lotto) QueryLottoUserCounts() *LottoUserCountQuery {
	return NewLottoClient(l.config).QueryLottoUserCounts(l)
}

// QueryLottoGetCountRecords queries the "lotto_get_count_records" edge of the Lotto entity.
func (l *Lotto) QueryLottoGetCountRecords() *LottoGetCountRecordQuery {
	return NewLottoClient(l.config).QueryLottoGetCountRecords(l)
}

// Update returns a builder for updating this Lotto.
// Note that you need to call Lotto.Unwrap() before calling this method if this Lotto
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Lotto) Update() *LottoUpdateOne {
	return NewLottoClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Lotto entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Lotto) Unwrap() *Lotto {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: Lotto is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Lotto) String() string {
	var builder strings.Builder
	builder.WriteString("Lotto(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", l.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", l.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(l.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(l.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(l.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(l.Name)
	builder.WriteString(", ")
	builder.WriteString("total_weight=")
	builder.WriteString(fmt.Sprintf("%v", l.TotalWeight))
	builder.WriteString(", ")
	builder.WriteString("started_at=")
	builder.WriteString(l.StartedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ended_at=")
	builder.WriteString(l.EndedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", l.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Lottos is a parsable slice of Lotto.
type Lottos []*Lotto
