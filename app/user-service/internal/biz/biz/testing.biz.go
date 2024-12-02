package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-micro-saas/user-service/app/user-service/internal/biz/bo"
	bizrepos "github.com/go-micro-saas/user-service/app/user-service/internal/biz/repo"
	"github.com/go-micro-saas/user-service/app/user-service/internal/data/po"
	datarepos "github.com/go-micro-saas/user-service/app/user-service/internal/data/repo"
	"time"
)

type testingBiz struct {
	log         *log.Helper
	testingData datarepos.TestingDataRepo
}

func NewTestingBiz(
	logger log.Logger,
	testingData datarepos.TestingDataRepo,
) bizrepos.TestingBizRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/biz/biz"))

	return &testingBiz{
		log:         logHelper,
		testingData: testingData,
	}
}

func (s *testingBiz) HelloWorld(ctx context.Context, param *bo.HelloWorldParam) (*bo.HelloWorldReply, error) {
	dataModel := s.tooHelloWorkModel(param)
	err := s.testingData.HelloWorld(ctx, dataModel)
	if err != nil {
		return nil, err
	}
	return &bo.HelloWorldReply{Message: dataModel.RequestMessage}, nil
}

func (s *testingBiz) tooHelloWorkModel(param *bo.HelloWorldParam) *po.HelloWorld {
	res := &po.HelloWorld{
		Id:             0,
		CreatedTime:    time.Now(),
		UpdatedTime:    time.Now(),
		RequestMessage: param.Message,
	}
	return res
}
