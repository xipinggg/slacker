package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Record struct {
	ID         string
	Type       string
	CreatorID  string
	BeginTime  time.Time
	UpdateTime time.Time
	EndTime    time.Time
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
	record := &Record{
		Type:      recordType,
		CreatorID: creatorID,
	}
	return uc.repo.SaveRecord(ctx, record)
}

func (uc *RecordUseCase) EndRecord(ctx context.Context, recordID string) (*Record, error) {
	record, err := uc.repo.GetRecord(ctx, recordID)
	if err != nil {
		return nil, err
	}

	record.EndTime = time.Now()
	return uc.repo.UpdateRecord(ctx, record)
}
