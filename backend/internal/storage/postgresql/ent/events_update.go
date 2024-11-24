// Code generated by ent, DO NOT EDIT.

package ent

import (
	"calendar-backend/internal/storage/postgresql/ent/events"
	"calendar-backend/internal/storage/postgresql/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EventsUpdate is the builder for updating Events entities.
type EventsUpdate struct {
	config
	hooks    []Hook
	mutation *EventsMutation
}

// Where appends a list predicates to the EventsUpdate builder.
func (eu *EventsUpdate) Where(ps ...predicate.Events) *EventsUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetName sets the "name" field.
func (eu *EventsUpdate) SetName(s string) *EventsUpdate {
	eu.mutation.SetName(s)
	return eu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (eu *EventsUpdate) SetNillableName(s *string) *EventsUpdate {
	if s != nil {
		eu.SetName(*s)
	}
	return eu
}

// SetSexAge sets the "sex_age" field.
func (eu *EventsUpdate) SetSexAge(s string) *EventsUpdate {
	eu.mutation.SetSexAge(s)
	return eu
}

// SetNillableSexAge sets the "sex_age" field if the given value is not nil.
func (eu *EventsUpdate) SetNillableSexAge(s *string) *EventsUpdate {
	if s != nil {
		eu.SetSexAge(*s)
	}
	return eu
}

// SetDiscipline sets the "discipline" field.
func (eu *EventsUpdate) SetDiscipline(s string) *EventsUpdate {
	eu.mutation.SetDiscipline(s)
	return eu
}

// SetNillableDiscipline sets the "discipline" field if the given value is not nil.
func (eu *EventsUpdate) SetNillableDiscipline(s *string) *EventsUpdate {
	if s != nil {
		eu.SetDiscipline(*s)
	}
	return eu
}

// SetTimeStart sets the "time_start" field.
func (eu *EventsUpdate) SetTimeStart(t time.Time) *EventsUpdate {
	eu.mutation.SetTimeStart(t)
	return eu
}

// SetNillableTimeStart sets the "time_start" field if the given value is not nil.
func (eu *EventsUpdate) SetNillableTimeStart(t *time.Time) *EventsUpdate {
	if t != nil {
		eu.SetTimeStart(*t)
	}
	return eu
}

// SetTimeEnd sets the "time_end" field.
func (eu *EventsUpdate) SetTimeEnd(t time.Time) *EventsUpdate {
	eu.mutation.SetTimeEnd(t)
	return eu
}

// SetNillableTimeEnd sets the "time_end" field if the given value is not nil.
func (eu *EventsUpdate) SetNillableTimeEnd(t *time.Time) *EventsUpdate {
	if t != nil {
		eu.SetTimeEnd(*t)
	}
	return eu
}

// SetVenue sets the "venue" field.
func (eu *EventsUpdate) SetVenue(s string) *EventsUpdate {
	eu.mutation.SetVenue(s)
	return eu
}

// SetNillableVenue sets the "venue" field if the given value is not nil.
func (eu *EventsUpdate) SetNillableVenue(s *string) *EventsUpdate {
	if s != nil {
		eu.SetVenue(*s)
	}
	return eu
}

// SetParticipants sets the "participants" field.
func (eu *EventsUpdate) SetParticipants(i int) *EventsUpdate {
	eu.mutation.ResetParticipants()
	eu.mutation.SetParticipants(i)
	return eu
}

// SetNillableParticipants sets the "participants" field if the given value is not nil.
func (eu *EventsUpdate) SetNillableParticipants(i *int) *EventsUpdate {
	if i != nil {
		eu.SetParticipants(*i)
	}
	return eu
}

// AddParticipants adds i to the "participants" field.
func (eu *EventsUpdate) AddParticipants(i int) *EventsUpdate {
	eu.mutation.AddParticipants(i)
	return eu
}

// SetSportID sets the "sport_id" field.
func (eu *EventsUpdate) SetSportID(s string) *EventsUpdate {
	eu.mutation.SetSportID(s)
	return eu
}

// SetNillableSportID sets the "sport_id" field if the given value is not nil.
func (eu *EventsUpdate) SetNillableSportID(s *string) *EventsUpdate {
	if s != nil {
		eu.SetSportID(*s)
	}
	return eu
}

// Mutation returns the EventsMutation object of the builder.
func (eu *EventsUpdate) Mutation() *EventsMutation {
	return eu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EventsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EventsUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EventsUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EventsUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eu *EventsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(events.Table, events.Columns, sqlgraph.NewFieldSpec(events.FieldID, field.TypeInt))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.SetField(events.FieldName, field.TypeString, value)
	}
	if value, ok := eu.mutation.SexAge(); ok {
		_spec.SetField(events.FieldSexAge, field.TypeString, value)
	}
	if value, ok := eu.mutation.Discipline(); ok {
		_spec.SetField(events.FieldDiscipline, field.TypeString, value)
	}
	if value, ok := eu.mutation.TimeStart(); ok {
		_spec.SetField(events.FieldTimeStart, field.TypeTime, value)
	}
	if value, ok := eu.mutation.TimeEnd(); ok {
		_spec.SetField(events.FieldTimeEnd, field.TypeTime, value)
	}
	if value, ok := eu.mutation.Venue(); ok {
		_spec.SetField(events.FieldVenue, field.TypeString, value)
	}
	if value, ok := eu.mutation.Participants(); ok {
		_spec.SetField(events.FieldParticipants, field.TypeInt, value)
	}
	if value, ok := eu.mutation.AddedParticipants(); ok {
		_spec.AddField(events.FieldParticipants, field.TypeInt, value)
	}
	if value, ok := eu.mutation.SportID(); ok {
		_spec.SetField(events.FieldSportID, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{events.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EventsUpdateOne is the builder for updating a single Events entity.
type EventsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EventsMutation
}

// SetName sets the "name" field.
func (euo *EventsUpdateOne) SetName(s string) *EventsUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (euo *EventsUpdateOne) SetNillableName(s *string) *EventsUpdateOne {
	if s != nil {
		euo.SetName(*s)
	}
	return euo
}

// SetSexAge sets the "sex_age" field.
func (euo *EventsUpdateOne) SetSexAge(s string) *EventsUpdateOne {
	euo.mutation.SetSexAge(s)
	return euo
}

// SetNillableSexAge sets the "sex_age" field if the given value is not nil.
func (euo *EventsUpdateOne) SetNillableSexAge(s *string) *EventsUpdateOne {
	if s != nil {
		euo.SetSexAge(*s)
	}
	return euo
}

// SetDiscipline sets the "discipline" field.
func (euo *EventsUpdateOne) SetDiscipline(s string) *EventsUpdateOne {
	euo.mutation.SetDiscipline(s)
	return euo
}

// SetNillableDiscipline sets the "discipline" field if the given value is not nil.
func (euo *EventsUpdateOne) SetNillableDiscipline(s *string) *EventsUpdateOne {
	if s != nil {
		euo.SetDiscipline(*s)
	}
	return euo
}

// SetTimeStart sets the "time_start" field.
func (euo *EventsUpdateOne) SetTimeStart(t time.Time) *EventsUpdateOne {
	euo.mutation.SetTimeStart(t)
	return euo
}

// SetNillableTimeStart sets the "time_start" field if the given value is not nil.
func (euo *EventsUpdateOne) SetNillableTimeStart(t *time.Time) *EventsUpdateOne {
	if t != nil {
		euo.SetTimeStart(*t)
	}
	return euo
}

// SetTimeEnd sets the "time_end" field.
func (euo *EventsUpdateOne) SetTimeEnd(t time.Time) *EventsUpdateOne {
	euo.mutation.SetTimeEnd(t)
	return euo
}

// SetNillableTimeEnd sets the "time_end" field if the given value is not nil.
func (euo *EventsUpdateOne) SetNillableTimeEnd(t *time.Time) *EventsUpdateOne {
	if t != nil {
		euo.SetTimeEnd(*t)
	}
	return euo
}

// SetVenue sets the "venue" field.
func (euo *EventsUpdateOne) SetVenue(s string) *EventsUpdateOne {
	euo.mutation.SetVenue(s)
	return euo
}

// SetNillableVenue sets the "venue" field if the given value is not nil.
func (euo *EventsUpdateOne) SetNillableVenue(s *string) *EventsUpdateOne {
	if s != nil {
		euo.SetVenue(*s)
	}
	return euo
}

// SetParticipants sets the "participants" field.
func (euo *EventsUpdateOne) SetParticipants(i int) *EventsUpdateOne {
	euo.mutation.ResetParticipants()
	euo.mutation.SetParticipants(i)
	return euo
}

// SetNillableParticipants sets the "participants" field if the given value is not nil.
func (euo *EventsUpdateOne) SetNillableParticipants(i *int) *EventsUpdateOne {
	if i != nil {
		euo.SetParticipants(*i)
	}
	return euo
}

// AddParticipants adds i to the "participants" field.
func (euo *EventsUpdateOne) AddParticipants(i int) *EventsUpdateOne {
	euo.mutation.AddParticipants(i)
	return euo
}

// SetSportID sets the "sport_id" field.
func (euo *EventsUpdateOne) SetSportID(s string) *EventsUpdateOne {
	euo.mutation.SetSportID(s)
	return euo
}

// SetNillableSportID sets the "sport_id" field if the given value is not nil.
func (euo *EventsUpdateOne) SetNillableSportID(s *string) *EventsUpdateOne {
	if s != nil {
		euo.SetSportID(*s)
	}
	return euo
}

// Mutation returns the EventsMutation object of the builder.
func (euo *EventsUpdateOne) Mutation() *EventsMutation {
	return euo.mutation
}

// Where appends a list predicates to the EventsUpdate builder.
func (euo *EventsUpdateOne) Where(ps ...predicate.Events) *EventsUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EventsUpdateOne) Select(field string, fields ...string) *EventsUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Events entity.
func (euo *EventsUpdateOne) Save(ctx context.Context) (*Events, error) {
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EventsUpdateOne) SaveX(ctx context.Context) *Events {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EventsUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EventsUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (euo *EventsUpdateOne) sqlSave(ctx context.Context) (_node *Events, err error) {
	_spec := sqlgraph.NewUpdateSpec(events.Table, events.Columns, sqlgraph.NewFieldSpec(events.FieldID, field.TypeInt))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Events.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, events.FieldID)
		for _, f := range fields {
			if !events.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != events.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Name(); ok {
		_spec.SetField(events.FieldName, field.TypeString, value)
	}
	if value, ok := euo.mutation.SexAge(); ok {
		_spec.SetField(events.FieldSexAge, field.TypeString, value)
	}
	if value, ok := euo.mutation.Discipline(); ok {
		_spec.SetField(events.FieldDiscipline, field.TypeString, value)
	}
	if value, ok := euo.mutation.TimeStart(); ok {
		_spec.SetField(events.FieldTimeStart, field.TypeTime, value)
	}
	if value, ok := euo.mutation.TimeEnd(); ok {
		_spec.SetField(events.FieldTimeEnd, field.TypeTime, value)
	}
	if value, ok := euo.mutation.Venue(); ok {
		_spec.SetField(events.FieldVenue, field.TypeString, value)
	}
	if value, ok := euo.mutation.Participants(); ok {
		_spec.SetField(events.FieldParticipants, field.TypeInt, value)
	}
	if value, ok := euo.mutation.AddedParticipants(); ok {
		_spec.AddField(events.FieldParticipants, field.TypeInt, value)
	}
	if value, ok := euo.mutation.SportID(); ok {
		_spec.SetField(events.FieldSportID, field.TypeString, value)
	}
	_node = &Events{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{events.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}
