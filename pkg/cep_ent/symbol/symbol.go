// Code generated by ent, DO NOT EDIT.

package symbol

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the symbol type in the database.
	Label = "symbol"
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
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeWallets holds the string denoting the wallets edge name in mutations.
	EdgeWallets = "wallets"
	// EdgeBills holds the string denoting the bills edge name in mutations.
	EdgeBills = "bills"
	// EdgeMissionOrders holds the string denoting the mission_orders edge name in mutations.
	EdgeMissionOrders = "mission_orders"
	// EdgeTransferOrders holds the string denoting the transfer_orders edge name in mutations.
	EdgeTransferOrders = "transfer_orders"
	// EdgeExtraServiceOrder holds the string denoting the extra_service_order edge name in mutations.
	EdgeExtraServiceOrder = "extra_service_order"
	// Table holds the table name of the symbol in the database.
	Table = "symbols"
	// WalletsTable is the table that holds the wallets relation/edge.
	WalletsTable = "wallets"
	// WalletsInverseTable is the table name for the Wallet entity.
	// It exists in this package in order to avoid circular dependency with the "wallet" package.
	WalletsInverseTable = "wallets"
	// WalletsColumn is the table column denoting the wallets relation/edge.
	WalletsColumn = "symbol_id"
	// BillsTable is the table that holds the bills relation/edge.
	BillsTable = "bills"
	// BillsInverseTable is the table name for the Bill entity.
	// It exists in this package in order to avoid circular dependency with the "bill" package.
	BillsInverseTable = "bills"
	// BillsColumn is the table column denoting the bills relation/edge.
	BillsColumn = "symbol_id"
	// MissionOrdersTable is the table that holds the mission_orders relation/edge.
	MissionOrdersTable = "mission_orders"
	// MissionOrdersInverseTable is the table name for the MissionOrder entity.
	// It exists in this package in order to avoid circular dependency with the "missionorder" package.
	MissionOrdersInverseTable = "mission_orders"
	// MissionOrdersColumn is the table column denoting the mission_orders relation/edge.
	MissionOrdersColumn = "symbol_id"
	// TransferOrdersTable is the table that holds the transfer_orders relation/edge.
	TransferOrdersTable = "transfer_orders"
	// TransferOrdersInverseTable is the table name for the TransferOrder entity.
	// It exists in this package in order to avoid circular dependency with the "transferorder" package.
	TransferOrdersInverseTable = "transfer_orders"
	// TransferOrdersColumn is the table column denoting the transfer_orders relation/edge.
	TransferOrdersColumn = "symbol_id"
	// ExtraServiceOrderTable is the table that holds the extra_service_order relation/edge.
	ExtraServiceOrderTable = "extra_service_orders"
	// ExtraServiceOrderInverseTable is the table name for the ExtraServiceOrder entity.
	// It exists in this package in order to avoid circular dependency with the "extraserviceorder" package.
	ExtraServiceOrderInverseTable = "extra_service_orders"
	// ExtraServiceOrderColumn is the table column denoting the extra_service_order relation/edge.
	ExtraServiceOrderColumn = "symbol_id"
)

// Columns holds all SQL columns for symbol fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
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
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

// OrderOption defines the ordering options for the Symbol queries.
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

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByWalletsCount orders the results by wallets count.
func ByWalletsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWalletsStep(), opts...)
	}
}

// ByWallets orders the results by wallets terms.
func ByWallets(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWalletsStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByMissionOrdersCount orders the results by mission_orders count.
func ByMissionOrdersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMissionOrdersStep(), opts...)
	}
}

// ByMissionOrders orders the results by mission_orders terms.
func ByMissionOrders(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionOrdersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTransferOrdersCount orders the results by transfer_orders count.
func ByTransferOrdersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTransferOrdersStep(), opts...)
	}
}

// ByTransferOrders orders the results by transfer_orders terms.
func ByTransferOrders(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTransferOrdersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByExtraServiceOrderCount orders the results by extra_service_order count.
func ByExtraServiceOrderCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newExtraServiceOrderStep(), opts...)
	}
}

// ByExtraServiceOrder orders the results by extra_service_order terms.
func ByExtraServiceOrder(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExtraServiceOrderStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newWalletsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WalletsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WalletsTable, WalletsColumn),
	)
}
func newBillsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BillsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BillsTable, BillsColumn),
	)
}
func newMissionOrdersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionOrdersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MissionOrdersTable, MissionOrdersColumn),
	)
}
func newTransferOrdersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TransferOrdersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TransferOrdersTable, TransferOrdersColumn),
	)
}
func newExtraServiceOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExtraServiceOrderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ExtraServiceOrderTable, ExtraServiceOrderColumn),
	)
}
