// Code generated by ent, DO NOT EDIT.

package mission

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
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
	// FieldMissionKindID holds the string denoting the mission_kind_id field in the database.
	FieldMissionKindID = "mission_kind_id"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// FieldCallBackURL holds the string denoting the call_back_url field in the database.
	FieldCallBackURL = "call_back_url"
	// FieldCallBackInfo holds the string denoting the call_back_info field in the database.
	FieldCallBackInfo = "call_back_info"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldResult holds the string denoting the result field in the database.
	FieldResult = "result"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldResultUrls holds the string denoting the result_urls field in the database.
	FieldResultUrls = "result_urls"
	// FieldUrls holds the string denoting the urls field in the database.
	FieldUrls = "urls"
	// FieldKeyPairID holds the string denoting the key_pair_id field in the database.
	FieldKeyPairID = "key_pair_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldMissionBatchID holds the string denoting the mission_batch_id field in the database.
	FieldMissionBatchID = "mission_batch_id"
	// FieldMissionBatchNumber holds the string denoting the mission_batch_number field in the database.
	FieldMissionBatchNumber = "mission_batch_number"
	// FieldGpuVersion holds the string denoting the gpu_version field in the database.
	FieldGpuVersion = "gpu_version"
	// FieldUnitCep holds the string denoting the unit_cep field in the database.
	FieldUnitCep = "unit_cep"
	// FieldRespStatusCode holds the string denoting the resp_status_code field in the database.
	FieldRespStatusCode = "resp_status_code"
	// FieldRespBody holds the string denoting the resp_body field in the database.
	FieldRespBody = "resp_body"
	// FieldInnerURI holds the string denoting the inner_uri field in the database.
	FieldInnerURI = "inner_uri"
	// FieldInnerMethod holds the string denoting the inner_method field in the database.
	FieldInnerMethod = "inner_method"
	// FieldTempHmacKey holds the string denoting the temp_hmac_key field in the database.
	FieldTempHmacKey = "temp_hmac_key"
	// FieldTempHmacSecret holds the string denoting the temp_hmac_secret field in the database.
	FieldTempHmacSecret = "temp_hmac_secret"
	// FieldSecondHmacKey holds the string denoting the second_hmac_key field in the database.
	FieldSecondHmacKey = "second_hmac_key"
	// EdgeMissionKind holds the string denoting the mission_kind edge name in mutations.
	EdgeMissionKind = "mission_kind"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeMissionKeyPairs holds the string denoting the mission_key_pairs edge name in mutations.
	EdgeMissionKeyPairs = "mission_key_pairs"
	// EdgeKeyPair holds the string denoting the key_pair edge name in mutations.
	EdgeKeyPair = "key_pair"
	// EdgeMissionConsumeOrder holds the string denoting the mission_consume_order edge name in mutations.
	EdgeMissionConsumeOrder = "mission_consume_order"
	// EdgeMissionProduceOrders holds the string denoting the mission_produce_orders edge name in mutations.
	EdgeMissionProduceOrders = "mission_produce_orders"
	// EdgeMissionBatch holds the string denoting the mission_batch edge name in mutations.
	EdgeMissionBatch = "mission_batch"
	// EdgeMissionProductions holds the string denoting the mission_productions edge name in mutations.
	EdgeMissionProductions = "mission_productions"
	// EdgeMissionOrders holds the string denoting the mission_orders edge name in mutations.
	EdgeMissionOrders = "mission_orders"
	// EdgeRenewalAgreement holds the string denoting the renewal_agreement edge name in mutations.
	EdgeRenewalAgreement = "renewal_agreement"
	// Table holds the table name of the mission in the database.
	Table = "missions"
	// MissionKindTable is the table that holds the mission_kind relation/edge.
	MissionKindTable = "missions"
	// MissionKindInverseTable is the table name for the MissionKind entity.
	// It exists in this package in order to avoid circular dependency with the "missionkind" package.
	MissionKindInverseTable = "mission_kinds"
	// MissionKindColumn is the table column denoting the mission_kind relation/edge.
	MissionKindColumn = "mission_kind_id"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "missions"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// MissionKeyPairsTable is the table that holds the mission_key_pairs relation/edge.
	MissionKeyPairsTable = "mission_key_pairs"
	// MissionKeyPairsInverseTable is the table name for the MissionKeyPair entity.
	// It exists in this package in order to avoid circular dependency with the "missionkeypair" package.
	MissionKeyPairsInverseTable = "mission_key_pairs"
	// MissionKeyPairsColumn is the table column denoting the mission_key_pairs relation/edge.
	MissionKeyPairsColumn = "mission_id"
	// KeyPairTable is the table that holds the key_pair relation/edge.
	KeyPairTable = "missions"
	// KeyPairInverseTable is the table name for the HmacKeyPair entity.
	// It exists in this package in order to avoid circular dependency with the "hmackeypair" package.
	KeyPairInverseTable = "hmac_key_pairs"
	// KeyPairColumn is the table column denoting the key_pair relation/edge.
	KeyPairColumn = "key_pair_id"
	// MissionConsumeOrderTable is the table that holds the mission_consume_order relation/edge.
	MissionConsumeOrderTable = "mission_consume_orders"
	// MissionConsumeOrderInverseTable is the table name for the MissionConsumeOrder entity.
	// It exists in this package in order to avoid circular dependency with the "missionconsumeorder" package.
	MissionConsumeOrderInverseTable = "mission_consume_orders"
	// MissionConsumeOrderColumn is the table column denoting the mission_consume_order relation/edge.
	MissionConsumeOrderColumn = "mission_id"
	// MissionProduceOrdersTable is the table that holds the mission_produce_orders relation/edge.
	MissionProduceOrdersTable = "mission_produce_orders"
	// MissionProduceOrdersInverseTable is the table name for the MissionProduceOrder entity.
	// It exists in this package in order to avoid circular dependency with the "missionproduceorder" package.
	MissionProduceOrdersInverseTable = "mission_produce_orders"
	// MissionProduceOrdersColumn is the table column denoting the mission_produce_orders relation/edge.
	MissionProduceOrdersColumn = "mission_mission_produce_orders"
	// MissionBatchTable is the table that holds the mission_batch relation/edge.
	MissionBatchTable = "missions"
	// MissionBatchInverseTable is the table name for the MissionBatch entity.
	// It exists in this package in order to avoid circular dependency with the "missionbatch" package.
	MissionBatchInverseTable = "mission_batches"
	// MissionBatchColumn is the table column denoting the mission_batch relation/edge.
	MissionBatchColumn = "mission_batch_id"
	// MissionProductionsTable is the table that holds the mission_productions relation/edge.
	MissionProductionsTable = "mission_productions"
	// MissionProductionsInverseTable is the table name for the MissionProduction entity.
	// It exists in this package in order to avoid circular dependency with the "missionproduction" package.
	MissionProductionsInverseTable = "mission_productions"
	// MissionProductionsColumn is the table column denoting the mission_productions relation/edge.
	MissionProductionsColumn = "mission_id"
	// MissionOrdersTable is the table that holds the mission_orders relation/edge.
	MissionOrdersTable = "mission_orders"
	// MissionOrdersInverseTable is the table name for the MissionOrder entity.
	// It exists in this package in order to avoid circular dependency with the "missionorder" package.
	MissionOrdersInverseTable = "mission_orders"
	// MissionOrdersColumn is the table column denoting the mission_orders relation/edge.
	MissionOrdersColumn = "mission_id"
	// RenewalAgreementTable is the table that holds the renewal_agreement relation/edge.
	RenewalAgreementTable = "renewal_agreements"
	// RenewalAgreementInverseTable is the table name for the RenewalAgreement entity.
	// It exists in this package in order to avoid circular dependency with the "renewalagreement" package.
	RenewalAgreementInverseTable = "renewal_agreements"
	// RenewalAgreementColumn is the table column denoting the renewal_agreement relation/edge.
	RenewalAgreementColumn = "mission_id"
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
	FieldMissionKindID,
	FieldBody,
	FieldCallBackURL,
	FieldCallBackInfo,
	FieldStatus,
	FieldResult,
	FieldState,
	FieldResultUrls,
	FieldUrls,
	FieldKeyPairID,
	FieldUserID,
	FieldMissionBatchID,
	FieldMissionBatchNumber,
	FieldGpuVersion,
	FieldUnitCep,
	FieldRespStatusCode,
	FieldRespBody,
	FieldInnerURI,
	FieldInnerMethod,
	FieldTempHmacKey,
	FieldTempHmacSecret,
	FieldSecondHmacKey,
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
	// DefaultMissionKindID holds the default value on creation for the "mission_kind_id" field.
	DefaultMissionKindID int64
	// DefaultBody holds the default value on creation for the "body" field.
	DefaultBody string
	// DefaultCallBackURL holds the default value on creation for the "call_back_url" field.
	DefaultCallBackURL string
	// DefaultUrls holds the default value on creation for the "urls" field.
	DefaultUrls string
	// DefaultKeyPairID holds the default value on creation for the "key_pair_id" field.
	DefaultKeyPairID int64
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID int64
	// DefaultMissionBatchID holds the default value on creation for the "mission_batch_id" field.
	DefaultMissionBatchID int64
	// DefaultMissionBatchNumber holds the default value on creation for the "mission_batch_number" field.
	DefaultMissionBatchNumber string
	// DefaultUnitCep holds the default value on creation for the "unit_cep" field.
	DefaultUnitCep int64
	// DefaultRespStatusCode holds the default value on creation for the "resp_status_code" field.
	DefaultRespStatusCode int32
	// DefaultRespBody holds the default value on creation for the "resp_body" field.
	DefaultRespBody string
	// DefaultInnerURI holds the default value on creation for the "inner_uri" field.
	DefaultInnerURI string
	// DefaultTempHmacKey holds the default value on creation for the "temp_hmac_key" field.
	DefaultTempHmacKey string
	// DefaultTempHmacSecret holds the default value on creation for the "temp_hmac_secret" field.
	DefaultTempHmacSecret string
	// DefaultSecondHmacKey holds the default value on creation for the "second_hmac_key" field.
	DefaultSecondHmacKey string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

const DefaultType enums.MissionType = "unknown"

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type enums.MissionType) error {
	switch _type {
	case "unknown", "sd_time", "txt2img", "img2img", "jp_time", "wt_time", "extra-single-image", "sd_api", "key_pair", "jp_dk_time", "ssh_time", "sd_time_plan_hour", "sd_time_plan_day", "sd_time_plan_week", "sd_time_plan_month":
		return nil
	default:
		return fmt.Errorf("mission: invalid enum value for type field: %q", _type)
	}
}

const DefaultStatus enums.MissionStatus = "waiting"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s enums.MissionStatus) error {
	switch s {
	case "waiting", "canceled", "doing", "supplying", "closing", "done":
		return nil
	default:
		return fmt.Errorf("mission: invalid enum value for status field: %q", s)
	}
}

const DefaultResult enums.MissionResult = "pending"

// ResultValidator is a validator for the "result" field enum values. It is called by the builders before save.
func ResultValidator(r enums.MissionResult) error {
	switch r {
	case "pending", "succeed", "failed":
		return nil
	default:
		return fmt.Errorf("mission: invalid enum value for result field: %q", r)
	}
}

const DefaultState enums.MissionState = "unknown"

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(s enums.MissionState) error {
	switch s {
	case "unknown", "waiting", "canceled", "doing", "supplying", "closing", "succeed", "failed":
		return nil
	default:
		return fmt.Errorf("mission: invalid enum value for state field: %q", s)
	}
}

const DefaultGpuVersion enums.GpuVersion = "unknown"

// GpuVersionValidator is a validator for the "gpu_version" field enum values. It is called by the builders before save.
func GpuVersionValidator(gv enums.GpuVersion) error {
	switch gv {
	case "unknown", "RTX2060", "RTX2060Ti", "RTX2070", "RTX2070Ti", "RTX2080", "RTX2080Ti", "RTX3060", "RTX3060Ti", "RTX3070", "RTX3070Ti", "RTX3080", "RTX3080Ti", "RTX3090", "RTX3090Ti", "RTX4060", "RTX4060Ti", "RTX4070", "RTX4070Ti", "RTX4080", "RTX4090", "A800", "A100", "V100":
		return nil
	default:
		return fmt.Errorf("mission: invalid enum value for gpu_version field: %q", gv)
	}
}

const DefaultInnerMethod enums.InnerMethod = "POST"

// InnerMethodValidator is a validator for the "inner_method" field enum values. It is called by the builders before save.
func InnerMethodValidator(im enums.InnerMethod) error {
	switch im {
	case "POST", "GET", "HEAD":
		return nil
	default:
		return fmt.Errorf("mission: invalid enum value for inner_method field: %q", im)
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

// ByMissionKindID orders the results by the mission_kind_id field.
func ByMissionKindID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionKindID, opts...).ToFunc()
}

// ByBody orders the results by the body field.
func ByBody(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBody, opts...).ToFunc()
}

// ByCallBackURL orders the results by the call_back_url field.
func ByCallBackURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCallBackURL, opts...).ToFunc()
}

// ByCallBackInfo orders the results by the call_back_info field.
func ByCallBackInfo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCallBackInfo, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByResult orders the results by the result field.
func ByResult(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldResult, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// ByUrls orders the results by the urls field.
func ByUrls(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUrls, opts...).ToFunc()
}

// ByKeyPairID orders the results by the key_pair_id field.
func ByKeyPairID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKeyPairID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByMissionBatchID orders the results by the mission_batch_id field.
func ByMissionBatchID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionBatchID, opts...).ToFunc()
}

// ByMissionBatchNumber orders the results by the mission_batch_number field.
func ByMissionBatchNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionBatchNumber, opts...).ToFunc()
}

// ByGpuVersion orders the results by the gpu_version field.
func ByGpuVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGpuVersion, opts...).ToFunc()
}

// ByUnitCep orders the results by the unit_cep field.
func ByUnitCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUnitCep, opts...).ToFunc()
}

// ByRespStatusCode orders the results by the resp_status_code field.
func ByRespStatusCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRespStatusCode, opts...).ToFunc()
}

// ByRespBody orders the results by the resp_body field.
func ByRespBody(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRespBody, opts...).ToFunc()
}

// ByInnerURI orders the results by the inner_uri field.
func ByInnerURI(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInnerURI, opts...).ToFunc()
}

// ByInnerMethod orders the results by the inner_method field.
func ByInnerMethod(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInnerMethod, opts...).ToFunc()
}

// ByTempHmacKey orders the results by the temp_hmac_key field.
func ByTempHmacKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTempHmacKey, opts...).ToFunc()
}

// ByTempHmacSecret orders the results by the temp_hmac_secret field.
func ByTempHmacSecret(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTempHmacSecret, opts...).ToFunc()
}

// BySecondHmacKey orders the results by the second_hmac_key field.
func BySecondHmacKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSecondHmacKey, opts...).ToFunc()
}

// ByMissionKindField orders the results by mission_kind field.
func ByMissionKindField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionKindStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByMissionKeyPairsCount orders the results by mission_key_pairs count.
func ByMissionKeyPairsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMissionKeyPairsStep(), opts...)
	}
}

// ByMissionKeyPairs orders the results by mission_key_pairs terms.
func ByMissionKeyPairs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionKeyPairsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByKeyPairField orders the results by key_pair field.
func ByKeyPairField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newKeyPairStep(), sql.OrderByField(field, opts...))
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

// ByMissionBatchField orders the results by mission_batch field.
func ByMissionBatchField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionBatchStep(), sql.OrderByField(field, opts...))
	}
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

// ByRenewalAgreementField orders the results by renewal_agreement field.
func ByRenewalAgreementField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRenewalAgreementStep(), sql.OrderByField(field, opts...))
	}
}
func newMissionKindStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionKindInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MissionKindTable, MissionKindColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newMissionKeyPairsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionKeyPairsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MissionKeyPairsTable, MissionKeyPairsColumn),
	)
}
func newKeyPairStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(KeyPairInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, KeyPairTable, KeyPairColumn),
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
func newMissionBatchStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionBatchInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MissionBatchTable, MissionBatchColumn),
	)
}
func newMissionProductionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionProductionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MissionProductionsTable, MissionProductionsColumn),
	)
}
func newMissionOrdersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionOrdersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MissionOrdersTable, MissionOrdersColumn),
	)
}
func newRenewalAgreementStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RenewalAgreementInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, RenewalAgreementTable, RenewalAgreementColumn),
	)
}
