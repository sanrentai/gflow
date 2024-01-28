package enum

/**
抄送规则
*/
type CCRoleEnum uint8

const (
	/**
	不能抄送
	*/
	UnCC CCRoleEnum = iota
	/**
	手工抄送
	*/
	HandCC
	/**
	自动抄送
	*/
	AutoCC
	/**
	手工与自动并存
	*/
	HandAndAuto
	/**
	按字段
	*/
	BySysCCEmps
	/**
	在发送前打开
	*/
	WhenSend
)

/**
抄送状态
*/
type CCSta uint8

const (
	/**
	未读
	*/
	UnRead CCSta = iota
	/**
	已读
	*/
	Read
	/**
	已回复
	*/
	CheckOver
	/**
	已删除
	*/
	Del
)

/**
抄送到角色计算方式
*/
type CCStaWay uint8

const (
	/**
	仅按角色计算
	*/
	StationOnly CCStaWay = iota
	/**
	按角色智能计算(当前节点的人员角色)
	*/
	StationSmartCurrNodeWorker
	/**
	按角色智能计算(接受节点的人员角色)
	*/
	StationSmartNextNodeWorker
	/**
	按角色与部门的交集
	*/
	StationAndDept
	/**
	按直线部门找角色下的人员(当前节点)
	*/
	StationDeptUpLevelCurrNodeWorker
	/**
	按直线部门找角色下的人员
	*/
	StationDeptUpLevelNextNodeWorker
)

/**
抄送方式
*/
type CCWay uint8

const (
	/**
	按照信息发送
	*/
	ByMsg CCWay = iota
	/**
	按照e-mail
	*/
	ByEmail_CCWay
	/**
	按照电话
	*/
	ByPhone
	/**
	按照数据库功能
	*/
	ByDBFunc
)

/**
抄送数据写入规则
*/
type CCWriteTo uint8

const (
	/**
	抄送列表
	*/
	CCList CCWriteTo = iota
	/**
	待办列表
	*/
	Todolist
	/**
	抄送与待办列表
	*/
	All_CCWriteTo
)

/**
工作提醒规则
*/
type CHAlertRole uint8

const (
	/**
	不提醒
	*/
	None_CHAlertRole CHAlertRole = iota
	/**
	一天一次
	*/
	OneDayOneTime
	/**
	一天两次
	*/
	OneDayTowTime
)

/**
工作提醒方式
*/
type CHAlertWay uint8

const (
	/**
	邮件
	*/
	ByEmail_CHAlertWay CHAlertWay = iota
	/**
	短消息
	*/
	BySMS
	/**
	即时通讯
	*/
	ByCCIM
)

/**
考核规则
*/
type CHWay uint8

const (
	/**
	不考核
	*/
	None_CHWay CHWay = iota
	/**
	按照时效考核
	*/
	ByTime
	/**
	按照工作量考核
	*/
	ByWorkNum
	/**
	是否是考核质量点
	*/
	IsQuality
)

/**
撤销规则
*/
type CancelRole uint8

const (
	/**
	仅上一步
	*/
	OnlyNextStep CancelRole = iota
	/**
	不能撤销
	*/
	None_CancelRole
	/**
	上一步与开始节点.
	*/
	NextStepAndStartNode
	/**
	可以撤销指定的节点
	*/
	SpecNodes
)

/**
删除流程规则
@0=不能删除
@1=逻辑删除
@2=记录日志方式删除: 数据删除后，记录到WF_DeleteWorkFlow中。
@3=彻底删除：
@4=让用户决定删除方式
*/
type DelWorkFlowRole uint8

const (
	/**
	不能删除
	*/
	None_DelWorkFlowRole DelWorkFlowRole = iota
	/**
	按照标记删除(需要交互,填写删除原因)
	*/
	DeleteByFlag
	/**
	删除到日志库(需要交互,填写删除原因)
	*/
	DeleteAndWriteToLog
	/**
	彻底的删除(不需要交互，直接干净彻底的删除)
	*/
	DeleteReal
	/**
	让用户决定删除方式(需要交互)
	*/
	ByUser
)

/**
投递方式
*/
type DeliveryWay uint8

const (
	/**
	按角色(以部门为纬度)
	*/
	ByStation_DeliveryWay DeliveryWay = iota
	/**
	按部门
	*/
	ByDept_DeliveryWay
	/**
	按SQL
	*/
	BySQL_DeliveryWay
	/**
	按本节点绑定的人员
	*/
	ByBindEmp
	/**
	由上一步发送人选择
	*/
	BySelected

	/**
	按表单选择人员
	*/
	ByPreviousNodeFormEmpsField
	/**
	与上一节点的人员相同
	*/
	ByPreviousNodeEmp
	/**
	与开始节点的人员相同
	*/
	ByStarter
	/**
	与指定节点的人员相同
	*/
	BySpecNodeEmp
	/**
	按角色与部门交集计算
	*/
	ByDeptAndStation
	/**
	按角色计算(以部门集合为纬度)
	*/
	ByStationAndEmpDept
	/**
	按指定节点的人员或者指定字段作为人员的角色计算
	*/
	BySpecNodeEmpStation
	/**
	按SQL确定子线程接受人与数据源.
	*/
	BySQLAsSubThreadEmpsAndData
	/**
	按明细表确定子线程接受人.
	*/
	ByDtlAsSubThreadEmps
	/**
	仅按角色计算
	*/
	ByStationOnly
	/**
	FEE计算.
	*/
	ByFEE
	/**
	按绑定部门计算,该部门一人处理标识该工作结束(子线程).
	*/
	BySetDeptAsSubthread
	/**
	按SQL模版计算
	*/
	BySQLTemplate
	/**
	从人员到人员
	*/
	ByFromEmpToEmp
	/**
	本部门内的人员
	*/
	FindSpecDeptEmpsInStationlist
	/**
	按照角色计算-范围内的
	*/
	ByStationForPrj
	/**
	按照选择模式计算.
	*/
	BySelectedForPrj
	/**
	按照设置的组织计算
	*/
	BySelectedOrgs
	/**
	按照部门领导计算
	*/
	ByDeptLeader
	/**
	按照用户组计算(本组织范围内)
	*/
	ByTeamOrgOnly
	/**
	按照用户组计算(全集团)
	*/
	ByTeamOnly
	/**
	按照用户组计算(本部门范围内)
	*/
	ByTeamDeptOnly
	/**
	按照绑定角色的用户组人员
	*/
	ByBindTeamEmp
	/**
	按照部门分管领导计算
	*/
	ByDeptShipLeader

	/**
	所有人员都可以发起
	*/
	BySelected_1 DeliveryWay = iota + 12

	/**
	按照组织模式人员选择器
	*/
	BySelectedEmpsOrgModel DeliveryWay = iota + 13
	/**
	按照自定义Url模式的人员选择器
	*/
	BySelfUrl
	/**
	按照设置的WebAPI接口获取的数据计算
	*/
	ByAPIUrl
	/**
	发送人的上级部门的负责人
	就是找上级领导主管.
	*/
	BySenderParentDeptLeader
	/**
	发送人上级部门指定的角色
	*/
	BySenderParentDeptStations
	_
	_
	/**
	找自己的直属领导.
	*/
	ByEmpLeader

	/**
	外部用户
	*/
	ByGuest
	/**
	按照部门计算
	*/
	ByPreviousNodeFormDepts
	/**
	按照角色计算
	*/
	ByPreviousNodeFormStationsAI
	/**
	智能计算
	*/
	ByPreviousNodeFormStationsOnly
	/**
	选择其他组织的联络员
	*/
	BySelectEmpByOfficer

	/**
	按指定的部门集合与设置的岗位交集
	*/
	ByStationSpecDepts
	/**
	按指定的角色集合与设置的部门交集
	*/
	ByStationSpecStas
	/**
	绑定字典表
	*/
	BySFTable
	_
	/**
	固定范围的选择
	*/
	BySelected_2
	/**
	按照ccflow的BPM模式处理
	*/
	ByCCFlowBPM DeliveryWay = 100
)

/**
方向条件控制规则
*/
type DirCondModel uint8

const (
	/**
	由连接线控制
	*/
	ByLineCond DirCondModel = iota
	/**
	发送后手工选择到达节点与接受人（用于子线程节点）
	*/
	ByPopSelect
	/**
	下拉框模式
	*/
	ByDDLSelected
	/**
	按照按钮选择
	*/
	ByButtonSelected
)

/**
附件开放类型
*/
type FJOpen uint8

const (
	/**
	不开放
	*/
	None_FJOpen FJOpen = iota
	/**
	对操作员开放
	*/
	ForEmp_FJOpen
	/**
	对工作ID开放
	*/
	ForWorkID_FJOpen
	/**
	对流程ID开放
	*/
	ForFID_FJOpen
)

/**
分流规则
*/
type FLRole uint8

const (
	/**
	按照接受人
	*/
	ByEmp_FLRole FLRole = iota
	/**
	按照部门
	*/
	ByDept_FLRole
	/**
	按照角色
	*/
	ByStation_FLRole
)

/**
流程启动类型
*/
type FlowRunWay uint8

const (
	/**
	手工启动
	*/
	HandWork FlowRunWay = iota
	/**
	指定人员按时启动
	*/
	SpecEmp
	/**
	数据集按时启动
	*/
	SelectSQLModel
	/**
	触发式启动
	*/
	WF_TaskTableInsertModel
	/**
	指定人员按时启动高级模式
	*/
	SpecEmpAdv
	/**
	  让管理员启动流程发送到指定
	*/
	LetAdminSendSpecEmp
)

/**
// 流程类型
// */
// type FlowType_del uint8

// const (
// 	/**
// 	平面流程
// 	*/
// 	Panel FlowType_del = iota
// 	/**
// 	分合流
// 	*/
// 	FHL
// )

type GuestFlowRole uint8

const (
	/**
	不参与
	*/
	None_GuestFlowRole GuestFlowRole = iota
	/**
	开始节点参与
	*/
	StartNodeJoin
	/**
	中间节点参与
	*/
	MiddleNodeJoin
)

/**
组长会签规则
*/
type HuiQianLeaderRole uint8

const (
	/**
	仅有一个组长
	*/
	OnlyOne HuiQianLeaderRole = iota
	/**
	最后一个组长为主
	*/
	LastOneMain
	/**
	任意组长为主
	*/
	EveryOneMain
)

/**
会签模式
*/
type HuiQianRole uint8

const (
	None_HuiQianRole HuiQianRole = iota
	/**
	队列(按照顺序处理，有最后一个人发送到下一个节点)
	*/
	Teamup_HuiQianRole
	/**
	协作组长模式
	*/
	TeamupGroupLeader_HuiQianRole
)

/**
会签任务状态
*/
type HuiQianTaskSta uint8

const (
	/**
	无
	*/
	None_HuiQianTaskSta HuiQianTaskSta = iota
	/**
	会签中
	*/
	HuiQianing
	/**
	会签完成
	*/
	HuiQianOver
)

type HungupSta uint8

const (
	/**
	申请
	*/
	Apply HungupSta = iota
	/**
	同意
	*/
	Agree
	/**
	拒绝
	*/
	Reject
)

/**
挂起方式
*/
type HungupWay uint8

const (
	/**
	永久挂起
	*/
	Forever HungupWay = iota
	/**
	在指定的日期解除
	*/
	SpecDataRel
)

/**
导入流程的模式
*/
type ImpFlowTempleteModel uint8

const (
	/**
	按新的流程导入
	*/
	AsNewFlow ImpFlowTempleteModel = iota
	/**
	按模版的流程编号
	*/
	AsTempleteFlowNo
	/**
	覆盖当前的流程
	*/
	OvrewaiteCurrFlowNo
	/**
	按指定的流程编号导入
	*/
	AsSpecFlowNo
)

/**
节点工作退回规则
*/
type JumpWay uint8

const (
	/**
	不能跳转
	*/
	CanNotJump JumpWay = iota
	/**
	向后跳转
	*/
	Next
	/**
	向前跳转
	*/
	Previous
	/**
	任何节点
	*/
	AnyNode
	/**
	任意点
	*/
	JumpSpecifiedNodes
)

/**
消息状态
*/
type MsgSta uint8

const (
	/**
	未开始
	*/
	UnRun MsgSta = iota
	/**
	成功
	*/
	RunOK
	/**
	失败
	*/
	RunError
	/**
	禁止发送
	*/
	Disable_MsgSta
)

/**
节点表单类型
*/
type NodeFormType uint8

const (
	/**
	傻瓜表单.
	*/
	FoolForm_NodeFormType NodeFormType = iota
	_
	/**
	嵌入式表单.
	*/
	SelfForm
	/**
	SDKForm
	*/
	SDKForm
	/**
	表单树
	*/
	SheetTree
	/**
	动态表单树
	*/
	SheetAutoTree
	/**
	公文表单
	*/
	WebOffice
	/**
	Excel表单
	*/
	ExcelForm
	/**
	Word表单
	*/
	WordForm
	/**
	累加表单
	*/
	FoolTruck
	/**
	节点单表单
	*/
	RefOneFrmTree
	/**
	开发者表单
	*/
	Develop_NodeFormType
	/**
	开发者表单
	*/
	ChapterFrm_NodeFormType
	/**
	引用指定节点上的表单
	*/
	RefNodeFrm
	/**
	禁用(对多表单流程有效)
	*/
	DisableIt NodeFormType = 100
)

/**
位置类型
*/
type NodePosType uint8

const (
	/**
	开始
	*/
	Start_NodePosType NodePosType = iota
	/**
	中间
	*/
	Mid
	/**
	结束
	*/
	End
)

type NodeType uint8

const (
	/**
	用户节点
	*/
	UserNode NodeType = iota
	/**
	路由节点
	*/
	RouteNode
	/**
	抄送节点
	*/
	CCNode
	/**
	子流程节点
	*/
	SubFlowNode
)

/**
节点工作类型
*/
type NodeWorkType uint8

const (
	Work NodeWorkType = iota
	/**
	开始节点
	*/
	StartWork
	/**
	开始节点分流
	*/
	StartWorkFL
	/**
	合流节点
	*/
	WorkHL
	/**
	分流节点
	*/
	WorkFL
	/**
	分合流
	*/
	WorkFHL
	/**
	子流程
	*/
	SubThreadWork
)

/**
待办工作超时处理方式
*/
type OutTimeDeal uint8

const (
	/**
	不处理
	*/
	None_OutTimeDeal OutTimeDeal = iota
	/**
	自动的转向下一步骤
	*/
	AutoTurntoNextStep
	/**
	自动跳转到指定的点
	*/
	AutoJumpToSpecNode
	/**
	自动移交到指定的人员
	*/
	AutoShiftToSpecUser
	/**
	向指定的人员发送消息
	*/
	SendMsgToSpecUser
	/**
	删除流程
	*/
	DeleteFlow
	/**
	执行SQL
	*/
	RunSQL_OutTimeDeal
)

/**
打印方式
@0=不打印@1=打印网页@2=打印RTF模板
*/
type PrintDocEnable uint8

const (
	/**
	不打印
	*/
	None_PrintDocEnable PrintDocEnable = iota
	/**
	打印网页
	*/
	PrintHtml
	/**
	打印RTF模板
	*/
	PrintRTF
	/**
	打印word
	*/
	PrintWord
)

/**
抢办发送后执行规则
*/
type QiangBanSendAfterRole uint8

const (
	/**
	不处理
	*/
	None_QiangBanSendAfterRole QiangBanSendAfterRole = iota
	/**
	抄送给其他人
	*/
	CCToEtcEmps
	/**
	发送消息给其他人
	*/
	SendMsgToEtcEmps
)

/**
已读回执类型
*/
type ReadReceipts uint8

const (
	/**
	不回执
	*/
	None_ReadReceipts ReadReceipts = iota
	/**
	自动回执
	*/
	Auto_ReadReceipts
	/**
	由系统字段决定
	*/
	BySysField
	/**
	按开发者参数
	*/
	BySDKPara
)

/**
节点工作退回规则
*/
type ReturnRole uint8

const (
	/**
	不能退回
	*/
	CanNotReturn_ReturnRole ReturnRole = iota
	/**
	只能退回上一个节点
	*/
	ReturnPreviousNode
	/**
	可退回以前任意节点(默认)
	*/
	ReturnAnyNodes
	/**
	可退回指定的节点
	*/
	ReturnSpecifiedNodes
	/**
	由流程图设计的退回路线来决定
	*/
	ByReturnLine
)

/**
运行模式
*/
type RunModel uint8

const (
	/**
	普通
	*/
	Ordinary_RunModel RunModel = iota
	/**
	合流
	*/
	HL
	/**
	分流
	*/
	FL
	/**
	分合流
	*/
	FHL
	/**
	同表单子线程
	*/
	SubThreadSameWorkID
	/**
	异表单子线程
	*/
	SubThreadUnSameWorkID
)

/**
保存模式
*/
type SaveModel uint8

const (
	/**
	仅节点表.
	*/
	NDOnly SaveModel = iota
	/**
	节点表与Rpt表.
	*/
	NDAndRpt
)

type SelectorModel uint8

const (
	/**
	角色
	*/
	Station_SelectorModel SelectorModel = iota
	/**
	部门
	*/
	Dept
	/**
	操作员
	*/
	Emp
	/**
	SQL
	*/
	SQL_SelectorModel
	/**
	SQL模版计算
	*/
	SQLTemplate
	/**
	通用的人员选择器.
	*/
	GenerUserSelecter
	/**
	按部门与角色的交集
	*/
	DeptAndStation
	/**
	自定义链接
	*/
	Url_SelectorModel
	/**
	通用部门角色人员选择器
	*/
	AccepterOfDeptStationEmp
	/**
	按角色智能计算(操作员所在部门)
	*/
	AccepterOfDeptStationOfCurrentOper
	/**
	按本组织用户计算.
	*/
	TeamOrgOnly
	/**
	按全组织用户计算
	*/
	TeamOnly
	/**
	按本组织部门用户计算
	*/
	TeamDeptOnly
	/**
	按照角色智能计算
	*/
	ByStationAI
	/**
	按照webapi接口计算
	*/
	ByWebAPI
	/**
	按照webapi接口计算
	*/
	ByMyDeptEmps
)

/**
消息类型
*/
type SendReturnMsgType uint8

const (
	/**
	消息
	*/
	Info_SendReturnMsgType SendReturnMsgType = iota
	/**
	系统消息
	*/
	SystemMsg
)

/**
短消息写入规则（当系统产生短消息的时候，需要写入的方式。）
*/
type ShortMessageWriteTo uint8

const (
	/**
	写入Sys_SMS表
	*/
	ToSMSTable ShortMessageWriteTo = iota
	/**
	写入WebServices.
	WS地址: \DataUser\PortalInterface.asmx 的 WriteShortMessage
	*/
	ToWebservices
	/**
	写入丁丁
	*/
	ToDingDing
	/**
	写入微信.
	*/
	ToWeiXin
	/**
	写入CCIM
	*/
	CCIM
)

/**
流程发起限制
*/
type StartLimitRole uint8

const (
	/**
	不限制
	*/
	None_StartLimitRole StartLimitRole = iota
	/**
	一人一天一次
	*/
	Day
	/**
	一人一周一次
	*/
	Week
	/**
	一人一月一次
	*/
	Month
	/**
	一人一季度一次
	*/
	JD
	/**
	一人一年一次
	*/
	Year
	/**
	发起的列不能重复,(多个列可以用逗号分开)
	*/
	ColNotExit
	/**
	设置的SQL数据源为空,或者返回结果为零时可以启动.
	*/
	ResultIsZero
	/**
	设置的SQL数据源为空,或者返回结果为零时不可以启动.
	*/
	ResultIsNotZero
	/**
	为子流程时仅仅只能被调用1次.
	*/
	OnlyOneSubFlow
)

/**
装在前提示
*/
type StartLimitWhen uint8

const (
	/**
	表单装载后
	*/
	StartFlow StartLimitWhen = iota
	/**
	发送前检查
	*/
	SendWhen
)

/**
根据子流程运行模式控制父流程的运行到下一个节点/结束
*/
type SubFlowRunModel uint8

const (
	/**
	无，不设置
	*/
	None_SubFlowRunModel SubFlowRunModel = iota
	/**
	子流程结束
	*/
	FlowOver_SubFlowRunModel
	/**
	子流程运行到指定节点
	*/
	SpecifiedNodes
)

/**
时间段
*/
type TSpan uint8

const (
	/**
	本周
	*/
	ThisWeek TSpan = iota
	/**
	上周
	*/
	NextWeek
	/**
	上上周
	*/
	TowWeekAgo
	/**
	更早
	*/
	More_TSpan
)

/**
任务状态
*/
type TaskSta uint8

const (
	/**
	无
	*/
	None_TaskSta TaskSta = iota
	/**
	共享
	*/
	Sharing_TaskSta
	/**
	已经取走
	*/
	Takeback
)

/**
组长确认规则
*/
type TeamLeaderConfirmRole uint8

const (
	/**
	按照部门表的字段 Leader 模式计算.
	*/
	ByDeptFieldLeader TeamLeaderConfirmRole = iota
	/**
	按照SQL计算.
	*/
	BySQL_TeamLeaderConfirmRole
	/**
	会签时主持人计算
	*/
	HuiQianLeader
)

/**
子线程删除规则
*/
type ThreadKillRole uint8

const (
	/**
	不能删除，不许等到全部完成才可以向下运动。
	*/
	None_ThreadKillRole ThreadKillRole = iota
	/**
	需要手工的删除才可以向下运动。
	*/
	ByHand
	/**
	自动删除未完成的子线程。
	*/
	ByAuto
)

/**
普通工作节点处理模式
*/
type TodolistModel uint8

const (
	/**
	抢办(谁抢到谁来办理,办理完后其他人就不能办理.)
	*/
	QiangBan TodolistModel = iota
	/**
	协作(没有处理顺序，接受的人都要去处理,由最后一个人发送到下一个节点)
	*/
	Teamup_TodolistModel
	/**
	队列(按照顺序处理，有最后一个人发送到下一个节点)
	*/
	Order_TodolistModel
	/**
	共享模式(需要申请，申请后才能执行)
	*/
	Sharing_TodolistModel
	/**
	协作组长模式
	*/
	TeamupGroupLeader_TodolistModel
)

/**
流程运行类型
*/
type TransferCustomType uint8

const (
	/**
	按照流程定义的模式执行(自动模式)
	*/
	ByCCBPMDefine TransferCustomType = iota
	/**
	按照工作人员的设置执行(人工干涉模式,人工定义模式.)
	*/
	ByWorkerSet
)

/**
节点完成转向处理
*/
type TurnToDeal uint8

const (
	/**
	按系统默认的提示
	*/
	CCFlowMsg TurnToDeal = iota
	/**
	指定消息
	*/
	SpecMsg
	/**
	指定Url
	*/
	SpecUrl
	/**
	发送后关闭
	*/
	TurntoClose
	/**
	按条件转向
	*/
	TurnToByCond
)

/**
用户信息显示格式
*/
type UserInfoShowModel uint8

const (
	/**
	用户ID,用户名
	*/
	UserIDUserName UserInfoShowModel = iota
	/**
	用户ID
	*/
	UserIDOnly
	/**
	用户名
	*/
	UserNameOnly
)

/**
流程状态(简)
*/
type WFSta uint8

const (
	/**
	运行中
	*/
	Runing_WFSta WFSta = iota
	/**
	已完成
	*/
	Complete_WFSta
	/**
	其他
	*/
	Etc_WFSta
)

/**
 流程状态(详)
 ccflow根据是否启用草稿分两种工作模式,它的设置是在web.config 是 IsEnableDraft 节点来配置的.
 1, 不启用草稿  IsEnableDraft = 0.
	这种模式下，就没有草稿状态, 一个用户进入工作界面后就生成一个Blank, 用户保存时，也是存储blank状态。
 2, 启用草稿.
*/
type WFState uint8

const (
	/**
	空白
	*/
	Blank WFState = iota
	/**
	草稿
	*/
	Draft_WFState
	/**
	运行中
	*/
	Runing_WFState
	/**
	已完成
	*/
	Complete_WFState
	/**
	挂起
	*/
	Hungup_WFState
	/**
	退回
	*/
	ReturnSta
	/**
	转发(移交)
	*/
	Shift_WFState
	/**
	删除(逻辑删除状态)
	*/
	Delete_WFState
	/**
	加签
	*/
	Askfor
	/**
	冻结
	*/
	Fix
	/**
	加签回复状态
	*/
	AskForReplay
)

/**
周末休息类型
*/
type WeekResetType uint8

const (
	/**
	双休
	*/
	Double WeekResetType = iota
	/**
	单休
	*/
	Single_WeekResetType
	/**
	不
	*/
	None_WeekResetType
)

/**
谁执行它
*/
type WhoDoIt uint8

const (
	/**
	操作员
	*/
	Operator WhoDoIt = iota
	/**
	机器
	*/
	MachtionOnly
	/**
	混合
	*/
	Mux
)

/**
工作类型
*/
type WorkType uint8

const (
	/**
	普通的
	*/
	Ordinary_WorkType WorkType = iota
	/**
	自动的
	*/
	Auto_WorkType
)

/**
完成状态
*/
type CHSta uint8

const (
	/**
	按期完成
	*/
	AnQi CHSta = iota
	/**
	预期完成
	*/
	YuQi
)
