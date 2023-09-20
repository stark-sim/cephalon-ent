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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionkeypair"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// 原任务执行情况，被 mission_productions 取代
type MissionKeyPair struct {
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
	// 任务 ID
	MissionID int64 `json:"mission_id"`
	// 密钥对 ID
	KeyPairID int64 `json:"key_pair_id"`
	// 任务开始时刻
	StartedAt time.Time `json:"started_at"`
	// 任务完成时刻
	FinishedAt time.Time `json:"finished_at"`
	// 任务结果
	Result enums.MissionResult `json:"result"`
	// 领到任务的设备 ID
	DeviceID int64 `json:"device_id"`
	// 任务结果资源位置列表序列化
	ResultUrls []string `json:"-"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MissionKeyPairQuery when eager-loading is set.
	Edges        MissionKeyPairEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MissionKeyPairEdges holds the relations/edges for other nodes in the graph.
type MissionKeyPairEdges struct {
	// Mission holds the value of the mission edge.
	Mission *Mission `json:"mission,omitempty"`
	// KeyPair holds the value of the key_pair edge.
	KeyPair *HmacKeyPair `json:"key_pair,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MissionOrErr returns the Mission value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionKeyPairEdges) MissionOrErr() (*Mission, error) {
	if e.loadedTypes[0] {
		if e.Mission == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: mission.Label}
		}
		return e.Mission, nil
	}
	return nil, &NotLoadedError{edge: "mission"}
}

// KeyPairOrErr returns the KeyPair value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionKeyPairEdges) KeyPairOrErr() (*HmacKeyPair, error) {
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
func (*MissionKeyPair) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case missionkeypair.FieldResultUrls:
			values[i] = new([]byte)
		case missionkeypair.FieldID, missionkeypair.FieldCreatedBy, missionkeypair.FieldUpdatedBy, missionkeypair.FieldMissionID, missionkeypair.FieldKeyPairID, missionkeypair.FieldDeviceID:
			values[i] = new(sql.NullInt64)
		case missionkeypair.FieldResult:
			values[i] = new(sql.NullString)
		case missionkeypair.FieldCreatedAt, missionkeypair.FieldUpdatedAt, missionkeypair.FieldDeletedAt, missionkeypair.FieldStartedAt, missionkeypair.FieldFinishedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MissionKeyPair fields.
func (mkp *MissionKeyPair) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case missionkeypair.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mkp.ID = int64(value.Int64)
		case missionkeypair.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				mkp.CreatedBy = value.Int64
			}
		case missionkeypair.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				mkp.UpdatedBy = value.Int64
			}
		case missionkeypair.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mkp.CreatedAt = value.Time
			}
		case missionkeypair.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mkp.UpdatedAt = value.Time
			}
		case missionkeypair.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				mkp.DeletedAt = value.Time
			}
		case missionkeypair.FieldMissionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mission_id", values[i])
			} else if value.Valid {
				mkp.MissionID = value.Int64
			}
		case missionkeypair.FieldKeyPairID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field key_pair_id", values[i])
			} else if value.Valid {
				mkp.KeyPairID = value.Int64
			}
		case missionkeypair.FieldStartedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field started_at", values[i])
			} else if value.Valid {
				mkp.StartedAt = value.Time
			}
		case missionkeypair.FieldFinishedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field finished_at", values[i])
			} else if value.Valid {
				mkp.FinishedAt = value.Time
			}
		case missionkeypair.FieldResult:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field result", values[i])
			} else if value.Valid {
				mkp.Result = enums.MissionResult(value.String)
			}
		case missionkeypair.FieldDeviceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field device_id", values[i])
			} else if value.Valid {
				mkp.DeviceID = value.Int64
			}
		case missionkeypair.FieldResultUrls:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field result_urls", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &mkp.ResultUrls); err != nil {
					return fmt.Errorf("unmarshal field result_urls: %w", err)
				}
			}
		default:
			mkp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MissionKeyPair.
// This includes values selected through modifiers, order, etc.
func (mkp *MissionKeyPair) Value(name string) (ent.Value, error) {
	return mkp.selectValues.Get(name)
}

// QueryMission queries the "mission" edge of the MissionKeyPair entity.
func (mkp *MissionKeyPair) QueryMission() *MissionQuery {
	return NewMissionKeyPairClient(mkp.config).QueryMission(mkp)
}

// QueryKeyPair queries the "key_pair" edge of the MissionKeyPair entity.
func (mkp *MissionKeyPair) QueryKeyPair() *HmacKeyPairQuery {
	return NewMissionKeyPairClient(mkp.config).QueryKeyPair(mkp)
}

// Update returns a builder for updating this MissionKeyPair.
// Note that you need to call MissionKeyPair.Unwrap() before calling this method if this MissionKeyPair
// was returned from a transaction, and the transaction was committed or rolled back.
func (mkp *MissionKeyPair) Update() *MissionKeyPairUpdateOne {
	return NewMissionKeyPairClient(mkp.config).UpdateOne(mkp)
}

// Unwrap unwraps the MissionKeyPair entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mkp *MissionKeyPair) Unwrap() *MissionKeyPair {
	_tx, ok := mkp.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: MissionKeyPair is not a transactional entity")
	}
	mkp.config.driver = _tx.drv
	return mkp
}

// String implements the fmt.Stringer.
func (mkp *MissionKeyPair) String() string {
	var builder strings.Builder
	builder.WriteString("MissionKeyPair(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mkp.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", mkp.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", mkp.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(mkp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(mkp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(mkp.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("mission_id=")
	builder.WriteString(fmt.Sprintf("%v", mkp.MissionID))
	builder.WriteString(", ")
	builder.WriteString("key_pair_id=")
	builder.WriteString(fmt.Sprintf("%v", mkp.KeyPairID))
	builder.WriteString(", ")
	builder.WriteString("started_at=")
	builder.WriteString(mkp.StartedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("finished_at=")
	builder.WriteString(mkp.FinishedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("result=")
	builder.WriteString(fmt.Sprintf("%v", mkp.Result))
	builder.WriteString(", ")
	builder.WriteString("device_id=")
	builder.WriteString(fmt.Sprintf("%v", mkp.DeviceID))
	builder.WriteString(", ")
	builder.WriteString("result_urls=<sensitive>")
	builder.WriteByte(')')
	return builder.String()
}

// MissionKeyPairs is a parsable slice of MissionKeyPair.
type MissionKeyPairs []*MissionKeyPair
