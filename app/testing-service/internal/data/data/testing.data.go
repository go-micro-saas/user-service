package data

import (
	"github.com/go-kratos/kratos/log"
	datarepos "github.com/go-micro-saas/service-layout/app/testing-service/internal/data/repo"
)

type testingData struct {
	log *log.Helper
}

func NewTestingData(
	logger log.Logger,
) datarepos.TestingDataRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/data/data"))

	return &testingData{
		log: logHelper,
	}
}
