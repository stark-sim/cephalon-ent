// Code generated by ent, DO NOT EDIT.

package extraserviceorder

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

const (
	// Label holds the string label denoting the extraserviceorder type in the database.
	Label = "extra_service_order"
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
	// FieldMissionOrderID holds the string denoting the mission_order_id field in the database.
	FieldMissionOrderID = "mission_order_id"
	// FieldExtraServiceBillingType holds the string denoting the extra_service_billing_type field in the database.
	FieldExtraServiceBillingType = "extra_service_billing_type"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldSymbolID holds the string denoting the symbol_id field in the database.
	FieldSymbolID = "symbol_id"
	// FieldUnitCep holds the string denoting the unit_cep field in the database.
	FieldUnitCep = "unit_cep"
	// FieldExtraServiceType holds the string denoting the extra_service_type field in the database.
	FieldExtraServiceType = "extra_service_type"
	// FieldBuyDuration holds the string denoting the buy_duration field in the database.
	FieldBuyDuration = "buy_duration"
	// FieldStartedAt holds the string denoting the started_at field in the database.
	FieldStartedAt = "started_at"
	// FieldFinishedAt holds the string denoting the finished_at field in the database.
	FieldFinishedAt = "finished_at"
	// FieldPlanStartedAt holds the string denoting the plan_started_at field in the database.
	FieldPlanStartedAt = "plan_started_at"
	// FieldPlanFinishedAt holds the string denoting the plan_finished_at field in the database.
	FieldPlanFinishedAt = "plan_finished_at"
	// FieldMissionBatchID holds the string denoting the mission_batch_id field in the database.
	FieldMissionBatchID = "mission_batch_id"
	// FieldSettledAmount holds the string denoting the settled_amount field in the database.
	FieldSettledAmount = "settled_amount"
	// FieldSettledCount holds the string denoting the settled_count field in the database.
	FieldSettledCount = "settled_count"
	// FieldTotalSettleCount holds the string denoting the total_settle_count field in the database.
	FieldTotalSettleCount = "total_settle_count"
	// FieldLatelySettledAt holds the string denoting the lately_settled_at field in the database.
	FieldLatelySettledAt = "lately_settled_at"
	// EdgeMission holds the string denoting the mission edge name in mutations.
	EdgeMission = "mission"
	// EdgeMissionOrder holds the string denoting the mission_order edge name in mutations.
	EdgeMissionOrder = "mission_order"
	// EdgeSymbol holds the string denoting the symbol edge name in mutations.
	EdgeSymbol = "symbol"
	// EdgeMissionBatch holds the string denoting the mission_batch edge name in mutations.
	EdgeMissionBatch = "mission_batch"
	// Table holds the table name of the extraserviceorder in the database.
	Table = "extra_service_orders"
	// MissionTable is the table that holds the mission relation/edge.
	MissionTable = "extra_service_orders"
	// MissionInverseTable is the table name for the Mission entity.
	// It exists in this package in order to avoid circular dependency with the "mission" package.
	MissionInverseTable = "missions"
	// MissionColumn is the table column denoting the mission relation/edge.
	MissionColumn = "mission_id"
	// MissionOrderTable is the table that holds the mission_order relation/edge.
	MissionOrderTable = "extra_service_orders"
	// MissionOrderInverseTable is the table name for the MissionOrder entity.
	// It exists in this package in order to avoid circular dependency with the "missionorder" package.
	MissionOrderInverseTable = "mission_orders"
	// MissionOrderColumn is the table column denoting the mission_order relation/edge.
	MissionOrderColumn = "mission_order_id"
	// SymbolTable is the table that holds the symbol relation/edge.
	SymbolTable = "extra_service_orders"
	// SymbolInverseTable is the table name for the Symbol entity.
	// It exists in this package in order to avoid circular dependency with the "symbol" package.
	SymbolInverseTable = "symbols"
	// SymbolColumn is the table column denoting the symbol relation/edge.
	SymbolColumn = "symbol_id"
	// MissionBatchTable is the table that holds the mission_batch relation/edge.
	MissionBatchTable = "extra_service_orders"
	// MissionBatchInverseTable is the table name for the MissionBatch entity.
	// It exists in this package in order to avoid circular dependency with the "missionbatch" package.
	MissionBatchInverseTable = "mission_batches"
	// MissionBatchColumn is the table column denoting the mission_batch relation/edge.
	MissionBatchColumn = "mission_batch_id"
)

// Columns holds all SQL columns for extraserviceorder fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldMissionID,
	FieldMissionOrderID,
	FieldExtraServiceBillingType,
	FieldAmount,
	FieldSymbolID,
	FieldUnitCep,
	FieldExtraServiceType,
	FieldBuyDuration,
	FieldStartedAt,
	FieldFinishedAt,
	FieldPlanStartedAt,
	FieldPlanFinishedAt,
	FieldMissionBatchID,
	FieldSettledAmount,
	FieldSettledCount,
	FieldTotalSettleCount,
	FieldLatelySettledAt,
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
	// DefaultMissionID holds the default value on creation for the "mission_id" field.
	DefaultMissionID int64
	// DefaultMissionOrderID holds the default value on creation for the "mission_order_id" field.
	DefaultMissionOrderID int64
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount int64
	// DefaultSymbolID holds the default value on creation for the "symbol_id" field.
	DefaultSymbolID int64
	// DefaultUnitCep holds the default value on creation for the "unit_cep" field.
	DefaultUnitCep int64
	// DefaultBuyDuration holds the default value on creation for the "buy_duration" field.
	DefaultBuyDuration int64
	// DefaultStartedAt holds the default value on creation for the "started_at" field.
	DefaultStartedAt time.Time
	// DefaultFinishedAt holds the default value on creation for the "finished_at" field.
	DefaultFinishedAt time.Time
	// DefaultPlanStartedAt holds the default value on creation for the "plan_started_at" field.
	DefaultPlanStartedAt time.Time
	// DefaultPlanFinishedAt holds the default value on creation for the "plan_finished_at" field.
	DefaultPlanFinishedAt time.Time
	// DefaultMissionBatchID holds the default value on creation for the "mission_batch_id" field.
	DefaultMissionBatchID int64
	// DefaultSettledAmount holds the default value on creation for the "settled_amount" field.
	DefaultSettledAmount int64
	// DefaultSettledCount holds the default value on creation for the "settled_count" field.
	DefaultSettledCount int64
	// DefaultTotalSettleCount holds the default value on creation for the "total_settle_count" field.
	DefaultTotalSettleCount int64
	// DefaultLatelySettledAt holds the default value on creation for the "lately_settled_at" field.
	DefaultLatelySettledAt time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

const DefaultExtraServiceBillingType enums.ExtraServiceBillingType = "unknown"

// ExtraServiceBillingTypeValidator is a validator for the "extra_service_billing_type" field enum values. It is called by the builders before save.
func ExtraServiceBillingTypeValidator(esbt enums.ExtraServiceBillingType) error {
	switch esbt {
	case "unknown", "time_plan_hour", "time_plan_day", "time_plan_week", "time_plan_month", "time_plan_volume", "hold", "time":
		return nil
	default:
		return fmt.Errorf("extraserviceorder: invalid enum value for extra_service_billing_type field: %q", esbt)
	}
}

const DefaultExtraServiceType enums.ExtraServiceType = "unknown"

// ExtraServiceTypeValidator is a validator for the "extra_service_type" field enum values. It is called by the builders before save.
func ExtraServiceTypeValidator(est enums.ExtraServiceType) error {
	switch est {
	case "unknown", "vpn":
		return nil
	default:
		return fmt.Errorf("extraserviceorder: invalid enum value for extra_service_type field: %q", est)
	}
}

// OrderOption defines the ordering options for the ExtraServiceOrder queries.
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

// ByMissionOrderID orders the results by the mission_order_id field.
func ByMissionOrderID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionOrderID, opts...).ToFunc()
}

// ByExtraServiceBillingType orders the results by the extra_service_billing_type field.
func ByExtraServiceBillingType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExtraServiceBillingType, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// BySymbolID orders the results by the symbol_id field.
func BySymbolID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSymbolID, opts...).ToFunc()
}

// ByUnitCep orders the results by the unit_cep field.
func ByUnitCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUnitCep, opts...).ToFunc()
}

// ByExtraServiceType orders the results by the extra_service_type field.
func ByExtraServiceType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExtraServiceType, opts...).ToFunc()
}

// ByBuyDuration orders the results by the buy_duration field.
func ByBuyDuration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBuyDuration, opts...).ToFunc()
}

// ByStartedAt orders the results by the started_at field.
func ByStartedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartedAt, opts...).ToFunc()
}

// ByFinishedAt orders the results by the finished_at field.
func ByFinishedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFinishedAt, opts...).ToFunc()
}

// ByPlanStartedAt orders the results by the plan_started_at field.
func ByPlanStartedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlanStartedAt, opts...).ToFunc()
}

// ByPlanFinishedAt orders the results by the plan_finished_at field.
func ByPlanFinishedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlanFinishedAt, opts...).ToFunc()
}

// ByMissionBatchID orders the results by the mission_batch_id field.
func ByMissionBatchID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionBatchID, opts...).ToFunc()
}

// BySettledAmount orders the results by the settled_amount field.
func BySettledAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSettledAmount, opts...).ToFunc()
}

// BySettledCount orders the results by the settled_count field.
func BySettledCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSettledCount, opts...).ToFunc()
}

// ByTotalSettleCount orders the results by the total_settle_count field.
func ByTotalSettleCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalSettleCount, opts...).ToFunc()
}

// ByLatelySettledAt orders the results by the lately_settled_at field.
func ByLatelySettledAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLatelySettledAt, opts...).ToFunc()
}

// ByMissionField orders the results by mission field.
func ByMissionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionStep(), sql.OrderByField(field, opts...))
	}
}

// ByMissionOrderField orders the results by mission_order field.
func ByMissionOrderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionOrderStep(), sql.OrderByField(field, opts...))
	}
}

// BySymbolField orders the results by symbol field.
func BySymbolField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSymbolStep(), sql.OrderByField(field, opts...))
	}
}

// ByMissionBatchField orders the results by mission_batch field.
func ByMissionBatchField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionBatchStep(), sql.OrderByField(field, opts...))
	}
}
func newMissionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MissionTable, MissionColumn),
	)
}
func newMissionOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionOrderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MissionOrderTable, MissionOrderColumn),
	)
}
func newSymbolStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SymbolInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SymbolTable, SymbolColumn),
	)
}
func newMissionBatchStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionBatchInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MissionBatchTable, MissionBatchColumn),
	)
}
