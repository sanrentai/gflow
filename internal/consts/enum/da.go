package enum

/**
图表类型
*/
type ChartType uint8

const (
	/**
	柱状图
	*/
	Histogram ChartType = iota
	/**
	丙状图
	*/
	Pie_ChartType
	/**
	折线图
	*/
	Line_ChartType
)

/**
数据检查级别
*/
type DBCheckLevel uint8

const (
	/**
	低,只出报告,不操作任何数据
	*/
	Low DBCheckLevel = iota + 1
	/**
	中,出检查报告,删除外键的左右空格.
	*/
	Middle
	/**
	高,删除对应不上的数据.
	*/
	High
)

/**
数据库部署类型
*/
type DBModel uint8

const (
	/**
	独立（集中模式）
	*/
	Single_DBModel DBModel = iota
	/**
	域模式
	*/
	Domain
)

/**
数据库类型
*/
type DBType uint8

const (
	/**
	sqlserver
	*/
	MSSQL DBType = iota
	/**
	oracle
	*/
	Oracle
	/**
	Access
	*/
	Access
	/**
	PostgreSQL
	*/
	PostgreSQL
	/**
	DB2
	*/
	DB2
	/**
	MySQL
	*/
	MySQL
	/**
	Informix
	*/
	Informix
	/**
	达梦
	*/
	DM
	/**
	 * 人大金仓R3
	 */
	KingBaseR3
	/**
	 * 人大金仓R6
	 */
	KingBaseR6
	/**
	 * 优炫
	 */
	UX
	/**
	 * 瀚高
	 */
	HGDB
)

/**
　连接到哪个库上．
  他们存放在 web.config 的列表内．
*/
type DBUrlType uint8

const (
	/**
	主应用程序
	*/
	AppCenterDSN DBUrlType = iota
	/**
	1号连接
	*/
	DBAccessOfOracle1
	/**
	2号连接
	*/
	DBAccessOfOracle2
	/**
	1号连接
	*/
	DBAccessOfMSSQL1
	/**
	2号连接
	*/
	DBAccessOfMSSQL2
	/**
	access的连接．
	*/
	DBAccessOfOLE
	/**
	ODBC连接
	*/
	DBAccessOfODBC
	/**
	数据源连接
	*/
	DBSrc
)

/**
保管位置
*/
type Depositary uint8

const (
	/**
	不保管
	*/
	None_Depositary Depositary = iota
	/**
	全体
	*/
	Application_Depositary
)

/**
时间计算方式
*/
type FieldCaseModel uint8

const (
	/// <summary>
	/// 无
	/// </summary>
	None_FieldCaseModel FieldCaseModel = iota
	/// <summary>
	/// 大写
	/// </summary>
	UpperCase
	/// <summary>
	/// 小写
	/// </summary>
	Lowercase
)

/**
分组方式
*/
type GroupWay uint8

const (
	/**
	求合
	*/
	BySum GroupWay = iota
	/**
	求平均
	*/
	ByAvg
)

/**
信息类型
*/
type LogType uint8

const (
	/**
	 * 调试
	 */
	Debug LogType = iota
	/**
	提示
	*/
	Info
	/**
	警告
	*/
	Warning
	/**
	错误
	*/
	Error
)

/**
排序方式
*/
type OrderWay uint8

const (
	/**
	升序
	*/
	OrderByUp OrderWay = iota
	/**
	降序
	*/
	OrderByDown
)

/**
时间计算方式
*/
type TWay uint8

const (
	/**
	计算节假日
	*/
	Holiday TWay = iota
	/**
	不计算节假日
	*/
	AllDays
)

type XmlWriteMode uint8

const (
	/**
	 * 以 XML 数据形式写入 bp.da.DataSet 的当前内容，以关系结构作为内联 XSD 架构。如果
	 * bp.da.DataSet 只有架构而无数据，那么只写入内联架构。如果 bp.da.DataSet
	 * 没有当前架构，则不写入任何内容。
	 */
	WriteSchema XmlWriteMode = iota
	/**
	 * 以 XML 数据形式写入 bp.da.DataSet 的当前内容，不带 XSD 架构。如果无数据加载到
	 * bp.da.DataSet 中，则不写入任何内容。
	 */
	IgnoreSchema
	/**
	 * 作为 DiffGram 写入整个 bp.da.DataSet，包括原始值和当前值。若要生成只包含已更改的值的 DiffGram，请调用
	 * bp.da.DataSet.GetChanges()，然后在返回的 bp.da.DataSet 上作为 DiffGram
	 * 调用 bp.da.DataSet.WriteXml(System.IO.Stream)。
	 */
	// DiffGram
)
