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
  - 生成授权链接 [ URL(clt *core.SDKClient, req *URLRequest) string ]
  - 获取Access Token [ TokenCreate(clt *core.SDKClient, code string) (*Token, error) ]
  - 刷新Access Token [ TokenRefresh(clt *core.SDKClient, refreshToken string) (*Token, error) ]
- 多多客(api/ddk)
  - 创建多多进宝推广位 [ GoodsPidGenerate(clt *core.SDKClient, req *GoodsPidGenerateRequest, accessToken string) ([]Pid, error) ]
  - 多多礼金状态更新 [ CashgiftStatusUpdate(clt *core.SDKClient, req *CashgiftStatusUpdateRequest, accessToken string) (uint64, error) ]
  - 查询已经生成的推广位信息 [ GoodsPidQuery(clt *core.SDKClient, req *GoodsPidQueryRequest, accessToken string) (int, []Pid, error) ]
  - 批量绑定推广位的媒体id [ PidMediaIDBind(clt *core.SDKClient, mediaID uint64, pidList []string, accessToken string) (*PidMediaIDBindResult, error) ]
  - 多多进宝转链接口 [ GoodsZsUnitUrlGen(clt *core.SDKClient, req *GoodsZsUnitUrlGenRequest, accessToken string) (*PromURL, error) ]
  - 查询是否绑定备案 [ MemberAuthorityQuery(clt *core.SDKClient, req *MemberAuthorityQueryRequest, accessToken string) (bool, error) ]
  - 多多进宝数据统计查询接口 [ StatisticDataQuery(clt *core.SDKClient, req *StatisticDataQueryRequest, accessToken string) ([]StatisticData, error) ]
  - 生成商城-频道推广链接 [ CmsPromUrlGenerate(clt *core.SDKClient, req *CmsPromUrlGenerateRequest, accessToken string) (int, []PromURL, error) ]
  - 多多进宝推广链接生成 [ GoodsPromotionUrlGenerate(clt *core.SDKClient, req *GoodsPromotionUrlGenerateRequest, accessToken string) ([]PromURL, error) ]
  - 用时间段查询推广订单接口 [ OrderListRangeGet(clt *core.SDKClient, req *OrderListRangeGetRequest, accessToken string) (string, []Order, error) ]
  - 查询订单详情 [ OrderDetailGet(clt *core.SDKClient, req *OrderDetailGetRequest, accessToken string) (*Order, error) ]
  - 多多进宝商品推荐API [ GoodsRecommendGet(clt *core.SDKClient, req *GoodsRecommendGetRequest, accessToken string) (*GoodsRecommendGetResult, error) ]
  - 生成营销工具推广链接 [ RpPromUrlGenerate(clt *core.SDKClient, req *RpPromUrlGenerateRequest, accessToken string) (*RpPromUrlGenerateResult, error) ]
  - 最后更新时间段增量同步推广订单信息 [ OrderListIncrementGet(clt *core.SDKClient, req *OrderListIncrementGetRequest, accessToken string) (int, []Order, error) ]
  - 多多进宝商品查询 [ GoodsSearch(clt *core.SDKClient, req *GoodsSearchRequest, accessToken string) (*GoodsSearchResult, error) ]
  - 多多进宝商品详情查询 [ GoodsDetail(clt *core.SDKClient, req *GoodsDetailRequest, accessToken string) ([]Goods, error) ]
  - 生成多多进宝频道推广 [ ResourceUrlGen(clt *core.SDKClient, req *ResourceUrlGenRequest, accessToken string) (*PromURL, error) ]
  - 创建多多礼金 [ CashgiftCreate(clt *core.SDKClient, req *CashgiftCreateRequest, accessToken string) (*CashgiftCreateResult, error) ]
- 多多客工具(api/ddk/oauth)
  - 查询所有授权的多多客订单 [ AllOrderListIncrementGet(clt *core.SDKClient, req *AllOrderListIncrementGetRequest, accessToken string) (int, []ddk.Order, error) ]
  - 创建多多礼金 [ CashgiftCreate(clt *core.SDKClient, req *CashgiftCreateRequest, accessToken string) (*ddk.CashgiftCreateResult, error) ]
  - 多多礼金状态更新 [ CashgiftStatusUpdate(clt *core.SDKClient, req *CashgiftStatusUpdateRequest, accessToken string) (uint64, error) ]
  - 生成商城推广链接接口 [ CmsPromUrlGenerate(clt *core.SDKClient, req *CmsPromUrlGenerateRequest, accessToken string) (int, []ddk.PromURL, error) ]
  - 多多进宝商品详情查询 [ GoodsDetail(clt *core.SDKClient, req *GoodsDetailRequest, accessToken string) ([]ddk.Goods, error) ]
  - 多多进宝推广位创建接口 [ GoodsPidGenerate(clt *core.SDKClient, req *GoodsPidGenerateRequest, accessToken string) ([]Pid, error) ]
  - 多多客已生成推广位信息查询 [ GoodsPidQuery(clt *core.SDKClient, req *GoodsPidQueryRequest, accessToken string) (int, []ddk.Pid, error) ]
  - 生成多多进宝推广链接 [ GoodsPromotionUrlGenerate(clt *core.SDKClient, req *GoodsPromotionUrlGenerateRequest, accessToken string) ([]ddk.PromURL, error) ]
  - 运营频道商品查询API [ GoodsRecommendGet(clt *core.SDKClient, req *GoodsRecommendGetRequest, accessToken string) (*ddk.GoodsRecommendGetResult, error) ]
  - 多多进宝商品查询 [ GoodsSearch(clt *core.SDKClient, req *GoodsSearchRequest, accessToken string) (*ddk.GoodsSearchResult, error) ]
  - 多多进宝转链接口 [ GoodsZsUnitUrlGen(clt *core.SDKClient, req *GoodsZsUnitUrlGenRequest, accessToken string) (*ddk.PromURL, error) ]
  - 查询是否绑定备案 [ MemberAuthorityQuery(clt *core.SDKClient, req *MemberAuthorityQueryRequest, accessToken string) (bool, error) ]
  - 获取订单详情 [ OrderDetailGet(clt *core.SDKClient, req *OrderDetailGetRequest, accessToken string) (*ddk.Order, error) ]
  - 批量绑定推广位的媒体id [ PidMediaIDBind(clt *core.SDKClient, mediaID uint64, pidList []string, accessToken string) (*ddk.PidMediaIDBindResult, error) ]
  - 拼多多主站频道推广接口 [ ResourceUrlGen(clt *core.SDKClient, req *ResourceUrlGenRequest, accessToken string) (*ddk.PromURL, error) ]
  - 生成营销工具推广链接 [ RpPromUrlGenerate(clt *core.SDKClient, req *RpPromUrlGenerateRequest, accessToken string) (*ddk.RpPromUrlGenerateResult, error) ]

