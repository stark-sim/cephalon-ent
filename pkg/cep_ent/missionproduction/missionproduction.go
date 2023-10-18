// Code generated by ent, DO NOT EDIT.

package missionproduction

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

const (
	// Label holds the string label denoting the missionproduction type in the database.
	Label = "mission_production"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldMissionID holds the string denoting the mission_id field in the database.
	FieldMissionID = "mission_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldStartedAt holds the string denoting the started_at field in the database.
	FieldStartedAt = "started_at"
	// FieldFinishedAt holds the string denoting the finished_at field in the database.
	FieldFinishedAt = "finished_at"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldDeviceID holds the string denoting the device_id field in the database.
	FieldDeviceID = "device_id"
	// FieldGpuVersion holds the string denoting the gpu_version field in the database.
	FieldGpuVersion = "gpu_version"
	// FieldUrls holds the string denoting the urls field in the database.
	FieldUrls = "urls"
	// FieldRespStatusCode holds the string denoting the resp_status_code field in the database.
	FieldRespStatusCode = "resp_status_code"
	// FieldRespBody holds the string denoting the resp_body field in the database.
	FieldRespBody = "resp_body"
	// EdgeMission holds the string denoting the mission edge name in mutations.
	EdgeMission = "mission"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeMissionProduceOrder holds the string denoting the mission_produce_order edge name in mutations.
	EdgeMissionProduceOrder = "mission_produce_order"
	// Table holds the table name of the missionproduction in the database.
	Table = "mission_productions"
	// MissionTable is the table that holds the mission relation/edge.
	MissionTable = "mission_productions"
	// MissionInverseTable is the table name for the Mission entity.
	// It exists in this package in order to avoid circular dependency with the "mission" package.
	MissionInverseTable = "missions"
	// MissionColumn is the table column denoting the mission relation/edge.
	MissionColumn = "mission_id"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "mission_productions"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// MissionProduceOrderTable is the table that holds the mission_produce_order relation/edge.
	MissionProduceOrderTable = "mission_produce_orders"
	// MissionProduceOrderInverseTable is the table name for the MissionProduceOrder entity.
	// It exists in this package in order to avoid circular dependency with the "missionproduceorder" package.
	MissionProduceOrderInverseTable = "mission_produce_orders"
	// MissionProduceOrderColumn is the table column denoting the mission_produce_order relation/edge.
	MissionProduceOrderColumn = "mission_production_id"
)

// Columns holds all SQL columns for missionproduction fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldMissionID,
	FieldUserID,
	FieldStartedAt,
	FieldFinishedAt,
	FieldState,
	FieldDeviceID,
	FieldGpuVersion,
	FieldUrls,
	FieldRespStatusCode,
	FieldRespBody,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedBy holds the default value on creation for the "created_by" field.
	DefaultCreatedBy int64
	// DefaultUpdatedBy holds the default value on creation for the "updated_by" field.
	DefaultUpdatedBy int64
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt time.Time
	// DefaultStartedAt holds the default value on creation for the "started_at" field.
	DefaultStartedAt time.Time
	// DefaultFinishedAt holds the default value on creation for the "finished_at" field.
	DefaultFinishedAt time.Time
	// DefaultDeviceID holds the default value on creation for the "device_id" field.
	DefaultDeviceID int64
	// DefaultUrls holds the default value on creation for the "urls" field.
	DefaultUrls string
	// DefaultRespStatusCode holds the default value on creation for the "resp_status_code" field.
	DefaultRespStatusCode int32
	// DefaultRespBody holds the default value on creation for the "resp_body" field.
	DefaultRespBody string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

const DefaultState enums.MissionState = "unknown"

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(s enums.MissionState) error {
	switch s {
	case "unknown", "waiting", "canceled", "doing", "supplying", "closing", "succeed", "failed":
		return nil
	default:
		return fmt.Errorf("missionproduction: invalid enum value for state field: %q", s)
	}
}

const DefaultGpuVersion enums.GpuVersion = "unknown"

// GpuVersionValidator is a validator for the "gpu_version" field enum values. It is called by the builders before save.
func GpuVersionValidator(gv enums.GpuVersion) error {
	switch gv {
	case "unknown", "RTX2060", "RTX2060Ti", "RTX2070", "RTX2070Ti", "RTX2080", "RTX2080Ti", "RTX3060", "RTX3060Ti", "RTX3070", "RTX3070Ti", "RTX3080", "RTX3080Ti", "RTX3090", "RTX3090Ti", "RTX4060", "RTX4060Ti", "RTX4070", "RTX4070Ti", "RTX4080", "RTX4090", "A800", "A100", "V100":
		return nil
	default:
		return fmt.Errorf("missionproduction: invalid enum value for gpu_version field: %q", gv)
	}
}

// OrderOption defines the ordering options for the MissionProduction queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByMissionID orders the results by the mission_id field.
func ByMissionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByStartedAt orders the results by the started_at field.
func ByStartedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartedAt, opts...).ToFunc()
}

// ByFinishedAt orders the results by the finished_at field.
func ByFinishedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFinishedAt, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// ByDeviceID orders the results by the device_id field.
func ByDeviceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeviceID, opts...).ToFunc()
}

// ByGpuVersion orders the results by the gpu_version field.
func ByGpuVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGpuVersion, opts...).ToFunc()
}

// ByUrls orders the results by the urls field.
func ByUrls(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUrls, opts...).ToFunc()
}

// ByRespStatusCode orders the results by the resp_status_code field.
func ByRespStatusCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRespStatusCode, opts...).ToFunc()
}

// ByRespBody orders the results by the resp_body field.
func ByRespBody(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRespBody, opts...).ToFunc()
}

// ByMissionField orders the results by mission field.
func ByMissionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByMissionProduceOrderField orders the results by mission_produce_order field.
func ByMissionProduceOrderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionProduceOrderStep(), sql.OrderByField(field, opts...))
	}
}
func newMissionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MissionTable, MissionColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newMissionProduceOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionProduceOrderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, MissionProduceOrderTable, MissionProduceOrderColumn),
	)
}