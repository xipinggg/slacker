package record

import (
	"context"
	"math"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"

	"slacker/internal/pkg/util/errutil"
)

type GetGetTotalRecordTimeFilter struct {
	BeginTime *time.Time // 开始时间，和开始打卡时间比较，为 nil 时不过滤
	EndTime   *time.Time // 截止时间，和结束打卡时间比较，为 nil 时不过滤
}

type GetRecordByCreatorIDsFilter GetGetTotalRecordTimeFilter

type Repo interface {
	SaveRecord(ctx context.Context, record *Record) (*Record, error)
	UpdateRecord(ctx context.Context, record *Record) (*Record, error)
	GetRecord(ctx context.Context, recordID string) (*Record, error)
	GetRecordNotEnd(ctx context.Context, recordID string) ([]*Record, error)
	GetRecordByCreatorIDs(ctx context.Context, recordIDs []string, filter GetRecordByCreatorIDsFilter) (map[string][]*Record, error)
}

type UseCase struct {
	logger *log.Helper
	repo   Repo
}

func NewUseCase(l log.Logger, repo Repo) *UseCase {
	return &UseCase{
		logger: log.NewHelper(l),
		repo:   repo,
	}
}

func (uc *UseCase) BeginRecord(ctx context.Context, creatorID, recordType string) (*Root, error) {
	uc.logger.WithContext(ctx).Infof("start to begin record, creatorID: %s, recordType: %s", creatorID, recordType)
	record := &Record{
		Type:      recordType,
		CreatorID: creatorID,
		BeginTime: time.Now(),
	}

	record, err := uc.repo.SaveRecord(ctx, record)
	if err != nil {
		return nil, err
	}

	return &Root{
		record: record,
	}, nil
}

func (uc *UseCase) EndRecord(ctx context.Context, recordID string, duration time.Duration) (*Root, error) {
	uc.logger.WithContext(ctx).Infof("start to end record, recordID: %s", recordID)
	record, err := uc.repo.GetRecord(ctx, recordID)
	if err != nil {
		return nil, err
	}

	// 检查是否已结束
	if record.EndTime != nil {
		return nil, errutil.WithStack(errors.BadRequest("INVALID_END_RECORD", "record already end"))
	}

	// 检查结束时间是否在误差内
	endTime := record.BeginTime.Add(duration)
	record.EndTime = &endTime
	now := time.Now().Unix()
	if int(math.Abs(float64(record.EndTime.Unix()-now))) > 10 {
		return nil, errutil.Wrap(errors.BadRequest("INVALID_END_TIME", "end time is invalid"),
			"begin time: %v, duration: %v", record.BeginTime, duration)
	}

	record, err = uc.repo.UpdateRecord(ctx, record)
	if err != nil {
		return nil, err
	}

	return &Root{
		record: record,
	}, nil
}

func (uc *UseCase) GetRecordNotEnd(ctx context.Context, creatorID string) ([]*Record, error) {
	records, err := uc.repo.GetRecordNotEnd(ctx, creatorID)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (uc *UseCase) GetTotalRecordTime(ctx context.Context, creatorIDs []string, filter GetGetTotalRecordTimeFilter) (map[string]time.Duration, error) {
	idToRecords, err := uc.repo.GetRecordByCreatorIDs(ctx, creatorIDs, GetRecordByCreatorIDsFilter(filter))
	if err != nil {
		return nil, err
	}

	idToTime := make(map[string]time.Duration, len(idToRecords))
	for id, records := range idToRecords {
		for _, rec := range records {
			idToTime[id] += rec.RecordTime()
		}
	}

	return idToTime, nil
}
