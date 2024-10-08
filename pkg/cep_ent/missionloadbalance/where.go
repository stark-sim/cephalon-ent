// Code generated by ent, DO NOT EDIT.

package missionloadbalance

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldDeletedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldUserID, v))
}

// StartedAt applies equality check predicate on the "started_at" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldStartedAt, v))
}

// FinishedAt applies equality check predicate on the "finished_at" field. It's identical to FinishedAtEQ.
func FinishedAt(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldFinishedAt, v))
}

// GpuNum applies equality check predicate on the "gpu_num" field. It's identical to GpuNumEQ.
func GpuNum(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldGpuNum, v))
}

// MaxMissionCount applies equality check predicate on the "max_mission_count" field. It's identical to MaxMissionCountEQ.
func MaxMissionCount(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldMaxMissionCount, v))
}

// MinMissionCount applies equality check predicate on the "min_mission_count" field. It's identical to MinMissionCountEQ.
func MinMissionCount(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldMinMissionCount, v))
}

// CurrentMissionCount applies equality check predicate on the "current_mission_count" field. It's identical to CurrentMissionCountEQ.
func CurrentMissionCount(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldCurrentMissionCount, v))
}

// MissionBatchID applies equality check predicate on the "mission_batch_id" field. It's identical to MissionBatchIDEQ.
func MissionBatchID(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldMissionBatchID, v))
}

// MissionBatchNumber applies equality check predicate on the "mission_batch_number" field. It's identical to MissionBatchNumberEQ.
func MissionBatchNumber(v string) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldMissionBatchNumber, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLTE(FieldDeletedAt, v))
}

// MissionTypeEQ applies the EQ predicate on the "mission_type" field.
func MissionTypeEQ(v enums.MissionType) predicate.MissionLoadBalance {
	vc := v
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldMissionType, vc))
}

// MissionTypeNEQ applies the NEQ predicate on the "mission_type" field.
func MissionTypeNEQ(v enums.MissionType) predicate.MissionLoadBalance {
	vc := v
	return predicate.MissionLoadBalance(sql.FieldNEQ(FieldMissionType, vc))
}

// MissionTypeIn applies the In predicate on the "mission_type" field.
func MissionTypeIn(vs ...enums.MissionType) predicate.MissionLoadBalance {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MissionLoadBalance(sql.FieldIn(FieldMissionType, v...))
}

// MissionTypeNotIn applies the NotIn predicate on the "mission_type" field.
func MissionTypeNotIn(vs ...enums.MissionType) predicate.MissionLoadBalance {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MissionLoadBalance(sql.FieldNotIn(FieldMissionType, v...))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLTE(FieldUserID, v))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v enums.MissionLoadBalanceState) predicate.MissionLoadBalance {
	vc := v
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldState, vc))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v enums.MissionLoadBalanceState) predicate.MissionLoadBalance {
	vc := v
	return predicate.MissionLoadBalance(sql.FieldNEQ(FieldState, vc))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...enums.MissionLoadBalanceState) predicate.MissionLoadBalance {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MissionLoadBalance(sql.FieldIn(FieldState, v...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...enums.MissionLoadBalanceState) predicate.MissionLoadBalance {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MissionLoadBalance(sql.FieldNotIn(FieldState, v...))
}

// StartedAtEQ applies the EQ predicate on the "started_at" field.
func StartedAtEQ(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldStartedAt, v))
}

// StartedAtNEQ applies the NEQ predicate on the "started_at" field.
func StartedAtNEQ(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNEQ(FieldStartedAt, v))
}

// StartedAtIn applies the In predicate on the "started_at" field.
func StartedAtIn(vs ...time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldIn(FieldStartedAt, vs...))
}

// StartedAtNotIn applies the NotIn predicate on the "started_at" field.
func StartedAtNotIn(vs ...time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNotIn(FieldStartedAt, vs...))
}

// StartedAtGT applies the GT predicate on the "started_at" field.
func StartedAtGT(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGT(FieldStartedAt, v))
}

// StartedAtGTE applies the GTE predicate on the "started_at" field.
func StartedAtGTE(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGTE(FieldStartedAt, v))
}

// StartedAtLT applies the LT predicate on the "started_at" field.
func StartedAtLT(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLT(FieldStartedAt, v))
}

// StartedAtLTE applies the LTE predicate on the "started_at" field.
func StartedAtLTE(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLTE(FieldStartedAt, v))
}

// FinishedAtEQ applies the EQ predicate on the "finished_at" field.
func FinishedAtEQ(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldFinishedAt, v))
}

// FinishedAtNEQ applies the NEQ predicate on the "finished_at" field.
func FinishedAtNEQ(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNEQ(FieldFinishedAt, v))
}

// FinishedAtIn applies the In predicate on the "finished_at" field.
func FinishedAtIn(vs ...time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldIn(FieldFinishedAt, vs...))
}

// FinishedAtNotIn applies the NotIn predicate on the "finished_at" field.
func FinishedAtNotIn(vs ...time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNotIn(FieldFinishedAt, vs...))
}

// FinishedAtGT applies the GT predicate on the "finished_at" field.
func FinishedAtGT(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGT(FieldFinishedAt, v))
}

// FinishedAtGTE applies the GTE predicate on the "finished_at" field.
func FinishedAtGTE(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGTE(FieldFinishedAt, v))
}

// FinishedAtLT applies the LT predicate on the "finished_at" field.
func FinishedAtLT(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLT(FieldFinishedAt, v))
}

// FinishedAtLTE applies the LTE predicate on the "finished_at" field.
func FinishedAtLTE(v time.Time) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLTE(FieldFinishedAt, v))
}

// GpuVersionEQ applies the EQ predicate on the "gpu_version" field.
func GpuVersionEQ(v enums.GpuVersion) predicate.MissionLoadBalance {
	vc := v
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldGpuVersion, vc))
}

// GpuVersionNEQ applies the NEQ predicate on the "gpu_version" field.
func GpuVersionNEQ(v enums.GpuVersion) predicate.MissionLoadBalance {
	vc := v
	return predicate.MissionLoadBalance(sql.FieldNEQ(FieldGpuVersion, vc))
}

// GpuVersionIn applies the In predicate on the "gpu_version" field.
func GpuVersionIn(vs ...enums.GpuVersion) predicate.MissionLoadBalance {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MissionLoadBalance(sql.FieldIn(FieldGpuVersion, v...))
}

// GpuVersionNotIn applies the NotIn predicate on the "gpu_version" field.
func GpuVersionNotIn(vs ...enums.GpuVersion) predicate.MissionLoadBalance {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MissionLoadBalance(sql.FieldNotIn(FieldGpuVersion, v...))
}

// GpuNumEQ applies the EQ predicate on the "gpu_num" field.
func GpuNumEQ(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldGpuNum, v))
}

// GpuNumNEQ applies the NEQ predicate on the "gpu_num" field.
func GpuNumNEQ(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNEQ(FieldGpuNum, v))
}

// GpuNumIn applies the In predicate on the "gpu_num" field.
func GpuNumIn(vs ...int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldIn(FieldGpuNum, vs...))
}

// GpuNumNotIn applies the NotIn predicate on the "gpu_num" field.
func GpuNumNotIn(vs ...int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNotIn(FieldGpuNum, vs...))
}

// GpuNumGT applies the GT predicate on the "gpu_num" field.
func GpuNumGT(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGT(FieldGpuNum, v))
}

// GpuNumGTE applies the GTE predicate on the "gpu_num" field.
func GpuNumGTE(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGTE(FieldGpuNum, v))
}

// GpuNumLT applies the LT predicate on the "gpu_num" field.
func GpuNumLT(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLT(FieldGpuNum, v))
}

// GpuNumLTE applies the LTE predicate on the "gpu_num" field.
func GpuNumLTE(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLTE(FieldGpuNum, v))
}

// MaxMissionCountEQ applies the EQ predicate on the "max_mission_count" field.
func MaxMissionCountEQ(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldMaxMissionCount, v))
}

// MaxMissionCountNEQ applies the NEQ predicate on the "max_mission_count" field.
func MaxMissionCountNEQ(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNEQ(FieldMaxMissionCount, v))
}

// MaxMissionCountIn applies the In predicate on the "max_mission_count" field.
func MaxMissionCountIn(vs ...int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldIn(FieldMaxMissionCount, vs...))
}

// MaxMissionCountNotIn applies the NotIn predicate on the "max_mission_count" field.
func MaxMissionCountNotIn(vs ...int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNotIn(FieldMaxMissionCount, vs...))
}

// MaxMissionCountGT applies the GT predicate on the "max_mission_count" field.
func MaxMissionCountGT(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGT(FieldMaxMissionCount, v))
}

// MaxMissionCountGTE applies the GTE predicate on the "max_mission_count" field.
func MaxMissionCountGTE(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGTE(FieldMaxMissionCount, v))
}

// MaxMissionCountLT applies the LT predicate on the "max_mission_count" field.
func MaxMissionCountLT(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLT(FieldMaxMissionCount, v))
}

// MaxMissionCountLTE applies the LTE predicate on the "max_mission_count" field.
func MaxMissionCountLTE(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLTE(FieldMaxMissionCount, v))
}

// MinMissionCountEQ applies the EQ predicate on the "min_mission_count" field.
func MinMissionCountEQ(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldMinMissionCount, v))
}

// MinMissionCountNEQ applies the NEQ predicate on the "min_mission_count" field.
func MinMissionCountNEQ(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNEQ(FieldMinMissionCount, v))
}

// MinMissionCountIn applies the In predicate on the "min_mission_count" field.
func MinMissionCountIn(vs ...int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldIn(FieldMinMissionCount, vs...))
}

// MinMissionCountNotIn applies the NotIn predicate on the "min_mission_count" field.
func MinMissionCountNotIn(vs ...int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNotIn(FieldMinMissionCount, vs...))
}

// MinMissionCountGT applies the GT predicate on the "min_mission_count" field.
func MinMissionCountGT(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGT(FieldMinMissionCount, v))
}

// MinMissionCountGTE applies the GTE predicate on the "min_mission_count" field.
func MinMissionCountGTE(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGTE(FieldMinMissionCount, v))
}

// MinMissionCountLT applies the LT predicate on the "min_mission_count" field.
func MinMissionCountLT(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLT(FieldMinMissionCount, v))
}

// MinMissionCountLTE applies the LTE predicate on the "min_mission_count" field.
func MinMissionCountLTE(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLTE(FieldMinMissionCount, v))
}

// CurrentMissionCountEQ applies the EQ predicate on the "current_mission_count" field.
func CurrentMissionCountEQ(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldCurrentMissionCount, v))
}

// CurrentMissionCountNEQ applies the NEQ predicate on the "current_mission_count" field.
func CurrentMissionCountNEQ(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNEQ(FieldCurrentMissionCount, v))
}

// CurrentMissionCountIn applies the In predicate on the "current_mission_count" field.
func CurrentMissionCountIn(vs ...int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldIn(FieldCurrentMissionCount, vs...))
}

// CurrentMissionCountNotIn applies the NotIn predicate on the "current_mission_count" field.
func CurrentMissionCountNotIn(vs ...int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNotIn(FieldCurrentMissionCount, vs...))
}

// CurrentMissionCountGT applies the GT predicate on the "current_mission_count" field.
func CurrentMissionCountGT(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGT(FieldCurrentMissionCount, v))
}

// CurrentMissionCountGTE applies the GTE predicate on the "current_mission_count" field.
func CurrentMissionCountGTE(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGTE(FieldCurrentMissionCount, v))
}

// CurrentMissionCountLT applies the LT predicate on the "current_mission_count" field.
func CurrentMissionCountLT(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLT(FieldCurrentMissionCount, v))
}

// CurrentMissionCountLTE applies the LTE predicate on the "current_mission_count" field.
func CurrentMissionCountLTE(v int8) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLTE(FieldCurrentMissionCount, v))
}

// MissionBatchIDEQ applies the EQ predicate on the "mission_batch_id" field.
func MissionBatchIDEQ(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldMissionBatchID, v))
}

// MissionBatchIDNEQ applies the NEQ predicate on the "mission_batch_id" field.
func MissionBatchIDNEQ(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNEQ(FieldMissionBatchID, v))
}

// MissionBatchIDIn applies the In predicate on the "mission_batch_id" field.
func MissionBatchIDIn(vs ...int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldIn(FieldMissionBatchID, vs...))
}

// MissionBatchIDNotIn applies the NotIn predicate on the "mission_batch_id" field.
func MissionBatchIDNotIn(vs ...int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNotIn(FieldMissionBatchID, vs...))
}

// MissionBatchIDGT applies the GT predicate on the "mission_batch_id" field.
func MissionBatchIDGT(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGT(FieldMissionBatchID, v))
}

// MissionBatchIDGTE applies the GTE predicate on the "mission_batch_id" field.
func MissionBatchIDGTE(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGTE(FieldMissionBatchID, v))
}

// MissionBatchIDLT applies the LT predicate on the "mission_batch_id" field.
func MissionBatchIDLT(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLT(FieldMissionBatchID, v))
}

// MissionBatchIDLTE applies the LTE predicate on the "mission_batch_id" field.
func MissionBatchIDLTE(v int64) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLTE(FieldMissionBatchID, v))
}

// MissionBatchNumberEQ applies the EQ predicate on the "mission_batch_number" field.
func MissionBatchNumberEQ(v string) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEQ(FieldMissionBatchNumber, v))
}

// MissionBatchNumberNEQ applies the NEQ predicate on the "mission_batch_number" field.
func MissionBatchNumberNEQ(v string) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNEQ(FieldMissionBatchNumber, v))
}

// MissionBatchNumberIn applies the In predicate on the "mission_batch_number" field.
func MissionBatchNumberIn(vs ...string) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldIn(FieldMissionBatchNumber, vs...))
}

// MissionBatchNumberNotIn applies the NotIn predicate on the "mission_batch_number" field.
func MissionBatchNumberNotIn(vs ...string) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldNotIn(FieldMissionBatchNumber, vs...))
}

// MissionBatchNumberGT applies the GT predicate on the "mission_batch_number" field.
func MissionBatchNumberGT(v string) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGT(FieldMissionBatchNumber, v))
}

// MissionBatchNumberGTE applies the GTE predicate on the "mission_batch_number" field.
func MissionBatchNumberGTE(v string) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldGTE(FieldMissionBatchNumber, v))
}

// MissionBatchNumberLT applies the LT predicate on the "mission_batch_number" field.
func MissionBatchNumberLT(v string) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLT(FieldMissionBatchNumber, v))
}

// MissionBatchNumberLTE applies the LTE predicate on the "mission_batch_number" field.
func MissionBatchNumberLTE(v string) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldLTE(FieldMissionBatchNumber, v))
}

// MissionBatchNumberContains applies the Contains predicate on the "mission_batch_number" field.
func MissionBatchNumberContains(v string) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldContains(FieldMissionBatchNumber, v))
}

// MissionBatchNumberHasPrefix applies the HasPrefix predicate on the "mission_batch_number" field.
func MissionBatchNumberHasPrefix(v string) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldHasPrefix(FieldMissionBatchNumber, v))
}

// MissionBatchNumberHasSuffix applies the HasSuffix predicate on the "mission_batch_number" field.
func MissionBatchNumberHasSuffix(v string) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldHasSuffix(FieldMissionBatchNumber, v))
}

// MissionBatchNumberEqualFold applies the EqualFold predicate on the "mission_batch_number" field.
func MissionBatchNumberEqualFold(v string) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldEqualFold(FieldMissionBatchNumber, v))
}

// MissionBatchNumberContainsFold applies the ContainsFold predicate on the "mission_batch_number" field.
func MissionBatchNumberContainsFold(v string) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.FieldContainsFold(FieldMissionBatchNumber, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MissionLoadBalance) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MissionLoadBalance) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MissionLoadBalance) predicate.MissionLoadBalance {
	return predicate.MissionLoadBalance(sql.NotPredicates(p))
}
