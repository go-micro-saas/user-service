package service

import (
	"github.com/go-kratos/kratos/log"
	servicev1 "github.com/go-micro-saas/service-layout/api/testdata-service/v1/services"
	bizrepos "github.com/go-micro-saas/service-layout/app/testing-service/internal/biz/repo"
)

type testingV1Service struct {
	servicev1.UnimplementedSrvTestdataServer

	log        *log.Helper
	testingBiz bizrepos.TestingBizRepo
}

func NewTestingV1Service(logger log.Logger, testingBiz bizrepos.TestingBizRepo) servicev1.SrvTestdataServer {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/service/service"))
	return &testingV1Service{
		log:        logHelper,
		testingBiz: testingBiz,
	}
}
