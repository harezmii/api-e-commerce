package controller

type QueryArg struct {
	Offset       string `query:"offset"`
	Sort         string `query:"sort"`
	SortField    string `query:"sortField"`
	SelectFields string `query:"selectFields"`
	Search       string `query:"search"`
}
