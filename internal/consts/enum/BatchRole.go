package enum

/**
节点工作批处理
*/
type BatchRole uint8

const (
	/**
	不可以
	*/
	None_BatchRole BatchRole = iota
	/**
	批量审批
	*/
	WorkCheckModel
	/**
	分组批量审核
	*/
	Group_BatchRole
)
