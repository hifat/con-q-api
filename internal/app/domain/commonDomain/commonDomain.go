package commonDomain

type ReqQuery struct {
	Page     *int    `form:"page"`
	PerPage  *int    `form:"perPage"`
	Sort     *string `form:"sort"`
	Fields   *string `form:"fields"`
	SearchBy *string `form:"searchBy"`
	Search   *string `form:"search"`
}
