// Code generated by ent, DO NOT EDIT.

package withdrawrecord

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

const (
	// Label holds the string label denoting the withdrawrecord type in the database.
	Label = "withdraw_record"
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
	// FieldWithdrawAccount holds the string denoting the withdraw_account field in the database.
	FieldWithdrawAccount = "withdraw_account"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldRemainAmount holds the string denoting the remain_amount field in the database.
	FieldRemainAmount = "remain_amount"
	// FieldRate holds the string denoting the rate field in the database.
	FieldRate = "rate"
	// FieldRealAmount holds the string denoting the real_amount field in the database.
	FieldRealAmount = "real_amount"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldRejectReason holds the string denoting the reject_reason field in the database.
	FieldRejectReason = "reject_reason"
	// FieldOperateUserID holds the string denoting the operate_user_id field in the database.
	FieldOperateUserID = "operate_user_id"
	// FieldTransferOrderID holds the string denoting the transfer_order_id field in the database.
	FieldTransferOrderID = "transfer_order_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeOperateUser holds the string denoting the operate_user edge name in mutations.
	EdgeOperateUser = "operate_user"
	// EdgeTransferOrder holds the string denoting the transfer_order edge name in mutations.
	EdgeTransferOrder = "transfer_order"
	// Table holds the table name of the withdrawrecord in the database.
	Table = "withdraw_records"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "withdraw_records"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// OperateUserTable is the table that holds the operate_user relation/edge.
	OperateUserTable = "withdraw_records"
	// OperateUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OperateUserInverseTable = "users"
	// OperateUserColumn is the table column denoting the operate_user relation/edge.
	OperateUserColumn = "operate_user_id"
	// TransferOrderTable is the table that holds the transfer_order relation/edge.
	TransferOrderTable = "withdraw_records"
	// TransferOrderInverseTable is the table name for the TransferOrder entity.
	// It exists in this package in order to avoid circular dependency with the "transferorder" package.
	TransferOrderInverseTable = "transfer_orders"
	// TransferOrderColumn is the table column denoting the transfer_order relation/edge.
	TransferOrderColumn = "transfer_order_id"
)

// Columns holds all SQL columns for withdrawrecord fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldUserID,
	FieldWithdrawAccount,
	FieldType,
	FieldAmount,
	FieldRemainAmount,
	FieldRate,
	FieldRealAmount,
	FieldStatus,
	FieldRejectReason,
	FieldOperateUserID,
	FieldTransferOrderID,
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
	// DefaultWithdrawAccount holds the default value on creation for the "withdraw_account" field.
	DefaultWithdrawAccount string
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount int64
	// DefaultRemainAmount holds the default value on creation for the "remain_amount" field.
	DefaultRemainAmount int64
	// DefaultRate holds the default value on creation for the "rate" field.
	DefaultRate int64
	// DefaultRealAmount holds the default value on creation for the "real_amount" field.
	DefaultRealAmount int64
	// DefaultRejectReason holds the default value on creation for the "reject_reason" field.
	DefaultRejectReason string
	// DefaultOperateUserID holds the default value on creation for the "operate_user_id" field.
	DefaultOperateUserID int64
	// DefaultTransferOrderID holds the default value on creation for the "transfer_order_id" field.
	DefaultTransferOrderID int64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

const DefaultType enums.WithdrawType = "unknown"

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type enums.WithdrawType) error {
	switch _type {
	case "unknown", "withdraw", "withdraw_vx", "withdraw_alipay", "withdraw_bank_card":
		return nil
	default:
		return fmt.Errorf("withdrawrecord: invalid enum value for type field: %q", _type)
	}
}

const DefaultStatus enums.WithdrawStatus = "pending"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s enums.WithdrawStatus) error {
	switch s {
	case "pending", "canceled", "succeed", "failed", "reexchange", "pending_order", "approved", "reject":
		return nil
	default:
		return fmt.Errorf("withdrawrecord: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the WithdrawRecord queries.
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

// ByWithdrawAccount orders the results by the withdraw_account field.
func ByWithdrawAccount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWithdrawAccount, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByRemainAmount orders the results by the remain_amount field.
func ByRemainAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemainAmount, opts...).ToFunc()
}

// ByRate orders the results by the rate field.
func ByRate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRate, opts...).ToFunc()
}

// ByRealAmount orders the results by the real_amount field.
func ByRealAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRealAmount, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByRejectReason orders the results by the reject_reason field.
func ByRejectReason(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRejectReason, opts...).ToFunc()
}

// ByOperateUserID orders the results by the operate_user_id field.
func ByOperateUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOperateUserID, opts...).ToFunc()
}

// ByTransferOrderID orders the results by the transfer_order_id field.
func ByTransferOrderID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTransferOrderID, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByOperateUserField orders the results by operate_user field.
func ByOperateUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOperateUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByTransferOrderField orders the results by transfer_order field.
func ByTransferOrderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTransferOrderStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newOperateUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OperateUserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OperateUserTable, OperateUserColumn),
	)
}
func newTransferOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TransferOrderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, TransferOrderTable, TransferOrderColumn),
	)
}
