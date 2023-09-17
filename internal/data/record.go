package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"slacker/internal/biz"
	"slacker/internal/data/ent"
	"slacker/internal/pkg/util/errutil"
)

type RecordRepo struct {
	data *Data
	log  *log.Helper
}

func NewRecordRepo(data *Data, logger log.Logger) biz.RecordRepo {
	return &RecordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *RecordRepo) SaveRecord(ctx context.Context, record *biz.Record) (*biz.Record, error) {
	rc, err := r.data.DBClient.Record.
		Create().
		SetCreatorID(record.CreatorID).
		SetType(record.Type).
		SetBeginTime(record.BeginTime).
		Save(ctx)
	if err != nil {
		return nil, errutil.WithStack(err)
	}

	return toRecordEntity(rc), nil
}

func (r *RecordRepo) UpdateRecord(ctx context.Context, record *biz.Record) (*biz.Record, error) {
	id, err := uuid.Parse(record.ID)
	if err != nil {
		return nil, errutil.Wrap(err, "parse id failed: %s", record.ID)
	}

	rc, err := r.data.DBClient.Record.
		UpdateOneID(id).
		SetEndTime(*record.EndTime).
		Save(ctx)
	if err != nil {
		return nil, errutil.Wrap(err, "update by id failed: %s", record.ID)
	}

	return toRecordEntity(rc), nil
}

func (r *RecordRepo) GetRecord(ctx context.Context, recordID string) (*biz.Record, error) {
	id, err := uuid.Parse(recordID)
	if err != nil {
		return nil, errutil.Wrap(err, "parse id failed: %s", recordID)
	}

	rc, err := r.data.DBClient.Record.Get(ctx, id)
	if err != nil {
		return nil, errutil.Wrap(err, "get by id failed: %s", recordID)
	}

	return toRecordEntity(rc), nil
}

func toRecordEntity(rc *ent.Record) *biz.Record {
	return &biz.Record{
		ID:         rc.ID.String(),
		Type:       rc.Type,
		CreatorID:  rc.CreatorID,
		BeginTime:  rc.BeginTime,
		UpdateTime: rc.UpdatedAt,
		EndTime:    rc.EndTime,
	}
}
