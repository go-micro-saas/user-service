package dto

import (
	resourcev1 "github.com/go-micro-saas/user-service/api/user-service/v1/resources"
	"github.com/go-micro-saas/user-service/app/user-service/internal/biz/bo"
	"github.com/go-micro-saas/user-service/app/user-service/internal/data/po"
	timepkg "github.com/ikaiguang/go-srv-kit/kit/time"
)

var (
	UserDto userDto
)

type userDto struct{}

func (s *userDto) ToPbLoginResp(userModel *po.User, tokenResp *bo.SignTokenResp) *resourcev1.LoginResp {
	res := &resourcev1.LoginResp{
		UserInfo: &resourcev1.UserInfo{
			// Id:           userModel.Id,
			UserId:       userModel.UserId,
			UserNickname: userModel.UserNickname,
			UserAvatar:   userModel.UserAvatar,
			// UserGender:   userModel.UserGender,
			// UserStatus:   userModel.UserStatus,
		},
		AccessToken:           tokenResp.AccessToken,
		AccessTokenExpiredAt:  tokenResp.AccessTokenItem.ExpiredAt,
		RefreshToken:          tokenResp.RefreshToken,
		RefreshTokenExpiredAt: tokenResp.RefreshTokenItem.ExpiredAt,
	}
	return res
}

func (s *userDto) ToPbUser(dataModel *po.User) *resourcev1.User {
	newDataModel := &resourcev1.User{
		Id:          dataModel.Id,                                 // ID
		UserId:      dataModel.UserId,                             // uid
		CreatedTime: dataModel.CreatedTime.Format(timepkg.YmdHms), // 创建时间
		UpdatedTime: dataModel.UpdatedTime.Format(timepkg.YmdHms), // 最后修改时间
		// DeletedTime:  dataModel.DeletedTime,                        // 删除时间
		UserEmail:    dataModel.UserEmail,    // 邮箱
		UserNickname: dataModel.UserNickname, // 昵称
		UserAvatar:   dataModel.UserAvatar,   // 头像
		UserGender:   dataModel.UserGender,   // 性别
		UserStatus:   dataModel.UserStatus,   // 状态
		//RegisterType: dataModel.RegisterType, // 注册类型
		//PasswordHash: dataModel.PasswordHash, // 密码
	}
	// ActiveBeginTime 激活开始时间
	//if dataModel.ActiveBeginTime != nil {
	//	newDataModel.ActiveBeginTime = dataModel.ActiveBeginTime.Format(timepkg.YmdHms)
	//}
	// ActiveEndTime 激活结束时间
	//if dataModel.ActiveEndTime != nil {
	//	newDataModel.ActiveEndTime = dataModel.ActiveEndTime.Format(timepkg.YmdHms)
	//}
	// DisableTime 禁用时间
	//if dataModel.DisableTime != nil {
	//	newDataModel.DisableTime = dataModel.DisableTime.Format(timepkg.YmdHms)
	//}
	// BlacklistTime 黑名单时间
	//if dataModel.BlacklistTime != nil {
	//	newDataModel.BlacklistTime = dataModel.BlacklistTime.Format(timepkg.YmdHms)
	//}

	return newDataModel
}
