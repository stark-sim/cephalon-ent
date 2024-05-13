// Code generated by ent, DO NOT EDIT.

package missionkind

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

const (
	// Label holds the string label denoting the missionkind type in the database.
	Label = "mission_kind"
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
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldBillingType holds the string denoting the billing_type field in the database.
	FieldBillingType = "billing_type"
	// EdgeMissions holds the string denoting the missions edge name in mutations.
	EdgeMissions = "missions"
	// Table holds the table name of the missionkind in the database.
	Table = "mission_kinds"
	// MissionsTable is the table that holds the missions relation/edge.
	MissionsTable = "missions"
	// MissionsInverseTable is the table name for the Mission entity.
	// It exists in this package in order to avoid circular dependency with the "mission" package.
	MissionsInverseTable = "missions"
	// MissionsColumn is the table column denoting the missions relation/edge.
	MissionsColumn = "mission_kind_id"
)

// Columns holds all SQL columns for missionkind fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldType,
	FieldCategory,
	FieldBillingType,
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

const DefaultType enums.MissionType = "unknown"

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type enums.MissionType) error {
	switch _type {
	case "unknown", "sd_time", "txt2img", "img2img", "jp_time", "wt_time", "extra-single-image", "sd_api", "key_pair", "jp_dk_time", "ssh_time", "sd_time_plan", "wt_time_plan", "jp_time_plan", "jp_dk_time_plan", "ssh_time_plan", "sd_tomato_time", "sd_tomato_time_plan", "sd_cmd_time", "sd_cmd_time_plan", "sd_bingo_time", "sd_bingo_time_plan", "fooocus_time", "fooocus_time_plan", "fooocus_lan_que_time", "fooocus_lan_que_time_plan", "tabby_time", "tabby_time_plan", "jp_conda_time", "jp_conda_time_plan", "jp_ml_time", "jp_ml_time_plan", "sd_cat_time", "sd_cat_time_plan", "sd_fire_time", "sd_fire_time_plan", "comfyui_time", "comfyui_time_plan", "comfyui_pro_time", "comfyui_pro_time_plan", "jp_dk_3_time", "jp_dk_3_time_plan", "sd_xl_time", "sd_xl_time_plan", "sd_chick_time", "sd_chick_time_plan", "ascend_time", "ascend_time_plan", "sd_wu_shan_time", "sd_wu_shan_time_plan", "sd_lang_time", "sd_lang_time_plan", "comfyui_ke_time", "comfyui_ke_time_plan", "chatchat_time", "chatchat_time_plan", "lora_time", "lora_time_plan", "fooocus_wu_time", "fooocus_wu_time_plan", "svd_back_time", "svd_back_time_plan", "sd_ji_time", "sd_ji_time_plan", "sd_shang_jin_time", "sd_shang_jin_time_plan", "sd_xiao_chun_time", "sd_xiao_chun_time_plan", "comfyui_wu_time", "comfyui_wu_time_plan", "comfyui_liu_time", "comfyui_liu_time_plan", "waiting_time", "waiting_time_plan":
		return nil
	default:
		return fmt.Errorf("missionkind: invalid enum value for type field: %q", _type)
	}
}

const DefaultCategory enums.MissionCategory = "unknown"

// CategoryValidator is a validator for the "category" field enum values. It is called by the builders before save.
func CategoryValidator(c enums.MissionCategory) error {
	switch c {
	case "unknown", "SD", "JP", "WT", "JP_DK", "SSH", "SD_TOMATO", "SD_CMD", "SD_BINGO", "FOOOCUS", "FOOOCUS_LAN_QUE", "TABBY", "JP_CONDA", "SD_CAT", "SD_FIRE", "COMFYUI", "SD_XL", "SD_CHICK", "ASCEND", "SD_WU_SHAN", "SD_LANG", "COMFYUI_KE", "CHATCHAT", "LORA", "FOOOCUS_WU", "SVD_BACK", "SD_JI", "SD_SHANG_JIN", "SD_XIAO_CHUN", "COMFYUI_WU", "COMFYUI_LIU", "WAITING":
		return nil
	default:
		return fmt.Errorf("missionkind: invalid enum value for category field: %q", c)
	}
}

const DefaultBillingType enums.MissionBillingType = "unknown"

// BillingTypeValidator is a validator for the "billing_type" field enum values. It is called by the builders before save.
func BillingTypeValidator(bt enums.MissionBillingType) error {
	switch bt {
	case "unknown", "time", "count", "hold", "volume", "time_plan_hour", "time_plan_day", "time_plan_week", "time_plan_month":
		return nil
	default:
		return fmt.Errorf("missionkind: invalid enum value for billing_type field: %q", bt)
	}
}

// OrderOption defines the ordering options for the MissionKind queries.
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

// ByCategory orders the results by the category field.
func ByCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategory, opts...).ToFunc()
}

// ByBillingType orders the results by the billing_type field.
func ByBillingType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBillingType, opts...).ToFunc()
}

// ByMissionsCount orders the results by missions count.
func ByMissionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMissionsStep(), opts...)
	}
}

// ByMissions orders the results by missions terms.
func ByMissions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newMissionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MissionsTable, MissionsColumn),
	)
}
