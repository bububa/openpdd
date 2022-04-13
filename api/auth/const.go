package auth

const (
	// SHOP_AUTH_WEB_URL 拼多多店铺WEB端网页授权
	SHOP_AUTH_WEB_URL = "https://fuwu.pinduoduo.com/service-market/auth"
	// SHOP_AUTH_H5_URL 拼多多店铺H5移动端网页授权
	SHOP_AUTH_H5_URL = "https://mai.pinduoduo.com/h5-login.html"
	// DDK_AUTH_URL 多多进宝推手WEB端网页授权
	DDK_AUTH_URL = "https://jinbao.pinduoduo.com/open.html"
	// KTT_AUTH_URL 快团团团长WEB端网页授权
	KTT_AUTH_URL = "https://oauth.pinduoduo.com/authorize/ktt"
	// LOGISTIC_AUTH_URL 拼多多电子面单用户WEB端网页授权
	LOGISTIC_AUTH_URL = "https://wb.pinduoduo.com/logistics/auth"
)

// AuthType 授权类型
type AuthType int

const (
	// AuthType_SHOP_WEB 拼多多店铺WEB端
	AuthType_SHOP_WEB AuthType = iota
	// AuthType_SHOP_H5 拼多多店铺H5移动端网页授权
	AuthType_SHOP_H5
	// AuthType_DDK 多多进宝推手WEB端网页授权
	AuthType_DDK
	// AuthType_KTT 快团团团长WEB端网页授权
	AuthType_KTT
	// AuthType_LOGISTIC 拼多多电子面单用户WEB端网页授权
	AuthType_LOGISTIC
)

const (
	// H5 view h5
	H5 string = "h5"
	// WEB view web
	WEB string = "web"
)
