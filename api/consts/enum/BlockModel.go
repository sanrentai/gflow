package enum

/**
阻塞模式
*/
type BlockModel uint8

const (
	/**
	不阻塞
	*/
	None_BlockModel BlockModel = iota
	/**
	当前节点的有未完成的子线程
	*/
	CurrNodeAll
	/**
	按照约定的格式阻塞.
	*/
	SpecSubFlow
	/**
	按照配置的sql阻塞,返回大于等于1表示阻塞,否则不阻塞.
	*/
	BySQL_BlockModel
	/**
	按照表达式阻塞，表达式类似方向条件的表达式.
	*/
	ByExp
	/**
	为父流程时，指定的子流程未运行到指定节点，则阻塞
	*/
	SpecSubFlowNode
	/**
	为平级子流程时，指定的子流程未运行到指定节点，则阻塞
	*/
	SameLevelSubFlow
)
