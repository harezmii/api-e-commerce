// Code generated by entc, DO NOT EDIT.

package ent

import (
	"api/ent/faq"
	"api/ent/predicate"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FaqDelete is the builder for deleting a Faq entity.
type FaqDelete struct {
	config
	hooks    []Hook
	mutation *FaqMutation
}

// Where appends a list predicates to the FaqDelete builder.
func (fd *FaqDelete) Where(ps ...predicate.Faq) *FaqDelete {
	fd.mutation.Where(ps...)
	return fd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fd *FaqDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fd.hooks) == 0 {
		affected, err = fd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FaqMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fd.mutation = mutation
			affected, err = fd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fd.hooks) - 1; i >= 0; i-- {
			if fd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (fd *FaqDelete) ExecX(ctx context.Context) int {
	n, err := fd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fd *FaqDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: faq.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: faq.FieldID,
			},
		},
	}
	if ps := fd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, fd.driver, _spec)
}

// FaqDeleteOne is the builder for deleting a single Faq entity.
type FaqDeleteOne struct {
	fd *FaqDelete
}

// Exec executes the deletion query.
func (fdo *FaqDeleteOne) Exec(ctx context.Context) error {
	n, err := fdo.fd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{faq.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fdo *FaqDeleteOne) ExecX(ctx context.Context) {
	fdo.fd.ExecX(ctx)
}