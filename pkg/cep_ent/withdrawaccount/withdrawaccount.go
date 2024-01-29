// Code generated by ent, DO NOT EDIT.

package withdrawaccount

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

const (
	// Label holds the string label denoting the withdrawaccount type in the database.
	Label = "withdraw_account"
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
	// FieldBusinessName holds the string denoting the business_name field in the database.
	FieldBusinessName = "business_name"
	// FieldBusinessID holds the string denoting the business_id field in the database.
	FieldBusinessID = "business_id"
	// FieldBusinessType holds the string denoting the business_type field in the database.
	FieldBusinessType = "business_type"
	// FieldIDCard holds the string denoting the id_card field in the database.
	FieldIDCard = "id_card"
	// FieldPersonalName holds the string denoting the personal_name field in the database.
	FieldPersonalName = "personal_name"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldBankCardNumber holds the string denoting the bank_card_number field in the database.
	FieldBankCardNumber = "bank_card_number"
	// FieldBank holds the string denoting the bank field in the database.
	FieldBank = "bank"
	// FieldWay holds the string denoting the way field in the database.
	FieldWay = "way"
	// FieldAlipayCardNo holds the string denoting the alipay_card_no field in the database.
	FieldAlipayCardNo = "alipay_card_no"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the withdrawaccount in the database.
	Table = "withdraw_accounts"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "withdraw_accounts"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for withdrawaccount fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldUserID,
	FieldBusinessName,
	FieldBusinessID,
	FieldBusinessType,
	FieldIDCard,
	FieldPersonalName,
	FieldPhone,
	FieldBankCardNumber,
	FieldBank,
	FieldWay,
	FieldAlipayCardNo,
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
	// DefaultBusinessName holds the default value on creation for the "business_name" field.
	DefaultBusinessName string
	// DefaultBusinessID holds the default value on creation for the "business_id" field.
	DefaultBusinessID int64
	// DefaultIDCard holds the default value on creation for the "id_card" field.
	DefaultIDCard string
	// DefaultPersonalName holds the default value on creation for the "personal_name" field.
	DefaultPersonalName string
	// DefaultPhone holds the default value on creation for the "phone" field.
	DefaultPhone string
	// DefaultBankCardNumber holds the default value on creation for the "bank_card_number" field.
	DefaultBankCardNumber string
	// DefaultBank holds the default value on creation for the "bank" field.
	DefaultBank string
	// DefaultAlipayCardNo holds the default value on creation for the "alipay_card_no" field.
	DefaultAlipayCardNo string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

const DefaultBusinessType enums.BusinessType = "yun"

// BusinessTypeValidator is a validator for the "business_type" field enum values. It is called by the builders before save.
func BusinessTypeValidator(bt enums.BusinessType) error {
	switch bt {
	case "yun", "wft":
		return nil
	default:
		return fmt.Errorf("withdrawaccount: invalid enum value for business_type field: %q", bt)
	}
}

const DefaultWay enums.TransferOrderType = "unknown"

// WayValidator is a validator for the "way" field enum values. It is called by the builders before save.
func WayValidator(w enums.TransferOrderType) error {
	switch w {
	case "withdraw_vx", "withdraw_alipay", "withdraw_bank_card", "unknown", "recharge", "recharge_vx", "recharge_alipay", "manual", "withdraw", "recharge_refund":
		return nil
	default:
		return fmt.Errorf("withdrawaccount: invalid enum value for way field: %q", w)
	}
}

// OrderOption defines the ordering options for the WithdrawAccount queries.
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

// ByBusinessName orders the results by the business_name field.
func ByBusinessName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBusinessName, opts...).ToFunc()
}

// ByBusinessID orders the results by the business_id field.
func ByBusinessID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBusinessID, opts...).ToFunc()
}

// ByBusinessType orders the results by the business_type field.
func ByBusinessType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBusinessType, opts...).ToFunc()
}

// ByIDCard orders the results by the id_card field.
func ByIDCard(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIDCard, opts...).ToFunc()
}

// ByPersonalName orders the results by the personal_name field.
func ByPersonalName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPersonalName, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByBankCardNumber orders the results by the bank_card_number field.
func ByBankCardNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBankCardNumber, opts...).ToFunc()
}

// ByBank orders the results by the bank field.
func ByBank(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBank, opts...).ToFunc()
}

// ByWay orders the results by the way field.
func ByWay(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWay, opts...).ToFunc()
}

// ByAlipayCardNo orders the results by the alipay_card_no field.
func ByAlipayCardNo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAlipayCardNo, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
	)
}
