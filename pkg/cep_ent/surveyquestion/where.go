// Code generated by ent, DO NOT EDIT.

package surveyquestion

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldEQ(FieldDeletedAt, v))
}

// SurveyID applies equality check predicate on the "survey_id" field. It's identical to SurveyIDEQ.
func SurveyID(v int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldEQ(FieldSurveyID, v))
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldEQ(FieldText, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldLTE(FieldDeletedAt, v))
}

// SurveyIDEQ applies the EQ predicate on the "survey_id" field.
func SurveyIDEQ(v int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldEQ(FieldSurveyID, v))
}

// SurveyIDNEQ applies the NEQ predicate on the "survey_id" field.
func SurveyIDNEQ(v int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldNEQ(FieldSurveyID, v))
}

// SurveyIDIn applies the In predicate on the "survey_id" field.
func SurveyIDIn(vs ...int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldIn(FieldSurveyID, vs...))
}

// SurveyIDNotIn applies the NotIn predicate on the "survey_id" field.
func SurveyIDNotIn(vs ...int64) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldNotIn(FieldSurveyID, vs...))
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldEQ(FieldText, v))
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldNEQ(FieldText, v))
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldIn(FieldText, vs...))
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldNotIn(FieldText, vs...))
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldGT(FieldText, v))
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldGTE(FieldText, v))
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldLT(FieldText, v))
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldLTE(FieldText, v))
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldContains(FieldText, v))
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldHasPrefix(FieldText, v))
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldHasSuffix(FieldText, v))
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldEqualFold(FieldText, v))
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldContainsFold(FieldText, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v enums.SurveyQuestionType) predicate.SurveyQuestion {
	vc := v
	return predicate.SurveyQuestion(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v enums.SurveyQuestionType) predicate.SurveyQuestion {
	vc := v
	return predicate.SurveyQuestion(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...enums.SurveyQuestionType) predicate.SurveyQuestion {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SurveyQuestion(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...enums.SurveyQuestionType) predicate.SurveyQuestion {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SurveyQuestion(sql.FieldNotIn(FieldType, v...))
}

// OptionsIsNil applies the IsNil predicate on the "options" field.
func OptionsIsNil() predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldIsNull(FieldOptions))
}

// OptionsNotNil applies the NotNil predicate on the "options" field.
func OptionsNotNil() predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.FieldNotNull(FieldOptions))
}

// HasSurvey applies the HasEdge predicate on the "survey" edge.
func HasSurvey() predicate.SurveyQuestion {
	return predicate.SurveyQuestion(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SurveyTable, SurveyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSurveyWith applies the HasEdge predicate on the "survey" edge with a given conditions (other predicates).
func HasSurveyWith(preds ...predicate.Survey) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(func(s *sql.Selector) {
		step := newSurveyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSurveyAnswers applies the HasEdge predicate on the "survey_answers" edge.
func HasSurveyAnswers() predicate.SurveyQuestion {
	return predicate.SurveyQuestion(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SurveyAnswersTable, SurveyAnswersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSurveyAnswersWith applies the HasEdge predicate on the "survey_answers" edge with a given conditions (other predicates).
func HasSurveyAnswersWith(preds ...predicate.SurveyAnswer) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(func(s *sql.Selector) {
		step := newSurveyAnswersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SurveyQuestion) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SurveyQuestion) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SurveyQuestion) predicate.SurveyQuestion {
	return predicate.SurveyQuestion(sql.NotPredicates(p))
}