// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/hmackeypair"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/mission"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// Mission is the model entity for the Mission schema.
type Mission struct {
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
	// 任务类型
	Type enums.MissionType `json:"type"`
	// 任务的请求参数体
	Body string `json:"body"`
	// 回调地址，空字符串表示不回调
	CallBackURL string `json:"call_back_url"`
	// 任务状态
	Status enums.MissionStatus `json:"status"`
	// 任务结果，pending 表示还没有结果
	Result enums.MissionResult `json:"result"`
	// 任务结果资源位置列表序列化
	ResultUrls []string `json:"-"`
	// 任务创建者的密钥对 ID
	KeyPairID int64 `json:"key_pair_id"`
	// 任务批次号
	MissionBatchNumber string `json:"mission_batch_number"`
	// 最低可接显卡
	GpuVersion enums.GpuVersion `json:"gpu_version"`
	// 任务单价，按次就是 unit_cep/次，按时就是 unit_cep/分钟
	UnitCep int64 `json:"unit_cep"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MissionQuery when eager-loading is set.
	Edges        MissionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MissionEdges holds the relations/edges for other nodes in the graph.
type MissionEdges struct {
	// MissionKeyPairs holds the value of the mission_key_pairs edge.
	MissionKeyPairs []*MissionKeyPair `json:"mission_key_pairs,omitempty"`
	// KeyPair holds the value of the key_pair edge.
	KeyPair *HmacKeyPair `json:"key_pair,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MissionKeyPairsOrErr returns the MissionKeyPairs value or an error if the edge
// was not loaded in eager-loading.
func (e MissionEdges) MissionKeyPairsOrErr() ([]*MissionKeyPair, error) {
	if e.loadedTypes[0] {
		return e.MissionKeyPairs, nil
	}
	return nil, &NotLoadedError{edge: "mission_key_pairs"}
}

// KeyPairOrErr returns the KeyPair value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionEdges) KeyPairOrErr() (*HmacKeyPair, error) {
	if e.loadedTypes[1] {
		if e.KeyPair == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: hmackeypair.Label}
		}
		return e.KeyPair, nil
	}
	return nil, &NotLoadedError{edge: "key_pair"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Mission) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case mission.FieldResultUrls:
			values[i] = new([]byte)
		case mission.FieldID, mission.FieldCreatedBy, mission.FieldUpdatedBy, mission.FieldKeyPairID, mission.FieldUnitCep:
			values[i] = new(sql.NullInt64)
		case mission.FieldType, mission.FieldBody, mission.FieldCallBackURL, mission.FieldStatus, mission.FieldResult, mission.FieldMissionBatchNumber, mission.FieldGpuVersion:
			values[i] = new(sql.NullString)
		case mission.FieldCreatedAt, mission.FieldUpdatedAt, mission.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Mission fields.
func (m *Mission) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int64(value.Int64)
		case mission.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				m.CreatedBy = value.Int64
			}
		case mission.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				m.UpdatedBy = value.Int64
			}
		case mission.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case mission.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case mission.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				m.DeletedAt = value.Time
			}
		case mission.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				m.Type = enums.MissionType(value.String)
			}
		case mission.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				m.Body = value.String
			}
		case mission.FieldCallBackURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field call_back_url", values[i])
			} else if value.Valid {
				m.CallBackURL = value.String
			}
		case mission.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				m.Status = enums.MissionStatus(value.String)
			}
		case mission.FieldResult:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field result", values[i])
			} else if value.Valid {
				m.Result = enums.MissionResult(value.String)
			}
		case mission.FieldResultUrls:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field result_urls", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &m.ResultUrls); err != nil {
					return fmt.Errorf("unmarshal field result_urls: %w", err)
				}
			}
		case mission.FieldKeyPairID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field key_pair_id", values[i])
			} else if value.Valid {
				m.KeyPairID = value.Int64
			}
		case mission.FieldMissionBatchNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mission_batch_number", values[i])
			} else if value.Valid {
				m.MissionBatchNumber = value.String
			}
		case mission.FieldGpuVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gpu_version", values[i])
			} else if value.Valid {
				m.GpuVersion = enums.GpuVersion(value.String)
			}
		case mission.FieldUnitCep:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field unit_cep", values[i])
			} else if value.Valid {
				m.UnitCep = value.Int64
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Mission.
// This includes values selected through modifiers, order, etc.
func (m *Mission) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryMissionKeyPairs queries the "mission_key_pairs" edge of the Mission entity.
func (m *Mission) QueryMissionKeyPairs() *MissionKeyPairQuery {
	return NewMissionClient(m.config).QueryMissionKeyPairs(m)
}

// QueryKeyPair queries the "key_pair" edge of the Mission entity.
func (m *Mission) QueryKeyPair() *HmacKeyPairQuery {
	return NewMissionClient(m.config).QueryKeyPair(m)
}

// Update returns a builder for updating this Mission.
// Note that you need to call Mission.Unwrap() before calling this method if this Mission
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Mission) Update() *MissionUpdateOne {
	return NewMissionClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Mission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Mission) Unwrap() *Mission {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: Mission is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Mission) String() string {
	var builder strings.Builder
	builder.WriteString("Mission(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", m.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", m.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(m.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", m.Type))
	builder.WriteString(", ")
	builder.WriteString("body=")
	builder.WriteString(m.Body)
	builder.WriteString(", ")
	builder.WriteString("call_back_url=")
	builder.WriteString(m.CallBackURL)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", m.Status))
	builder.WriteString(", ")
	builder.WriteString("result=")
	builder.WriteString(fmt.Sprintf("%v", m.Result))
	builder.WriteString(", ")
	builder.WriteString("result_urls=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("key_pair_id=")
	builder.WriteString(fmt.Sprintf("%v", m.KeyPairID))
	builder.WriteString(", ")
	builder.WriteString("mission_batch_number=")
	builder.WriteString(m.MissionBatchNumber)
	builder.WriteString(", ")
	builder.WriteString("gpu_version=")
	builder.WriteString(fmt.Sprintf("%v", m.GpuVersion))
	builder.WriteString(", ")
	builder.WriteString("unit_cep=")
	builder.WriteString(fmt.Sprintf("%v", m.UnitCep))
	builder.WriteByte(')')
	return builder.String()
}

// Missions is a parsable slice of Mission.
type Missions []*Mission
