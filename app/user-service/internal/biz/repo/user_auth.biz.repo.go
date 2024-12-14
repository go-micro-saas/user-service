package bizrepos

import (
	"context"
	"github.com/go-micro-saas/user-service/app/user-service/internal/biz/bo"
	"github.com/go-micro-saas/user-service/app/user-service/internal/data/po"
)

type UserAuthBizRepo interface {
	GenSignTokenRequestByUserModel(ctx context.Context, userModel *po.User) (*bo.SignTokenReq, error)
	GenSignTokenRequestByAdminUserModel(ctx context.Context, userModel *po.User) (*bo.SignTokenReq, error)
	SignToken(ctx context.Context, req *bo.SignTokenReq) (*bo.SignTokenResp, error)
	RefreshToken(ctx context.Context, refreshToken string) (*bo.SignTokenResp, error)

	CheckAndGetByRegisterEmail(ctx context.Context, email string) (*po.UserRegEmail, error)
	CheckAndGetUserByUserId(ctx context.Context, userId int64) (*po.User, error)
	ValidateLoginUser(userModel *po.User, plaintextPassword string) error
	ComparePassword(hashPassword, plaintextPassword string) error
}
