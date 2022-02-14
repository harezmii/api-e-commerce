// Code generated by entc, DO NOT EDIT.

package ent

import (
	"api/ent/image"
	"api/ent/predicate"
	"api/ent/product"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ImageUpdate is the builder for updating Image entities.
type ImageUpdate struct {
	config
	hooks    []Hook
	mutation *ImageMutation
}

// Where appends a list predicates to the ImageUpdate builder.
func (iu *ImageUpdate) Where(ps ...predicate.Image) *ImageUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetTitle sets the "title" field.
func (iu *ImageUpdate) SetTitle(s string) *ImageUpdate {
	iu.mutation.SetTitle(s)
	return iu
}

// SetImage sets the "image" field.
func (iu *ImageUpdate) SetImage(s string) *ImageUpdate {
	iu.mutation.SetImage(s)
	return iu
}

// SetStatus sets the "status" field.
func (iu *ImageUpdate) SetStatus(b bool) *ImageUpdate {
	iu.mutation.SetStatus(b)
	return iu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (iu *ImageUpdate) SetNillableStatus(b *bool) *ImageUpdate {
	if b != nil {
		iu.SetStatus(*b)
	}
	return iu
}

// SetUpdatedAt sets the "updated_at" field.
func (iu *ImageUpdate) SetUpdatedAt(t time.Time) *ImageUpdate {
	iu.mutation.SetUpdatedAt(t)
	return iu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (iu *ImageUpdate) SetNillableUpdatedAt(t *time.Time) *ImageUpdate {
	if t != nil {
		iu.SetUpdatedAt(*t)
	}
	return iu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (iu *ImageUpdate) ClearUpdatedAt() *ImageUpdate {
	iu.mutation.ClearUpdatedAt()
	return iu
}

// SetDeletedAt sets the "deleted_at" field.
func (iu *ImageUpdate) SetDeletedAt(t time.Time) *ImageUpdate {
	iu.mutation.SetDeletedAt(t)
	return iu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (iu *ImageUpdate) SetNillableDeletedAt(t *time.Time) *ImageUpdate {
	if t != nil {
		iu.SetDeletedAt(*t)
	}
	return iu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (iu *ImageUpdate) ClearDeletedAt() *ImageUpdate {
	iu.mutation.ClearDeletedAt()
	return iu
}

// AddOwnerIDs adds the "owner" edge to the Product entity by IDs.
func (iu *ImageUpdate) AddOwnerIDs(ids ...int) *ImageUpdate {
	iu.mutation.AddOwnerIDs(ids...)
	return iu
}

// AddOwner adds the "owner" edges to the Product entity.
func (iu *ImageUpdate) AddOwner(p ...*Product) *ImageUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iu.AddOwnerIDs(ids...)
}

// Mutation returns the ImageMutation object of the builder.
func (iu *ImageUpdate) Mutation() *ImageMutation {
	return iu.mutation
}

// ClearOwner clears all "owner" edges to the Product entity.
func (iu *ImageUpdate) ClearOwner() *ImageUpdate {
	iu.mutation.ClearOwner()
	return iu
}

// RemoveOwnerIDs removes the "owner" edge to Product entities by IDs.
func (iu *ImageUpdate) RemoveOwnerIDs(ids ...int) *ImageUpdate {
	iu.mutation.RemoveOwnerIDs(ids...)
	return iu
}

// RemoveOwner removes "owner" edges to Product entities.
func (iu *ImageUpdate) RemoveOwner(p ...*Product) *ImageUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iu.RemoveOwnerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ImageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(iu.hooks) == 0 {
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ImageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			if iu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ImageUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ImageUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ImageUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iu *ImageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   image.Table,
			Columns: image.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: image.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: image.FieldTitle,
		})
	}
	if value, ok := iu.mutation.Image(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: image.FieldImage,
		})
	}
	if value, ok := iu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: image.FieldStatus,
		})
	}
	if value, ok := iu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: image.FieldUpdatedAt,
		})
	}
	if iu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: image.FieldUpdatedAt,
		})
	}
	if value, ok := iu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: image.FieldDeletedAt,
		})
	}
	if iu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: image.FieldDeletedAt,
		})
	}
	if iu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   image.OwnerTable,
			Columns: image.OwnerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedOwnerIDs(); len(nodes) > 0 && !iu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   image.OwnerTable,
			Columns: image.OwnerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   image.OwnerTable,
			Columns: image.OwnerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{image.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ImageUpdateOne is the builder for updating a single Image entity.
type ImageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ImageMutation
}

// SetTitle sets the "title" field.
func (iuo *ImageUpdateOne) SetTitle(s string) *ImageUpdateOne {
	iuo.mutation.SetTitle(s)
	return iuo
}

// SetImage sets the "image" field.
func (iuo *ImageUpdateOne) SetImage(s string) *ImageUpdateOne {
	iuo.mutation.SetImage(s)
	return iuo
}

// SetStatus sets the "status" field.
func (iuo *ImageUpdateOne) SetStatus(b bool) *ImageUpdateOne {
	iuo.mutation.SetStatus(b)
	return iuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (iuo *ImageUpdateOne) SetNillableStatus(b *bool) *ImageUpdateOne {
	if b != nil {
		iuo.SetStatus(*b)
	}
	return iuo
}

// SetUpdatedAt sets the "updated_at" field.
func (iuo *ImageUpdateOne) SetUpdatedAt(t time.Time) *ImageUpdateOne {
	iuo.mutation.SetUpdatedAt(t)
	return iuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (iuo *ImageUpdateOne) SetNillableUpdatedAt(t *time.Time) *ImageUpdateOne {
	if t != nil {
		iuo.SetUpdatedAt(*t)
	}
	return iuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (iuo *ImageUpdateOne) ClearUpdatedAt() *ImageUpdateOne {
	iuo.mutation.ClearUpdatedAt()
	return iuo
}

// SetDeletedAt sets the "deleted_at" field.
func (iuo *ImageUpdateOne) SetDeletedAt(t time.Time) *ImageUpdateOne {
	iuo.mutation.SetDeletedAt(t)
	return iuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (iuo *ImageUpdateOne) SetNillableDeletedAt(t *time.Time) *ImageUpdateOne {
	if t != nil {
		iuo.SetDeletedAt(*t)
	}
	return iuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (iuo *ImageUpdateOne) ClearDeletedAt() *ImageUpdateOne {
	iuo.mutation.ClearDeletedAt()
	return iuo
}

// AddOwnerIDs adds the "owner" edge to the Product entity by IDs.
func (iuo *ImageUpdateOne) AddOwnerIDs(ids ...int) *ImageUpdateOne {
	iuo.mutation.AddOwnerIDs(ids...)
	return iuo
}

// AddOwner adds the "owner" edges to the Product entity.
func (iuo *ImageUpdateOne) AddOwner(p ...*Product) *ImageUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iuo.AddOwnerIDs(ids...)
}

// Mutation returns the ImageMutation object of the builder.
func (iuo *ImageUpdateOne) Mutation() *ImageMutation {
	return iuo.mutation
}

// ClearOwner clears all "owner" edges to the Product entity.
func (iuo *ImageUpdateOne) ClearOwner() *ImageUpdateOne {
	iuo.mutation.ClearOwner()
	return iuo
}

// RemoveOwnerIDs removes the "owner" edge to Product entities by IDs.
func (iuo *ImageUpdateOne) RemoveOwnerIDs(ids ...int) *ImageUpdateOne {
	iuo.mutation.RemoveOwnerIDs(ids...)
	return iuo
}

// RemoveOwner removes "owner" edges to Product entities.
func (iuo *ImageUpdateOne) RemoveOwner(p ...*Product) *ImageUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iuo.RemoveOwnerIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ImageUpdateOne) Select(field string, fields ...string) *ImageUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Image entity.
func (iuo *ImageUpdateOne) Save(ctx context.Context) (*Image, error) {
	var (
		err  error
		node *Image
	)
	if len(iuo.hooks) == 0 {
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ImageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			if iuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ImageUpdateOne) SaveX(ctx context.Context) *Image {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ImageUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ImageUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iuo *ImageUpdateOne) sqlSave(ctx context.Context) (_node *Image, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   image.Table,
			Columns: image.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: image.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Image.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, image.FieldID)
		for _, f := range fields {
			if !image.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != image.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: image.FieldTitle,
		})
	}
	if value, ok := iuo.mutation.Image(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: image.FieldImage,
		})
	}
	if value, ok := iuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: image.FieldStatus,
		})
	}
	if value, ok := iuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: image.FieldUpdatedAt,
		})
	}
	if iuo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: image.FieldUpdatedAt,
		})
	}
	if value, ok := iuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: image.FieldDeletedAt,
		})
	}
	if iuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: image.FieldDeletedAt,
		})
	}
	if iuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   image.OwnerTable,
			Columns: image.OwnerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedOwnerIDs(); len(nodes) > 0 && !iuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   image.OwnerTable,
			Columns: image.OwnerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   image.OwnerTable,
			Columns: image.OwnerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Image{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{image.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
