package enum

/**
自动跳转规则
*/
type AutoJumpRole uint8

const (
	/**
	处理人就是提交人
	*/
	DealerIsDealer AutoJumpRole = iota
	/**
	处理人已经出现过
	*/
	DealerIsInWorkerList
	/**
	处理人与上一步相同
	*/
	DealerAsNextStepWorker
)
