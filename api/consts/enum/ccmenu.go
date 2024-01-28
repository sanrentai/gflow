package enum

/**
控制方式
*/
type CtrlWay uint8

const (
	/**
	游客
	*/
	Guest CtrlWay = iota
	/**
	任何人
	*/
	AnyOne
	/**
	按角色
	*/
	ByStation_CtrlWay
	/**
	按部门
	*/
	ByDeptCtrlWay
	/**
	按人员
	*/
	ByEmp_CtrlWay
	/**
	按sql
	*/
	BySQL_CtrlWay
)

/**
菜单控制类型
*/
type MenuCtrlWay uint8

const (
	/**
	按照设置的控制
	*/
	BySetting MenuCtrlWay = iota
	/**
	任何人
	*/
	Anyone
	/**
	仅仅管理员
	*/
	AdminOnly
)

/**
菜单类型
*/
type MenuType uint8

const (
	/**
	系统根目录
	*/
	Root MenuType = iota
	/**
	系统类别
	*/
	AppSort
	/**
	系统
	*/
	App_MenuType
	/**
	目录
	*/
	Dir
	/**
	菜单
	*/
	Menu
	/**
	功能控制点
	*/
	Function
)
