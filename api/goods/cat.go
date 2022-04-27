package goods

// Cat 类目
type Cat struct {
	// ID 商品类目ID
	ID uint64 `json:"cat_id,omitempty"`
	// Name 商品类目名称
	Name string `json:"cat_name,omitempty"`
	// Level 类目层级，1-一级类目，2-二级类目，3-三级类目，4-四级类目
	Level int `json:"level,omitempty"`
	// ParentID id所属父类目ID，其中，parent_id=0时为顶级节点
	ParentID uint64 `json:"parent_cat_id,omitempty"`
}
