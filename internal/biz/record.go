package biz

import (
	"context"
	"math"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"slacker/internal/pkg/util/errutil"
)

type Record struct {
	ID         string
	Type       string
	CreatorID  string
	BeginTime  time.Time
	UpdateTime *time.Time
	EndTime    *time.Time
}

type RecordRepo interface {
	SaveRecord(ctx context.Context, record *Record) (*Record, error)
	UpdateRecord(ctx context.Context, record *Record) (*Record, error)
	GetRecord(ctx context.Context, recordID string) (*Record, error)
}

type RecordUseCase struct {
	logger *log.Helper
	repo   RecordRepo
}

func NewRecordUseCase(l log.Logger, repo RecordRepo) *RecordUseCase {
	return &RecordUseCase{
		logger: log.NewHelper(l),
		repo:   repo,
	}
}

func (uc *RecordUseCase) BeginRecord(ctx context.Context, creatorID, recordType string) (*Record, error) {
	uc.logger.WithContext(ctx).Infof("start to begin record, creatorID: %s, recordType: %s", creatorID, recordType)
	record := &Record{
		Type:      recordType,
		CreatorID: creatorID,
		BeginTime: time.Now(),
	}
	return uc.repo.SaveRecord(ctx, record)
}

func (uc *RecordUseCase) EndRecord(ctx context.Context, recordID string, duration time.Duration) (*Record, error) {
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
	return uc.repo.UpdateRecord(ctx, record)
}
