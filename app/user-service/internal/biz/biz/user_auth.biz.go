package biz

import (
	"context"
	errorv1 "github.com/go-micro-saas/user-service/api/user-service/v1/errors"
	"github.com/go-micro-saas/user-service/app/user-service/internal/biz/bo"
	bizrepos "github.com/go-micro-saas/user-service/app/user-service/internal/biz/repo"
	"github.com/go-micro-saas/user-service/app/user-service/internal/data/po"
	datarepos "github.com/go-micro-saas/user-service/app/user-service/internal/data/repo"
	passwordutil "github.com/ikaiguang/go-srv-kit/kit/password"
	uuidpkg "github.com/ikaiguang/go-srv-kit/kit/uuid"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

// userAuthBiz ...
type userAuthBiz struct {
	authRepo              authpkg.AuthRepo
	userDataRepo          datarepos.UserDataRepo
	userBindEmailDataRepo datarepos.UserRegEmailDataRepo
}

// NewUserAuthBiz ...
func NewUserAuthBiz(
	authRepo authpkg.AuthRepo,
	userDataRepo datarepos.UserDataRepo,
	userRegEmailDataRepo datarepos.UserRegEmailDataRepo,
) bizrepos.UserAuthBizRepo {
	return &userAuthBiz{
		authRepo:              authRepo,
		userDataRepo:          userDataRepo,
		userBindEmailDataRepo: userRegEmailDataRepo,
	}
}

// CheckAndGetByRegisterEmail 邮箱是否存在
func (s *userAuthBiz) CheckAndGetByRegisterEmail(ctx context.Context, email string) (*po.UserRegEmail, error) {
	// 注册邮箱
	regEmailModel, isNotFound, err := s.userBindEmailDataRepo.QueryOneByUserEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if isNotFound {
		e := errorv1.ErrorS103UserNotExist("用户不存在")
		return nil, errorpkg.WithStack(e)
	}
	return regEmailModel, nil
}

// CheckAndGetUserByUserId 用户是否存在
func (s *userAuthBiz) CheckAndGetUserByUserId(ctx context.Context, userId int64) (*po.User, error) {
	userModel, isNotFound, err := s.userDataRepo.QueryOneByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	if isNotFound {
		e := errorv1.ErrorS103UserNotExist("用户不存在")
		return nil, errorpkg.WithStack(e)
	}
	return userModel, nil
}

// ValidateLoginUser 对比密码
func (s *userAuthBiz) ValidateLoginUser(userModel *po.User, plaintextPassword string) error {
	err := s.ComparePassword(userModel.PasswordHash, plaintextPassword)
	if err != nil {
		return err
	}
	if !userModel.IsValidStatus() {
		e := errorv1.ErrorS103UserStatusNotAllow("无效的登录状态")
		return errorpkg.WithStack(e)
	}
	return nil
}

// ComparePassword 对比密码
// plaintextPassword plaintext || md5
func (s *userAuthBiz) ComparePassword(hashPassword, plaintextPassword string) error {
	// 验证密码
	err := passwordutil.Compare(hashPassword, plaintextPassword)
	if err != nil {
		e := errorv1.ErrorS103UserPasswordIncorrect("密码不正确")
		return errorpkg.Wrap(e, err)
	}
	return nil
}

// GenSignTokenRequestByUserModel auth claims
func (s *userAuthBiz) GenSignTokenRequestByUserModel(ctx context.Context, userModel *po.User) (*bo.SignTokenReq, error) {
	payload := &authpkg.Payload{
		TokenID:       uuidpkg.New(),
		UserID:        uint64(userModel.UserId),
		UserUuid:      "",
		LoginPlatform: authpkg.LoginPlatformEnum_UNSPECIFIED,
		LoginType:     authpkg.LoginTypeEnum_UNSPECIFIED,
		LoginLimit:    authpkg.LoginLimitEnum_UNLIMITED,
		TokenType:     authpkg.TokenTypeEnum_USER,
	}
	res := &bo.SignTokenReq{
		Claims: *authpkg.GenAuthClaimsByAuthPayload(payload, authpkg.AccessTokenExpire),
	}
	return res, nil
}

// GenSignTokenRequestByAdminUserModel auth claims
func (s *userAuthBiz) GenSignTokenRequestByAdminUserModel(ctx context.Context, userModel *po.User) (*bo.SignTokenReq, error) {
	payload := &authpkg.Payload{
		TokenID:       uuidpkg.New(),
		UserID:        uint64(userModel.UserId),
		UserUuid:      "",
		LoginPlatform: authpkg.LoginPlatformEnum_UNSPECIFIED,
		LoginType:     authpkg.LoginTypeEnum_UNSPECIFIED,
		LoginLimit:    authpkg.LoginLimitEnum_UNLIMITED,
		TokenType:     authpkg.TokenTypeEnum_ADMIN,
	}
	res := &bo.SignTokenReq{
		Claims: *authpkg.GenAuthClaimsByAuthPayload(payload, authpkg.AccessTokenExpire),
	}
	return res, nil
}

// SignToken token
func (s *userAuthBiz) SignToken(ctx context.Context, req *bo.SignTokenReq) (*bo.SignTokenResp, error) {
	tokenResp, err := s.authRepo.SignToken(ctx, &req.Claims)
	if err != nil {
		return nil, err
	}

	res := &bo.SignTokenResp{}
	res.SetByAuthTokenResponse(tokenResp)
	return res, nil
}

// RefreshToken token
func (s *userAuthBiz) RefreshToken(ctx context.Context, refreshToken string) (*bo.SignTokenResp, error) {
	if refreshToken == "" {
		e := errorpkg.ErrorInvalidParameter("refresh token is empty")
		return nil, errorpkg.WithStack(e)
	}
	authClaims, err := s.authRepo.DecodeRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	if err = s.authRepo.VerifyRefreshToken(ctx, authClaims); err != nil {
		return nil, err
	}

	tokenResp, err := s.authRepo.RefreshToken(ctx, authClaims)
	if err != nil {
		return nil, err
	}

	res := &bo.SignTokenResp{}
	res.SetByAuthTokenResponse(tokenResp)
	return res, nil
}
