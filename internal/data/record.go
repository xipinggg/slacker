package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	recordbiz "slacker/internal/biz/record"
	"slacker/internal/data/ent"
	recordent "slacker/internal/data/ent/record"
	"slacker/internal/pkg/util/errutil"
)

func toRecordEntity(rc *ent.Record) *recordbiz.Record {
	return &recordbiz.Record{
		ID:         rc.ID,
		Type:       rc.Type,
		CreatorID:  rc.CreatorID,
		BeginTime:  rc.BeginTime,
		UpdateTime: rc.UpdatedAt,
		EndTime:    rc.EndTime,
	}
}

type RecordRepo struct {
	data *Data
	log  *log.Helper
}

func NewRecordRepo(data *Data, logger log.Logger) recordbiz.Repo {
	return &RecordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *RecordRepo) SaveRecord(ctx context.Context, record *recordbiz.Record) (*recordbiz.Record, error) {
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

func (r *RecordRepo) UpdateRecord(ctx context.Context, record *recordbiz.Record) (*recordbiz.Record, error) {
	rc, err := r.data.DBClient.Record.
		UpdateOneID(record.ID).
		SetEndTime(*record.EndTime).
		Save(ctx)
	if err != nil {
		return nil, errutil.Wrap(err, "update by id failed: %s", record.ID)
	}

	return toRecordEntity(rc), nil
}

func (r *RecordRepo) GetRecord(ctx context.Context, recordID string) (*recordbiz.Record, error) {
	rc, err := r.data.DBClient.Record.Get(ctx, recordID)
	if err != nil {
		return nil, errutil.Wrap(err, "get by id failed: %s", recordID)
	}

	return toRecordEntity(rc), nil
}

func (r *RecordRepo) GetRecordNotEnd(ctx context.Context, recordID string) ([]*recordbiz.Record, error) {
	records, err := r.data.DBClient.Record.Query().Where(recordent.ID(recordID), recordent.EndTimeIsNil()).All(ctx)
	if err != nil {
		return nil, errutil.Wrap(err, "get by id failed: %s", recordID)
	}

	result := make([]*recordbiz.Record, len(records))
	for _, rec := range records {
		result = append(result, toRecordEntity(rec))
	}
	return result, nil
}

func (r *RecordRepo) GetRecordByCreatorIDs(ctx context.Context, recordIDs []string, filter recordbiz.GetRecordByCreatorIDsFilter) (map[string][]*recordbiz.Record, error) {
	//TODO implement me
	panic("implement me")
}
