package enum

/**
自动发起
*/
type AutoStart uint8

const (
	/**
	手工启动（默认）
	*/
	None_AutoStart AutoStart = iota
	/**
	按照指定的人员
	*/
	ByDesignee
	/**
	数据集按时启动
	*/
	ByTineData
	/**
	触发试启动
	*/
	ByTrigger
)
