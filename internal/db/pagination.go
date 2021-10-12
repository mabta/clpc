package db

import "math"

// Pagination 分页
type Pagination struct {
	// Page 当前页码
	Page int
	// TotalRows 总记录数
	TotalRows int
	// TotalPages 总页数
	TotalPages int
	// PageSize 每页记录数
	PageSize int
	// PageList 分页列表
	PageList []int
}

// NewPagination 构建新的分页对象
func NewPagination(page, totalRows, pageSize int) *Pagination {
	totalPages := int(math.Ceil(float64(totalRows) / float64(pageSize)))
	pageList := make([]int, totalPages)
	for i := 0; i < totalPages; i++ {
		pageList[i] = i
	}
	return &Pagination{
		Page:       page,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		PageSize:   pageSize,
		PageList:   pageList,
	}
}

// PageListToView 将分页列表转换成可易读
func (p *Pagination) PageListToView() []int {
	pageList := make([]int, 0, len(p.PageList))
	for _, i := range p.PageList {
		pageList = append(pageList, i+1)
	}
	return pageList
}
