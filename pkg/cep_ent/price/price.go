// Code generated by ent, DO NOT EDIT.

package price

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

const (
	// Label holds the string label denoting the price type in the database.
	Label = "price"
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
	// FieldGpuVersion holds the string denoting the gpu_version field in the database.
	FieldGpuVersion = "gpu_version"
	// FieldMissionType holds the string denoting the mission_type field in the database.
	FieldMissionType = "mission_type"
	// FieldMissionCategory holds the string denoting the mission_category field in the database.
	FieldMissionCategory = "mission_category"
	// FieldMissionBillingType holds the string denoting the mission_billing_type field in the database.
	FieldMissionBillingType = "mission_billing_type"
	// FieldCep holds the string denoting the cep field in the database.
	FieldCep = "cep"
	// FieldStartedAt holds the string denoting the started_at field in the database.
	FieldStartedAt = "started_at"
	// FieldFinishedAt holds the string denoting the finished_at field in the database.
	FieldFinishedAt = "finished_at"
	// Table holds the table name of the price in the database.
	Table = "prices"
)

// Columns holds all SQL columns for price fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldGpuVersion,
	FieldMissionType,
	FieldMissionCategory,
	FieldMissionBillingType,
	FieldCep,
	FieldStartedAt,
	FieldFinishedAt,
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
	// DefaultCep holds the default value on creation for the "cep" field.
	DefaultCep int64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

const DefaultGpuVersion enums.GpuVersion = "RTX2060"

// GpuVersionValidator is a validator for the "gpu_version" field enum values. It is called by the builders before save.
func GpuVersionValidator(gv enums.GpuVersion) error {
	switch gv {
	case "RTX2060", "RTX2060Ti", "RTX2070", "RTX2070Ti", "RTX2080", "RTX2080Ti", "RTX3060", "RTX3060Ti", "RTX3070", "RTX3070Ti", "RTX3080", "RTX3080Ti", "RTX3090", "RTX3090Ti", "RTX4060", "RTX4060Ti", "RTX4070", "RTX4070Ti", "RTX4080", "RTX4090":
		return nil
	default:
		return fmt.Errorf("price: invalid enum value for gpu_version field: %q", gv)
	}
}

const DefaultMissionType enums.MissionType = "txt2img"

// MissionTypeValidator is a validator for the "mission_type" field enum values. It is called by the builders before save.
func MissionTypeValidator(mt enums.MissionType) error {
	switch mt {
	case "sd_time", "txt2img", "img2img", "jp_time", "wt_time", "extra-single-image":
		return nil
	default:
		return fmt.Errorf("price: invalid enum value for mission_type field: %q", mt)
	}
}

const DefaultMissionCategory enums.MissionCategory = "SD"

// MissionCategoryValidator is a validator for the "mission_category" field enum values. It is called by the builders before save.
func MissionCategoryValidator(mc enums.MissionCategory) error {
	switch mc {
	case "SD", "JP", "WT":
		return nil
	default:
		return fmt.Errorf("price: invalid enum value for mission_category field: %q", mc)
	}
}

const DefaultMissionBillingType enums.MissionBillingType = "count"

// MissionBillingTypeValidator is a validator for the "mission_billing_type" field enum values. It is called by the builders before save.
func MissionBillingTypeValidator(mbt enums.MissionBillingType) error {
	switch mbt {
	case "time", "count":
		return nil
	default:
		return fmt.Errorf("price: invalid enum value for mission_billing_type field: %q", mbt)
	}
}

// OrderOption defines the ordering options for the Price queries.
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

// ByGpuVersion orders the results by the gpu_version field.
func ByGpuVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGpuVersion, opts...).ToFunc()
}

// ByMissionType orders the results by the mission_type field.
func ByMissionType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionType, opts...).ToFunc()
}

// ByMissionCategory orders the results by the mission_category field.
func ByMissionCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionCategory, opts...).ToFunc()
}

// ByMissionBillingType orders the results by the mission_billing_type field.
func ByMissionBillingType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionBillingType, opts...).ToFunc()
}

// ByCep orders the results by the cep field.
func ByCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCep, opts...).ToFunc()
}

// ByStartedAt orders the results by the started_at field.
func ByStartedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartedAt, opts...).ToFunc()
}

// ByFinishedAt orders the results by the finished_at field.
func ByFinishedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFinishedAt, opts...).ToFunc()
}
