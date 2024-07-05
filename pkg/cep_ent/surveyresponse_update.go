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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/surveyresponse"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// SurveyResponseUpdate is the builder for updating SurveyResponse entities.
type SurveyResponseUpdate struct {
	config
	hooks     []Hook
	mutation  *SurveyResponseMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SurveyResponseUpdate builder.
func (sru *SurveyResponseUpdate) Where(ps ...predicate.SurveyResponse) *SurveyResponseUpdate {
	sru.mutation.Where(ps...)
	return sru
}

// SetCreatedBy sets the "created_by" field.
func (sru *SurveyResponseUpdate) SetCreatedBy(i int64) *SurveyResponseUpdate {
	sru.mutation.ResetCreatedBy()
	sru.mutation.SetCreatedBy(i)
	return sru
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (sru *SurveyResponseUpdate) SetNillableCreatedBy(i *int64) *SurveyResponseUpdate {
	if i != nil {
		sru.SetCreatedBy(*i)
	}
	return sru
}

// AddCreatedBy adds i to the "created_by" field.
func (sru *SurveyResponseUpdate) AddCreatedBy(i int64) *SurveyResponseUpdate {
	sru.mutation.AddCreatedBy(i)
	return sru
}

// SetUpdatedBy sets the "updated_by" field.
func (sru *SurveyResponseUpdate) SetUpdatedBy(i int64) *SurveyResponseUpdate {
	sru.mutation.ResetUpdatedBy()
	sru.mutation.SetUpdatedBy(i)
	return sru
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (sru *SurveyResponseUpdate) SetNillableUpdatedBy(i *int64) *SurveyResponseUpdate {
	if i != nil {
		sru.SetUpdatedBy(*i)
	}
	return sru
}

// AddUpdatedBy adds i to the "updated_by" field.
func (sru *SurveyResponseUpdate) AddUpdatedBy(i int64) *SurveyResponseUpdate {
	sru.mutation.AddUpdatedBy(i)
	return sru
}

// SetUpdatedAt sets the "updated_at" field.
func (sru *SurveyResponseUpdate) SetUpdatedAt(t time.Time) *SurveyResponseUpdate {
	sru.mutation.SetUpdatedAt(t)
	return sru
}

// SetDeletedAt sets the "deleted_at" field.
func (sru *SurveyResponseUpdate) SetDeletedAt(t time.Time) *SurveyResponseUpdate {
	sru.mutation.SetDeletedAt(t)
	return sru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sru *SurveyResponseUpdate) SetNillableDeletedAt(t *time.Time) *SurveyResponseUpdate {
	if t != nil {
		sru.SetDeletedAt(*t)
	}
	return sru
}

// SetUserID sets the "user_id" field.
func (sru *SurveyResponseUpdate) SetUserID(i int64) *SurveyResponseUpdate {
	sru.mutation.SetUserID(i)
	return sru
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (sru *SurveyResponseUpdate) SetNillableUserID(i *int64) *SurveyResponseUpdate {
	if i != nil {
		sru.SetUserID(*i)
	}
	return sru
}

// SetSurveyID sets the "survey_id" field.
func (sru *SurveyResponseUpdate) SetSurveyID(i int64) *SurveyResponseUpdate {
	sru.mutation.SetSurveyID(i)
	return sru
}

// SetNillableSurveyID sets the "survey_id" field if the given value is not nil.
func (sru *SurveyResponseUpdate) SetNillableSurveyID(i *int64) *SurveyResponseUpdate {
	if i != nil {
		sru.SetSurveyID(*i)
	}
	return sru
}

// SetUser sets the "user" edge to the User entity.
func (sru *SurveyResponseUpdate) SetUser(u *User) *SurveyResponseUpdate {
	return sru.SetUserID(u.ID)
}

// SetSurvey sets the "survey" edge to the Survey entity.
func (sru *SurveyResponseUpdate) SetSurvey(s *Survey) *SurveyResponseUpdate {
	return sru.SetSurveyID(s.ID)
}

// AddSurveyAnswerIDs adds the "survey_answers" edge to the SurveyAnswer entity by IDs.
func (sru *SurveyResponseUpdate) AddSurveyAnswerIDs(ids ...int64) *SurveyResponseUpdate {
	sru.mutation.AddSurveyAnswerIDs(ids...)
	return sru
}

// AddSurveyAnswers adds the "survey_answers" edges to the SurveyAnswer entity.
func (sru *SurveyResponseUpdate) AddSurveyAnswers(s ...*SurveyAnswer) *SurveyResponseUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sru.AddSurveyAnswerIDs(ids...)
}

// Mutation returns the SurveyResponseMutation object of the builder.
func (sru *SurveyResponseUpdate) Mutation() *SurveyResponseMutation {
	return sru.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (sru *SurveyResponseUpdate) ClearUser() *SurveyResponseUpdate {
	sru.mutation.ClearUser()
	return sru
}

// ClearSurvey clears the "survey" edge to the Survey entity.
func (sru *SurveyResponseUpdate) ClearSurvey() *SurveyResponseUpdate {
	sru.mutation.ClearSurvey()
	return sru
}

// ClearSurveyAnswers clears all "survey_answers" edges to the SurveyAnswer entity.
func (sru *SurveyResponseUpdate) ClearSurveyAnswers() *SurveyResponseUpdate {
	sru.mutation.ClearSurveyAnswers()
	return sru
}

// RemoveSurveyAnswerIDs removes the "survey_answers" edge to SurveyAnswer entities by IDs.
func (sru *SurveyResponseUpdate) RemoveSurveyAnswerIDs(ids ...int64) *SurveyResponseUpdate {
	sru.mutation.RemoveSurveyAnswerIDs(ids...)
	return sru
}

// RemoveSurveyAnswers removes "survey_answers" edges to SurveyAnswer entities.
func (sru *SurveyResponseUpdate) RemoveSurveyAnswers(s ...*SurveyAnswer) *SurveyResponseUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sru.RemoveSurveyAnswerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sru *SurveyResponseUpdate) Save(ctx context.Context) (int, error) {
	sru.defaults()
	return withHooks(ctx, sru.sqlSave, sru.mutation, sru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sru *SurveyResponseUpdate) SaveX(ctx context.Context) int {
	affected, err := sru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sru *SurveyResponseUpdate) Exec(ctx context.Context) error {
	_, err := sru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sru *SurveyResponseUpdate) ExecX(ctx context.Context) {
	if err := sru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sru *SurveyResponseUpdate) defaults() {
	if _, ok := sru.mutation.UpdatedAt(); !ok {
		v := surveyresponse.UpdateDefaultUpdatedAt()
		sru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sru *SurveyResponseUpdate) check() error {
	if _, ok := sru.mutation.UserID(); sru.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "SurveyResponse.user"`)
	}
	if _, ok := sru.mutation.SurveyID(); sru.mutation.SurveyCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "SurveyResponse.survey"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sru *SurveyResponseUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SurveyResponseUpdate {
	sru.modifiers = append(sru.modifiers, modifiers...)
	return sru
}

func (sru *SurveyResponseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := sru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(surveyresponse.Table, surveyresponse.Columns, sqlgraph.NewFieldSpec(surveyresponse.FieldID, field.TypeInt64))
	if ps := sru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sru.mutation.CreatedBy(); ok {
		_spec.SetField(surveyresponse.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := sru.mutation.AddedCreatedBy(); ok {
		_spec.AddField(surveyresponse.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := sru.mutation.UpdatedBy(); ok {
		_spec.SetField(surveyresponse.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := sru.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(surveyresponse.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := sru.mutation.UpdatedAt(); ok {
		_spec.SetField(surveyresponse.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := sru.mutation.DeletedAt(); ok {
		_spec.SetField(surveyresponse.FieldDeletedAt, field.TypeTime, value)
	}
	if sru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyresponse.UserTable,
			Columns: []string{surveyresponse.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyresponse.UserTable,
			Columns: []string{surveyresponse.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sru.mutation.SurveyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyresponse.SurveyTable,
			Columns: []string{surveyresponse.SurveyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(survey.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.SurveyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyresponse.SurveyTable,
			Columns: []string{surveyresponse.SurveyColumn},
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
	if sru.mutation.SurveyAnswersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   surveyresponse.SurveyAnswersTable,
			Columns: []string{surveyresponse.SurveyAnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyanswer.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.RemovedSurveyAnswersIDs(); len(nodes) > 0 && !sru.mutation.SurveyAnswersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   surveyresponse.SurveyAnswersTable,
			Columns: []string{surveyresponse.SurveyAnswersColumn},
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
	if nodes := sru.mutation.SurveyAnswersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   surveyresponse.SurveyAnswersTable,
			Columns: []string{surveyresponse.SurveyAnswersColumn},
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
	_spec.AddModifiers(sru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, sru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{surveyresponse.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sru.mutation.done = true
	return n, nil
}

// SurveyResponseUpdateOne is the builder for updating a single SurveyResponse entity.
type SurveyResponseUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SurveyResponseMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (sruo *SurveyResponseUpdateOne) SetCreatedBy(i int64) *SurveyResponseUpdateOne {
	sruo.mutation.ResetCreatedBy()
	sruo.mutation.SetCreatedBy(i)
	return sruo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (sruo *SurveyResponseUpdateOne) SetNillableCreatedBy(i *int64) *SurveyResponseUpdateOne {
	if i != nil {
		sruo.SetCreatedBy(*i)
	}
	return sruo
}

// AddCreatedBy adds i to the "created_by" field.
func (sruo *SurveyResponseUpdateOne) AddCreatedBy(i int64) *SurveyResponseUpdateOne {
	sruo.mutation.AddCreatedBy(i)
	return sruo
}

// SetUpdatedBy sets the "updated_by" field.
func (sruo *SurveyResponseUpdateOne) SetUpdatedBy(i int64) *SurveyResponseUpdateOne {
	sruo.mutation.ResetUpdatedBy()
	sruo.mutation.SetUpdatedBy(i)
	return sruo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (sruo *SurveyResponseUpdateOne) SetNillableUpdatedBy(i *int64) *SurveyResponseUpdateOne {
	if i != nil {
		sruo.SetUpdatedBy(*i)
	}
	return sruo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (sruo *SurveyResponseUpdateOne) AddUpdatedBy(i int64) *SurveyResponseUpdateOne {
	sruo.mutation.AddUpdatedBy(i)
	return sruo
}

// SetUpdatedAt sets the "updated_at" field.
func (sruo *SurveyResponseUpdateOne) SetUpdatedAt(t time.Time) *SurveyResponseUpdateOne {
	sruo.mutation.SetUpdatedAt(t)
	return sruo
}

// SetDeletedAt sets the "deleted_at" field.
func (sruo *SurveyResponseUpdateOne) SetDeletedAt(t time.Time) *SurveyResponseUpdateOne {
	sruo.mutation.SetDeletedAt(t)
	return sruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sruo *SurveyResponseUpdateOne) SetNillableDeletedAt(t *time.Time) *SurveyResponseUpdateOne {
	if t != nil {
		sruo.SetDeletedAt(*t)
	}
	return sruo
}

// SetUserID sets the "user_id" field.
func (sruo *SurveyResponseUpdateOne) SetUserID(i int64) *SurveyResponseUpdateOne {
	sruo.mutation.SetUserID(i)
	return sruo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (sruo *SurveyResponseUpdateOne) SetNillableUserID(i *int64) *SurveyResponseUpdateOne {
	if i != nil {
		sruo.SetUserID(*i)
	}
	return sruo
}

// SetSurveyID sets the "survey_id" field.
func (sruo *SurveyResponseUpdateOne) SetSurveyID(i int64) *SurveyResponseUpdateOne {
	sruo.mutation.SetSurveyID(i)
	return sruo
}

// SetNillableSurveyID sets the "survey_id" field if the given value is not nil.
func (sruo *SurveyResponseUpdateOne) SetNillableSurveyID(i *int64) *SurveyResponseUpdateOne {
	if i != nil {
		sruo.SetSurveyID(*i)
	}
	return sruo
}

// SetUser sets the "user" edge to the User entity.
func (sruo *SurveyResponseUpdateOne) SetUser(u *User) *SurveyResponseUpdateOne {
	return sruo.SetUserID(u.ID)
}

// SetSurvey sets the "survey" edge to the Survey entity.
func (sruo *SurveyResponseUpdateOne) SetSurvey(s *Survey) *SurveyResponseUpdateOne {
	return sruo.SetSurveyID(s.ID)
}

// AddSurveyAnswerIDs adds the "survey_answers" edge to the SurveyAnswer entity by IDs.
func (sruo *SurveyResponseUpdateOne) AddSurveyAnswerIDs(ids ...int64) *SurveyResponseUpdateOne {
	sruo.mutation.AddSurveyAnswerIDs(ids...)
	return sruo
}

// AddSurveyAnswers adds the "survey_answers" edges to the SurveyAnswer entity.
func (sruo *SurveyResponseUpdateOne) AddSurveyAnswers(s ...*SurveyAnswer) *SurveyResponseUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sruo.AddSurveyAnswerIDs(ids...)
}

// Mutation returns the SurveyResponseMutation object of the builder.
func (sruo *SurveyResponseUpdateOne) Mutation() *SurveyResponseMutation {
	return sruo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (sruo *SurveyResponseUpdateOne) ClearUser() *SurveyResponseUpdateOne {
	sruo.mutation.ClearUser()
	return sruo
}

// ClearSurvey clears the "survey" edge to the Survey entity.
func (sruo *SurveyResponseUpdateOne) ClearSurvey() *SurveyResponseUpdateOne {
	sruo.mutation.ClearSurvey()
	return sruo
}

// ClearSurveyAnswers clears all "survey_answers" edges to the SurveyAnswer entity.
func (sruo *SurveyResponseUpdateOne) ClearSurveyAnswers() *SurveyResponseUpdateOne {
	sruo.mutation.ClearSurveyAnswers()
	return sruo
}

// RemoveSurveyAnswerIDs removes the "survey_answers" edge to SurveyAnswer entities by IDs.
func (sruo *SurveyResponseUpdateOne) RemoveSurveyAnswerIDs(ids ...int64) *SurveyResponseUpdateOne {
	sruo.mutation.RemoveSurveyAnswerIDs(ids...)
	return sruo
}

// RemoveSurveyAnswers removes "survey_answers" edges to SurveyAnswer entities.
func (sruo *SurveyResponseUpdateOne) RemoveSurveyAnswers(s ...*SurveyAnswer) *SurveyResponseUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sruo.RemoveSurveyAnswerIDs(ids...)
}

// Where appends a list predicates to the SurveyResponseUpdate builder.
func (sruo *SurveyResponseUpdateOne) Where(ps ...predicate.SurveyResponse) *SurveyResponseUpdateOne {
	sruo.mutation.Where(ps...)
	return sruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sruo *SurveyResponseUpdateOne) Select(field string, fields ...string) *SurveyResponseUpdateOne {
	sruo.fields = append([]string{field}, fields...)
	return sruo
}

// Save executes the query and returns the updated SurveyResponse entity.
func (sruo *SurveyResponseUpdateOne) Save(ctx context.Context) (*SurveyResponse, error) {
	sruo.defaults()
	return withHooks(ctx, sruo.sqlSave, sruo.mutation, sruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sruo *SurveyResponseUpdateOne) SaveX(ctx context.Context) *SurveyResponse {
	node, err := sruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sruo *SurveyResponseUpdateOne) Exec(ctx context.Context) error {
	_, err := sruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sruo *SurveyResponseUpdateOne) ExecX(ctx context.Context) {
	if err := sruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sruo *SurveyResponseUpdateOne) defaults() {
	if _, ok := sruo.mutation.UpdatedAt(); !ok {
		v := surveyresponse.UpdateDefaultUpdatedAt()
		sruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sruo *SurveyResponseUpdateOne) check() error {
	if _, ok := sruo.mutation.UserID(); sruo.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "SurveyResponse.user"`)
	}
	if _, ok := sruo.mutation.SurveyID(); sruo.mutation.SurveyCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "SurveyResponse.survey"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sruo *SurveyResponseUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SurveyResponseUpdateOne {
	sruo.modifiers = append(sruo.modifiers, modifiers...)
	return sruo
}

func (sruo *SurveyResponseUpdateOne) sqlSave(ctx context.Context) (_node *SurveyResponse, err error) {
	if err := sruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(surveyresponse.Table, surveyresponse.Columns, sqlgraph.NewFieldSpec(surveyresponse.FieldID, field.TypeInt64))
	id, ok := sruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "SurveyResponse.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, surveyresponse.FieldID)
		for _, f := range fields {
			if !surveyresponse.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != surveyresponse.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sruo.mutation.CreatedBy(); ok {
		_spec.SetField(surveyresponse.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := sruo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(surveyresponse.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := sruo.mutation.UpdatedBy(); ok {
		_spec.SetField(surveyresponse.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := sruo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(surveyresponse.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := sruo.mutation.UpdatedAt(); ok {
		_spec.SetField(surveyresponse.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := sruo.mutation.DeletedAt(); ok {
		_spec.SetField(surveyresponse.FieldDeletedAt, field.TypeTime, value)
	}
	if sruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyresponse.UserTable,
			Columns: []string{surveyresponse.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyresponse.UserTable,
			Columns: []string{surveyresponse.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sruo.mutation.SurveyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyresponse.SurveyTable,
			Columns: []string{surveyresponse.SurveyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(survey.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.SurveyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyresponse.SurveyTable,
			Columns: []string{surveyresponse.SurveyColumn},
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
	if sruo.mutation.SurveyAnswersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   surveyresponse.SurveyAnswersTable,
			Columns: []string{surveyresponse.SurveyAnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyanswer.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.RemovedSurveyAnswersIDs(); len(nodes) > 0 && !sruo.mutation.SurveyAnswersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   surveyresponse.SurveyAnswersTable,
			Columns: []string{surveyresponse.SurveyAnswersColumn},
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
	if nodes := sruo.mutation.SurveyAnswersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   surveyresponse.SurveyAnswersTable,
			Columns: []string{surveyresponse.SurveyAnswersColumn},
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
	_spec.AddModifiers(sruo.modifiers...)
	_node = &SurveyResponse{config: sruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{surveyresponse.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sruo.mutation.done = true
	return _node, nil
}
