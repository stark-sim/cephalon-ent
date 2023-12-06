// Code generated by ent, DO NOT EDIT.

package transferorder

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

const (
	// Label holds the string label denoting the transferorder type in the database.
	Label = "transfer_order"
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
	// FieldSourceUserID holds the string denoting the source_user_id field in the database.
	FieldSourceUserID = "source_user_id"
	// FieldTargetUserID holds the string denoting the target_user_id field in the database.
	FieldTargetUserID = "target_user_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldSymbolID holds the string denoting the symbol_id field in the database.
	FieldSymbolID = "symbol_id"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldSerialNumber holds the string denoting the serial_number field in the database.
	FieldSerialNumber = "serial_number"
	// FieldSocialID holds the string denoting the social_id field in the database.
	FieldSocialID = "social_id"
	// FieldThirdAPIResp holds the string denoting the third_api_resp field in the database.
	FieldThirdAPIResp = "third_api_resp"
	// FieldOutTransactionID holds the string denoting the out_transaction_id field in the database.
	FieldOutTransactionID = "out_transaction_id"
	// EdgeSourceUser holds the string denoting the source_user edge name in mutations.
	EdgeSourceUser = "source_user"
	// EdgeTargetUser holds the string denoting the target_user edge name in mutations.
	EdgeTargetUser = "target_user"
	// EdgeBills holds the string denoting the bills edge name in mutations.
	EdgeBills = "bills"
	// EdgeVxSocial holds the string denoting the vx_social edge name in mutations.
	EdgeVxSocial = "vx_social"
	// EdgeSymbol holds the string denoting the symbol edge name in mutations.
	EdgeSymbol = "symbol"
	// Table holds the table name of the transferorder in the database.
	Table = "transfer_orders"
	// SourceUserTable is the table that holds the source_user relation/edge.
	SourceUserTable = "transfer_orders"
	// SourceUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	SourceUserInverseTable = "users"
	// SourceUserColumn is the table column denoting the source_user relation/edge.
	SourceUserColumn = "source_user_id"
	// TargetUserTable is the table that holds the target_user relation/edge.
	TargetUserTable = "transfer_orders"
	// TargetUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	TargetUserInverseTable = "users"
	// TargetUserColumn is the table column denoting the target_user relation/edge.
	TargetUserColumn = "target_user_id"
	// BillsTable is the table that holds the bills relation/edge.
	BillsTable = "bills"
	// BillsInverseTable is the table name for the Bill entity.
	// It exists in this package in order to avoid circular dependency with the "bill" package.
	BillsInverseTable = "bills"
	// BillsColumn is the table column denoting the bills relation/edge.
	BillsColumn = "order_id"
	// VxSocialTable is the table that holds the vx_social relation/edge.
	VxSocialTable = "transfer_orders"
	// VxSocialInverseTable is the table name for the VXSocial entity.
	// It exists in this package in order to avoid circular dependency with the "vxsocial" package.
	VxSocialInverseTable = "vx_socials"
	// VxSocialColumn is the table column denoting the vx_social relation/edge.
	VxSocialColumn = "social_id"
	// SymbolTable is the table that holds the symbol relation/edge.
	SymbolTable = "transfer_orders"
	// SymbolInverseTable is the table name for the Symbol entity.
	// It exists in this package in order to avoid circular dependency with the "symbol" package.
	SymbolInverseTable = "symbols"
	// SymbolColumn is the table column denoting the symbol relation/edge.
	SymbolColumn = "symbol_id"
)

// Columns holds all SQL columns for transferorder fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldSourceUserID,
	FieldTargetUserID,
	FieldStatus,
	FieldSymbolID,
	FieldAmount,
	FieldType,
	FieldSerialNumber,
	FieldSocialID,
	FieldThirdAPIResp,
	FieldOutTransactionID,
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
	// DefaultSourceUserID holds the default value on creation for the "source_user_id" field.
	DefaultSourceUserID int64
	// DefaultTargetUserID holds the default value on creation for the "target_user_id" field.
	DefaultTargetUserID int64
	// DefaultSymbolID holds the default value on creation for the "symbol_id" field.
	DefaultSymbolID int64
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount int64
	// DefaultSerialNumber holds the default value on creation for the "serial_number" field.
	DefaultSerialNumber string
	// DefaultSocialID holds the default value on creation for the "social_id" field.
	DefaultSocialID int64
	// DefaultThirdAPIResp holds the default value on creation for the "third_api_resp" field.
	DefaultThirdAPIResp string
	// DefaultOutTransactionID holds the default value on creation for the "out_transaction_id" field.
	DefaultOutTransactionID string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

// Status defines the type for the "status" enum field.
type Status string

// StatusPending is the default value of the Status enum.
const DefaultStatus = StatusPending

// Status values.
const (
	StatusPending  Status = "pending"
	StatusCanceled Status = "canceled"
	StatusSucceed  Status = "succeed"
	StatusFailed   Status = "failed"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusPending, StatusCanceled, StatusSucceed, StatusFailed:
		return nil
	default:
		return fmt.Errorf("transferorder: invalid enum value for status field: %q", s)
	}
}

const DefaultType enums.TransferOrderType = "unknown"

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type enums.TransferOrderType) error {
	switch _type {
	case "unknown", "recharge", "recharge_vx", "recharge_alipay", "manual", "withdraw":
		return nil
	default:
		return fmt.Errorf("transferorder: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the TransferOrder queries.
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

// BySourceUserID orders the results by the source_user_id field.
func BySourceUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSourceUserID, opts...).ToFunc()
}

// ByTargetUserID orders the results by the target_user_id field.
func ByTargetUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTargetUserID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// BySymbolID orders the results by the symbol_id field.
func BySymbolID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSymbolID, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// BySerialNumber orders the results by the serial_number field.
func BySerialNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSerialNumber, opts...).ToFunc()
}

// BySocialID orders the results by the social_id field.
func BySocialID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSocialID, opts...).ToFunc()
}

// ByThirdAPIResp orders the results by the third_api_resp field.
func ByThirdAPIResp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldThirdAPIResp, opts...).ToFunc()
}

// ByOutTransactionID orders the results by the out_transaction_id field.
func ByOutTransactionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOutTransactionID, opts...).ToFunc()
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

// ByBillsCount orders the results by bills count.
func ByBillsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBillsStep(), opts...)
	}
}

// ByBills orders the results by bills terms.
func ByBills(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBillsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByVxSocialField orders the results by vx_social field.
func ByVxSocialField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVxSocialStep(), sql.OrderByField(field, opts...))
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
func newBillsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BillsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BillsTable, BillsColumn),
	)
}
func newVxSocialStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VxSocialInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, VxSocialTable, VxSocialColumn),
	)
}
func newSymbolStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SymbolInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SymbolTable, SymbolColumn),
	)
}
