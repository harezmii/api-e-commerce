// Code generated by entc, DO NOT EDIT.

package ent

import (
	"api/ent/predicate"
	"api/ent/profile"
	"api/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProfileUpdate is the builder for updating Profile entities.
type ProfileUpdate struct {
	config
	hooks    []Hook
	mutation *ProfileMutation
}

// Where appends a list predicates to the ProfileUpdate builder.
func (pu *ProfileUpdate) Where(ps ...predicate.Profile) *ProfileUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetAddress sets the "address" field.
func (pu *ProfileUpdate) SetAddress(s string) *ProfileUpdate {
	pu.mutation.SetAddress(s)
	return pu
}

// SetPhone sets the "phone" field.
func (pu *ProfileUpdate) SetPhone(s string) *ProfileUpdate {
	pu.mutation.SetPhone(s)
	return pu
}

// SetImage sets the "image" field.
func (pu *ProfileUpdate) SetImage(s string) *ProfileUpdate {
	pu.mutation.SetImage(s)
	return pu
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableImage(s *string) *ProfileUpdate {
	if s != nil {
		pu.SetImage(*s)
	}
	return pu
}

// ClearImage clears the value of the "image" field.
func (pu *ProfileUpdate) ClearImage() *ProfileUpdate {
	pu.mutation.ClearImage()
	return pu
}

// SetURL sets the "url" field.
func (pu *ProfileUpdate) SetURL(s string) *ProfileUpdate {
	pu.mutation.SetURL(s)
	return pu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableURL(s *string) *ProfileUpdate {
	if s != nil {
		pu.SetURL(*s)
	}
	return pu
}

// ClearURL clears the value of the "url" field.
func (pu *ProfileUpdate) ClearURL() *ProfileUpdate {
	pu.mutation.ClearURL()
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *ProfileUpdate) SetCreatedAt(t time.Time) *ProfileUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableCreatedAt(t *time.Time) *ProfileUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *ProfileUpdate) SetUpdatedAt(t time.Time) *ProfileUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableUpdatedAt(t *time.Time) *ProfileUpdate {
	if t != nil {
		pu.SetUpdatedAt(*t)
	}
	return pu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (pu *ProfileUpdate) ClearUpdatedAt() *ProfileUpdate {
	pu.mutation.ClearUpdatedAt()
	return pu
}

// SetDeletedAt sets the "deleted_at" field.
func (pu *ProfileUpdate) SetDeletedAt(t time.Time) *ProfileUpdate {
	pu.mutation.SetDeletedAt(t)
	return pu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableDeletedAt(t *time.Time) *ProfileUpdate {
	if t != nil {
		pu.SetDeletedAt(*t)
	}
	return pu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pu *ProfileUpdate) ClearDeletedAt() *ProfileUpdate {
	pu.mutation.ClearDeletedAt()
	return pu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (pu *ProfileUpdate) SetOwnerID(id int) *ProfileUpdate {
	pu.mutation.SetOwnerID(id)
	return pu
}

// SetOwner sets the "owner" edge to the User entity.
func (pu *ProfileUpdate) SetOwner(u *User) *ProfileUpdate {
	return pu.SetOwnerID(u.ID)
}

// Mutation returns the ProfileMutation object of the builder.
func (pu *ProfileUpdate) Mutation() *ProfileMutation {
	return pu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (pu *ProfileUpdate) ClearOwner() *ProfileUpdate {
	pu.mutation.ClearOwner()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProfileUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProfileUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProfileUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProfileUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProfileUpdate) check() error {
	if _, ok := pu.mutation.OwnerID(); pu.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Profile.owner"`)
	}
	return nil
}

func (pu *ProfileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   profile.Table,
			Columns: profile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: profile.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldAddress,
		})
	}
	if value, ok := pu.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldPhone,
		})
	}
	if value, ok := pu.mutation.Image(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldImage,
		})
	}
	if pu.mutation.ImageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: profile.FieldImage,
		})
	}
	if value, ok := pu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldURL,
		})
	}
	if pu.mutation.URLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: profile.FieldURL,
		})
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldCreatedAt,
		})
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldUpdatedAt,
		})
	}
	if pu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: profile.FieldUpdatedAt,
		})
	}
	if value, ok := pu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldDeletedAt,
		})
	}
	if pu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: profile.FieldDeletedAt,
		})
	}
	if pu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   profile.OwnerTable,
			Columns: []string{profile.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   profile.OwnerTable,
			Columns: []string{profile.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ProfileUpdateOne is the builder for updating a single Profile entity.
type ProfileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProfileMutation
}

// SetAddress sets the "address" field.
func (puo *ProfileUpdateOne) SetAddress(s string) *ProfileUpdateOne {
	puo.mutation.SetAddress(s)
	return puo
}

// SetPhone sets the "phone" field.
func (puo *ProfileUpdateOne) SetPhone(s string) *ProfileUpdateOne {
	puo.mutation.SetPhone(s)
	return puo
}

// SetImage sets the "image" field.
func (puo *ProfileUpdateOne) SetImage(s string) *ProfileUpdateOne {
	puo.mutation.SetImage(s)
	return puo
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableImage(s *string) *ProfileUpdateOne {
	if s != nil {
		puo.SetImage(*s)
	}
	return puo
}

// ClearImage clears the value of the "image" field.
func (puo *ProfileUpdateOne) ClearImage() *ProfileUpdateOne {
	puo.mutation.ClearImage()
	return puo
}

// SetURL sets the "url" field.
func (puo *ProfileUpdateOne) SetURL(s string) *ProfileUpdateOne {
	puo.mutation.SetURL(s)
	return puo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableURL(s *string) *ProfileUpdateOne {
	if s != nil {
		puo.SetURL(*s)
	}
	return puo
}

// ClearURL clears the value of the "url" field.
func (puo *ProfileUpdateOne) ClearURL() *ProfileUpdateOne {
	puo.mutation.ClearURL()
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *ProfileUpdateOne) SetCreatedAt(t time.Time) *ProfileUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableCreatedAt(t *time.Time) *ProfileUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *ProfileUpdateOne) SetUpdatedAt(t time.Time) *ProfileUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableUpdatedAt(t *time.Time) *ProfileUpdateOne {
	if t != nil {
		puo.SetUpdatedAt(*t)
	}
	return puo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (puo *ProfileUpdateOne) ClearUpdatedAt() *ProfileUpdateOne {
	puo.mutation.ClearUpdatedAt()
	return puo
}

// SetDeletedAt sets the "deleted_at" field.
func (puo *ProfileUpdateOne) SetDeletedAt(t time.Time) *ProfileUpdateOne {
	puo.mutation.SetDeletedAt(t)
	return puo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableDeletedAt(t *time.Time) *ProfileUpdateOne {
	if t != nil {
		puo.SetDeletedAt(*t)
	}
	return puo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (puo *ProfileUpdateOne) ClearDeletedAt() *ProfileUpdateOne {
	puo.mutation.ClearDeletedAt()
	return puo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (puo *ProfileUpdateOne) SetOwnerID(id int) *ProfileUpdateOne {
	puo.mutation.SetOwnerID(id)
	return puo
}

// SetOwner sets the "owner" edge to the User entity.
func (puo *ProfileUpdateOne) SetOwner(u *User) *ProfileUpdateOne {
	return puo.SetOwnerID(u.ID)
}

// Mutation returns the ProfileMutation object of the builder.
func (puo *ProfileUpdateOne) Mutation() *ProfileMutation {
	return puo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (puo *ProfileUpdateOne) ClearOwner() *ProfileUpdateOne {
	puo.mutation.ClearOwner()
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProfileUpdateOne) Select(field string, fields ...string) *ProfileUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Profile entity.
func (puo *ProfileUpdateOne) Save(ctx context.Context) (*Profile, error) {
	var (
		err  error
		node *Profile
	)
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProfileUpdateOne) SaveX(ctx context.Context) *Profile {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProfileUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProfileUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProfileUpdateOne) check() error {
	if _, ok := puo.mutation.OwnerID(); puo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Profile.owner"`)
	}
	return nil
}

func (puo *ProfileUpdateOne) sqlSave(ctx context.Context) (_node *Profile, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   profile.Table,
			Columns: profile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: profile.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Profile.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, profile.FieldID)
		for _, f := range fields {
			if !profile.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != profile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldAddress,
		})
	}
	if value, ok := puo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldPhone,
		})
	}
	if value, ok := puo.mutation.Image(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldImage,
		})
	}
	if puo.mutation.ImageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: profile.FieldImage,
		})
	}
	if value, ok := puo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldURL,
		})
	}
	if puo.mutation.URLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: profile.FieldURL,
		})
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldCreatedAt,
		})
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldUpdatedAt,
		})
	}
	if puo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: profile.FieldUpdatedAt,
		})
	}
	if value, ok := puo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldDeletedAt,
		})
	}
	if puo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: profile.FieldDeletedAt,
		})
	}
	if puo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   profile.OwnerTable,
			Columns: []string{profile.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   profile.OwnerTable,
			Columns: []string{profile.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Profile{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
