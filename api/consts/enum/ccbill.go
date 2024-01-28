package enum

/**
单据状态
*/
type BillState uint8

const (
	/**
	空白
	*/
	None_BillState BillState = iota
	/**
	草稿
	*/
	Draft_BillState
	/**
	编辑中
	*/
	Editing
	/**
	归档
	*/
	Over BillState = 100
)

type SearchDataRole uint8

const (
	/**
	只查询自己创建的
	*/
	ByOnlySelf_SearchDataRole SearchDataRole = iota
	/**
	查询本部门创建的包含兼职部门
	*/
	ByDept_SearchDataRole
	/**
	查询本部门（包含兼职部门）及子级部门
	*/
	ByDeptAndSSubLevel
	/**
	根据角色设定的部门的集合
	*/
	ByStationDept
	/**
	查询所有用户创建的数据信息
	*/
	SearchAll
)

type SearchDataRoleByDeptStation uint8

const (
	/**
	只查询自己创建的
	*/
	ByOnlySelf_SearchDataRoleByDeptStation SearchDataRoleByDeptStation = iota
	/**
	查询本部门创建的包含兼职部门
	*/
	ByDept_SearchDataRoleByDeptStation
)
