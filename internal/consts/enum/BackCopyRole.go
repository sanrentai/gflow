package enum

/**
子流程数据反填父流程数据规则
*/
type BackCopyRole uint8

const (

	/**
	不反填
	*/
	None_BackCopyRole BackCopyRole = iota
	/**
	字段自动匹配
	*/
	AutoFieldMatch
	/**
	按照设置的格式匹配
	*/
	FollowSetFormat
	/**
	混合模式
	*/
	MixedMode
)
