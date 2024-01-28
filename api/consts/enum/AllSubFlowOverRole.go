package enum

/**
所有子流程结束，父流程处理规则
*/
type AllSubFlowOverRole uint8

const (
	/**
	无
	*/
	None_AllSubFlowOverRole AllSubFlowOverRole = iota
	/**
	发送父流程到下一个节点
	*/
	SendParentFlowToNextStep
	/**
	结束父流程
	*/
	OverParentFlow
)
