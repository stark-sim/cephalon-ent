// Code generated by ent, DO NOT EDIT.

package surveyresponse

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldEQ(FieldDeletedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldEQ(FieldUserID, v))
}

// SurveyID applies equality check predicate on the "survey_id" field. It's identical to SurveyIDEQ.
func SurveyID(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldEQ(FieldSurveyID, v))
}

// ApprovedBy applies equality check predicate on the "approved_by" field. It's identical to ApprovedByEQ.
func ApprovedBy(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldEQ(FieldApprovedBy, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldLTE(FieldDeletedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldNotIn(FieldUserID, vs...))
}

// SurveyIDEQ applies the EQ predicate on the "survey_id" field.
func SurveyIDEQ(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldEQ(FieldSurveyID, v))
}

// SurveyIDNEQ applies the NEQ predicate on the "survey_id" field.
func SurveyIDNEQ(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldNEQ(FieldSurveyID, v))
}

// SurveyIDIn applies the In predicate on the "survey_id" field.
func SurveyIDIn(vs ...int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldIn(FieldSurveyID, vs...))
}

// SurveyIDNotIn applies the NotIn predicate on the "survey_id" field.
func SurveyIDNotIn(vs ...int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldNotIn(FieldSurveyID, vs...))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v enums.SurveyResponseStatus) predicate.SurveyResponse {
	vc := v
	return predicate.SurveyResponse(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v enums.SurveyResponseStatus) predicate.SurveyResponse {
	vc := v
	return predicate.SurveyResponse(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...enums.SurveyResponseStatus) predicate.SurveyResponse {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SurveyResponse(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...enums.SurveyResponseStatus) predicate.SurveyResponse {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SurveyResponse(sql.FieldNotIn(FieldStatus, v...))
}

// ApprovedByEQ applies the EQ predicate on the "approved_by" field.
func ApprovedByEQ(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldEQ(FieldApprovedBy, v))
}

// ApprovedByNEQ applies the NEQ predicate on the "approved_by" field.
func ApprovedByNEQ(v int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldNEQ(FieldApprovedBy, v))
}

// ApprovedByIn applies the In predicate on the "approved_by" field.
func ApprovedByIn(vs ...int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldIn(FieldApprovedBy, vs...))
}

// ApprovedByNotIn applies the NotIn predicate on the "approved_by" field.
func ApprovedByNotIn(vs ...int64) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.FieldNotIn(FieldApprovedBy, vs...))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.SurveyResponse {
	return predicate.SurveyResponse(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.SurveyResponse {
	return predicate.SurveyResponse(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSurvey applies the HasEdge predicate on the "survey" edge.
func HasSurvey() predicate.SurveyResponse {
	return predicate.SurveyResponse(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SurveyTable, SurveyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSurveyWith applies the HasEdge predicate on the "survey" edge with a given conditions (other predicates).
func HasSurveyWith(preds ...predicate.Survey) predicate.SurveyResponse {
	return predicate.SurveyResponse(func(s *sql.Selector) {
		step := newSurveyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasApprovedUser applies the HasEdge predicate on the "approved_user" edge.
func HasApprovedUser() predicate.SurveyResponse {
	return predicate.SurveyResponse(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ApprovedUserTable, ApprovedUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasApprovedUserWith applies the HasEdge predicate on the "approved_user" edge with a given conditions (other predicates).
func HasApprovedUserWith(preds ...predicate.User) predicate.SurveyResponse {
	return predicate.SurveyResponse(func(s *sql.Selector) {
		step := newApprovedUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSurveyAnswers applies the HasEdge predicate on the "survey_answers" edge.
func HasSurveyAnswers() predicate.SurveyResponse {
	return predicate.SurveyResponse(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SurveyAnswersTable, SurveyAnswersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSurveyAnswersWith applies the HasEdge predicate on the "survey_answers" edge with a given conditions (other predicates).
func HasSurveyAnswersWith(preds ...predicate.SurveyAnswer) predicate.SurveyResponse {
	return predicate.SurveyResponse(func(s *sql.Selector) {
		step := newSurveyAnswersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SurveyResponse) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SurveyResponse) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SurveyResponse) predicate.SurveyResponse {
	return predicate.SurveyResponse(sql.NotPredicates(p))
}
