// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/survey"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/surveyanswer"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/surveyquestion"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// SurveyQuestionUpdate is the builder for updating SurveyQuestion entities.
type SurveyQuestionUpdate struct {
	config
	hooks     []Hook
	mutation  *SurveyQuestionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SurveyQuestionUpdate builder.
func (squ *SurveyQuestionUpdate) Where(ps ...predicate.SurveyQuestion) *SurveyQuestionUpdate {
	squ.mutation.Where(ps...)
	return squ
}

// SetCreatedBy sets the "created_by" field.
func (squ *SurveyQuestionUpdate) SetCreatedBy(i int64) *SurveyQuestionUpdate {
	squ.mutation.ResetCreatedBy()
	squ.mutation.SetCreatedBy(i)
	return squ
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (squ *SurveyQuestionUpdate) SetNillableCreatedBy(i *int64) *SurveyQuestionUpdate {
	if i != nil {
		squ.SetCreatedBy(*i)
	}
	return squ
}

// AddCreatedBy adds i to the "created_by" field.
func (squ *SurveyQuestionUpdate) AddCreatedBy(i int64) *SurveyQuestionUpdate {
	squ.mutation.AddCreatedBy(i)
	return squ
}

// SetUpdatedBy sets the "updated_by" field.
func (squ *SurveyQuestionUpdate) SetUpdatedBy(i int64) *SurveyQuestionUpdate {
	squ.mutation.ResetUpdatedBy()
	squ.mutation.SetUpdatedBy(i)
	return squ
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (squ *SurveyQuestionUpdate) SetNillableUpdatedBy(i *int64) *SurveyQuestionUpdate {
	if i != nil {
		squ.SetUpdatedBy(*i)
	}
	return squ
}

// AddUpdatedBy adds i to the "updated_by" field.
func (squ *SurveyQuestionUpdate) AddUpdatedBy(i int64) *SurveyQuestionUpdate {
	squ.mutation.AddUpdatedBy(i)
	return squ
}

// SetUpdatedAt sets the "updated_at" field.
func (squ *SurveyQuestionUpdate) SetUpdatedAt(t time.Time) *SurveyQuestionUpdate {
	squ.mutation.SetUpdatedAt(t)
	return squ
}

// SetDeletedAt sets the "deleted_at" field.
func (squ *SurveyQuestionUpdate) SetDeletedAt(t time.Time) *SurveyQuestionUpdate {
	squ.mutation.SetDeletedAt(t)
	return squ
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (squ *SurveyQuestionUpdate) SetNillableDeletedAt(t *time.Time) *SurveyQuestionUpdate {
	if t != nil {
		squ.SetDeletedAt(*t)
	}
	return squ
}

// SetSurveyID sets the "survey_id" field.
func (squ *SurveyQuestionUpdate) SetSurveyID(i int64) *SurveyQuestionUpdate {
	squ.mutation.SetSurveyID(i)
	return squ
}

// SetNillableSurveyID sets the "survey_id" field if the given value is not nil.
func (squ *SurveyQuestionUpdate) SetNillableSurveyID(i *int64) *SurveyQuestionUpdate {
	if i != nil {
		squ.SetSurveyID(*i)
	}
	return squ
}

// SetText sets the "text" field.
func (squ *SurveyQuestionUpdate) SetText(s string) *SurveyQuestionUpdate {
	squ.mutation.SetText(s)
	return squ
}

// SetNillableText sets the "text" field if the given value is not nil.
func (squ *SurveyQuestionUpdate) SetNillableText(s *string) *SurveyQuestionUpdate {
	if s != nil {
		squ.SetText(*s)
	}
	return squ
}

// SetType sets the "type" field.
func (squ *SurveyQuestionUpdate) SetType(eqt enums.SurveyQuestionType) *SurveyQuestionUpdate {
	squ.mutation.SetType(eqt)
	return squ
}

// SetNillableType sets the "type" field if the given value is not nil.
func (squ *SurveyQuestionUpdate) SetNillableType(eqt *enums.SurveyQuestionType) *SurveyQuestionUpdate {
	if eqt != nil {
		squ.SetType(*eqt)
	}
	return squ
}

// SetOptions sets the "options" field.
func (squ *SurveyQuestionUpdate) SetOptions(s []string) *SurveyQuestionUpdate {
	squ.mutation.SetOptions(s)
	return squ
}

// ClearOptions clears the value of the "options" field.
func (squ *SurveyQuestionUpdate) ClearOptions() *SurveyQuestionUpdate {
	squ.mutation.ClearOptions()
	return squ
}

// SetSurvey sets the "survey" edge to the Survey entity.
func (squ *SurveyQuestionUpdate) SetSurvey(s *Survey) *SurveyQuestionUpdate {
	return squ.SetSurveyID(s.ID)
}

// AddSurveyAnswerIDs adds the "survey_answers" edge to the SurveyAnswer entity by IDs.
func (squ *SurveyQuestionUpdate) AddSurveyAnswerIDs(ids ...int64) *SurveyQuestionUpdate {
	squ.mutation.AddSurveyAnswerIDs(ids...)
	return squ
}

// AddSurveyAnswers adds the "survey_answers" edges to the SurveyAnswer entity.
func (squ *SurveyQuestionUpdate) AddSurveyAnswers(s ...*SurveyAnswer) *SurveyQuestionUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return squ.AddSurveyAnswerIDs(ids...)
}

// Mutation returns the SurveyQuestionMutation object of the builder.
func (squ *SurveyQuestionUpdate) Mutation() *SurveyQuestionMutation {
	return squ.mutation
}

// ClearSurvey clears the "survey" edge to the Survey entity.
func (squ *SurveyQuestionUpdate) ClearSurvey() *SurveyQuestionUpdate {
	squ.mutation.ClearSurvey()
	return squ
}

// ClearSurveyAnswers clears all "survey_answers" edges to the SurveyAnswer entity.
func (squ *SurveyQuestionUpdate) ClearSurveyAnswers() *SurveyQuestionUpdate {
	squ.mutation.ClearSurveyAnswers()
	return squ
}

// RemoveSurveyAnswerIDs removes the "survey_answers" edge to SurveyAnswer entities by IDs.
func (squ *SurveyQuestionUpdate) RemoveSurveyAnswerIDs(ids ...int64) *SurveyQuestionUpdate {
	squ.mutation.RemoveSurveyAnswerIDs(ids...)
	return squ
}

// RemoveSurveyAnswers removes "survey_answers" edges to SurveyAnswer entities.
func (squ *SurveyQuestionUpdate) RemoveSurveyAnswers(s ...*SurveyAnswer) *SurveyQuestionUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return squ.RemoveSurveyAnswerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (squ *SurveyQuestionUpdate) Save(ctx context.Context) (int, error) {
	squ.defaults()
	return withHooks(ctx, squ.sqlSave, squ.mutation, squ.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (squ *SurveyQuestionUpdate) SaveX(ctx context.Context) int {
	affected, err := squ.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (squ *SurveyQuestionUpdate) Exec(ctx context.Context) error {
	_, err := squ.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (squ *SurveyQuestionUpdate) ExecX(ctx context.Context) {
	if err := squ.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (squ *SurveyQuestionUpdate) defaults() {
	if _, ok := squ.mutation.UpdatedAt(); !ok {
		v := surveyquestion.UpdateDefaultUpdatedAt()
		squ.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (squ *SurveyQuestionUpdate) check() error {
	if v, ok := squ.mutation.GetType(); ok {
		if err := surveyquestion.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`cep_ent: validator failed for field "SurveyQuestion.type": %w`, err)}
		}
	}
	if _, ok := squ.mutation.SurveyID(); squ.mutation.SurveyCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "SurveyQuestion.survey"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (squ *SurveyQuestionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SurveyQuestionUpdate {
	squ.modifiers = append(squ.modifiers, modifiers...)
	return squ
}

func (squ *SurveyQuestionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := squ.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(surveyquestion.Table, surveyquestion.Columns, sqlgraph.NewFieldSpec(surveyquestion.FieldID, field.TypeInt64))
	if ps := squ.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := squ.mutation.CreatedBy(); ok {
		_spec.SetField(surveyquestion.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := squ.mutation.AddedCreatedBy(); ok {
		_spec.AddField(surveyquestion.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := squ.mutation.UpdatedBy(); ok {
		_spec.SetField(surveyquestion.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := squ.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(surveyquestion.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := squ.mutation.UpdatedAt(); ok {
		_spec.SetField(surveyquestion.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := squ.mutation.DeletedAt(); ok {
		_spec.SetField(surveyquestion.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := squ.mutation.Text(); ok {
		_spec.SetField(surveyquestion.FieldText, field.TypeString, value)
	}
	if value, ok := squ.mutation.GetType(); ok {
		_spec.SetField(surveyquestion.FieldType, field.TypeEnum, value)
	}
	if value, ok := squ.mutation.Options(); ok {
		vv, err := surveyquestion.ValueScanner.Options.Value(value)
		if err != nil {
			return 0, err
		}
		_spec.SetField(surveyquestion.FieldOptions, field.TypeString, vv)
	}
	if squ.mutation.OptionsCleared() {
		_spec.ClearField(surveyquestion.FieldOptions, field.TypeString)
	}
	if squ.mutation.SurveyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyquestion.SurveyTable,
			Columns: []string{surveyquestion.SurveyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(survey.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := squ.mutation.SurveyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyquestion.SurveyTable,
			Columns: []string{surveyquestion.SurveyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(survey.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if squ.mutation.SurveyAnswersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   surveyquestion.SurveyAnswersTable,
			Columns: []string{surveyquestion.SurveyAnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyanswer.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := squ.mutation.RemovedSurveyAnswersIDs(); len(nodes) > 0 && !squ.mutation.SurveyAnswersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   surveyquestion.SurveyAnswersTable,
			Columns: []string{surveyquestion.SurveyAnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyanswer.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := squ.mutation.SurveyAnswersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   surveyquestion.SurveyAnswersTable,
			Columns: []string{surveyquestion.SurveyAnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyanswer.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(squ.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, squ.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{surveyquestion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	squ.mutation.done = true
	return n, nil
}

// SurveyQuestionUpdateOne is the builder for updating a single SurveyQuestion entity.
type SurveyQuestionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SurveyQuestionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (squo *SurveyQuestionUpdateOne) SetCreatedBy(i int64) *SurveyQuestionUpdateOne {
	squo.mutation.ResetCreatedBy()
	squo.mutation.SetCreatedBy(i)
	return squo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (squo *SurveyQuestionUpdateOne) SetNillableCreatedBy(i *int64) *SurveyQuestionUpdateOne {
	if i != nil {
		squo.SetCreatedBy(*i)
	}
	return squo
}

// AddCreatedBy adds i to the "created_by" field.
func (squo *SurveyQuestionUpdateOne) AddCreatedBy(i int64) *SurveyQuestionUpdateOne {
	squo.mutation.AddCreatedBy(i)
	return squo
}

// SetUpdatedBy sets the "updated_by" field.
func (squo *SurveyQuestionUpdateOne) SetUpdatedBy(i int64) *SurveyQuestionUpdateOne {
	squo.mutation.ResetUpdatedBy()
	squo.mutation.SetUpdatedBy(i)
	return squo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (squo *SurveyQuestionUpdateOne) SetNillableUpdatedBy(i *int64) *SurveyQuestionUpdateOne {
	if i != nil {
		squo.SetUpdatedBy(*i)
	}
	return squo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (squo *SurveyQuestionUpdateOne) AddUpdatedBy(i int64) *SurveyQuestionUpdateOne {
	squo.mutation.AddUpdatedBy(i)
	return squo
}

// SetUpdatedAt sets the "updated_at" field.
func (squo *SurveyQuestionUpdateOne) SetUpdatedAt(t time.Time) *SurveyQuestionUpdateOne {
	squo.mutation.SetUpdatedAt(t)
	return squo
}

// SetDeletedAt sets the "deleted_at" field.
func (squo *SurveyQuestionUpdateOne) SetDeletedAt(t time.Time) *SurveyQuestionUpdateOne {
	squo.mutation.SetDeletedAt(t)
	return squo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (squo *SurveyQuestionUpdateOne) SetNillableDeletedAt(t *time.Time) *SurveyQuestionUpdateOne {
	if t != nil {
		squo.SetDeletedAt(*t)
	}
	return squo
}

// SetSurveyID sets the "survey_id" field.
func (squo *SurveyQuestionUpdateOne) SetSurveyID(i int64) *SurveyQuestionUpdateOne {
	squo.mutation.SetSurveyID(i)
	return squo
}

// SetNillableSurveyID sets the "survey_id" field if the given value is not nil.
func (squo *SurveyQuestionUpdateOne) SetNillableSurveyID(i *int64) *SurveyQuestionUpdateOne {
	if i != nil {
		squo.SetSurveyID(*i)
	}
	return squo
}

// SetText sets the "text" field.
func (squo *SurveyQuestionUpdateOne) SetText(s string) *SurveyQuestionUpdateOne {
	squo.mutation.SetText(s)
	return squo
}

// SetNillableText sets the "text" field if the given value is not nil.
func (squo *SurveyQuestionUpdateOne) SetNillableText(s *string) *SurveyQuestionUpdateOne {
	if s != nil {
		squo.SetText(*s)
	}
	return squo
}

// SetType sets the "type" field.
func (squo *SurveyQuestionUpdateOne) SetType(eqt enums.SurveyQuestionType) *SurveyQuestionUpdateOne {
	squo.mutation.SetType(eqt)
	return squo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (squo *SurveyQuestionUpdateOne) SetNillableType(eqt *enums.SurveyQuestionType) *SurveyQuestionUpdateOne {
	if eqt != nil {
		squo.SetType(*eqt)
	}
	return squo
}

// SetOptions sets the "options" field.
func (squo *SurveyQuestionUpdateOne) SetOptions(s []string) *SurveyQuestionUpdateOne {
	squo.mutation.SetOptions(s)
	return squo
}

// ClearOptions clears the value of the "options" field.
func (squo *SurveyQuestionUpdateOne) ClearOptions() *SurveyQuestionUpdateOne {
	squo.mutation.ClearOptions()
	return squo
}

// SetSurvey sets the "survey" edge to the Survey entity.
func (squo *SurveyQuestionUpdateOne) SetSurvey(s *Survey) *SurveyQuestionUpdateOne {
	return squo.SetSurveyID(s.ID)
}

// AddSurveyAnswerIDs adds the "survey_answers" edge to the SurveyAnswer entity by IDs.
func (squo *SurveyQuestionUpdateOne) AddSurveyAnswerIDs(ids ...int64) *SurveyQuestionUpdateOne {
	squo.mutation.AddSurveyAnswerIDs(ids...)
	return squo
}

// AddSurveyAnswers adds the "survey_answers" edges to the SurveyAnswer entity.
func (squo *SurveyQuestionUpdateOne) AddSurveyAnswers(s ...*SurveyAnswer) *SurveyQuestionUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return squo.AddSurveyAnswerIDs(ids...)
}

// Mutation returns the SurveyQuestionMutation object of the builder.
func (squo *SurveyQuestionUpdateOne) Mutation() *SurveyQuestionMutation {
	return squo.mutation
}

// ClearSurvey clears the "survey" edge to the Survey entity.
func (squo *SurveyQuestionUpdateOne) ClearSurvey() *SurveyQuestionUpdateOne {
	squo.mutation.ClearSurvey()
	return squo
}

// ClearSurveyAnswers clears all "survey_answers" edges to the SurveyAnswer entity.
func (squo *SurveyQuestionUpdateOne) ClearSurveyAnswers() *SurveyQuestionUpdateOne {
	squo.mutation.ClearSurveyAnswers()
	return squo
}

// RemoveSurveyAnswerIDs removes the "survey_answers" edge to SurveyAnswer entities by IDs.
func (squo *SurveyQuestionUpdateOne) RemoveSurveyAnswerIDs(ids ...int64) *SurveyQuestionUpdateOne {
	squo.mutation.RemoveSurveyAnswerIDs(ids...)
	return squo
}

// RemoveSurveyAnswers removes "survey_answers" edges to SurveyAnswer entities.
func (squo *SurveyQuestionUpdateOne) RemoveSurveyAnswers(s ...*SurveyAnswer) *SurveyQuestionUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return squo.RemoveSurveyAnswerIDs(ids...)
}

// Where appends a list predicates to the SurveyQuestionUpdate builder.
func (squo *SurveyQuestionUpdateOne) Where(ps ...predicate.SurveyQuestion) *SurveyQuestionUpdateOne {
	squo.mutation.Where(ps...)
	return squo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (squo *SurveyQuestionUpdateOne) Select(field string, fields ...string) *SurveyQuestionUpdateOne {
	squo.fields = append([]string{field}, fields...)
	return squo
}

// Save executes the query and returns the updated SurveyQuestion entity.
func (squo *SurveyQuestionUpdateOne) Save(ctx context.Context) (*SurveyQuestion, error) {
	squo.defaults()
	return withHooks(ctx, squo.sqlSave, squo.mutation, squo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (squo *SurveyQuestionUpdateOne) SaveX(ctx context.Context) *SurveyQuestion {
	node, err := squo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (squo *SurveyQuestionUpdateOne) Exec(ctx context.Context) error {
	_, err := squo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (squo *SurveyQuestionUpdateOne) ExecX(ctx context.Context) {
	if err := squo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (squo *SurveyQuestionUpdateOne) defaults() {
	if _, ok := squo.mutation.UpdatedAt(); !ok {
		v := surveyquestion.UpdateDefaultUpdatedAt()
		squo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (squo *SurveyQuestionUpdateOne) check() error {
	if v, ok := squo.mutation.GetType(); ok {
		if err := surveyquestion.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`cep_ent: validator failed for field "SurveyQuestion.type": %w`, err)}
		}
	}
	if _, ok := squo.mutation.SurveyID(); squo.mutation.SurveyCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "SurveyQuestion.survey"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (squo *SurveyQuestionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SurveyQuestionUpdateOne {
	squo.modifiers = append(squo.modifiers, modifiers...)
	return squo
}

func (squo *SurveyQuestionUpdateOne) sqlSave(ctx context.Context) (_node *SurveyQuestion, err error) {
	if err := squo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(surveyquestion.Table, surveyquestion.Columns, sqlgraph.NewFieldSpec(surveyquestion.FieldID, field.TypeInt64))
	id, ok := squo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "SurveyQuestion.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := squo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, surveyquestion.FieldID)
		for _, f := range fields {
			if !surveyquestion.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != surveyquestion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := squo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := squo.mutation.CreatedBy(); ok {
		_spec.SetField(surveyquestion.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := squo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(surveyquestion.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := squo.mutation.UpdatedBy(); ok {
		_spec.SetField(surveyquestion.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := squo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(surveyquestion.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := squo.mutation.UpdatedAt(); ok {
		_spec.SetField(surveyquestion.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := squo.mutation.DeletedAt(); ok {
		_spec.SetField(surveyquestion.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := squo.mutation.Text(); ok {
		_spec.SetField(surveyquestion.FieldText, field.TypeString, value)
	}
	if value, ok := squo.mutation.GetType(); ok {
		_spec.SetField(surveyquestion.FieldType, field.TypeEnum, value)
	}
	if value, ok := squo.mutation.Options(); ok {
		vv, err := surveyquestion.ValueScanner.Options.Value(value)
		if err != nil {
			return nil, err
		}
		_spec.SetField(surveyquestion.FieldOptions, field.TypeString, vv)
	}
	if squo.mutation.OptionsCleared() {
		_spec.ClearField(surveyquestion.FieldOptions, field.TypeString)
	}
	if squo.mutation.SurveyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyquestion.SurveyTable,
			Columns: []string{surveyquestion.SurveyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(survey.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := squo.mutation.SurveyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyquestion.SurveyTable,
			Columns: []string{surveyquestion.SurveyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(survey.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if squo.mutation.SurveyAnswersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   surveyquestion.SurveyAnswersTable,
			Columns: []string{surveyquestion.SurveyAnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyanswer.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := squo.mutation.RemovedSurveyAnswersIDs(); len(nodes) > 0 && !squo.mutation.SurveyAnswersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   surveyquestion.SurveyAnswersTable,
			Columns: []string{surveyquestion.SurveyAnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyanswer.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := squo.mutation.SurveyAnswersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   surveyquestion.SurveyAnswersTable,
			Columns: []string{surveyquestion.SurveyAnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyanswer.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(squo.modifiers...)
	_node = &SurveyQuestion{config: squo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, squo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{surveyquestion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	squo.mutation.done = true
	return _node, nil
}
