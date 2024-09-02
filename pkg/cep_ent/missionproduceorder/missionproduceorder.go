// Code generated by ent, DO NOT EDIT.

package missionproduceorder

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

const (
	// Label holds the string label denoting the missionproduceorder type in the database.
	Label = "mission_produce_order"
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
	// FieldMissionProductionID holds the string denoting the mission_production_id field in the database.
	FieldMissionProductionID = "mission_production_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldPureCep holds the string denoting the pure_cep field in the database.
	FieldPureCep = "pure_cep"
	// FieldGiftCep holds the string denoting the gift_cep field in the database.
	FieldGiftCep = "gift_cep"
	// FieldSymbolID holds the string denoting the symbol_id field in the database.
	FieldSymbolID = "symbol_id"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldIsTime holds the string denoting the is_time field in the database.
	FieldIsTime = "is_time"
	// FieldDeviceID holds the string denoting the device_id field in the database.
	FieldDeviceID = "device_id"
	// FieldSerialNumber holds the string denoting the serial_number field in the database.
	FieldSerialNumber = "serial_number"
	// FieldMissionConsumeOrderID holds the string denoting the mission_consume_order_id field in the database.
	FieldMissionConsumeOrderID = "mission_consume_order_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeEarnBills holds the string denoting the earn_bills edge name in mutations.
	EdgeEarnBills = "earn_bills"
	// EdgeDevice holds the string denoting the device edge name in mutations.
	EdgeDevice = "device"
	// EdgeMissionConsumeOrder holds the string denoting the mission_consume_order edge name in mutations.
	EdgeMissionConsumeOrder = "mission_consume_order"
	// EdgeMissionProduction holds the string denoting the mission_production edge name in mutations.
	EdgeMissionProduction = "mission_production"
	// Table holds the table name of the missionproduceorder in the database.
	Table = "mission_produce_orders"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "mission_produce_orders"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// EarnBillsTable is the table that holds the earn_bills relation/edge.
	EarnBillsTable = "earn_bills"
	// EarnBillsInverseTable is the table name for the EarnBill entity.
	// It exists in this package in order to avoid circular dependency with the "earnbill" package.
	EarnBillsInverseTable = "earn_bills"
	// EarnBillsColumn is the table column denoting the earn_bills relation/edge.
	EarnBillsColumn = "reason_id"
	// DeviceTable is the table that holds the device relation/edge.
	DeviceTable = "mission_produce_orders"
	// DeviceInverseTable is the table name for the Device entity.
	// It exists in this package in order to avoid circular dependency with the "device" package.
	DeviceInverseTable = "devices"
	// DeviceColumn is the table column denoting the device relation/edge.
	DeviceColumn = "device_id"
	// MissionConsumeOrderTable is the table that holds the mission_consume_order relation/edge.
	MissionConsumeOrderTable = "mission_produce_orders"
	// MissionConsumeOrderInverseTable is the table name for the MissionConsumeOrder entity.
	// It exists in this package in order to avoid circular dependency with the "missionconsumeorder" package.
	MissionConsumeOrderInverseTable = "mission_consume_orders"
	// MissionConsumeOrderColumn is the table column denoting the mission_consume_order relation/edge.
	MissionConsumeOrderColumn = "mission_consume_order_id"
	// MissionProductionTable is the table that holds the mission_production relation/edge.
	MissionProductionTable = "mission_produce_orders"
	// MissionProductionInverseTable is the table name for the MissionProduction entity.
	// It exists in this package in order to avoid circular dependency with the "missionproduction" package.
	MissionProductionInverseTable = "mission_productions"
	// MissionProductionColumn is the table column denoting the mission_production relation/edge.
	MissionProductionColumn = "mission_production_id"
)

// Columns holds all SQL columns for missionproduceorder fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldUserID,
	FieldMissionID,
	FieldMissionProductionID,
	FieldStatus,
	FieldPureCep,
	FieldGiftCep,
	FieldSymbolID,
	FieldAmount,
	FieldType,
	FieldIsTime,
	FieldDeviceID,
	FieldSerialNumber,
	FieldMissionConsumeOrderID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "mission_produce_orders"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"mission_mission_produce_orders",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// DefaultSymbolID holds the default value on creation for the "symbol_id" field.
	DefaultSymbolID int64
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount int64
	// DefaultIsTime holds the default value on creation for the "is_time" field.
	DefaultIsTime bool
	// DefaultDeviceID holds the default value on creation for the "device_id" field.
	DefaultDeviceID int64
	// DefaultSerialNumber holds the default value on creation for the "serial_number" field.
	DefaultSerialNumber string
	// DefaultMissionConsumeOrderID holds the default value on creation for the "mission_consume_order_id" field.
	DefaultMissionConsumeOrderID int64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

const DefaultStatus enums.MissionOrderStatus = "waiting"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s enums.MissionOrderStatus) error {
	switch s {
	case "unknown", "waiting", "canceled", "doing", "supplying", "failed", "succeed":
		return nil
	default:
		return fmt.Errorf("missionproduceorder: invalid enum value for status field: %q", s)
	}
}

const DefaultType enums.MissionType = "unknown"

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type enums.MissionType) error {
	switch _type {
	case "unknown", "sd_time", "sd_pro_time", "txt2img", "img2img", "jp_time", "wt_time", "extra-single-image", "sd_api", "key_pair", "jp_dk_time", "ssh_time", "sd_time_plan", "sd_pro_time_plan", "wt_time_plan", "jp_time_plan", "jp_dk_time_plan", "ssh_time_plan", "sd_tomato_time", "sd_tomato_time_plan", "sd_cmd_time", "sd_cmd_time_plan", "sd_tian_time", "sd_tian_time_plan", "sd_bingo_time", "sd_bingo_time_plan", "fooocus_time", "fooocus_time_plan", "fooocus_lan_que_time", "fooocus_lan_que_time_plan", "fooocus_sha_api_time", "fooocus_sha_api_time_plan", "tabby_time", "tabby_time_plan", "ollama_time", "ollama_time_plan", "jp_conda_time", "jp_conda_time_plan", "jp_ml_time", "jp_ml_time_plan", "sd_cat_time", "sd_cat_time_plan", "sd_fire_time", "sd_fire_time_plan", "comfyui_time", "comfyui_time_plan", "comfyui_pro_time", "comfyui_pro_time_plan", "comfyui_advance_time", "comfyui_advance_time_plan", "jp_dk_3_time", "jp_dk_3_time_plan", "sd_xl_time", "sd_xl_time_plan", "sd_chick_time", "sd_chick_time_plan", "ascend_time", "ascend_time_plan", "sd_wu_shan_time", "sd_wu_shan_time_plan", "sd_lang_time", "sd_lang_time_plan", "sd_zhi_dao_time", "sd_zhi_dao_time_plan", "comfyui_ke_time", "comfyui_ke_time_plan", "chatchat_time", "chatchat_time_plan", "chat_tts_time", "chat_tts_time_plan", "lora_time", "lora_time_plan", "fooocus_wu_time", "fooocus_wu_time_plan", "svd_back_time", "svd_back_time_plan", "sd_ji_time", "sd_ji_time_plan", "sd_shang_jin_time", "sd_shang_jin_time_plan", "sd_xiao_chun_time", "sd_xiao_chun_time_plan", "comfyui_wu_time", "comfyui_wu_time_plan", "comfyui_liu_time", "comfyui_liu_time_plan", "comfyui_li_time", "comfyui_li_time_plan", "comfyui_nenly_time", "comfyui_nenly_time_plan", "comfyui_ou_time", "comfyui_ou_time_plan", "comfyui_lu_time", "comfyui_lu_time_plan", "comfyui_jiang_time", "comfyui_jiang_time_plan", "waiting_time", "waiting_time_plan", "opencl_core_time", "opencl_core_time_plan", "io_paint_time", "io_paint_time_plan":
		return nil
	default:
		return fmt.Errorf("missionproduceorder: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the MissionProduceOrder queries.
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

// ByMissionProductionID orders the results by the mission_production_id field.
func ByMissionProductionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionProductionID, opts...).ToFunc()
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

// BySymbolID orders the results by the symbol_id field.
func BySymbolID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSymbolID, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByIsTime orders the results by the is_time field.
func ByIsTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsTime, opts...).ToFunc()
}

// ByDeviceID orders the results by the device_id field.
func ByDeviceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeviceID, opts...).ToFunc()
}

// BySerialNumber orders the results by the serial_number field.
func BySerialNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSerialNumber, opts...).ToFunc()
}

// ByMissionConsumeOrderID orders the results by the mission_consume_order_id field.
func ByMissionConsumeOrderID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionConsumeOrderID, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
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

// ByDeviceField orders the results by device field.
func ByDeviceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeviceStep(), sql.OrderByField(field, opts...))
	}
}

// ByMissionConsumeOrderField orders the results by mission_consume_order field.
func ByMissionConsumeOrderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionConsumeOrderStep(), sql.OrderByField(field, opts...))
	}
}

// ByMissionProductionField orders the results by mission_production field.
func ByMissionProductionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionProductionStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newEarnBillsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EarnBillsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EarnBillsTable, EarnBillsColumn),
	)
}
func newDeviceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeviceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DeviceTable, DeviceColumn),
	)
}
func newMissionConsumeOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionConsumeOrderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MissionConsumeOrderTable, MissionConsumeOrderColumn),
	)
}
func newMissionProductionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionProductionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, MissionProductionTable, MissionProductionColumn),
	)
}
