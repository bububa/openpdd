package goods

// Opt 商品标签
type Opt struct {
	// ID 商品标签ID
	ID uint64 `json:"opt_id,omitempty"`
	// Name 商品标签名
	Name string `json:"opt_name,omitempty"`
	// ParentID id所属父ID，其中，parent_id=0时为顶级节点
	ParentID uint64 `json:"parent_opt_id,omitempty"`
	// Level 层级，1-一级，2-二级，3-三级，4-四级
	Level int `json:"level,omitempty"`
}
