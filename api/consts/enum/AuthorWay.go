package enum

/**
授权方式
*/
type AuthorWay uint8

const (

	/**
	不授权
	*/
	None_AuthorWay AuthorWay = iota
	/**
	全部授权
	*/
	All_AuthorWay
	/**
	指定流程授权
	*/
	SpecFlows
)
