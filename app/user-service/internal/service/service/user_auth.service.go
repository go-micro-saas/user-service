package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	resourcev1 "github.com/go-micro-saas/user-service/api/user-service/v1/resources"
	servicev1 "github.com/go-micro-saas/user-service/api/user-service/v1/services"
	bizrepos "github.com/go-micro-saas/user-service/app/user-service/internal/biz/repo"
	"github.com/go-micro-saas/user-service/app/user-service/internal/service/dto"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
)

// userAuth ...
type userAuth struct {
	servicev1.UnimplementedSrvUserAuthV1Server

	log             *log.Helper
	userAuthBizRepo bizrepos.UserAuthBizRepo
}

// NewUserAuthService ...
func NewUserAuthService(
	logger log.Logger,
	userAuthBizRepo bizrepos.UserAuthBizRepo,
) servicev1.SrvUserAuthV1Server {
	return &userAuth{
		log:             log.NewHelper(log.With(logger, "module", "user-service/service/user_auth")),
		userAuthBizRepo: userAuthBizRepo,
	}
}

// LoginByEmail Email登录
func (s *userAuth) LoginByEmail(ctx context.Context, in *resourcev1.LoginByEmailReq) (*resourcev1.LoginResp, error) {
	userModel, signResp, err := s.userAuthBizRepo.LoginByEmail(ctx, in)
	if err != nil {
		return nil, err
	}

	out := dto.UserDto.ToPbLoginResp(userModel, signResp)
	return out, nil
}

// LoginByPhone 手机登录
func (s *userAuth) LoginByPhone(ctx context.Context, in *resourcev1.LoginByPhoneReq) (*resourcev1.LoginResp, error) {
	userModel, signResp, err := s.userAuthBizRepo.LoginByPhone(ctx, in)
	if err != nil {
		return nil, err
	}

	out := dto.UserDto.ToPbLoginResp(userModel, signResp)
	return out, nil
}

// RefreshToken 刷新token
func (s *userAuth) RefreshToken(ctx context.Context, in *resourcev1.RefreshTokenReq) (*resourcev1.LoginResp, error) {
	signResp, err := s.userAuthBizRepo.RefreshToken(ctx, in.RefreshToken)
	if err != nil {
		return nil, err
	}

	// user
	userId := signResp.TokenResponse.RefreshTokenItem.Payload.UserID
	userModel, err := s.userAuthBizRepo.CheckAndGetUserByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	out := dto.UserDto.ToPbLoginResp(userModel, signResp)
	return out, nil
}

// Ping ping pong
func (s *userAuth) Ping(ctx context.Context, in *resourcev1.PingReq) (out *resourcev1.PingResp, err error) {
	// 可以解析
	authClaims, tokenClaimsOK := authpkg.GetAuthClaimsFromContext(ctx)
	s.log.WithContext(ctx).Infow(
		"tokenClaimsOK", tokenClaimsOK,
		"authClaims", authClaims,
	)

	out = &resourcev1.PingResp{
		Message: in.Message,
	}
	return out, err
}
