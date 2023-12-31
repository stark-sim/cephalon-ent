// Code generated by ent, DO NOT EDIT.

package earnbill

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the earnbill type in the database.
	Label = "earn_bill"
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
	// FieldIsMinus holds the string denoting the is_minus field in the database.
	FieldIsMinus = "is_minus"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldSerialNumber holds the string denoting the serial_number field in the database.
	FieldSerialNumber = "serial_number"
	// FieldProfitAccountID holds the string denoting the profit_account_id field in the database.
	FieldProfitAccountID = "profit_account_id"
	// FieldPureCep holds the string denoting the pure_cep field in the database.
	FieldPureCep = "pure_cep"
	// FieldGiftCep holds the string denoting the gift_cep field in the database.
	FieldGiftCep = "gift_cep"
	// FieldPlatformAccountID holds the string denoting the platform_account_id field in the database.
	FieldPlatformAccountID = "platform_account_id"
	// FieldPlatformPureCep holds the string denoting the platform_pure_cep field in the database.
	FieldPlatformPureCep = "platform_pure_cep"
	// FieldPlatformGiftCep holds the string denoting the platform_gift_cep field in the database.
	FieldPlatformGiftCep = "platform_gift_cep"
	// FieldReasonID holds the string denoting the reason_id field in the database.
	FieldReasonID = "reason_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeProfitAccount holds the string denoting the profit_account edge name in mutations.
	EdgeProfitAccount = "profit_account"
	// EdgePlatformAccount holds the string denoting the platform_account edge name in mutations.
	EdgePlatformAccount = "platform_account"
	// EdgeMissionProduceOrders holds the string denoting the mission_produce_orders edge name in mutations.
	EdgeMissionProduceOrders = "mission_produce_orders"
	// Table holds the table name of the earnbill in the database.
	Table = "earn_bills"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "earn_bills"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// ProfitAccountTable is the table that holds the profit_account relation/edge.
	ProfitAccountTable = "earn_bills"
	// ProfitAccountInverseTable is the table name for the ProfitAccount entity.
	// It exists in this package in order to avoid circular dependency with the "profitaccount" package.
	ProfitAccountInverseTable = "profit_accounts"
	// ProfitAccountColumn is the table column denoting the profit_account relation/edge.
	ProfitAccountColumn = "profit_account_id"
	// PlatformAccountTable is the table that holds the platform_account relation/edge.
	PlatformAccountTable = "earn_bills"
	// PlatformAccountInverseTable is the table name for the PlatformAccount entity.
	// It exists in this package in order to avoid circular dependency with the "platformaccount" package.
	PlatformAccountInverseTable = "platform_accounts"
	// PlatformAccountColumn is the table column denoting the platform_account relation/edge.
	PlatformAccountColumn = "platform_account_id"
	// MissionProduceOrdersTable is the table that holds the mission_produce_orders relation/edge.
	MissionProduceOrdersTable = "earn_bills"
	// MissionProduceOrdersInverseTable is the table name for the MissionProduceOrder entity.
	// It exists in this package in order to avoid circular dependency with the "missionproduceorder" package.
	MissionProduceOrdersInverseTable = "mission_produce_orders"
	// MissionProduceOrdersColumn is the table column denoting the mission_produce_orders relation/edge.
	MissionProduceOrdersColumn = "reason_id"
)

// Columns holds all SQL columns for earnbill fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldType,
	FieldIsMinus,
	FieldUserID,
	FieldSerialNumber,
	FieldProfitAccountID,
	FieldPureCep,
	FieldGiftCep,
	FieldPlatformAccountID,
	FieldPlatformPureCep,
	FieldPlatformGiftCep,
	FieldReasonID,
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
	// DefaultIsMinus holds the default value on creation for the "is_minus" field.
	DefaultIsMinus bool
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID int64
	// DefaultSerialNumber holds the default value on creation for the "serial_number" field.
	DefaultSerialNumber string
	// DefaultProfitAccountID holds the default value on creation for the "profit_account_id" field.
	DefaultProfitAccountID int64
	// DefaultPureCep holds the default value on creation for the "pure_cep" field.
	DefaultPureCep int64
	// DefaultGiftCep holds the default value on creation for the "gift_cep" field.
	DefaultGiftCep int64
	// DefaultPlatformAccountID holds the default value on creation for the "platform_account_id" field.
	DefaultPlatformAccountID int64
	// DefaultPlatformPureCep holds the default value on creation for the "platform_pure_cep" field.
	DefaultPlatformPureCep int64
	// DefaultPlatformGiftCep holds the default value on creation for the "platform_gift_cep" field.
	DefaultPlatformGiftCep int64
	// DefaultReasonID holds the default value on creation for the "reason_id" field.
	DefaultReasonID int64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

// Type defines the type for the "type" enum field.
type Type string

// TypeMission is the default value of the Type enum.
const DefaultType = TypeMission

// Type values.
const (
	TypeMission  Type = "mission"
	TypeWithdraw Type = "withdraw"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeMission, TypeWithdraw:
		return nil
	default:
		return fmt.Errorf("earnbill: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the EarnBill queries.
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

// ByIsMinus orders the results by the is_minus field.
func ByIsMinus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsMinus, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// BySerialNumber orders the results by the serial_number field.
func BySerialNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSerialNumber, opts...).ToFunc()
}

// ByProfitAccountID orders the results by the profit_account_id field.
func ByProfitAccountID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProfitAccountID, opts...).ToFunc()
}

// ByPureCep orders the results by the pure_cep field.
func ByPureCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPureCep, opts...).ToFunc()
}

// ByGiftCep orders the results by the gift_cep field.
func ByGiftCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGiftCep, opts...).ToFunc()
}

// ByPlatformAccountID orders the results by the platform_account_id field.
func ByPlatformAccountID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlatformAccountID, opts...).ToFunc()
}

// ByPlatformPureCep orders the results by the platform_pure_cep field.
func ByPlatformPureCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlatformPureCep, opts...).ToFunc()
}

// ByPlatformGiftCep orders the results by the platform_gift_cep field.
func ByPlatformGiftCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlatformGiftCep, opts...).ToFunc()
}

// ByReasonID orders the results by the reason_id field.
func ByReasonID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReasonID, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByProfitAccountField orders the results by profit_account field.
func ByProfitAccountField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProfitAccountStep(), sql.OrderByField(field, opts...))
	}
}

// ByPlatformAccountField orders the results by platform_account field.
func ByPlatformAccountField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlatformAccountStep(), sql.OrderByField(field, opts...))
	}
}

// ByMissionProduceOrdersField orders the results by mission_produce_orders field.
func ByMissionProduceOrdersField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionProduceOrdersStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newProfitAccountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProfitAccountInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProfitAccountTable, ProfitAccountColumn),
	)
}
func newPlatformAccountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlatformAccountInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PlatformAccountTable, PlatformAccountColumn),
	)
}
func newMissionProduceOrdersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionProduceOrdersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MissionProduceOrdersTable, MissionProduceOrdersColumn),
	)
}
