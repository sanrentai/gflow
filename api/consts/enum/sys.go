package enum

/**
表单类型
*/
type AppType uint8

const (
	/**
	独立表单
	*/
	Application_AppType AppType = 0
	/**
	节点表单
	*/
	Node_AppType
)

type AthCtrlWay uint8

const (
	/**
	表单主键
	*/
	PK_AthCtrlWay AthCtrlWay = iota
	/**
	FID
	*/
	FID
	/**
	父流程ID
	*/
	PWorkID
	/**
	仅仅查看自己的
	*/
	MySelfOnly
	/**
	工作ID,对流程有效.
	*/
	WorkID
	/**
	P2流程
	*/
	P2WorkID
	/**
	P3流程
	*/
	P3WorkID
	/**
	根流程的WorkID
	*/
	RootFlowWorkID_AthCtrlWay
)

/**
附件删除规则
*/
type AthDeleteWay uint8

const (
	/**
	不删除 0
	*/
	None_AthDeleteWay AthDeleteWay = iota
	/**
	删除所有 1
	*/
	DelAll
	/**
	只删除自己上传 2
	*/
	DelSelf
)

/**
运行模式
*/
type AthRunModel uint8

const (
	/**
	记录模式
	*/
	RecordModel AthRunModel = iota
	/**
	固定模式
	*/
	FixModel
)

/**
文件保存方式
*/
type AthSaveWay uint8

const (
	/**
	IIS服务器
	*/
	IISServer AthSaveWay = iota
	/**
	保存到数据库
	*/
	DB
	/**
	ftp
	*/
	FTPServer
	/**
	 * OSS
	 */
	OSS
)

/**
附件在扩展控件里的显示方式
*/
type AthShowModel uint8

const (
	/**
	简单的
	*/
	Simple AthShowModel = iota
	/**
	只有文件名称
	*/
	FileNameOnly
)

/**
附件上传方式
*/
type AthUploadWay uint8

const (
	/**
	继承模式
	*/
	Inherit AthUploadWay = iota
	/**
	协作模式
	*/
	Interwork
)

/**
附件上传类型
*/
type AttachmentUploadType uint8

const (
	/**
	单个的
	*/
	Single_AttachmentUploadType AttachmentUploadType = iota
	/**
	多个的
	*/
	Multi_AttachmentUploadType
	/**
	指定的
	*/
	Specifically
)

/**
按钮事件类型 - 与sl 中设置的要相同。
*/
type BtnEventType uint8

const (
	/**
	禁用
	*/
	Disable_BtnEventType BtnEventType = iota
	/**
	运行存储过程
	*/
	RunSP
	/**
	运行sql
	*/
	RunSQL_BtnEventType
	/**
	执行URL
	*/
	RunURL
	/**
	运行webservices
	*/
	RunWS
	/**
	运行Exe文件.
	*/
	RunExe
	/**
	运行JS
	*/
	RunJS
)

/**
按钮类型
*/
type BtnType uint8

const (
	/**
	保存
	*/
	Save BtnType = iota
	/**
	打印
	*/
	Print
	/**
	删除
	*/
	Delete_BtnType
	/**
	增加
	*/
	Add
	/**
	自定义
	*/
	Self BtnType = 100
)

/**
按钮访问
*/
type BtnUAC uint8

const (
	/**
	不处理
	*/
	None_BtnUAC BtnUAC = iota
	/**
	按人员
	*/
	ByEmp
	/**
	按角色
	*/
	ByStation
	/**
	按部门
	*/
	ByDept
	/**
	按sql
	*/
	BySQL
)

/**
运行模式
*/
type CCBPMRunModel uint8

const (

	/**
	单机版
	*/
	Single CCBPMRunModel = iota
	/**
	集团模式
	*/
	GroupInc
	/**
	多租户模式
	*/
	SAAS
)

/**
编码表类型
*/
type CodeStruct uint8

const (
	/**
	普通的编码表
	*/
	NoName_CodeStruct CodeStruct = iota
	/**
	树编码表(No,Name,ParentNo)
	*/
	Tree_CodeStruct
	/**
	行政机构编码表
	*/
	GradeNoName
)

/**
按日期查询方式
*/
type DTSearchWay uint8

const (
	/**
	不设置
	*/
	None_DTSearchWay DTSearchWay = iota
	/**
	按日期
	*/
	ByDate
	/**
	按日期时间
	*/
	ByDateTime
	/**
	按照年月
	*/
	ByYearMonth
	/**
	按照年份
	*/
	ByYear
)

/**
明细表存盘方式
*/
type DtlAddRecModel uint8

const (
	/**
	自动初始化空白行
	*/
	ByBlank DtlAddRecModel = iota
	/**
	用按钮增加行
	*/
	ByButton
)

/**
明细表工作方式
*/
type DtlModel uint8

const (
	/**
	普通的
	*/
	Ordinary DtlModel = iota
	/**
	固定列
	*/
	FixRow
)

type DtlOpenType uint8

const (
	/**
	对人员开放
	*/
	ForEmp_DtlOpenType DtlOpenType = iota
	/**
	对工作开放
	*/
	ForWorkID_DtlOpenType
	/**
	对流程开放
	*/
	ForFID_DtlOpenType
	/**
	父工作ID
	*/
	ForPWorkID

	ForP2WorkID

	ForP3WorkID
	/**
	根流程的WorkID
	*/
	RootFlowWorkID_DtlOpenType
)

type DtlSaveModel uint8

const (
	/**
	失去焦点自动存盘
	*/
	AutoSave DtlSaveModel = iota
	/**
	由保存按钮触发存盘
	*/
	HandSave
)

/**
从表显示方式
*/
type EditModel uint8

const (
	/**
	表格模式
	*/
	TableModel EditModel = iota
	/**
	傻瓜表单模式
	*/
	FoolModel
	/**
	自由表单模式
	*/
	FreeModel
)

/**
实体类型
*/
type EntityType uint8

const (
	/**
	独立表单
	*/
	SingleFrm EntityType = iota
	/**
	单据
	*/
	FrmBill
	/**
	实体
	*/
	FrmDict
	/**
	实体树
	*/
	EntityTree
	/**
	数据源类型
	*/
	DBList_EntityType EntityType = 100
)

/**
事件执行内容
*/
type EventDoType uint8

const (
	/**
	禁用
	*/
	Disable EventDoType = iota
	/**
	执行存储过程
	*/
	SP
	/**
	运行SQL
	*/
	SQL_EventDoType
	/**
	自定义URL
	*/
	URLOfSelf
	/**
	自定义WS
	*/
	WSOfSelf
	/**
	执行ddl文件的类与方法
	*/
	SpecClass
	/**
	基类
	*/
	EventBase
	/**
	执行的业务单元
	*/
	BuessUnit
	/**
	自定义WebApi
	*/
	WebApi
	/**
	SFProduces
	*/
	SFProcedure
)

/**
事件类型
*/
type EventSource uint8

const (
	/**
	表单
	*/
	Frm EventSource = iota
	/**
	流程
	*/
	Flow_EventSource
	/**
	节点
	*/
	Node_EventSource
)

/**
文件展现方式
*/
type FileShowWay uint8

const (
	/**
	表格
	*/
	Table_FileShowWay FileShowWay = iota
	/**
	图片
	*/
	Pict
	/**
	自由模式
	*/
	Free
)

/**
表单事件列表
*/
type FrmEvenList11 uint8

const (
	/**
	创建OID
	*/
	CreateOID FrmEvenList11 = iota
	/**
	装载前
	*/
	FrmLoadBefore
	/**
	装载后
	*/
	FrmLoadAfter
	/**
	保存前
	*/
	SaveBefore
	/**
	保存后
	*/
	SaveAfter
)

type FrmFrom uint8

const (
	Flow FrmFrom = iota
	Node
	Dtl
)

/**
表单类型 @0=傻瓜表单@1=自由表单@3=嵌入式表单@4=Word表单@5=Excel表单@6=VSTOForExcel@7=Entity
*/
type FrmType uint8

const (
	/**
	傻瓜表单
	*/
	FoolForm_FrmType FrmType = iota
	_
	_
	/**
	URL表单(自定义)
	*/
	Url_FrmType
	/**
	Word类型表单
	*/
	WordFrm
	/**
	Excel表单
	*/
	ExcelFrm
	/**
	VSTOExccel模式.
	*/
	VSTOForExcel
	/**
	实体类
	*/
	Entity
	/**
	开发者表单
	*/
	Develop_FrmType
	/**
	WPS表单
	*/
	WPSFrm

	/**
	章节表单
	*/
	ChapterFrm_FrmType
	/**
	外部数据源列表
	*/
	DBList_FrmType FrmType = 100
)

/**
图片应用类型
*/
type ImgAppType uint8

const (
	/**
	图片
	*/
	Img ImgAppType = iota
	/**
	图片公章
	*/
	Seal
	/**
	北京安证通公章CA
	*/
	SealESA
	/**
	二维码
	*/
	QRCode
)

/**
导入模式
*/
type ImpModel uint8

const (
	/**
	不执行导入
	*/
	None_ImpModel ImpModel = iota
	/**
	表格模式
	*/
	Table_ImpModel
	/**
	按照Excel文件模式
	*/
	ExcelFile
	/**
	单据模式
	*/
	BillModel
)

/**
从表显示模式
*/
type ListShowModel uint8

const (
	/**
	表格模式
	*/
	Table_ListShowModel ListShowModel = iota
	/**
	傻瓜表单模式
	*/
	Card
)

/**
编号生成规则
*/
type NoGenerModel uint8

const (
	/**
	自定义
	*/
	None_NoGenerModel NoGenerModel = iota
	/**
	流水号
	*/
	ByLSH
	/**
	标签的全拼
	*/
	ByQuanPin
	/**
	标签的简拼
	*/
	ByJianPin
	/**
	按GUID生成
	*/
	ByGUID
)

/**
组织结构类型
*/
type OSModel uint8

const (
	/**
	一个人一个部门模式.
	*/
	OneOne OSModel = iota
	/**
	一个人多个部门模式
	*/
	OneMore
)

/**
百分比显示方式
*/
type PercentModel uint8

const (
	/**
	不显示
	*/
	None_PercentModel PercentModel = iota
	/**
	纵向
	*/
	Vertical
	/**
	横向
	*/
	Transverse
)

type PicType uint8

const (
	/**
	自动签名
	*/
	Auto PicType = iota
	/**
	手动签名
	*/
	ShouDong
)

/**
Pop返回值类型
*/
type PopValFormat uint8

const (
	/**
	编号
	*/
	OnlyNo PopValFormat = iota
	/**
	名称
	*/
	OnlyName
	/**
	编号与名称
	*/
	NoName_PopValFormat
)

/**
选择模式
*/
type PopValSelectModel uint8

const (
	/**
	单选
	*/
	One PopValSelectModel = iota
	/**
	多选
	*/
	More_PopValSelectModel
)

/**
PopVal - 工作方式
*/
type PopValWorkModel uint8

const (
	/**
	禁用
	*/
	None_PopValWorkModel PopValWorkModel = iota
	/**
	自定义URL
	*/
	SelfUrl
	/**
	表格模式
	*/
	TableOnly
	/**
	表格分页模式
	*/
	TablePage
	/**
	分组模式
	*/
	Group_PopValWorkModel
	/**
	树展现模式
	*/
	Tree_PopValWorkModel
	/**
	双实体树
	*/
	TreeDouble
)

/**
数字签名类型
*/
type SignType uint8

const (
	/**
	无
	*/
	None_SignType SignType = iota
	/**
	图片
	*/
	Pic
	/**
	山东CA签名.
	*/
	CA
	/**
	广东CA
	*/
	GDCA
	/**
	图片盖章
	*/
	GZCA
)

/**
文本框类型
*/
type TBModel uint8

const (
	/**
	正常的
	*/
	Normal TBModel = iota
	/**
	大文本
	*/
	BigDoc
	/**
	富文本
	*/
	RichText
	/**
	超大文本
	*/
	SupperText
)

/**
上传校验,方式.
*/
type UploadFileNumCheck uint8

const (
	/**
	不校验
	*/
	None_UploadFileNumCheck = iota
	/**
	不能为空
	*/
	NotEmpty
	/**
	每个类别不能为空.
	*/
	EverySortNoteEmpty
)

/**
棫行处理
*/
type WhenOverSize uint8

const (
	/**
	不处理
	*/
	None_WhenOverSize WhenOverSize = iota
	/**
	增加一行
	*/
	AddRow
	/**
	翻页
	*/
	TurnPage
)
