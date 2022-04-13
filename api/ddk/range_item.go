package ddk

// RangeItem 自定义红包抵后价和商品佣金区间对象
type RangeItem struct {
	// RangeFrom 区间的开始值
	RangeFrom int64 `json:"range_from,omitempty"`
	// RangeID range_id为1表示红包抵后价（单位分）， range_id为2表示佣金比例（单位千分之几)
	RangeID int `json:"range_id,omitempty"`
	// RangeTo 区间的结束值
	RangeTo int64 `json:"range_to,omitempty"`
}
