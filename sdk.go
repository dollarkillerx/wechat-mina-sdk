package wechat_mina_sdk

type WechatMinaSDK struct {
	appid  string
	secret string
}

func New(appid string, secret string) *WechatMinaSDK {
	return &WechatMinaSDK{
		appid:  appid,
		secret: secret,
	}
}
