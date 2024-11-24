// Code generated by ent, DO NOT EDIT.

package ent

import (
	"calendar-backend/internal/storage/postgresql/ent/predicate"
	"calendar-backend/internal/storage/postgresql/ent/sports"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SportsDelete is the builder for deleting a Sports entity.
type SportsDelete struct {
	config
	hooks    []Hook
	mutation *SportsMutation
}

// Where appends a list predicates to the SportsDelete builder.
func (sd *SportsDelete) Where(ps ...predicate.Sports) *SportsDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *SportsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sd.sqlExec, sd.mutation, sd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *SportsDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *SportsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(sports.Table, sqlgraph.NewFieldSpec(sports.FieldID, field.TypeUUID))
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sd.mutation.done = true
	return affected, err
}

// SportsDeleteOne is the builder for deleting a single Sports entity.
type SportsDeleteOne struct {
	sd *SportsDelete
}

// Where appends a list predicates to the SportsDelete builder.
func (sdo *SportsDeleteOne) Where(ps ...predicate.Sports) *SportsDeleteOne {
	sdo.sd.mutation.Where(ps...)
	return sdo
}

// Exec executes the deletion query.
func (sdo *SportsDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sports.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *SportsDeleteOne) ExecX(ctx context.Context) {
	if err := sdo.Exec(ctx); err != nil {
		panic(err)
	}
}
