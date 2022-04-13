package ddk

// PromURL 链接
type PromURL struct {
	// MobileShortURL 推广短链接（可唤起拼多多app）
	MobileShortURL string `json:"mobile_short_url,omitempty" xml:"mobile_short_url,omitempty"`
	// MobileURL 推广长链接（唤起拼多多app）
	MobileURL string `json:"mobile_url,omitempty" xml:"mobile_url,omitempty"`
	// MultiGroupMobileShortURL 推广短链接（唤起拼多多app）
	MultiGroupMobileShortURL string `json:"multi_group_mobile_short_url,omitempty" xml:"multi_group_mobile_short_url,omitempty"`
	// MultiGroupMobileURL 推广长链接（可唤起拼多多app）
	MultiGroupMobileURL string `json:"multi_group_mobile_url,omitempty" xml:"multi_group_mobile_url,omitempty"`
	// MultiGroupShortURL 双人团推广短链接
	MultiGroupShortURL string `json:"multi_group_short_url,omitempty" xml:"multi_group_short_url,omitempty"`
	// MultiGroupURL 双人团推广长链接
	MultiGroupURL string `json:"multi_group_url,omitempty" xml:"multi_group_url,omitempty"`
	// ShortURL 对应出参url的短链接，与url功能一致。
	ShortURL string `json:"short_url,omitempty" xml:"short_url,omitempty"`
	// URL 普通推广长链接，唤起H5页面
	URL string `json:"url,omitempty" xml:"url,omitempty"`
	// Sign CPSsign
	Sign string `json:"sign,omitempty" xml:"sign,omitempty"`
	// SchemeURL schema链接，用户安装拼多多APP的情况下会唤起APP（需客户端支持schema跳转协议）
	SchemeURL string `json:"scheme_url,omitempty" xml:"scheme_url,omitempty"`
	// TzSchemeURL 使用此推广链接，用户安装多多团长APP的情况下会唤起APP（需客户端支持schema跳转协议）
	TzSchemeURL string `json:"tz_scheme_url,omitempty" xml:"tz_shceme_url,omitempty"`
	// MultiURLLIst 双人团链接列表
	MultiURLList *PromURL `json:"multi_url_list,omitempty" xml:"multi_url_list,omitempty"`
	// SingleURLList 单人团链接列表
	SingleURLList *PromURL `json:"single_url_list,omitempty" xml:"single_url_list,omitempty"`
	// WeAppInfo 拼多多福利券微信小程序信息
	WeAppInfo *WeAppInfo `json:"we_app_info,omitempty" xml:"we_app_info,omitempty"`
	// QQAppInfo qq小程序信息
	QQAppInfo *WeAppInfo `json:"qq_app_info,omitempty" xml:"qq_app_info,omitempty"`
}

// WeAppInfo 拼多多福利券微信小程序信息
type WeAppInfo struct {
	// AppID 小程序id
	AppID string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// WeAppIconURL 小程序图片
	WeAppIconURl string `json:"we_app_icon_url,omitempty" xml:"we_app_icon_url,omitempty"`
	// BannerURL Banner图
	BannerURL string `json:"banner_url,omitempty" xml:"banner_url,omitempty"`
	// Desc 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// SourceDisplayName 来源名
	SourceDisplayName string `json:"source_display_name,omitempty" xml:"source_display_name,omitempty"`
	// PagePath 小程序path值
	PagePath string `json:"page_path,omitempty" xml:"page_path,omitempty"`
	// UserName 用户名
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// Title 小程序标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
}
