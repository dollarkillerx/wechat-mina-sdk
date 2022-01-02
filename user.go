package wechat_mina_sdk

import (
	"fmt"

	"github.com/dollarkillerx/urllib"
)

// UserLogin UserLogin `https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html`
func (w *WechatMinaSDK) UserLogin(jsCode string) (*Login, error) {
	var login Login
	err := urllib.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", w.appid, w.secret, jsCode)).
		FromJson(&login)
	if err != nil {
		return nil, err
	}
	return &login, nil
}

// GetPhoneNumber GetPhoneNumber `https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/phonenumber/phonenumber.getPhoneNumber.html#%E4%BA%91%E8%B0%83%E7%94%A8`
func (w *WechatMinaSDK) GetPhoneNumber(token string) (*PhoneNumber, error) {
	var phoneNumber PhoneNumber

	err := urllib.Post(fmt.Sprintf("https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s", token)).
		FromJson(&phoneNumber)
	if err != nil {
		return nil, err
	}

	return &phoneNumber, nil
}
