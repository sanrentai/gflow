package enum

/**
柱状图显示类型
*/
type ColumnChartShowType uint8

const (
	/**
	不显示
	*/
	None_ColumnChartShowType ColumnChartShowType = iota
	/**
	横向
	*/
	HengXiang_ColumnChartShowType
	/**
	竖向
	*/
	ShuXiang_ColumnChartShowType
	/**
	横向叠加
	*/
	HengXiangAdd
	/**
	竖向叠加
	*/
	ShuXiangAdd
)

/**
显示类型.
*/
type DBAChartType uint8

const (
	Table_DBAChartType DBAChartType = iota
	Column
	Pie_DBAChartType
	Line_DBAChartType
)

/**
折线图图显示类型
*/
type LineChartShowType uint8

const (
	/**
	不显示
	*/
	None_LineChartShowType LineChartShowType = iota
	/**
	横向
	*/
	HengXiang_LineChartShowType
	/**
	竖向
	*/
	ShuXiang_LineChartShowType
)
