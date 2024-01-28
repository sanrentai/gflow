package enum

// 附件类型
type AdjunctType uint8

const (
	/**
	不需要附件。
	*/
	None_AdjunctType AdjunctType = iota
	/**
	图片
	*/
	PhotoOnly
	/**
	word 文档。
	*/
	WordOnly
	/**
	所有的类型
	*/
	ExcelOnly
	/**
	所有的类型。
	*/
	AllType
)

/**
实体附件类型
*/
type BPEntityAthType uint8

const (
	/**
	无
	*/
	None_BPEntityAthType BPEntityAthType = iota
	/**
	单附件
	*/
	Single_BPEntityAthType
	/**
	多附件
	*/
	Multi_BPEntityAthType
)

/**
工作模式
*/
type Dot2DotModel uint8

const (
	/**
	默认模式
	*/
	Default Dot2DotModel = iota
	/**
	树模式
	*/
	TreeDept
	/**
	树叶子模式
	*/
	TreeDeptEmp
)

/**
 * 从表的编辑模式
 */
type DtlEditerModel uint8

const (
	//批量编辑
	DtlBatch DtlEditerModel = iota
	//查询编辑模式
	DtlSearch

	//自定义URL
	DtlURL

	//在EnOnly显示查询
	DtlBatchEnonly

	//在EnOnly
	DtlSearchEnonly

	//
	DtlURLEnonly
)

/**
编辑类型
*/
type EditType uint

const (
	/**
	可编辑
	*/
	Edit EditType = iota
	/**
	不可删除
	*/
	UnDel
	/**
	只读,不可删除。
	*/
	Readonly
)

/**
实体类型
*/
type EnType int

const (
	/**
	系统实体
	*/
	Sys EditType = iota
	/**
	管理员维护的实体
	*/
	Admin
	/**
	应用程序实体
	*/
	App_EnType
	/**
	第三方的实体（可以更新）
	*/
	ThirdPartApp
	/**
	视图(更新无效)
	*/
	View
	/**
	可以纳入权限管理
	*/
	PowerAble
	/**
	其他
	*/
	Etc_EnType
	/**
	明细或着点对点。
	*/
	Dtl_EnType
	/**
	点对点
	*/
	Dot2Dot
	/**
	XML　类型
	*/
	XML
	/**
	扩展类型，它用于查询的需要。
	*/
	Ext
)

/**
编辑器类型
*/
type EditerType uint8

const (
	/**
	无编辑器
	*/
	None_EditerType EditerType = iota
	/**
	Sina编辑器
	*/
	Sina
	/**
	FKEditer
	*/
	FKEditer
	/**
	KindEditor
	*/
	KindEditor
	/**
	百度的UEditor
	*/
	UEditor
)

/**
字段类型
*/
type FieldType uint8

const (
	/**
	正常的
	*/
	Normal_FieldType FieldType = iota
	/**
	主键
	*/
	PK_FieldType
	/**
	外键
	*/
	FK_FieldType
	/**
	枚举
	*/
	Enum_FieldType
	/**
	既是主键又是外键
	*/
	PKFK
	/**
	既是主键又是枚举
	*/
	PKEnum
	/**
	关连的文本.
	*/
	RefText
	/**
	虚拟的
	*/
	NormalVirtual
	/**
	多值的
	*/
	MultiValues
)

/**
逻辑类型
*/
type FieldTypeS uint8

const (

	/**
	普通类型
	*/
	Normal_FieldTypeS FieldTypeS = iota
	/**
	枚举类型
	*/
	Enum_FieldTypeS
	/**
	外键
	*/
	FK_FieldTypeS
)

/**
移动到显示方式
*/
type MoveToShowWay uint8

const (
	/**
	不显示
	*/
	None_MoveToShowWay MoveToShowWay = iota
	/**
	下拉列表
	*/
	DDL_MoveToShowWay
	/**
	平铺
	*/
	Panel
)

type MsgShowType uint8

const (
	/**
	本界面
	*/
	SelfAlert MsgShowType = iota
	/**
	提示框
	*/
	SelfMsgWindows
	/**
	新窗口
	*/
	Blank_MsgShowType
)

type OperatorSymbol uint8

const (
	/**
	大于
	*/
	DaYu OperatorSymbol = iota
	/**
	等于
	*/
	DengYu
	/**
	小于
	*/
	XiaoYu
	/**
	相似
	*/
	Like
)

/**
显示位置
*/
type RMShowWhere uint8

const (
	/**
	实例左侧
	*/
	EnLeft RMShowWhere = iota
	/**
	实例工具栏
	*/
	EnToolbar
	/**
	查询工具栏
	*/
	SearchToolbar
)

/**
相关功能类型
*/
type RefMethodType uint8

const (
	/**
	功能
	*/
	Func RefMethodType = iota
	/**
	模态窗口打开
	*/
	LinkModel
	/**
	新窗口打开
	*/
	LinkeWinOpen
	/**
	右侧窗口打开
	*/
	RightFrameOpen
)

/**
  控件类型
*/
type UIContralType uint8

const (
	/**
	文本框
	*/
	TB UIContralType = iota
	/**
	下拉框
	*/
	DDL_UIContralType
	/**
	CheckBok
	*/
	CheckBok
	/**
	单选择按钮
	*/
	RadioBtn
	/**
	地图定位
	*/
	MapPin
	/**
	录音控件
	*/
	MicHot
	/**
	附件展示控件
	*/
	AthShow
	/**
	手机拍照控件
	*/
	MobilePhoto
	/**
	手写签名版
	*/
	HandWriting
	/**
	超链接
	*/
	HyperLink
	/**
	文本
	*/
	Lab
	/**
	图片
	*/
	FrmImg
	/**
	图片附件
	*/
	FrmImgAth
	/**
	身份证号
	*/
	IDCard
	/**
	签批组件
	*/
	SignCheck
	/**
	评论组件
	*/
	FlowBBS_UIContralType
	/**
	系统定位
	*/
	Fixed
	/**
	发文字号
	*/
	DocWord

	/**
	 * 公文正文组件
	 */
	GovDocFile UIContralType = 110
	/**
	收文字号
	*/
	DocWordReceive UIContralType = 170
	/**
	流程进度图
	*/
	JobSchedule UIContralType = 50
	/**
	大块文本Html(说明性文字)
	*/
	BigText UIContralType = 60
	/**
	评分
	*/
	Score UIContralType = 101
)
