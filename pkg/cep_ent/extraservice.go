// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/extraservice"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// 附加服务表
type ExtraService struct {
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
	// 附加服务名称
	Name string `json:"name"`
	// 附加服务类型
	ExtraServiceType enums.ExtraServiceType `json:"extra_service_type"`
	// 附加服务购买时间开始，为空表示永久有效
	StartedAt *time.Time `json:"started_at"`
	// 附加服务购买时间结束，为空表示永久有效
	FinishedAt *time.Time `json:"finished_at"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ExtraServiceQuery when eager-loading is set.
	Edges                  ExtraServiceEdges `json:"edges"`
	mission_extra_services *int64
	selectValues           sql.SelectValues
}

// ExtraServiceEdges holds the relations/edges for other nodes in the graph.
type ExtraServiceEdges struct {
	// Missions holds the value of the missions edge.
	Missions []*Mission `json:"missions,omitempty"`
	// MissionExtraServices holds the value of the mission_extra_services edge.
	MissionExtraServices []*MissionExtraService `json:"mission_extra_services,omitempty"`
	// ExtraServicePrices holds the value of the extra_service_prices edge.
	ExtraServicePrices []*ExtraServicePrice `json:"extra_service_prices,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// MissionsOrErr returns the Missions value or an error if the edge
// was not loaded in eager-loading.
func (e ExtraServiceEdges) MissionsOrErr() ([]*Mission, error) {
	if e.loadedTypes[0] {
		return e.Missions, nil
	}
	return nil, &NotLoadedError{edge: "missions"}
}

// MissionExtraServicesOrErr returns the MissionExtraServices value or an error if the edge
// was not loaded in eager-loading.
func (e ExtraServiceEdges) MissionExtraServicesOrErr() ([]*MissionExtraService, error) {
	if e.loadedTypes[1] {
		return e.MissionExtraServices, nil
	}
	return nil, &NotLoadedError{edge: "mission_extra_services"}
}

// ExtraServicePricesOrErr returns the ExtraServicePrices value or an error if the edge
// was not loaded in eager-loading.
func (e ExtraServiceEdges) ExtraServicePricesOrErr() ([]*ExtraServicePrice, error) {
	if e.loadedTypes[2] {
		return e.ExtraServicePrices, nil
	}
	return nil, &NotLoadedError{edge: "extra_service_prices"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ExtraService) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case extraservice.FieldID, extraservice.FieldCreatedBy, extraservice.FieldUpdatedBy:
			values[i] = new(sql.NullInt64)
		case extraservice.FieldName, extraservice.FieldExtraServiceType:
			values[i] = new(sql.NullString)
		case extraservice.FieldCreatedAt, extraservice.FieldUpdatedAt, extraservice.FieldDeletedAt, extraservice.FieldStartedAt, extraservice.FieldFinishedAt:
			values[i] = new(sql.NullTime)
		case extraservice.ForeignKeys[0]: // mission_extra_services
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ExtraService fields.
func (es *ExtraService) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case extraservice.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			es.ID = int64(value.Int64)
		case extraservice.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				es.CreatedBy = value.Int64
			}
		case extraservice.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				es.UpdatedBy = value.Int64
			}
		case extraservice.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				es.CreatedAt = value.Time
			}
		case extraservice.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				es.UpdatedAt = value.Time
			}
		case extraservice.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				es.DeletedAt = value.Time
			}
		case extraservice.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				es.Name = value.String
			}
		case extraservice.FieldExtraServiceType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field extra_service_type", values[i])
			} else if value.Valid {
				es.ExtraServiceType = enums.ExtraServiceType(value.String)
			}
		case extraservice.FieldStartedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field started_at", values[i])
			} else if value.Valid {
				es.StartedAt = new(time.Time)
				*es.StartedAt = value.Time
			}
		case extraservice.FieldFinishedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field finished_at", values[i])
			} else if value.Valid {
				es.FinishedAt = new(time.Time)
				*es.FinishedAt = value.Time
			}
		case extraservice.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field mission_extra_services", value)
			} else if value.Valid {
				es.mission_extra_services = new(int64)
				*es.mission_extra_services = int64(value.Int64)
			}
		default:
			es.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ExtraService.
// This includes values selected through modifiers, order, etc.
func (es *ExtraService) Value(name string) (ent.Value, error) {
	return es.selectValues.Get(name)
}

// QueryMissions queries the "missions" edge of the ExtraService entity.
func (es *ExtraService) QueryMissions() *MissionQuery {
	return NewExtraServiceClient(es.config).QueryMissions(es)
}

// QueryMissionExtraServices queries the "mission_extra_services" edge of the ExtraService entity.
func (es *ExtraService) QueryMissionExtraServices() *MissionExtraServiceQuery {
	return NewExtraServiceClient(es.config).QueryMissionExtraServices(es)
}

// QueryExtraServicePrices queries the "extra_service_prices" edge of the ExtraService entity.
func (es *ExtraService) QueryExtraServicePrices() *ExtraServicePriceQuery {
	return NewExtraServiceClient(es.config).QueryExtraServicePrices(es)
}

// Update returns a builder for updating this ExtraService.
// Note that you need to call ExtraService.Unwrap() before calling this method if this ExtraService
// was returned from a transaction, and the transaction was committed or rolled back.
func (es *ExtraService) Update() *ExtraServiceUpdateOne {
	return NewExtraServiceClient(es.config).UpdateOne(es)
}

// Unwrap unwraps the ExtraService entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (es *ExtraService) Unwrap() *ExtraService {
	_tx, ok := es.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: ExtraService is not a transactional entity")
	}
	es.config.driver = _tx.drv
	return es
}

// String implements the fmt.Stringer.
func (es *ExtraService) String() string {
	var builder strings.Builder
	builder.WriteString("ExtraService(")
	builder.WriteString(fmt.Sprintf("id=%v, ", es.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", es.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", es.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(es.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(es.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(es.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(es.Name)
	builder.WriteString(", ")
	builder.WriteString("extra_service_type=")
	builder.WriteString(fmt.Sprintf("%v", es.ExtraServiceType))
	builder.WriteString(", ")
	if v := es.StartedAt; v != nil {
		builder.WriteString("started_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := es.FinishedAt; v != nil {
		builder.WriteString("finished_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// ExtraServices is a parsable slice of ExtraService.
type ExtraServices []*ExtraService
