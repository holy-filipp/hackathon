// Code generated by ent, DO NOT EDIT.

package ent

import (
	"calendar-backend/internal/storage/postgresql/ent/predicate"
	"calendar-backend/internal/storage/postgresql/ent/sports"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SportsUpdate is the builder for updating Sports entities.
type SportsUpdate struct {
	config
	hooks    []Hook
	mutation *SportsMutation
}

// Where appends a list predicates to the SportsUpdate builder.
func (su *SportsUpdate) Where(ps ...predicate.Sports) *SportsUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *SportsUpdate) SetName(s string) *SportsUpdate {
	su.mutation.SetName(s)
	return su
}

// SetNillableName sets the "name" field if the given value is not nil.
func (su *SportsUpdate) SetNillableName(s *string) *SportsUpdate {
	if s != nil {
		su.SetName(*s)
	}
	return su
}

// Mutation returns the SportsMutation object of the builder.
func (su *SportsUpdate) Mutation() *SportsMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SportsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SportsUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SportsUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SportsUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SportsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(sports.Table, sports.Columns, sqlgraph.NewFieldSpec(sports.FieldID, field.TypeUUID))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(sports.FieldName, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sports.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SportsUpdateOne is the builder for updating a single Sports entity.
type SportsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SportsMutation
}

// SetName sets the "name" field.
func (suo *SportsUpdateOne) SetName(s string) *SportsUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (suo *SportsUpdateOne) SetNillableName(s *string) *SportsUpdateOne {
	if s != nil {
		suo.SetName(*s)
	}
	return suo
}

// Mutation returns the SportsMutation object of the builder.
func (suo *SportsUpdateOne) Mutation() *SportsMutation {
	return suo.mutation
}

// Where appends a list predicates to the SportsUpdate builder.
func (suo *SportsUpdateOne) Where(ps ...predicate.Sports) *SportsUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SportsUpdateOne) Select(field string, fields ...string) *SportsUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Sports entity.
func (suo *SportsUpdateOne) Save(ctx context.Context) (*Sports, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SportsUpdateOne) SaveX(ctx context.Context) *Sports {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SportsUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SportsUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SportsUpdateOne) sqlSave(ctx context.Context) (_node *Sports, err error) {
	_spec := sqlgraph.NewUpdateSpec(sports.Table, sports.Columns, sqlgraph.NewFieldSpec(sports.FieldID, field.TypeUUID))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Sports.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sports.FieldID)
		for _, f := range fields {
			if !sports.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sports.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(sports.FieldName, field.TypeString, value)
	}
	_node = &Sports{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sports.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
