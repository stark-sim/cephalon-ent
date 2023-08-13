// Code generated by ent, DO NOT EDIT.

package mission

import (
	"cephalon-ent/pkg/enums"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the mission type in the database.
	Label = "mission"
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
	// FieldIsTime holds the string denoting the is_time field in the database.
	FieldIsTime = "is_time"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// FieldCallBackURL holds the string denoting the call_back_url field in the database.
	FieldCallBackURL = "call_back_url"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldResultUrls holds the string denoting the result_urls field in the database.
	FieldResultUrls = "result_urls"
	// FieldAdditionalResult holds the string denoting the additional_result field in the database.
	FieldAdditionalResult = "additional_result"
	// FieldHmacKeyPairID holds the string denoting the hmac_key_pair_id field in the database.
	FieldHmacKeyPairID = "hmac_key_pair_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldMissionBatchNumber holds the string denoting the mission_batch_number field in the database.
	FieldMissionBatchNumber = "mission_batch_number"
	// FieldMissionBatchID holds the string denoting the mission_batch_id field in the database.
	FieldMissionBatchID = "mission_batch_id"
	// EdgeMissionProductions holds the string denoting the mission_productions edge name in mutations.
	EdgeMissionProductions = "mission_productions"
	// EdgeMissionConsumeOrder holds the string denoting the mission_consume_order edge name in mutations.
	EdgeMissionConsumeOrder = "mission_consume_order"
	// EdgeMissionProduceOrders holds the string denoting the mission_produce_orders edge name in mutations.
	EdgeMissionProduceOrders = "mission_produce_orders"
	// EdgeHmacKeyPair holds the string denoting the hmac_key_pair edge name in mutations.
	EdgeHmacKeyPair = "hmac_key_pair"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeMissionBatch holds the string denoting the mission_batch edge name in mutations.
	EdgeMissionBatch = "mission_batch"
	// Table holds the table name of the mission in the database.
	Table = "missions"
	// MissionProductionsTable is the table that holds the mission_productions relation/edge.
	MissionProductionsTable = "mission_productions"
	// MissionProductionsInverseTable is the table name for the MissionProduction entity.
	// It exists in this package in order to avoid circular dependency with the "missionproduction" package.
	MissionProductionsInverseTable = "mission_productions"
	// MissionProductionsColumn is the table column denoting the mission_productions relation/edge.
	MissionProductionsColumn = "mission_id"
	// MissionConsumeOrderTable is the table that holds the mission_consume_order relation/edge.
	MissionConsumeOrderTable = "mission_consume_orders"
	// MissionConsumeOrderInverseTable is the table name for the MissionConsumeOrder entity.
	// It exists in this package in order to avoid circular dependency with the "missionconsumeorder" package.
	MissionConsumeOrderInverseTable = "mission_consume_orders"
	// MissionConsumeOrderColumn is the table column denoting the mission_consume_order relation/edge.
	MissionConsumeOrderColumn = "mission_mission_consume_order"
	// MissionProduceOrdersTable is the table that holds the mission_produce_orders relation/edge.
	MissionProduceOrdersTable = "mission_produce_orders"
	// MissionProduceOrdersInverseTable is the table name for the MissionProduceOrder entity.
	// It exists in this package in order to avoid circular dependency with the "missionproduceorder" package.
	MissionProduceOrdersInverseTable = "mission_produce_orders"
	// MissionProduceOrdersColumn is the table column denoting the mission_produce_orders relation/edge.
	MissionProduceOrdersColumn = "mission_id"
	// HmacKeyPairTable is the table that holds the hmac_key_pair relation/edge.
	HmacKeyPairTable = "missions"
	// HmacKeyPairInverseTable is the table name for the HmacKeyPair entity.
	// It exists in this package in order to avoid circular dependency with the "hmackeypair" package.
	HmacKeyPairInverseTable = "hmac_key_pairs"
	// HmacKeyPairColumn is the table column denoting the hmac_key_pair relation/edge.
	HmacKeyPairColumn = "hmac_key_pair_id"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "missions"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// MissionBatchTable is the table that holds the mission_batch relation/edge.
	MissionBatchTable = "missions"
	// MissionBatchInverseTable is the table name for the MissionBatch entity.
	// It exists in this package in order to avoid circular dependency with the "missionbatch" package.
	MissionBatchInverseTable = "mission_batches"
	// MissionBatchColumn is the table column denoting the mission_batch relation/edge.
	MissionBatchColumn = "mission_batch_id"
)

// Columns holds all SQL columns for mission fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldType,
	FieldIsTime,
	FieldBody,
	FieldCallBackURL,
	FieldStatus,
	FieldResultUrls,
	FieldAdditionalResult,
	FieldHmacKeyPairID,
	FieldUserID,
	FieldMissionBatchNumber,
	FieldMissionBatchID,
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
	// DefaultIsTime holds the default value on creation for the "is_time" field.
	DefaultIsTime bool
	// DefaultBody holds the default value on creation for the "body" field.
	DefaultBody string
	// DefaultCallBackURL holds the default value on creation for the "call_back_url" field.
	DefaultCallBackURL string
	// DefaultResultUrls holds the default value on creation for the "result_urls" field.
	DefaultResultUrls string
	// DefaultAdditionalResult holds the default value on creation for the "additional_result" field.
	DefaultAdditionalResult string
	// DefaultHmacKeyPairID holds the default value on creation for the "hmac_key_pair_id" field.
	DefaultHmacKeyPairID int64
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID int64
	// DefaultMissionBatchNumber holds the default value on creation for the "mission_batch_number" field.
	DefaultMissionBatchNumber string
	// DefaultMissionBatchID holds the default value on creation for the "mission_batch_id" field.
	DefaultMissionBatchID int64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

const DefaultType enums.MissionType = "txt2img"

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type enums.MissionType) error {
	switch _type {
	case "sd_time", "txt2img", "img2img", "jp_time", "wt_time":
		return nil
	default:
		return fmt.Errorf("mission: invalid enum value for type field: %q", _type)
	}
}

// Status defines the type for the "status" enum field.
type Status string

// StatusWaiting is the default value of the Status enum.
const DefaultStatus = StatusWaiting

// Status values.
const (
	StatusWaiting   Status = "waiting"
	StatusCanceled  Status = "canceled"
	StatusDoing     Status = "doing"
	StatusSupplying Status = "supplying"
	StatusClosing   Status = "closing"
	StatusSucceed   Status = "succeed"
	StatusFailed    Status = "failed"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusWaiting, StatusCanceled, StatusDoing, StatusSupplying, StatusClosing, StatusSucceed, StatusFailed:
		return nil
	default:
		return fmt.Errorf("mission: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Mission queries.
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

// ByIsTime orders the results by the is_time field.
func ByIsTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsTime, opts...).ToFunc()
}

// ByBody orders the results by the body field.
func ByBody(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBody, opts...).ToFunc()
}

// ByCallBackURL orders the results by the call_back_url field.
func ByCallBackURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCallBackURL, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByResultUrls orders the results by the result_urls field.
func ByResultUrls(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldResultUrls, opts...).ToFunc()
}

// ByAdditionalResult orders the results by the additional_result field.
func ByAdditionalResult(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAdditionalResult, opts...).ToFunc()
}

// ByHmacKeyPairID orders the results by the hmac_key_pair_id field.
func ByHmacKeyPairID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHmacKeyPairID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByMissionBatchNumber orders the results by the mission_batch_number field.
func ByMissionBatchNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionBatchNumber, opts...).ToFunc()
}

// ByMissionBatchID orders the results by the mission_batch_id field.
func ByMissionBatchID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionBatchID, opts...).ToFunc()
}

// ByMissionProductionsCount orders the results by mission_productions count.
func ByMissionProductionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMissionProductionsStep(), opts...)
	}
}

// ByMissionProductions orders the results by mission_productions terms.
func ByMissionProductions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionProductionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMissionConsumeOrderField orders the results by mission_consume_order field.
func ByMissionConsumeOrderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionConsumeOrderStep(), sql.OrderByField(field, opts...))
	}
}

// ByMissionProduceOrdersCount orders the results by mission_produce_orders count.
func ByMissionProduceOrdersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMissionProduceOrdersStep(), opts...)
	}
}

// ByMissionProduceOrders orders the results by mission_produce_orders terms.
func ByMissionProduceOrders(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionProduceOrdersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByHmacKeyPairField orders the results by hmac_key_pair field.
func ByHmacKeyPairField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHmacKeyPairStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByMissionBatchField orders the results by mission_batch field.
func ByMissionBatchField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionBatchStep(), sql.OrderByField(field, opts...))
	}
}
func newMissionProductionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionProductionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MissionProductionsTable, MissionProductionsColumn),
	)
}
func newMissionConsumeOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionConsumeOrderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, MissionConsumeOrderTable, MissionConsumeOrderColumn),
	)
}
func newMissionProduceOrdersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionProduceOrdersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MissionProduceOrdersTable, MissionProduceOrdersColumn),
	)
}
func newHmacKeyPairStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HmacKeyPairInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, HmacKeyPairTable, HmacKeyPairColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newMissionBatchStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionBatchInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MissionBatchTable, MissionBatchColumn),
	)
}
