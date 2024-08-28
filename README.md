# 拼多多 OpenAPI golang版

[![Go Reference](https://pkg.go.dev/badge/github.com/bububa/openpdd.svg)](https://pkg.go.dev/github.com/bububa/openpdd)
[![Go](https://github.com/bububa/openpdd/actions/workflows/go.yml/badge.svg)](https://github.com/bububa/openpdd/actions/workflows/go.yml)
[![goreleaser](https://github.com/bububa/openpdd/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/bububa/openpdd/actions/workflows/goreleaser.yml)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/bububa/openpdd.svg)](https://github.com/bububa/openpdd)
[![GoReportCard](https://goreportcard.com/badge/github.com/bububa/openpdd)](https://goreportcard.com/report/github.com/bububa/openpdd)
[![GitHub license](https://img.shields.io/github/license/bububa/openpdd.svg)](https://github.com/bububa/openpdd/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/bububa/openpdd.svg)](https://GitHub.com/bububa/openpdd/releases/)

## Implemented API

- 授权(api/oauth)
  - 生成授权链接 [ URL(ctx context.Context, clt *core.SDKClient, req *URLRequest) string ]
  - 获取Access Token [ TokenCreate(ctx context.Context, clt *core.SDKClient, code string) (*Token, error) ]
  - 刷新Access Token [ TokenRefresh(ctx context.Context, clt *core.SDKClient, refreshToken string) (*Token, error) ]
- 多多客(api/ddk)
  - 创建多多进宝推广位 [ GoodsPidGenerate(ctx context.Context, clt *core.SDKClient, req *GoodsPidGenerateRequest) ([]Pid, error) ]
  - 多多客生成单品推广小程序二维码url [ WeappQrcodeUrlGen(ctx context.Context, clt *core.SDKClient, req *WeappQrcodeUrlGenRequest) (string, error) ]
  - 多多礼金状态更新 [ CashgiftStatusUpdate(ctx context.Context, clt *core.SDKClient, req *CashgiftStatusUpdateRequest) (uint64, error) ]
  - 查询已经生成的推广位信息 [ GoodsPidQuery(ctx context.Context, clt *core.SDKClient, req *GoodsPidQueryRequest) (int, []Pid, error) ]
  - 批量绑定推广位的媒体id [ PidMediaIDBind(ctx context.Context, clt *core.SDKClient, mediaID uint64, pidList []string) (*PidMediaIDBindResult, error) ]
  - 多多进宝转链接口 [ GoodsZsUnitUrlGen(ctx context.Context, clt *core.SDKClient, req *GoodsZsUnitUrlGenRequest) (*PromURL, error) ]
  - 查询是否绑定备案 [ MemberAuthorityQuery(ctx context.Context, clt *core.SDKClient, req *MemberAuthorityQueryRequest) (bool, error) ]
  - 多多进宝数据统计查询接口 [ StatisticDataQuery(ctx context.Context, clt *core.SDKClient, req *StatisticDataQueryRequest) ([]StatisticData, error) ]
  - 生成商城-频道推广链接 [ CmsPromUrlGenerate(ctx context.Context, clt *core.SDKClient, req *CmsPromUrlGenerateRequest) (int, []PromURL, error) ]
  - 多多进宝推广链接生成 [ GoodsPromotionUrlGenerate(ctx context.Context, clt *core.SDKClient, req *GoodsPromotionUrlGenerateRequest) ([]PromURL, error) ]
  - 用时间段查询推广订单接口 [ OrderListRangeGet(ctx context.Context, clt *core.SDKClient, req *OrderListRangeGetRequest) (string, []Order, error) ]
  - 查询订单详情 [ OrderDetailGet(ctx context.Context, clt *core.SDKClient, req *OrderDetailGetRequest) (*Order, error) ]
  - 多多进宝商品推荐API [ GoodsRecommendGet(ctx context.Context, clt *core.SDKClient, req *GoodsRecommendGetRequest) (*GoodsRecommendGetResult, error) ]
  - 生成营销工具推广链接 [ RpPromUrlGenerate(ctx context.Context, clt *core.SDKClient, req *RpPromUrlGenerateRequest) (*RpPromUrlGenerateResult, error) ]
  - 最后更新时间段增量同步推广订单信息 [ OrderListIncrementGet(ctx context.Context, clt *core.SDKClient, req *OrderListIncrementGetRequest) (int, []Order, error) ]
  - 多多进宝商品查询 [ GoodsSearch(ctx context.Context, clt *core.SDKClient, req *GoodsSearchRequest) (*GoodsSearchResult, error) ]
  - 多多进宝商品详情查询 [ GoodsDetail(ctx context.Context, clt *core.SDKClient, req *GoodsDetailRequest) ([]Goods, error) ]
  - 生成多多进宝频道推广 [ ResourceUrlGen(ctx context.Context, clt *core.SDKClient, req *ResourceUrlGenRequest) (*PromURL, error) ]
  - 创建多多礼金 [ CashgiftCreate(ctx context.Context, clt *core.SDKClient, req *CashgiftCreateRequest) (*CashgiftCreateResult, error) ]
  - 多多进宝信息流渠道备案授权素材上传接口 [ GoodsPromotionRightAuth(ctx context.Context, clt *core.SDKClient, req *GoodsPromotionRightAuthRequest) error ]
  - 多多客信息流投放备案图片上传接口 [ ReportImgUpload(ctx context.Context, clt *core.SDKClient, req *ReportImgUploadRequest) (string, error) ]
  - 多多客信息流投放备案视频上传接口 [ ReportVideoUpload(ctx context.Context, clt *core.SDKClient, req *ReportVideoUploadRequest) (string, error) ]
  - 多多客信息流投放备案视频上传分片初始化接口 [ ReportVideoUploadPartInit(ctx context.Context, clt *core.SDKClient, contentType string) (string, error) ]
  - 多多客信息流投放备案视频上传分片上传接口 [ ReportVideoUploadPart(ctx context.Context, clt *core.SDKClient, req *ReportVideoUploadPartRequest) (string, error) ]
  - 多多客信息流投放备案视频上传分片完成接口 [ ReportVideoUploadPartComplete(ctx context.Context, clt *core.SDKClient, uploadSign string) (string, error) ]
- 多多客工具(api/ddk/oauth)
  - 查询所有授权的多多客订单 [ AllOrderListIncrementGet(ctx context.Context, clt *core.SDKClient, req *AllOrderListIncrementGetRequest, accessToken string) (int, []ddk.Order, error) ]
  - 创建多多礼金 [ CashgiftCreate(ctx context.Context, clt *core.SDKClient, req *CashgiftCreateRequest, accessToken string) (*ddk.CashgiftCreateResult, error) ]
  - 多多礼金状态更新 [ CashgiftStatusUpdate(ctx context.Context, clt *core.SDKClient, req *CashgiftStatusUpdateRequest, accessToken string) (uint64, error) ]
  - 生成商城推广链接接口 [ CmsPromUrlGenerate(ctx context.Context, clt *core.SDKClient, req *CmsPromUrlGenerateRequest, accessToken string) (int, []ddk.PromURL, error) ]
  - 多多进宝商品详情查询 [ GoodsDetail(ctx context.Context, clt *core.SDKClient, req *GoodsDetailRequest, accessToken string) ([]ddk.Goods, error) ]
  - 多多进宝推广位创建接口 [ GoodsPidGenerate(ctx context.Context, clt *core.SDKClient, req *GoodsPidGenerateRequest, accessToken string) ([]Pid, error) ]
  - 多多客已生成推广位信息查询 [ GoodsPidQuery(ctx context.Context, clt *core.SDKClient, req *GoodsPidQueryRequest, accessToken string) (int, []ddk.Pid, error) ]
  - 生成多多进宝推广链接 [ GoodsPromotionUrlGenerate(ctx context.Context, clt *core.SDKClient, req *GoodsPromotionUrlGenerateRequest, accessToken string) ([]ddk.PromURL, error) ]
  - 运营频道商品查询API [ GoodsRecommendGet(ctx context.Context, clt *core.SDKClient, req *GoodsRecommendGetRequest, accessToken string) (*ddk.GoodsRecommendGetResult, error) ]
  - 多多进宝商品查询 [ GoodsSearch(ctx context.Context, clt *core.SDKClient, req *GoodsSearchRequest, accessToken string) (*ddk.GoodsSearchResult, error) ]
  - 多多进宝转链接口 [ GoodsZsUnitUrlGen(ctx context.Context, clt *core.SDKClient, req *GoodsZsUnitUrlGenRequest, accessToken string) (*ddk.PromURL, error) ]
  - 查询是否绑定备案 [ MemberAuthorityQuery(ctx context.Context, clt *core.SDKClient, req *MemberAuthorityQueryRequest, accessToken string) (bool, error) ]
  - 获取订单详情 [ OrderDetailGet(ctx context.Context, clt *core.SDKClient, req *OrderDetailGetRequest, accessToken string) (*ddk.Order, error) ]
  - 批量绑定推广位的媒体id [ PidMediaIDBind(ctx context.Context, clt *core.SDKClient, mediaID uint64, pidList []string, accessToken string) (*ddk.PidMediaIDBindResult, error) ]
  - 拼多多主站频道推广接口 [ ResourceUrlGen(ctx context.Context, clt *core.SDKClient, req *ResourceUrlGenRequest, accessToken string) (*ddk.PromURL, error) ]
  - 生成营销工具推广链接 [ RpPromUrlGenerate(ctx context.Context, clt *core.SDKClient, req *RpPromUrlGenerateRequest, accessToken string) (*ddk.RpPromUrlGenerateResult, error) ]
  - 多多客生成单品推广小程序二维码url [ WeappQrcodeUrlGen(ctx context.Context, clt *core.SDKClient, req *WeappQrcodeUrlGenRequest, accessToken string) (string, error) ]
- 商品API (api/goods)
  - 商品标准类目接口 [ CatsGet(ctx context.Context, clt \*core.SDKClient, parentID uint64) ([]Cat, error) ]
  - 查询商品标签列表 [ OptGet(ctx context.Context, clt \*core.SDKClient, parentID uint64) ([]Opt, error) ]
- 消息服务API (api/pmc)
  - 消息队列积压数量查询 [ AccureQuery(ctx context.Context, clt *core.SDKClient) (int64, error) ]
  - 取消用户的消息服务 [ UserCancel(ctx context.Context, clt *core.SDKClient, ownerID string) (bool, error) ]
  - 获取用户已开通消息 [ UserGet(ctx context.Context, clt *core.SDKClient, ownerID string) (*User, error) ]
  - 为已授权的用户开通消息服务 [ UserPermit(ctx context.Context, clt \*core.SDKClient, topics []string, accessToken string) (bool, error) ]
  - 接收消息推送 [ Read(ctx context.Context, clt \*core.SDKClient) <-chan []byte ]
