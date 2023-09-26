// Code generated by ent, DO NOT EDIT.

package bill

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

const (
	// Label holds the string label denoting the bill type in the database.
	Label = "bill"
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
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldWay holds the string denoting the way field in the database.
	FieldWay = "way"
	// FieldSymbolID holds the string denoting the symbol_id field in the database.
	FieldSymbolID = "symbol_id"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldTargetUserID holds the string denoting the target_user_id field in the database.
	FieldTargetUserID = "target_user_id"
	// FieldTargetBeforeAmount holds the string denoting the target_before_amount field in the database.
	FieldTargetBeforeAmount = "target_before_amount"
	// FieldTargetAfterAmount holds the string denoting the target_after_amount field in the database.
	FieldTargetAfterAmount = "target_after_amount"
	// FieldSourceUserID holds the string denoting the source_user_id field in the database.
	FieldSourceUserID = "source_user_id"
	// FieldSourceBeforeAmount holds the string denoting the source_before_amount field in the database.
	FieldSourceBeforeAmount = "source_before_amount"
	// FieldSourceAfterAmount holds the string denoting the source_after_amount field in the database.
	FieldSourceAfterAmount = "source_after_amount"
	// FieldSerialNumber holds the string denoting the serial_number field in the database.
	FieldSerialNumber = "serial_number"
	// FieldInviteID holds the string denoting the invite_id field in the database.
	FieldInviteID = "invite_id"
	// EdgeSourceUser holds the string denoting the source_user edge name in mutations.
	EdgeSourceUser = "source_user"
	// EdgeTargetUser holds the string denoting the target_user edge name in mutations.
	EdgeTargetUser = "target_user"
	// EdgeTransferOrder holds the string denoting the transfer_order edge name in mutations.
	EdgeTransferOrder = "transfer_order"
	// EdgeMissionOrder holds the string denoting the mission_order edge name in mutations.
	EdgeMissionOrder = "mission_order"
	// EdgeInvite holds the string denoting the invite edge name in mutations.
	EdgeInvite = "invite"
	// EdgeSymbol holds the string denoting the symbol edge name in mutations.
	EdgeSymbol = "symbol"
	// Table holds the table name of the bill in the database.
	Table = "bills"
	// SourceUserTable is the table that holds the source_user relation/edge.
	SourceUserTable = "bills"
	// SourceUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	SourceUserInverseTable = "users"
	// SourceUserColumn is the table column denoting the source_user relation/edge.
	SourceUserColumn = "source_user_id"
	// TargetUserTable is the table that holds the target_user relation/edge.
	TargetUserTable = "bills"
	// TargetUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	TargetUserInverseTable = "users"
	// TargetUserColumn is the table column denoting the target_user relation/edge.
	TargetUserColumn = "target_user_id"
	// TransferOrderTable is the table that holds the transfer_order relation/edge.
	TransferOrderTable = "bills"
	// TransferOrderInverseTable is the table name for the TransferOrder entity.
	// It exists in this package in order to avoid circular dependency with the "transferorder" package.
	TransferOrderInverseTable = "transfer_orders"
	// TransferOrderColumn is the table column denoting the transfer_order relation/edge.
	TransferOrderColumn = "order_id"
	// MissionOrderTable is the table that holds the mission_order relation/edge.
	MissionOrderTable = "bills"
	// MissionOrderInverseTable is the table name for the MissionOrder entity.
	// It exists in this package in order to avoid circular dependency with the "missionorder" package.
	MissionOrderInverseTable = "mission_orders"
	// MissionOrderColumn is the table column denoting the mission_order relation/edge.
	MissionOrderColumn = "order_id"
	// InviteTable is the table that holds the invite relation/edge.
	InviteTable = "bills"
	// InviteInverseTable is the table name for the Invite entity.
	// It exists in this package in order to avoid circular dependency with the "invite" package.
	InviteInverseTable = "invites"
	// InviteColumn is the table column denoting the invite relation/edge.
	InviteColumn = "invite_id"
	// SymbolTable is the table that holds the symbol relation/edge.
	SymbolTable = "bills"
	// SymbolInverseTable is the table name for the Symbol entity.
	// It exists in this package in order to avoid circular dependency with the "symbol" package.
	SymbolInverseTable = "symbols"
	// SymbolColumn is the table column denoting the symbol relation/edge.
	SymbolColumn = "symbol_id"
)

// Columns holds all SQL columns for bill fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldType,
	FieldOrderID,
	FieldWay,
	FieldSymbolID,
	FieldAmount,
	FieldTargetUserID,
	FieldTargetBeforeAmount,
	FieldTargetAfterAmount,
	FieldSourceUserID,
	FieldSourceBeforeAmount,
	FieldSourceAfterAmount,
	FieldSerialNumber,
	FieldInviteID,
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
	// DefaultOrderID holds the default value on creation for the "order_id" field.
	DefaultOrderID int64
	// DefaultSymbolID holds the default value on creation for the "symbol_id" field.
	DefaultSymbolID int64
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount int64
	// DefaultTargetUserID holds the default value on creation for the "target_user_id" field.
	DefaultTargetUserID int64
	// DefaultTargetBeforeAmount holds the default value on creation for the "target_before_amount" field.
	DefaultTargetBeforeAmount int64
	// DefaultTargetAfterAmount holds the default value on creation for the "target_after_amount" field.
	DefaultTargetAfterAmount int64
	// DefaultSourceUserID holds the default value on creation for the "source_user_id" field.
	DefaultSourceUserID int64
	// DefaultSourceBeforeAmount holds the default value on creation for the "source_before_amount" field.
	DefaultSourceBeforeAmount int64
	// DefaultSourceAfterAmount holds the default value on creation for the "source_after_amount" field.
	DefaultSourceAfterAmount int64
	// DefaultSerialNumber holds the default value on creation for the "serial_number" field.
	DefaultSerialNumber string
	// DefaultInviteID holds the default value on creation for the "invite_id" field.
	DefaultInviteID int64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

const DefaultType enums.BillType = "unknown"

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type enums.BillType) error {
	switch _type {
	case "unknown", "recharge", "mission_consume", "mission_produce", "transfer", "active", "mission", "gas":
		return nil
	default:
		return fmt.Errorf("bill: invalid enum value for type field: %q", _type)
	}
}

const DefaultWay enums.BillWay = "unknown"

// WayValidator is a validator for the "way" field enum values. It is called by the builders before save.
func WayValidator(w enums.BillWay) error {
	switch w {
	case "unknown", "recharge_wechat", "recharge_alipay", "mission_time", "mission_count", "mission_hold", "mission_volume", "active_register", "active_share", "active_recharge", "transfer_manual", "first_invite_recharge":
		return nil
	default:
		return fmt.Errorf("bill: invalid enum value for way field: %q", w)
	}
}

// OrderOption defines the ordering options for the Bill queries.
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

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByOrderID orders the results by the order_id field.
func ByOrderID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderID, opts...).ToFunc()
}

// ByWay orders the results by the way field.
func ByWay(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWay, opts...).ToFunc()
}

// BySymbolID orders the results by the symbol_id field.
func BySymbolID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSymbolID, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByTargetUserID orders the results by the target_user_id field.
func ByTargetUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTargetUserID, opts...).ToFunc()
}

// ByTargetBeforeAmount orders the results by the target_before_amount field.
func ByTargetBeforeAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTargetBeforeAmount, opts...).ToFunc()
}

// ByTargetAfterAmount orders the results by the target_after_amount field.
func ByTargetAfterAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTargetAfterAmount, opts...).ToFunc()
}

// BySourceUserID orders the results by the source_user_id field.
func BySourceUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSourceUserID, opts...).ToFunc()
}

// BySourceBeforeAmount orders the results by the source_before_amount field.
func BySourceBeforeAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSourceBeforeAmount, opts...).ToFunc()
}

// BySourceAfterAmount orders the results by the source_after_amount field.
func BySourceAfterAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSourceAfterAmount, opts...).ToFunc()
}

// BySerialNumber orders the results by the serial_number field.
func BySerialNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSerialNumber, opts...).ToFunc()
}

// ByInviteID orders the results by the invite_id field.
func ByInviteID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInviteID, opts...).ToFunc()
}

// BySourceUserField orders the results by source_user field.
func BySourceUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSourceUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByTargetUserField orders the results by target_user field.
func ByTargetUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTargetUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByTransferOrderField orders the results by transfer_order field.
func ByTransferOrderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTransferOrderStep(), sql.OrderByField(field, opts...))
	}
}

// ByMissionOrderField orders the results by mission_order field.
func ByMissionOrderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionOrderStep(), sql.OrderByField(field, opts...))
	}
}

// ByInviteField orders the results by invite field.
func ByInviteField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInviteStep(), sql.OrderByField(field, opts...))
	}
}

// BySymbolField orders the results by symbol field.
func BySymbolField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSymbolStep(), sql.OrderByField(field, opts...))
	}
}
func newSourceUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SourceUserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SourceUserTable, SourceUserColumn),
	)
}
func newTargetUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TargetUserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TargetUserTable, TargetUserColumn),
	)
}
func newTransferOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TransferOrderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TransferOrderTable, TransferOrderColumn),
	)
}
func newMissionOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionOrderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MissionOrderTable, MissionOrderColumn),
	)
}
func newInviteStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InviteInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, InviteTable, InviteColumn),
	)
}
func newSymbolStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SymbolInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SymbolTable, SymbolColumn),
	)
}
