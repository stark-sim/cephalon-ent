// Code generated by ent, DO NOT EDIT.

package device

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldDeletedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldUserID, v))
}

// SerialNumber applies equality check predicate on the "serial_number" field. It's identical to SerialNumberEQ.
func SerialNumber(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldSerialNumber, v))
}

// SumCep applies equality check predicate on the "sum_cep" field. It's identical to SumCepEQ.
func SumCep(v int64) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldSumCep, v))
}

// Linking applies equality check predicate on the "linking" field. It's identical to LinkingEQ.
func Linking(v bool) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldLinking, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldName, v))
}

// CoresNumber applies equality check predicate on the "cores_number" field. It's identical to CoresNumberEQ.
func CoresNumber(v int64) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldCoresNumber, v))
}

// CPU applies equality check predicate on the "cpu" field. It's identical to CPUEQ.
func CPU(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldCPU, v))
}

// Memory applies equality check predicate on the "memory" field. It's identical to MemoryEQ.
func Memory(v int64) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldMemory, v))
}

// Disk applies equality check predicate on the "disk" field. It's identical to DiskEQ.
func Disk(v float32) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldDisk, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldDeletedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldUserID, vs...))
}

// SerialNumberEQ applies the EQ predicate on the "serial_number" field.
func SerialNumberEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldSerialNumber, v))
}

// SerialNumberNEQ applies the NEQ predicate on the "serial_number" field.
func SerialNumberNEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldSerialNumber, v))
}

// SerialNumberIn applies the In predicate on the "serial_number" field.
func SerialNumberIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldSerialNumber, vs...))
}

// SerialNumberNotIn applies the NotIn predicate on the "serial_number" field.
func SerialNumberNotIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldSerialNumber, vs...))
}

// SerialNumberGT applies the GT predicate on the "serial_number" field.
func SerialNumberGT(v string) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldSerialNumber, v))
}

// SerialNumberGTE applies the GTE predicate on the "serial_number" field.
func SerialNumberGTE(v string) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldSerialNumber, v))
}

// SerialNumberLT applies the LT predicate on the "serial_number" field.
func SerialNumberLT(v string) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldSerialNumber, v))
}

// SerialNumberLTE applies the LTE predicate on the "serial_number" field.
func SerialNumberLTE(v string) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldSerialNumber, v))
}

// SerialNumberContains applies the Contains predicate on the "serial_number" field.
func SerialNumberContains(v string) predicate.Device {
	return predicate.Device(sql.FieldContains(FieldSerialNumber, v))
}

// SerialNumberHasPrefix applies the HasPrefix predicate on the "serial_number" field.
func SerialNumberHasPrefix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasPrefix(FieldSerialNumber, v))
}

// SerialNumberHasSuffix applies the HasSuffix predicate on the "serial_number" field.
func SerialNumberHasSuffix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasSuffix(FieldSerialNumber, v))
}

// SerialNumberEqualFold applies the EqualFold predicate on the "serial_number" field.
func SerialNumberEqualFold(v string) predicate.Device {
	return predicate.Device(sql.FieldEqualFold(FieldSerialNumber, v))
}

// SerialNumberContainsFold applies the ContainsFold predicate on the "serial_number" field.
func SerialNumberContainsFold(v string) predicate.Device {
	return predicate.Device(sql.FieldContainsFold(FieldSerialNumber, v))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v State) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldState, v))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v State) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldState, v))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...State) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldState, vs...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...State) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldState, vs...))
}

// SumCepEQ applies the EQ predicate on the "sum_cep" field.
func SumCepEQ(v int64) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldSumCep, v))
}

// SumCepNEQ applies the NEQ predicate on the "sum_cep" field.
func SumCepNEQ(v int64) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldSumCep, v))
}

// SumCepIn applies the In predicate on the "sum_cep" field.
func SumCepIn(vs ...int64) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldSumCep, vs...))
}

// SumCepNotIn applies the NotIn predicate on the "sum_cep" field.
func SumCepNotIn(vs ...int64) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldSumCep, vs...))
}

// SumCepGT applies the GT predicate on the "sum_cep" field.
func SumCepGT(v int64) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldSumCep, v))
}

// SumCepGTE applies the GTE predicate on the "sum_cep" field.
func SumCepGTE(v int64) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldSumCep, v))
}

// SumCepLT applies the LT predicate on the "sum_cep" field.
func SumCepLT(v int64) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldSumCep, v))
}

// SumCepLTE applies the LTE predicate on the "sum_cep" field.
func SumCepLTE(v int64) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldSumCep, v))
}

// LinkingEQ applies the EQ predicate on the "linking" field.
func LinkingEQ(v bool) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldLinking, v))
}

// LinkingNEQ applies the NEQ predicate on the "linking" field.
func LinkingNEQ(v bool) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldLinking, v))
}

// BindingStatusEQ applies the EQ predicate on the "binding_status" field.
func BindingStatusEQ(v enums.DeviceBindingStatus) predicate.Device {
	vc := v
	return predicate.Device(sql.FieldEQ(FieldBindingStatus, vc))
}

// BindingStatusNEQ applies the NEQ predicate on the "binding_status" field.
func BindingStatusNEQ(v enums.DeviceBindingStatus) predicate.Device {
	vc := v
	return predicate.Device(sql.FieldNEQ(FieldBindingStatus, vc))
}

// BindingStatusIn applies the In predicate on the "binding_status" field.
func BindingStatusIn(vs ...enums.DeviceBindingStatus) predicate.Device {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(sql.FieldIn(FieldBindingStatus, v...))
}

// BindingStatusNotIn applies the NotIn predicate on the "binding_status" field.
func BindingStatusNotIn(vs ...enums.DeviceBindingStatus) predicate.Device {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(sql.FieldNotIn(FieldBindingStatus, v...))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldStatus, vs...))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Device {
	return predicate.Device(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Device {
	return predicate.Device(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Device {
	return predicate.Device(sql.FieldContainsFold(FieldName, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v enums.DeviceType) predicate.Device {
	vc := v
	return predicate.Device(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v enums.DeviceType) predicate.Device {
	vc := v
	return predicate.Device(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...enums.DeviceType) predicate.Device {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...enums.DeviceType) predicate.Device {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(sql.FieldNotIn(FieldType, v...))
}

// CoresNumberEQ applies the EQ predicate on the "cores_number" field.
func CoresNumberEQ(v int64) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldCoresNumber, v))
}

// CoresNumberNEQ applies the NEQ predicate on the "cores_number" field.
func CoresNumberNEQ(v int64) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldCoresNumber, v))
}

// CoresNumberIn applies the In predicate on the "cores_number" field.
func CoresNumberIn(vs ...int64) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldCoresNumber, vs...))
}

// CoresNumberNotIn applies the NotIn predicate on the "cores_number" field.
func CoresNumberNotIn(vs ...int64) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldCoresNumber, vs...))
}

// CoresNumberGT applies the GT predicate on the "cores_number" field.
func CoresNumberGT(v int64) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldCoresNumber, v))
}

// CoresNumberGTE applies the GTE predicate on the "cores_number" field.
func CoresNumberGTE(v int64) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldCoresNumber, v))
}

// CoresNumberLT applies the LT predicate on the "cores_number" field.
func CoresNumberLT(v int64) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldCoresNumber, v))
}

// CoresNumberLTE applies the LTE predicate on the "cores_number" field.
func CoresNumberLTE(v int64) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldCoresNumber, v))
}

// CPUEQ applies the EQ predicate on the "cpu" field.
func CPUEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldCPU, v))
}

// CPUNEQ applies the NEQ predicate on the "cpu" field.
func CPUNEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldCPU, v))
}

// CPUIn applies the In predicate on the "cpu" field.
func CPUIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldCPU, vs...))
}

// CPUNotIn applies the NotIn predicate on the "cpu" field.
func CPUNotIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldCPU, vs...))
}

// CPUGT applies the GT predicate on the "cpu" field.
func CPUGT(v string) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldCPU, v))
}

// CPUGTE applies the GTE predicate on the "cpu" field.
func CPUGTE(v string) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldCPU, v))
}

// CPULT applies the LT predicate on the "cpu" field.
func CPULT(v string) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldCPU, v))
}

// CPULTE applies the LTE predicate on the "cpu" field.
func CPULTE(v string) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldCPU, v))
}

// CPUContains applies the Contains predicate on the "cpu" field.
func CPUContains(v string) predicate.Device {
	return predicate.Device(sql.FieldContains(FieldCPU, v))
}

// CPUHasPrefix applies the HasPrefix predicate on the "cpu" field.
func CPUHasPrefix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasPrefix(FieldCPU, v))
}

// CPUHasSuffix applies the HasSuffix predicate on the "cpu" field.
func CPUHasSuffix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasSuffix(FieldCPU, v))
}

// CPUEqualFold applies the EqualFold predicate on the "cpu" field.
func CPUEqualFold(v string) predicate.Device {
	return predicate.Device(sql.FieldEqualFold(FieldCPU, v))
}

// CPUContainsFold applies the ContainsFold predicate on the "cpu" field.
func CPUContainsFold(v string) predicate.Device {
	return predicate.Device(sql.FieldContainsFold(FieldCPU, v))
}

// CpusIsNil applies the IsNil predicate on the "cpus" field.
func CpusIsNil() predicate.Device {
	return predicate.Device(sql.FieldIsNull(FieldCpus))
}

// CpusNotNil applies the NotNil predicate on the "cpus" field.
func CpusNotNil() predicate.Device {
	return predicate.Device(sql.FieldNotNull(FieldCpus))
}

// MemoryEQ applies the EQ predicate on the "memory" field.
func MemoryEQ(v int64) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldMemory, v))
}

// MemoryNEQ applies the NEQ predicate on the "memory" field.
func MemoryNEQ(v int64) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldMemory, v))
}

// MemoryIn applies the In predicate on the "memory" field.
func MemoryIn(vs ...int64) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldMemory, vs...))
}

// MemoryNotIn applies the NotIn predicate on the "memory" field.
func MemoryNotIn(vs ...int64) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldMemory, vs...))
}

// MemoryGT applies the GT predicate on the "memory" field.
func MemoryGT(v int64) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldMemory, v))
}

// MemoryGTE applies the GTE predicate on the "memory" field.
func MemoryGTE(v int64) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldMemory, v))
}

// MemoryLT applies the LT predicate on the "memory" field.
func MemoryLT(v int64) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldMemory, v))
}

// MemoryLTE applies the LTE predicate on the "memory" field.
func MemoryLTE(v int64) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldMemory, v))
}

// DiskEQ applies the EQ predicate on the "disk" field.
func DiskEQ(v float32) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldDisk, v))
}

// DiskNEQ applies the NEQ predicate on the "disk" field.
func DiskNEQ(v float32) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldDisk, v))
}

// DiskIn applies the In predicate on the "disk" field.
func DiskIn(vs ...float32) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldDisk, vs...))
}

// DiskNotIn applies the NotIn predicate on the "disk" field.
func DiskNotIn(vs ...float32) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldDisk, vs...))
}

// DiskGT applies the GT predicate on the "disk" field.
func DiskGT(v float32) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldDisk, v))
}

// DiskGTE applies the GTE predicate on the "disk" field.
func DiskGTE(v float32) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldDisk, v))
}

// DiskLT applies the LT predicate on the "disk" field.
func DiskLT(v float32) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldDisk, v))
}

// DiskLTE applies the LTE predicate on the "disk" field.
func DiskLTE(v float32) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldDisk, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMissionProduceOrders applies the HasEdge predicate on the "mission_produce_orders" edge.
func HasMissionProduceOrders() predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MissionProduceOrdersTable, MissionProduceOrdersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMissionProduceOrdersWith applies the HasEdge predicate on the "mission_produce_orders" edge with a given conditions (other predicates).
func HasMissionProduceOrdersWith(preds ...predicate.MissionProduceOrder) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := newMissionProduceOrdersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserDevices applies the HasEdge predicate on the "user_devices" edge.
func HasUserDevices() predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserDevicesTable, UserDevicesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserDevicesWith applies the HasEdge predicate on the "user_devices" edge with a given conditions (other predicates).
func HasUserDevicesWith(preds ...predicate.UserDevice) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := newUserDevicesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDeviceGpuMissions applies the HasEdge predicate on the "device_gpu_missions" edge.
func HasDeviceGpuMissions() predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DeviceGpuMissionsTable, DeviceGpuMissionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeviceGpuMissionsWith applies the HasEdge predicate on the "device_gpu_missions" edge with a given conditions (other predicates).
func HasDeviceGpuMissionsWith(preds ...predicate.DeviceGpuMission) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := newDeviceGpuMissionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFrpcInfos applies the HasEdge predicate on the "frpc_infos" edge.
func HasFrpcInfos() predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FrpcInfosTable, FrpcInfosColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFrpcInfosWith applies the HasEdge predicate on the "frpc_infos" edge with a given conditions (other predicates).
func HasFrpcInfosWith(preds ...predicate.FrpcInfo) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := newFrpcInfosStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMissionOrders applies the HasEdge predicate on the "mission_orders" edge.
func HasMissionOrders() predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MissionOrdersTable, MissionOrdersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMissionOrdersWith applies the HasEdge predicate on the "mission_orders" edge with a given conditions (other predicates).
func HasMissionOrdersWith(preds ...predicate.MissionOrder) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := newMissionOrdersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Device) predicate.Device {
	return predicate.Device(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Device) predicate.Device {
	return predicate.Device(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Device) predicate.Device {
	return predicate.Device(sql.NotPredicates(p))
}
