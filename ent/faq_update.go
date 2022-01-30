// Code generated by entc, DO NOT EDIT.

package ent

import (
	"api/ent/faq"
	"api/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FaqUpdate is the builder for updating Faq entities.
type FaqUpdate struct {
	config
	hooks    []Hook
	mutation *FaqMutation
}

// Where appends a list predicates to the FaqUpdate builder.
func (fu *FaqUpdate) Where(ps ...predicate.Faq) *FaqUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetQuestion sets the "question" field.
func (fu *FaqUpdate) SetQuestion(s string) *FaqUpdate {
	fu.mutation.SetQuestion(s)
	return fu
}

// SetAnswer sets the "answer" field.
func (fu *FaqUpdate) SetAnswer(s string) *FaqUpdate {
	fu.mutation.SetAnswer(s)
	return fu
}

// SetStatus sets the "status" field.
func (fu *FaqUpdate) SetStatus(b bool) *FaqUpdate {
	fu.mutation.SetStatus(b)
	return fu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (fu *FaqUpdate) SetNillableStatus(b *bool) *FaqUpdate {
	if b != nil {
		fu.SetStatus(*b)
	}
	return fu
}

// SetUpdatedAt sets the "updated_at" field.
func (fu *FaqUpdate) SetUpdatedAt(t time.Time) *FaqUpdate {
	fu.mutation.SetUpdatedAt(t)
	return fu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fu *FaqUpdate) SetNillableUpdatedAt(t *time.Time) *FaqUpdate {
	if t != nil {
		fu.SetUpdatedAt(*t)
	}
	return fu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (fu *FaqUpdate) ClearUpdatedAt() *FaqUpdate {
	fu.mutation.ClearUpdatedAt()
	return fu
}

// SetDeletedAt sets the "deleted_at" field.
func (fu *FaqUpdate) SetDeletedAt(t time.Time) *FaqUpdate {
	fu.mutation.SetDeletedAt(t)
	return fu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fu *FaqUpdate) SetNillableDeletedAt(t *time.Time) *FaqUpdate {
	if t != nil {
		fu.SetDeletedAt(*t)
	}
	return fu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (fu *FaqUpdate) ClearDeletedAt() *FaqUpdate {
	fu.mutation.ClearDeletedAt()
	return fu
}

// Mutation returns the FaqMutation object of the builder.
func (fu *FaqUpdate) Mutation() *FaqMutation {
	return fu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FaqUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fu.hooks) == 0 {
		if err = fu.check(); err != nil {
			return 0, err
		}
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FaqMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fu.check(); err != nil {
				return 0, err
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			if fu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FaqUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FaqUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FaqUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FaqUpdate) check() error {
	if v, ok := fu.mutation.Question(); ok {
		if err := faq.QuestionValidator(v); err != nil {
			return &ValidationError{Name: "question", err: fmt.Errorf(`ent: validator failed for field "Faq.question": %w`, err)}
		}
	}
	return nil
}

func (fu *FaqUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   faq.Table,
			Columns: faq.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: faq.FieldID,
			},
		},
	}
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Question(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: faq.FieldQuestion,
		})
	}
	if value, ok := fu.mutation.Answer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: faq.FieldAnswer,
		})
	}
	if value, ok := fu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: faq.FieldStatus,
		})
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: faq.FieldUpdatedAt,
		})
	}
	if fu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: faq.FieldUpdatedAt,
		})
	}
	if value, ok := fu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: faq.FieldDeletedAt,
		})
	}
	if fu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: faq.FieldDeletedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{faq.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// FaqUpdateOne is the builder for updating a single Faq entity.
type FaqUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FaqMutation
}

// SetQuestion sets the "question" field.
func (fuo *FaqUpdateOne) SetQuestion(s string) *FaqUpdateOne {
	fuo.mutation.SetQuestion(s)
	return fuo
}

// SetAnswer sets the "answer" field.
func (fuo *FaqUpdateOne) SetAnswer(s string) *FaqUpdateOne {
	fuo.mutation.SetAnswer(s)
	return fuo
}

// SetStatus sets the "status" field.
func (fuo *FaqUpdateOne) SetStatus(b bool) *FaqUpdateOne {
	fuo.mutation.SetStatus(b)
	return fuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (fuo *FaqUpdateOne) SetNillableStatus(b *bool) *FaqUpdateOne {
	if b != nil {
		fuo.SetStatus(*b)
	}
	return fuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fuo *FaqUpdateOne) SetUpdatedAt(t time.Time) *FaqUpdateOne {
	fuo.mutation.SetUpdatedAt(t)
	return fuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fuo *FaqUpdateOne) SetNillableUpdatedAt(t *time.Time) *FaqUpdateOne {
	if t != nil {
		fuo.SetUpdatedAt(*t)
	}
	return fuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (fuo *FaqUpdateOne) ClearUpdatedAt() *FaqUpdateOne {
	fuo.mutation.ClearUpdatedAt()
	return fuo
}

// SetDeletedAt sets the "deleted_at" field.
func (fuo *FaqUpdateOne) SetDeletedAt(t time.Time) *FaqUpdateOne {
	fuo.mutation.SetDeletedAt(t)
	return fuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fuo *FaqUpdateOne) SetNillableDeletedAt(t *time.Time) *FaqUpdateOne {
	if t != nil {
		fuo.SetDeletedAt(*t)
	}
	return fuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (fuo *FaqUpdateOne) ClearDeletedAt() *FaqUpdateOne {
	fuo.mutation.ClearDeletedAt()
	return fuo
}

// Mutation returns the FaqMutation object of the builder.
func (fuo *FaqUpdateOne) Mutation() *FaqMutation {
	return fuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FaqUpdateOne) Select(field string, fields ...string) *FaqUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Faq entity.
func (fuo *FaqUpdateOne) Save(ctx context.Context) (*Faq, error) {
	var (
		err  error
		node *Faq
	)
	if len(fuo.hooks) == 0 {
		if err = fuo.check(); err != nil {
			return nil, err
		}
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FaqMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fuo.check(); err != nil {
				return nil, err
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			if fuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FaqUpdateOne) SaveX(ctx context.Context) *Faq {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FaqUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FaqUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FaqUpdateOne) check() error {
	if v, ok := fuo.mutation.Question(); ok {
		if err := faq.QuestionValidator(v); err != nil {
			return &ValidationError{Name: "question", err: fmt.Errorf(`ent: validator failed for field "Faq.question": %w`, err)}
		}
	}
	return nil
}

func (fuo *FaqUpdateOne) sqlSave(ctx context.Context) (_node *Faq, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   faq.Table,
			Columns: faq.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: faq.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Faq.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, faq.FieldID)
		for _, f := range fields {
			if !faq.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != faq.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.Question(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: faq.FieldQuestion,
		})
	}
	if value, ok := fuo.mutation.Answer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: faq.FieldAnswer,
		})
	}
	if value, ok := fuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: faq.FieldStatus,
		})
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: faq.FieldUpdatedAt,
		})
	}
	if fuo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: faq.FieldUpdatedAt,
		})
	}
	if value, ok := fuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: faq.FieldDeletedAt,
		})
	}
	if fuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: faq.FieldDeletedAt,
		})
	}
	_node = &Faq{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{faq.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
