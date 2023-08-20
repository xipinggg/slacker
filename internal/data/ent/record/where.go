// Code generated by ent, DO NOT EDIT.

package record

import (
	"slacker/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Record {
	return predicate.Record(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Record {
	return predicate.Record(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Record {
	return predicate.Record(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Record {
	return predicate.Record(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Record {
	return predicate.Record(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Record {
	return predicate.Record(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Record {
	return predicate.Record(sql.FieldLTE(FieldID, id))
}

// CreatorID applies equality check predicate on the "creator_id" field. It's identical to CreatorIDEQ.
func CreatorID(v string) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldCreatorID, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldType, v))
}

// BeginTime applies equality check predicate on the "begin_time" field. It's identical to BeginTimeEQ.
func BeginTime(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldBeginTime, v))
}

// EndTime applies equality check predicate on the "end_time" field. It's identical to EndTimeEQ.
func EndTime(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldEndTime, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatorIDEQ applies the EQ predicate on the "creator_id" field.
func CreatorIDEQ(v string) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldCreatorID, v))
}

// CreatorIDNEQ applies the NEQ predicate on the "creator_id" field.
func CreatorIDNEQ(v string) predicate.Record {
	return predicate.Record(sql.FieldNEQ(FieldCreatorID, v))
}

// CreatorIDIn applies the In predicate on the "creator_id" field.
func CreatorIDIn(vs ...string) predicate.Record {
	return predicate.Record(sql.FieldIn(FieldCreatorID, vs...))
}

// CreatorIDNotIn applies the NotIn predicate on the "creator_id" field.
func CreatorIDNotIn(vs ...string) predicate.Record {
	return predicate.Record(sql.FieldNotIn(FieldCreatorID, vs...))
}

// CreatorIDGT applies the GT predicate on the "creator_id" field.
func CreatorIDGT(v string) predicate.Record {
	return predicate.Record(sql.FieldGT(FieldCreatorID, v))
}

// CreatorIDGTE applies the GTE predicate on the "creator_id" field.
func CreatorIDGTE(v string) predicate.Record {
	return predicate.Record(sql.FieldGTE(FieldCreatorID, v))
}

// CreatorIDLT applies the LT predicate on the "creator_id" field.
func CreatorIDLT(v string) predicate.Record {
	return predicate.Record(sql.FieldLT(FieldCreatorID, v))
}

// CreatorIDLTE applies the LTE predicate on the "creator_id" field.
func CreatorIDLTE(v string) predicate.Record {
	return predicate.Record(sql.FieldLTE(FieldCreatorID, v))
}

// CreatorIDContains applies the Contains predicate on the "creator_id" field.
func CreatorIDContains(v string) predicate.Record {
	return predicate.Record(sql.FieldContains(FieldCreatorID, v))
}

// CreatorIDHasPrefix applies the HasPrefix predicate on the "creator_id" field.
func CreatorIDHasPrefix(v string) predicate.Record {
	return predicate.Record(sql.FieldHasPrefix(FieldCreatorID, v))
}

// CreatorIDHasSuffix applies the HasSuffix predicate on the "creator_id" field.
func CreatorIDHasSuffix(v string) predicate.Record {
	return predicate.Record(sql.FieldHasSuffix(FieldCreatorID, v))
}

// CreatorIDEqualFold applies the EqualFold predicate on the "creator_id" field.
func CreatorIDEqualFold(v string) predicate.Record {
	return predicate.Record(sql.FieldEqualFold(FieldCreatorID, v))
}

// CreatorIDContainsFold applies the ContainsFold predicate on the "creator_id" field.
func CreatorIDContainsFold(v string) predicate.Record {
	return predicate.Record(sql.FieldContainsFold(FieldCreatorID, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Record {
	return predicate.Record(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Record {
	return predicate.Record(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Record {
	return predicate.Record(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Record {
	return predicate.Record(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Record {
	return predicate.Record(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Record {
	return predicate.Record(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Record {
	return predicate.Record(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Record {
	return predicate.Record(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Record {
	return predicate.Record(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Record {
	return predicate.Record(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Record {
	return predicate.Record(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Record {
	return predicate.Record(sql.FieldContainsFold(FieldType, v))
}

// BeginTimeEQ applies the EQ predicate on the "begin_time" field.
func BeginTimeEQ(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldBeginTime, v))
}

// BeginTimeNEQ applies the NEQ predicate on the "begin_time" field.
func BeginTimeNEQ(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldNEQ(FieldBeginTime, v))
}

// BeginTimeIn applies the In predicate on the "begin_time" field.
func BeginTimeIn(vs ...time.Time) predicate.Record {
	return predicate.Record(sql.FieldIn(FieldBeginTime, vs...))
}

// BeginTimeNotIn applies the NotIn predicate on the "begin_time" field.
func BeginTimeNotIn(vs ...time.Time) predicate.Record {
	return predicate.Record(sql.FieldNotIn(FieldBeginTime, vs...))
}

// BeginTimeGT applies the GT predicate on the "begin_time" field.
func BeginTimeGT(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldGT(FieldBeginTime, v))
}

// BeginTimeGTE applies the GTE predicate on the "begin_time" field.
func BeginTimeGTE(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldGTE(FieldBeginTime, v))
}

// BeginTimeLT applies the LT predicate on the "begin_time" field.
func BeginTimeLT(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldLT(FieldBeginTime, v))
}

// BeginTimeLTE applies the LTE predicate on the "begin_time" field.
func BeginTimeLTE(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldLTE(FieldBeginTime, v))
}

// EndTimeEQ applies the EQ predicate on the "end_time" field.
func EndTimeEQ(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldEndTime, v))
}

// EndTimeNEQ applies the NEQ predicate on the "end_time" field.
func EndTimeNEQ(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldNEQ(FieldEndTime, v))
}

// EndTimeIn applies the In predicate on the "end_time" field.
func EndTimeIn(vs ...time.Time) predicate.Record {
	return predicate.Record(sql.FieldIn(FieldEndTime, vs...))
}

// EndTimeNotIn applies the NotIn predicate on the "end_time" field.
func EndTimeNotIn(vs ...time.Time) predicate.Record {
	return predicate.Record(sql.FieldNotIn(FieldEndTime, vs...))
}

// EndTimeGT applies the GT predicate on the "end_time" field.
func EndTimeGT(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldGT(FieldEndTime, v))
}

// EndTimeGTE applies the GTE predicate on the "end_time" field.
func EndTimeGTE(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldGTE(FieldEndTime, v))
}

// EndTimeLT applies the LT predicate on the "end_time" field.
func EndTimeLT(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldLT(FieldEndTime, v))
}

// EndTimeLTE applies the LTE predicate on the "end_time" field.
func EndTimeLTE(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldLTE(FieldEndTime, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Record {
	return predicate.Record(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Record {
	return predicate.Record(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Record {
	return predicate.Record(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Record {
	return predicate.Record(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Record) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Record) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Record) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		p(s.Not())
	})
}