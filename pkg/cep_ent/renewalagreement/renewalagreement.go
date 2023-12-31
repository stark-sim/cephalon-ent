// Code generated by ent, DO NOT EDIT.

package renewalagreement

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

const (
	// Label holds the string label denoting the renewalagreement type in the database.
	Label = "renewal_agreement"
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
	// FieldNextPayTime holds the string denoting the next_pay_time field in the database.
	FieldNextPayTime = "next_pay_time"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldSubStatus holds the string denoting the sub_status field in the database.
	FieldSubStatus = "sub_status"
	// FieldPayStatus holds the string denoting the pay_status field in the database.
	FieldPayStatus = "pay_status"
	// FieldSymbolID holds the string denoting the symbol_id field in the database.
	FieldSymbolID = "symbol_id"
	// FieldFirstPay holds the string denoting the first_pay field in the database.
	FieldFirstPay = "first_pay"
	// FieldAfterPay holds the string denoting the after_pay field in the database.
	FieldAfterPay = "after_pay"
	// FieldLastWarningTime holds the string denoting the last_warning_time field in the database.
	FieldLastWarningTime = "last_warning_time"
	// FieldSubFinishedTime holds the string denoting the sub_finished_time field in the database.
	FieldSubFinishedTime = "sub_finished_time"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldMissionID holds the string denoting the mission_id field in the database.
	FieldMissionID = "mission_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeMission holds the string denoting the mission edge name in mutations.
	EdgeMission = "mission"
	// Table holds the table name of the renewalagreement in the database.
	Table = "renewal_agreements"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "renewal_agreements"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// MissionTable is the table that holds the mission relation/edge.
	MissionTable = "renewal_agreements"
	// MissionInverseTable is the table name for the Mission entity.
	// It exists in this package in order to avoid circular dependency with the "mission" package.
	MissionInverseTable = "missions"
	// MissionColumn is the table column denoting the mission relation/edge.
	MissionColumn = "mission_id"
)

// Columns holds all SQL columns for renewalagreement fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldNextPayTime,
	FieldType,
	FieldSubStatus,
	FieldPayStatus,
	FieldSymbolID,
	FieldFirstPay,
	FieldAfterPay,
	FieldLastWarningTime,
	FieldSubFinishedTime,
	FieldUserID,
	FieldMissionID,
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
	// DefaultNextPayTime holds the default value on creation for the "next_pay_time" field.
	DefaultNextPayTime time.Time
	// DefaultSymbolID holds the default value on creation for the "symbol_id" field.
	DefaultSymbolID int64
	// DefaultFirstPay holds the default value on creation for the "first_pay" field.
	DefaultFirstPay int64
	// DefaultAfterPay holds the default value on creation for the "after_pay" field.
	DefaultAfterPay int64
	// DefaultLastWarningTime holds the default value on creation for the "last_warning_time" field.
	DefaultLastWarningTime time.Time
	// DefaultSubFinishedTime holds the default value on creation for the "sub_finished_time" field.
	DefaultSubFinishedTime time.Time
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID int64
	// DefaultMissionID holds the default value on creation for the "mission_id" field.
	DefaultMissionID int64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

const DefaultType enums.RenewalType = "unknow"

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type enums.RenewalType) error {
	switch _type {
	case "unknow", "hour", "day", "week", "month":
		return nil
	default:
		return fmt.Errorf("renewalagreement: invalid enum value for type field: %q", _type)
	}
}

const DefaultSubStatus enums.RenewalSubStatus = "unknow"

// SubStatusValidator is a validator for the "sub_status" field enum values. It is called by the builders before save.
func SubStatusValidator(ss enums.RenewalSubStatus) error {
	switch ss {
	case "unknow", "subscribing", "finished":
		return nil
	default:
		return fmt.Errorf("renewalagreement: invalid enum value for sub_status field: %q", ss)
	}
}

const DefaultPayStatus enums.RenewalPayStatus = "unknow"

// PayStatusValidator is a validator for the "pay_status" field enum values. It is called by the builders before save.
func PayStatusValidator(ps enums.RenewalPayStatus) error {
	switch ps {
	case "unknow", "waiting", "succeed", "failed":
		return nil
	default:
		return fmt.Errorf("renewalagreement: invalid enum value for pay_status field: %q", ps)
	}
}

// OrderOption defines the ordering options for the RenewalAgreement queries.
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

// ByNextPayTime orders the results by the next_pay_time field.
func ByNextPayTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNextPayTime, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// BySubStatus orders the results by the sub_status field.
func BySubStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubStatus, opts...).ToFunc()
}

// ByPayStatus orders the results by the pay_status field.
func ByPayStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPayStatus, opts...).ToFunc()
}

// BySymbolID orders the results by the symbol_id field.
func BySymbolID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSymbolID, opts...).ToFunc()
}

// ByFirstPay orders the results by the first_pay field.
func ByFirstPay(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstPay, opts...).ToFunc()
}

// ByAfterPay orders the results by the after_pay field.
func ByAfterPay(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAfterPay, opts...).ToFunc()
}

// ByLastWarningTime orders the results by the last_warning_time field.
func ByLastWarningTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastWarningTime, opts...).ToFunc()
}

// BySubFinishedTime orders the results by the sub_finished_time field.
func BySubFinishedTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubFinishedTime, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByMissionID orders the results by the mission_id field.
func ByMissionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionID, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByMissionField orders the results by mission field.
func ByMissionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newMissionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MissionTable, MissionColumn),
	)
}
