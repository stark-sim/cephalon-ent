// Code generated by ent, DO NOT EDIT.

package missionconsumeorder

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

const (
	// Label holds the string label denoting the missionconsumeorder type in the database.
	Label = "mission_consume_order"
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
	// FieldMissionID holds the string denoting the mission_id field in the database.
	FieldMissionID = "mission_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldPureCep holds the string denoting the pure_cep field in the database.
	FieldPureCep = "pure_cep"
	// FieldGiftCep holds the string denoting the gift_cep field in the database.
	FieldGiftCep = "gift_cep"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldIsTime holds the string denoting the is_time field in the database.
	FieldIsTime = "is_time"
	// FieldCallWay holds the string denoting the call_way field in the database.
	FieldCallWay = "call_way"
	// FieldSerialNumber holds the string denoting the serial_number field in the database.
	FieldSerialNumber = "serial_number"
	// FieldStartedAt holds the string denoting the started_at field in the database.
	FieldStartedAt = "started_at"
	// FieldFinishedAt holds the string denoting the finished_at field in the database.
	FieldFinishedAt = "finished_at"
	// FieldMissionBatchID holds the string denoting the mission_batch_id field in the database.
	FieldMissionBatchID = "mission_batch_id"
	// FieldMissionBatchNumber holds the string denoting the mission_batch_number field in the database.
	FieldMissionBatchNumber = "mission_batch_number"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeCostBills holds the string denoting the cost_bills edge name in mutations.
	EdgeCostBills = "cost_bills"
	// EdgeMissionProduceOrders holds the string denoting the mission_produce_orders edge name in mutations.
	EdgeMissionProduceOrders = "mission_produce_orders"
	// EdgeMissionBatch holds the string denoting the mission_batch edge name in mutations.
	EdgeMissionBatch = "mission_batch"
	// EdgeMission holds the string denoting the mission edge name in mutations.
	EdgeMission = "mission"
	// Table holds the table name of the missionconsumeorder in the database.
	Table = "mission_consume_orders"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "mission_consume_orders"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// CostBillsTable is the table that holds the cost_bills relation/edge.
	CostBillsTable = "cost_bills"
	// CostBillsInverseTable is the table name for the CostBill entity.
	// It exists in this package in order to avoid circular dependency with the "costbill" package.
	CostBillsInverseTable = "cost_bills"
	// CostBillsColumn is the table column denoting the cost_bills relation/edge.
	CostBillsColumn = "reason_id"
	// MissionProduceOrdersTable is the table that holds the mission_produce_orders relation/edge.
	MissionProduceOrdersTable = "mission_produce_orders"
	// MissionProduceOrdersInverseTable is the table name for the MissionProduceOrder entity.
	// It exists in this package in order to avoid circular dependency with the "missionproduceorder" package.
	MissionProduceOrdersInverseTable = "mission_produce_orders"
	// MissionProduceOrdersColumn is the table column denoting the mission_produce_orders relation/edge.
	MissionProduceOrdersColumn = "mission_consume_order_id"
	// MissionBatchTable is the table that holds the mission_batch relation/edge.
	MissionBatchTable = "mission_consume_orders"
	// MissionBatchInverseTable is the table name for the MissionBatch entity.
	// It exists in this package in order to avoid circular dependency with the "missionbatch" package.
	MissionBatchInverseTable = "mission_batches"
	// MissionBatchColumn is the table column denoting the mission_batch relation/edge.
	MissionBatchColumn = "mission_batch_id"
	// MissionTable is the table that holds the mission relation/edge.
	MissionTable = "mission_consume_orders"
	// MissionInverseTable is the table name for the Mission entity.
	// It exists in this package in order to avoid circular dependency with the "mission" package.
	MissionInverseTable = "missions"
	// MissionColumn is the table column denoting the mission relation/edge.
	MissionColumn = "mission_id"
)

// Columns holds all SQL columns for missionconsumeorder fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldUserID,
	FieldMissionID,
	FieldStatus,
	FieldPureCep,
	FieldGiftCep,
	FieldType,
	FieldIsTime,
	FieldCallWay,
	FieldSerialNumber,
	FieldStartedAt,
	FieldFinishedAt,
	FieldMissionBatchID,
	FieldMissionBatchNumber,
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
	// DefaultMissionID holds the default value on creation for the "mission_id" field.
	DefaultMissionID int64
	// DefaultPureCep holds the default value on creation for the "pure_cep" field.
	DefaultPureCep int64
	// DefaultGiftCep holds the default value on creation for the "gift_cep" field.
	DefaultGiftCep int64
	// DefaultIsTime holds the default value on creation for the "is_time" field.
	DefaultIsTime bool
	// DefaultSerialNumber holds the default value on creation for the "serial_number" field.
	DefaultSerialNumber string
	// DefaultStartedAt holds the default value on creation for the "started_at" field.
	DefaultStartedAt time.Time
	// DefaultFinishedAt holds the default value on creation for the "finished_at" field.
	DefaultFinishedAt time.Time
	// DefaultMissionBatchID holds the default value on creation for the "mission_batch_id" field.
	DefaultMissionBatchID int64
	// DefaultMissionBatchNumber holds the default value on creation for the "mission_batch_number" field.
	DefaultMissionBatchNumber string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

const DefaultStatus enums.MissionOrderStatus = "unknown"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s enums.MissionOrderStatus) error {
	switch s {
	case "unknown", "waiting", "canceled", "doing", "supplying", "failed", "succeed":
		return nil
	default:
		return fmt.Errorf("missionconsumeorder: invalid enum value for status field: %q", s)
	}
}

const DefaultType enums.MissionType = "unknown"

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type enums.MissionType) error {
	switch _type {
	case "unknown", "sd_time", "sd_pro_time", "txt2img", "img2img", "jp_time", "wt_time", "extra-single-image", "sd_api", "key_pair", "jp_dk_time", "ssh_time", "sd_time_plan", "sd_pro_time_plan", "wt_time_plan", "jp_time_plan", "jp_dk_time_plan", "ssh_time_plan", "sd_tomato_time", "sd_tomato_time_plan", "sd_cmd_time", "sd_cmd_time_plan", "sd_tian_time", "sd_tian_time_plan", "sd_bingo_time", "sd_bingo_time_plan", "fooocus_time", "fooocus_time_plan", "fooocus_lan_que_time", "fooocus_lan_que_time_plan", "fooocus_sha_api_time", "fooocus_sha_api_time_plan", "tabby_time", "tabby_time_plan", "ollama_time", "ollama_time_plan", "jp_conda_time", "jp_conda_time_plan", "jp_ml_time", "jp_ml_time_plan", "sd_cat_time", "sd_cat_time_plan", "sd_fire_time", "sd_fire_time_plan", "comfyui_time", "comfyui_time_plan", "comfyui_pro_time", "comfyui_pro_time_plan", "comfyui_advance_time", "comfyui_advance_time_plan", "jp_dk_3_time", "jp_dk_3_time_plan", "sd_xl_time", "sd_xl_time_plan", "sd_chick_time", "sd_chick_time_plan", "ascend_time", "ascend_time_plan", "sd_wu_shan_time", "sd_wu_shan_time_plan", "sd_lang_time", "sd_lang_time_plan", "sd_zhi_dao_time", "sd_zhi_dao_time_plan", "comfyui_ke_time", "comfyui_ke_time_plan", "chatchat_time", "chatchat_time_plan", "chat_tts_time", "chat_tts_time_plan", "lora_time", "lora_time_plan", "lora_flux_time", "lora_flux_time_plan", "fooocus_wu_time", "fooocus_wu_time_plan", "svd_back_time", "svd_back_time_plan", "sd_ji_time", "sd_ji_time_plan", "sd_shang_jin_time", "sd_shang_jin_time_plan", "sd_xiao_chun_time", "sd_xiao_chun_time_plan", "comfyui_wu_time", "comfyui_wu_time_plan", "comfyui_liu_time", "comfyui_liu_time_plan", "comfyui_li_time", "comfyui_li_time_plan", "comfyui_nenly_time", "comfyui_nenly_time_plan", "comfyui_ou_time", "comfyui_ou_time_plan", "comfyui_lu_time", "comfyui_lu_time_plan", "comfyui_jiang_time", "comfyui_jiang_time_plan", "waiting_time", "waiting_time_plan", "opencl_core_time", "opencl_core_time_plan", "io_paint_time", "io_paint_time_plan":
		return nil
	default:
		return fmt.Errorf("missionconsumeorder: invalid enum value for type field: %q", _type)
	}
}

const DefaultCallWay enums.MissionCallWay = "api"

// CallWayValidator is a validator for the "call_way" field enum values. It is called by the builders before save.
func CallWayValidator(cw enums.MissionCallWay) error {
	switch cw {
	case "unknown", "api", "yuan_hui", "dev_platform":
		return nil
	default:
		return fmt.Errorf("missionconsumeorder: invalid enum value for call_way field: %q", cw)
	}
}

// OrderOption defines the ordering options for the MissionConsumeOrder queries.
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

// ByMissionID orders the results by the mission_id field.
func ByMissionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByPureCep orders the results by the pure_cep field.
func ByPureCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPureCep, opts...).ToFunc()
}

// ByGiftCep orders the results by the gift_cep field.
func ByGiftCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGiftCep, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByIsTime orders the results by the is_time field.
func ByIsTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsTime, opts...).ToFunc()
}

// ByCallWay orders the results by the call_way field.
func ByCallWay(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCallWay, opts...).ToFunc()
}

// BySerialNumber orders the results by the serial_number field.
func BySerialNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSerialNumber, opts...).ToFunc()
}

// ByStartedAt orders the results by the started_at field.
func ByStartedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartedAt, opts...).ToFunc()
}

// ByFinishedAt orders the results by the finished_at field.
func ByFinishedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFinishedAt, opts...).ToFunc()
}

// ByMissionBatchID orders the results by the mission_batch_id field.
func ByMissionBatchID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionBatchID, opts...).ToFunc()
}

// ByMissionBatchNumber orders the results by the mission_batch_number field.
func ByMissionBatchNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionBatchNumber, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByCostBillsCount orders the results by cost_bills count.
func ByCostBillsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCostBillsStep(), opts...)
	}
}

// ByCostBills orders the results by cost_bills terms.
func ByCostBills(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCostBillsStep(), append([]sql.OrderTerm{term}, terms...)...)
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
func newCostBillsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CostBillsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CostBillsTable, CostBillsColumn),
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
func newMissionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, MissionTable, MissionColumn),
	)
}
