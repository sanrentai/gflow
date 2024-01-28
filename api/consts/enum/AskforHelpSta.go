package enum

/**
加签模式
*/
type AskforHelpSta uint8

const (
	/**
	加签后直接发送
	*/
	AfterDealSend AskforHelpSta = 5
	/**
	加签后由我直接发送
	*/
	AfterDealSendByWorker AskforHelpSta = 6
)
