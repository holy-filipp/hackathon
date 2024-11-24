// Code generated by ent, DO NOT EDIT.

package ent

import (
	"calendar-backend/internal/storage/postgresql/ent/events"
	"calendar-backend/internal/storage/postgresql/ent/predicate"
	"calendar-backend/internal/storage/postgresql/ent/sports"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeEvents = "Events"
	TypeSports = "Sports"
)

// EventsMutation represents an operation that mutates the Events nodes in the graph.
type EventsMutation struct {
	config
	op              Op
	typ             string
	id              *int
	name            *string
	sex_age         *string
	discipline      *string
	time_start      *time.Time
	time_end        *time.Time
	venue           *string
	participants    *int
	addparticipants *int
	sport_id        *string
	clearedFields   map[string]struct{}
	done            bool
	oldValue        func(context.Context) (*Events, error)
	predicates      []predicate.Events
}

var _ ent.Mutation = (*EventsMutation)(nil)

// eventsOption allows management of the mutation configuration using functional options.
type eventsOption func(*EventsMutation)

// newEventsMutation creates new mutation for the Events entity.
func newEventsMutation(c config, op Op, opts ...eventsOption) *EventsMutation {
	m := &EventsMutation{
		config:        c,
		op:            op,
		typ:           TypeEvents,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withEventsID sets the ID field of the mutation.
func withEventsID(id int) eventsOption {
	return func(m *EventsMutation) {
		var (
			err   error
			once  sync.Once
			value *Events
		)
		m.oldValue = func(ctx context.Context) (*Events, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Events.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withEvents sets the old Events of the mutation.
func withEvents(node *Events) eventsOption {
	return func(m *EventsMutation) {
		m.oldValue = func(context.Context) (*Events, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m EventsMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m EventsMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Events entities.
func (m *EventsMutation) SetID(id int) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *EventsMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *EventsMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Events.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *EventsMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *EventsMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Events entity.
// If the Events object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EventsMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *EventsMutation) ResetName() {
	m.name = nil
}

// SetSexAge sets the "sex_age" field.
func (m *EventsMutation) SetSexAge(s string) {
	m.sex_age = &s
}

// SexAge returns the value of the "sex_age" field in the mutation.
func (m *EventsMutation) SexAge() (r string, exists bool) {
	v := m.sex_age
	if v == nil {
		return
	}
	return *v, true
}

// OldSexAge returns the old "sex_age" field's value of the Events entity.
// If the Events object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EventsMutation) OldSexAge(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSexAge is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSexAge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSexAge: %w", err)
	}
	return oldValue.SexAge, nil
}

// ResetSexAge resets all changes to the "sex_age" field.
func (m *EventsMutation) ResetSexAge() {
	m.sex_age = nil
}

// SetDiscipline sets the "discipline" field.
func (m *EventsMutation) SetDiscipline(s string) {
	m.discipline = &s
}

// Discipline returns the value of the "discipline" field in the mutation.
func (m *EventsMutation) Discipline() (r string, exists bool) {
	v := m.discipline
	if v == nil {
		return
	}
	return *v, true
}

// OldDiscipline returns the old "discipline" field's value of the Events entity.
// If the Events object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EventsMutation) OldDiscipline(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDiscipline is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDiscipline requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDiscipline: %w", err)
	}
	return oldValue.Discipline, nil
}

// ResetDiscipline resets all changes to the "discipline" field.
func (m *EventsMutation) ResetDiscipline() {
	m.discipline = nil
}

// SetTimeStart sets the "time_start" field.
func (m *EventsMutation) SetTimeStart(t time.Time) {
	m.time_start = &t
}

// TimeStart returns the value of the "time_start" field in the mutation.
func (m *EventsMutation) TimeStart() (r time.Time, exists bool) {
	v := m.time_start
	if v == nil {
		return
	}
	return *v, true
}

// OldTimeStart returns the old "time_start" field's value of the Events entity.
// If the Events object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EventsMutation) OldTimeStart(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTimeStart is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTimeStart requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTimeStart: %w", err)
	}
	return oldValue.TimeStart, nil
}

// ResetTimeStart resets all changes to the "time_start" field.
func (m *EventsMutation) ResetTimeStart() {
	m.time_start = nil
}

// SetTimeEnd sets the "time_end" field.
func (m *EventsMutation) SetTimeEnd(t time.Time) {
	m.time_end = &t
}

// TimeEnd returns the value of the "time_end" field in the mutation.
func (m *EventsMutation) TimeEnd() (r time.Time, exists bool) {
	v := m.time_end
	if v == nil {
		return
	}
	return *v, true
}

// OldTimeEnd returns the old "time_end" field's value of the Events entity.
// If the Events object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EventsMutation) OldTimeEnd(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTimeEnd is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTimeEnd requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTimeEnd: %w", err)
	}
	return oldValue.TimeEnd, nil
}

// ResetTimeEnd resets all changes to the "time_end" field.
func (m *EventsMutation) ResetTimeEnd() {
	m.time_end = nil
}

// SetVenue sets the "venue" field.
func (m *EventsMutation) SetVenue(s string) {
	m.venue = &s
}

// Venue returns the value of the "venue" field in the mutation.
func (m *EventsMutation) Venue() (r string, exists bool) {
	v := m.venue
	if v == nil {
		return
	}
	return *v, true
}

// OldVenue returns the old "venue" field's value of the Events entity.
// If the Events object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EventsMutation) OldVenue(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldVenue is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldVenue requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldVenue: %w", err)
	}
	return oldValue.Venue, nil
}

// ResetVenue resets all changes to the "venue" field.
func (m *EventsMutation) ResetVenue() {
	m.venue = nil
}

// SetParticipants sets the "participants" field.
func (m *EventsMutation) SetParticipants(i int) {
	m.participants = &i
	m.addparticipants = nil
}

// Participants returns the value of the "participants" field in the mutation.
func (m *EventsMutation) Participants() (r int, exists bool) {
	v := m.participants
	if v == nil {
		return
	}
	return *v, true
}

// OldParticipants returns the old "participants" field's value of the Events entity.
// If the Events object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EventsMutation) OldParticipants(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldParticipants is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldParticipants requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldParticipants: %w", err)
	}
	return oldValue.Participants, nil
}

// AddParticipants adds i to the "participants" field.
func (m *EventsMutation) AddParticipants(i int) {
	if m.addparticipants != nil {
		*m.addparticipants += i
	} else {
		m.addparticipants = &i
	}
}

// AddedParticipants returns the value that was added to the "participants" field in this mutation.
func (m *EventsMutation) AddedParticipants() (r int, exists bool) {
	v := m.addparticipants
	if v == nil {
		return
	}
	return *v, true
}

// ResetParticipants resets all changes to the "participants" field.
func (m *EventsMutation) ResetParticipants() {
	m.participants = nil
	m.addparticipants = nil
}

// SetSportID sets the "sport_id" field.
func (m *EventsMutation) SetSportID(s string) {
	m.sport_id = &s
}

// SportID returns the value of the "sport_id" field in the mutation.
func (m *EventsMutation) SportID() (r string, exists bool) {
	v := m.sport_id
	if v == nil {
		return
	}
	return *v, true
}

// OldSportID returns the old "sport_id" field's value of the Events entity.
// If the Events object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EventsMutation) OldSportID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSportID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSportID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSportID: %w", err)
	}
	return oldValue.SportID, nil
}

// ResetSportID resets all changes to the "sport_id" field.
func (m *EventsMutation) ResetSportID() {
	m.sport_id = nil
}

// Where appends a list predicates to the EventsMutation builder.
func (m *EventsMutation) Where(ps ...predicate.Events) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the EventsMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *EventsMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Events, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *EventsMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *EventsMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Events).
func (m *EventsMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *EventsMutation) Fields() []string {
	fields := make([]string, 0, 8)
	if m.name != nil {
		fields = append(fields, events.FieldName)
	}
	if m.sex_age != nil {
		fields = append(fields, events.FieldSexAge)
	}
	if m.discipline != nil {
		fields = append(fields, events.FieldDiscipline)
	}
	if m.time_start != nil {
		fields = append(fields, events.FieldTimeStart)
	}
	if m.time_end != nil {
		fields = append(fields, events.FieldTimeEnd)
	}
	if m.venue != nil {
		fields = append(fields, events.FieldVenue)
	}
	if m.participants != nil {
		fields = append(fields, events.FieldParticipants)
	}
	if m.sport_id != nil {
		fields = append(fields, events.FieldSportID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *EventsMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case events.FieldName:
		return m.Name()
	case events.FieldSexAge:
		return m.SexAge()
	case events.FieldDiscipline:
		return m.Discipline()
	case events.FieldTimeStart:
		return m.TimeStart()
	case events.FieldTimeEnd:
		return m.TimeEnd()
	case events.FieldVenue:
		return m.Venue()
	case events.FieldParticipants:
		return m.Participants()
	case events.FieldSportID:
		return m.SportID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *EventsMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case events.FieldName:
		return m.OldName(ctx)
	case events.FieldSexAge:
		return m.OldSexAge(ctx)
	case events.FieldDiscipline:
		return m.OldDiscipline(ctx)
	case events.FieldTimeStart:
		return m.OldTimeStart(ctx)
	case events.FieldTimeEnd:
		return m.OldTimeEnd(ctx)
	case events.FieldVenue:
		return m.OldVenue(ctx)
	case events.FieldParticipants:
		return m.OldParticipants(ctx)
	case events.FieldSportID:
		return m.OldSportID(ctx)
	}
	return nil, fmt.Errorf("unknown Events field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EventsMutation) SetField(name string, value ent.Value) error {
	switch name {
	case events.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case events.FieldSexAge:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSexAge(v)
		return nil
	case events.FieldDiscipline:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDiscipline(v)
		return nil
	case events.FieldTimeStart:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTimeStart(v)
		return nil
	case events.FieldTimeEnd:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTimeEnd(v)
		return nil
	case events.FieldVenue:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetVenue(v)
		return nil
	case events.FieldParticipants:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetParticipants(v)
		return nil
	case events.FieldSportID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSportID(v)
		return nil
	}
	return fmt.Errorf("unknown Events field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *EventsMutation) AddedFields() []string {
	var fields []string
	if m.addparticipants != nil {
		fields = append(fields, events.FieldParticipants)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *EventsMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case events.FieldParticipants:
		return m.AddedParticipants()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EventsMutation) AddField(name string, value ent.Value) error {
	switch name {
	case events.FieldParticipants:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddParticipants(v)
		return nil
	}
	return fmt.Errorf("unknown Events numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *EventsMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *EventsMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *EventsMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Events nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *EventsMutation) ResetField(name string) error {
	switch name {
	case events.FieldName:
		m.ResetName()
		return nil
	case events.FieldSexAge:
		m.ResetSexAge()
		return nil
	case events.FieldDiscipline:
		m.ResetDiscipline()
		return nil
	case events.FieldTimeStart:
		m.ResetTimeStart()
		return nil
	case events.FieldTimeEnd:
		m.ResetTimeEnd()
		return nil
	case events.FieldVenue:
		m.ResetVenue()
		return nil
	case events.FieldParticipants:
		m.ResetParticipants()
		return nil
	case events.FieldSportID:
		m.ResetSportID()
		return nil
	}
	return fmt.Errorf("unknown Events field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *EventsMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *EventsMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *EventsMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *EventsMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *EventsMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *EventsMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *EventsMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Events unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *EventsMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Events edge %s", name)
}

// SportsMutation represents an operation that mutates the Sports nodes in the graph.
type SportsMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	name          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Sports, error)
	predicates    []predicate.Sports
}

var _ ent.Mutation = (*SportsMutation)(nil)

// sportsOption allows management of the mutation configuration using functional options.
type sportsOption func(*SportsMutation)

// newSportsMutation creates new mutation for the Sports entity.
func newSportsMutation(c config, op Op, opts ...sportsOption) *SportsMutation {
	m := &SportsMutation{
		config:        c,
		op:            op,
		typ:           TypeSports,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withSportsID sets the ID field of the mutation.
func withSportsID(id uuid.UUID) sportsOption {
	return func(m *SportsMutation) {
		var (
			err   error
			once  sync.Once
			value *Sports
		)
		m.oldValue = func(ctx context.Context) (*Sports, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Sports.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withSports sets the old Sports of the mutation.
func withSports(node *Sports) sportsOption {
	return func(m *SportsMutation) {
		m.oldValue = func(context.Context) (*Sports, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m SportsMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m SportsMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Sports entities.
func (m *SportsMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *SportsMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *SportsMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Sports.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *SportsMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *SportsMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Sports entity.
// If the Sports object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SportsMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *SportsMutation) ResetName() {
	m.name = nil
}

// Where appends a list predicates to the SportsMutation builder.
func (m *SportsMutation) Where(ps ...predicate.Sports) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the SportsMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *SportsMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Sports, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *SportsMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *SportsMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Sports).
func (m *SportsMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *SportsMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, sports.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *SportsMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case sports.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *SportsMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case sports.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Sports field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SportsMutation) SetField(name string, value ent.Value) error {
	switch name {
	case sports.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Sports field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *SportsMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *SportsMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SportsMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Sports numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *SportsMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *SportsMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *SportsMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Sports nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *SportsMutation) ResetField(name string) error {
	switch name {
	case sports.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Sports field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *SportsMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *SportsMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *SportsMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *SportsMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *SportsMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *SportsMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *SportsMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Sports unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *SportsMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Sports edge %s", name)
}
