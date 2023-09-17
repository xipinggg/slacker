package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "slacker/api/slacker/v1"
	"slacker/internal/biz"
	"slacker/internal/pkg/auth"
)

type RecordService struct {
	pb.UnimplementedRecordServer

	logger *log.Helper
	uc     *biz.RecordUseCase
}

func NewRecordService(l log.Logger, uc *biz.RecordUseCase) *RecordService {
	return &RecordService{
		logger: log.NewHelper(l),
		uc:     uc,
	}
}

func (s *RecordService) BeginRecord(ctx context.Context, req *pb.BeginRecordRequest) (*pb.BeginRecordReply, error) {
	user, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	record, err := s.uc.BeginRecord(ctx, user.ID, req.Type)
	if err != nil {
		return nil, err
	}

	return &pb.BeginRecordReply{
		Id:        record.ID,
		BeginTime: timestamppb.New(record.BeginTime),
	}, nil
}

func (s *RecordService) EndRecord(ctx context.Context, req *pb.EndRecordRequest) (*pb.EndRecordReply, error) {
	record, err := s.uc.EndRecord(ctx, req.Id, req.GetDuration().AsDuration())
	if err != nil {
		return nil, err
	}

	return &pb.EndRecordReply{
		Id:        record.ID,
		BeginTime: timestamppb.New(record.BeginTime),
		EndTime:   timestamppb.New(*record.EndTime),
	}, nil
}
