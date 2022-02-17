// Code generated by entc, DO NOT EDIT.

package ent

import (
	"api/ent/profile"
	"api/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProfileCreate is the builder for creating a Profile entity.
type ProfileCreate struct {
	config
	mutation *ProfileMutation
	hooks    []Hook
}

// SetAddress sets the "address" field.
func (pc *ProfileCreate) SetAddress(s string) *ProfileCreate {
	pc.mutation.SetAddress(s)
	return pc
}

// SetPhone sets the "phone" field.
func (pc *ProfileCreate) SetPhone(s string) *ProfileCreate {
	pc.mutation.SetPhone(s)
	return pc
}

// SetImage sets the "image" field.
func (pc *ProfileCreate) SetImage(s string) *ProfileCreate {
	pc.mutation.SetImage(s)
	return pc
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableImage(s *string) *ProfileCreate {
	if s != nil {
		pc.SetImage(*s)
	}
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *ProfileCreate) SetCreatedAt(t time.Time) *ProfileCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableCreatedAt(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *ProfileCreate) SetUpdatedAt(t time.Time) *ProfileCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableUpdatedAt(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetDeletedAt sets the "deleted_at" field.
func (pc *ProfileCreate) SetDeletedAt(t time.Time) *ProfileCreate {
	pc.mutation.SetDeletedAt(t)
	return pc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableDeletedAt(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetDeletedAt(*t)
	}
	return pc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (pc *ProfileCreate) SetOwnerID(id int) *ProfileCreate {
	pc.mutation.SetOwnerID(id)
	return pc
}

// SetOwner sets the "owner" edge to the User entity.
func (pc *ProfileCreate) SetOwner(u *User) *ProfileCreate {
	return pc.SetOwnerID(u.ID)
}

// Mutation returns the ProfileMutation object of the builder.
func (pc *ProfileCreate) Mutation() *ProfileMutation {
	return pc.mutation
}

// Save creates the Profile in the database.
func (pc *ProfileCreate) Save(ctx context.Context) (*Profile, error) {
	var (
		err  error
		node *Profile
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProfileCreate) SaveX(ctx context.Context) *Profile {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProfileCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProfileCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProfileCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := profile.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProfileCreate) check() error {
	if _, ok := pc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Profile.address"`)}
	}
	if _, ok := pc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent: missing required field "Profile.phone"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Profile.created_at"`)}
	}
	if _, ok := pc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Profile.owner"`)}
	}
	return nil
}

func (pc *ProfileCreate) sqlSave(ctx context.Context) (*Profile, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *ProfileCreate) createSpec() (*Profile, *sqlgraph.CreateSpec) {
	var (
		_node = &Profile{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: profile.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: profile.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldAddress,
		})
		_node.Address = value
	}
	if value, ok := pc.mutation.Phone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldPhone,
		})
		_node.Phone = value
	}
	if value, ok := pc.mutation.Image(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldImage,
		})
		_node.Image = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if nodes := pc.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_node.user_profile = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProfileCreateBulk is the builder for creating many Profile entities in bulk.
type ProfileCreateBulk struct {
	config
	builders []*ProfileCreate
}

// Save creates the Profile entities in the database.
func (pcb *ProfileCreateBulk) Save(ctx context.Context) ([]*Profile, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Profile, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProfileMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProfileCreateBulk) SaveX(ctx context.Context) []*Profile {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProfileCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProfileCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
