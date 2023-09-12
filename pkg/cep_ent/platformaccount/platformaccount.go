// Code generated by ent, DO NOT EDIT.

package platformaccount

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the platformaccount type in the database.
	Label = "platform_account"
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
	// FieldSumTotalCep holds the string denoting the sum_total_cep field in the database.
	FieldSumTotalCep = "sum_total_cep"
	// FieldTotalCep holds the string denoting the total_cep field in the database.
	FieldTotalCep = "total_cep"
	// FieldSumPureCep holds the string denoting the sum_pure_cep field in the database.
	FieldSumPureCep = "sum_pure_cep"
	// FieldPureCep holds the string denoting the pure_cep field in the database.
	FieldPureCep = "pure_cep"
	// FieldSumGiftCep holds the string denoting the sum_gift_cep field in the database.
	FieldSumGiftCep = "sum_gift_cep"
	// FieldGiftCep holds the string denoting the gift_cep field in the database.
	FieldGiftCep = "gift_cep"
	// EdgeEarnBills holds the string denoting the earn_bills edge name in mutations.
	EdgeEarnBills = "earn_bills"
	// Table holds the table name of the platformaccount in the database.
	Table = "platform_accounts"
	// EarnBillsTable is the table that holds the earn_bills relation/edge.
	EarnBillsTable = "earn_bills"
	// EarnBillsInverseTable is the table name for the EarnBill entity.
	// It exists in this package in order to avoid circular dependency with the "earnbill" package.
	EarnBillsInverseTable = "earn_bills"
	// EarnBillsColumn is the table column denoting the earn_bills relation/edge.
	EarnBillsColumn = "platform_account_id"
)

// Columns holds all SQL columns for platformaccount fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldType,
	FieldSumTotalCep,
	FieldTotalCep,
	FieldSumPureCep,
	FieldPureCep,
	FieldSumGiftCep,
	FieldGiftCep,
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
	// DefaultSumTotalCep holds the default value on creation for the "sum_total_cep" field.
	DefaultSumTotalCep int64
	// DefaultTotalCep holds the default value on creation for the "total_cep" field.
	DefaultTotalCep int64
	// DefaultSumPureCep holds the default value on creation for the "sum_pure_cep" field.
	DefaultSumPureCep int64
	// DefaultPureCep holds the default value on creation for the "pure_cep" field.
	DefaultPureCep int64
	// DefaultSumGiftCep holds the default value on creation for the "sum_gift_cep" field.
	DefaultSumGiftCep int64
	// DefaultGiftCep holds the default value on creation for the "gift_cep" field.
	DefaultGiftCep int64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

// Type defines the type for the "type" enum field.
type Type string

// TypeProfit is the default value of the Type enum.
const DefaultType = TypeProfit

// Type values.
const (
	TypeProfit Type = "profit"
	TypeMarket Type = "market"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeProfit, TypeMarket:
		return nil
	default:
		return fmt.Errorf("platformaccount: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the PlatformAccount queries.
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

// BySumTotalCep orders the results by the sum_total_cep field.
func BySumTotalCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSumTotalCep, opts...).ToFunc()
}

// ByTotalCep orders the results by the total_cep field.
func ByTotalCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalCep, opts...).ToFunc()
}

// BySumPureCep orders the results by the sum_pure_cep field.
func BySumPureCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSumPureCep, opts...).ToFunc()
}

// ByPureCep orders the results by the pure_cep field.
func ByPureCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPureCep, opts...).ToFunc()
}

// BySumGiftCep orders the results by the sum_gift_cep field.
func BySumGiftCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSumGiftCep, opts...).ToFunc()
}

// ByGiftCep orders the results by the gift_cep field.
func ByGiftCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGiftCep, opts...).ToFunc()
}

// ByEarnBillsCount orders the results by earn_bills count.
func ByEarnBillsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEarnBillsStep(), opts...)
	}
}

// ByEarnBills orders the results by earn_bills terms.
func ByEarnBills(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEarnBillsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newEarnBillsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EarnBillsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EarnBillsTable, EarnBillsColumn),
	)
}
