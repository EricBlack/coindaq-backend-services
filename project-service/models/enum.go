package models

//项目状态
const (
	NotStart  int = 0
	Started int = 1
	Completed int = 2
)

//成员类型
const (
	Admin int = 1
	Partner int =2
	Investor int =3
)

//投票类型
const (
	NotVote int = 0
	Approve int = 1
	DisApprove int = 2
	Abstention int = 3
)

//媒体类型
const (
	ImageType int = 1
	VideoType int = 2
)

//公告类型
const (
	NewsType int = 1
	VoteType int = 2
)

//公告审核类型
const (
	AuditSubmit int = 0
	AuditPass int = 1
	AuditReject int = 2
)

//锁仓类型
const (
	VoteLock int = 1
	AdminLock int = 2
)

//锁仓状态
const (
	Locked int = 1
	Unlocked int = 2
)

//收款类型
const (
	BankType int = 1
	WeChatType int = 2
)

//是否删除
const (
	TrueValue int = 1
	FalseValue int = 2
)

//订单状态
const (
	OrderCreated int = 0 //创建阶段
	OrderTimedOut int = 1 //未付款过期阶段
	OrderPayed int = 2 //已付款阶段
	OrderCoined int = 3 //已打币阶段
	OrderCanceled int = 4 //取消阶段
	OrderAppealed int = 5 //申诉阶段
	OrderCompleted int = 6 //完成
)

//申诉方
const (
	Initiator int = 1 //发起方
	Acceptor int = 2 //接受方
)

//交易类型
const (
	FiatDeal int = 1 //法币交易
	CoinDeail int = 2 //币币交易
)
