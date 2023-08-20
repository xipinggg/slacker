// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"slacker/internal/data/ent/predicate"
	"slacker/internal/data/ent/record"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RecordUpdate is the builder for updating Record entities.
type RecordUpdate struct {
	config
	hooks    []Hook
	mutation *RecordMutation
}

// Where appends a list predicates to the RecordUpdate builder.
func (ru *RecordUpdate) Where(ps ...predicate.Record) *RecordUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetEndTime sets the "end_time" field.
func (ru *RecordUpdate) SetEndTime(t time.Time) *RecordUpdate {
	ru.mutation.SetEndTime(t)
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RecordUpdate) SetUpdatedAt(t time.Time) *RecordUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ru *RecordUpdate) SetNillableUpdatedAt(t *time.Time) *RecordUpdate {
	if t != nil {
		ru.SetUpdatedAt(*t)
	}
	return ru
}

// Mutation returns the RecordMutation object of the builder.
func (ru *RecordUpdate) Mutation() *RecordMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RecordUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RecordUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RecordUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RecordUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RecordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(record.Table, record.Columns, sqlgraph.NewFieldSpec(record.FieldID, field.TypeUUID))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.EndTime(); ok {
		_spec.SetField(record.FieldEndTime, field.TypeTime, value)
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(record.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{record.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RecordUpdateOne is the builder for updating a single Record entity.
type RecordUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RecordMutation
}

// SetEndTime sets the "end_time" field.
func (ruo *RecordUpdateOne) SetEndTime(t time.Time) *RecordUpdateOne {
	ruo.mutation.SetEndTime(t)
	return ruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RecordUpdateOne) SetUpdatedAt(t time.Time) *RecordUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ruo *RecordUpdateOne) SetNillableUpdatedAt(t *time.Time) *RecordUpdateOne {
	if t != nil {
		ruo.SetUpdatedAt(*t)
	}
	return ruo
}

// Mutation returns the RecordMutation object of the builder.
func (ruo *RecordUpdateOne) Mutation() *RecordMutation {
	return ruo.mutation
}

// Where appends a list predicates to the RecordUpdate builder.
func (ruo *RecordUpdateOne) Where(ps ...predicate.Record) *RecordUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RecordUpdateOne) Select(field string, fields ...string) *RecordUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Record entity.
func (ruo *RecordUpdateOne) Save(ctx context.Context) (*Record, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RecordUpdateOne) SaveX(ctx context.Context) *Record {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RecordUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RecordUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RecordUpdateOne) sqlSave(ctx context.Context) (_node *Record, err error) {
	_spec := sqlgraph.NewUpdateSpec(record.Table, record.Columns, sqlgraph.NewFieldSpec(record.FieldID, field.TypeUUID))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Record.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, record.FieldID)
		for _, f := range fields {
			if !record.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != record.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.EndTime(); ok {
		_spec.SetField(record.FieldEndTime, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(record.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Record{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{record.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}