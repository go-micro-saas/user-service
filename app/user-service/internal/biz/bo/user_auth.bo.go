package bo

import authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"

type SignTokenReq struct {
	authpkg.Claims
}

func (s *SignTokenReq) SetByAuthClaims(authClaims *authpkg.Claims) {
	s.Claims = *authClaims
}

type SignTokenResp struct {
	authpkg.TokenResponse
}

func (s *SignTokenResp) SetByAuthTokenResponse(authResp *authpkg.TokenResponse) {
	s.TokenResponse = *authResp
}
