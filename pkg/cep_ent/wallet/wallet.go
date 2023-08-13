// Code generated by ent, DO NOT EDIT.

package wallet

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the wallet type in the database.
	Label = "wallet"
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
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldCep holds the string denoting the cep field in the database.
	FieldCep = "cep"
	// FieldSumCep holds the string denoting the sum_cep field in the database.
	FieldSumCep = "sum_cep"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeBills holds the string denoting the bills edge name in mutations.
	EdgeBills = "bills"
	// Table holds the table name of the wallet in the database.
	Table = "wallets"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "wallets"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// BillsTable is the table that holds the bills relation/edge.
	BillsTable = "bills"
	// BillsInverseTable is the table name for the Bill entity.
	// It exists in this package in order to avoid circular dependency with the "bill" package.
	BillsInverseTable = "bills"
	// BillsColumn is the table column denoting the bills relation/edge.
	BillsColumn = "wallet_id"
)

// Columns holds all SQL columns for wallet fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldUserID,
	FieldCep,
	FieldSumCep,
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
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID int64
	// DefaultCep holds the default value on creation for the "cep" field.
	DefaultCep int64
	// DefaultSumCep holds the default value on creation for the "sum_cep" field.
	DefaultSumCep int64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

// OrderOption defines the ordering options for the Wallet queries.
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

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByCep orders the results by the cep field.
func ByCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCep, opts...).ToFunc()
}

// BySumCep orders the results by the sum_cep field.
func BySumCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSumCep, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
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
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
	)
}
func newBillsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BillsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BillsTable, BillsColumn),
	)
}
