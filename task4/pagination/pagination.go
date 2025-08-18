package pagination

import (
	"gorm.io/gorm"
)

// Param 分页请求参数
type Param struct {
	Page     int `form:"page" json:"page"`         // 页码，默认1
	PageSize int `form:"pageSize" json:"pageSize"` // 每页条数，默认10，最大100
}

// Result 分页返回结果
type Result struct {
	Page      int         `json:"page"`      // 当前页码
	PageSize  int         `json:"pageSize"`  // 每页条数
	Total     int64       `json:"total"`     // 总记录数
	TotalPage int         `json:"totalPage"` // 总页数
	Items     interface{} `json:"items"`     // 分页数据
}

// 处理分页参数，设置默认值
func (p *Param) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 || p.PageSize > 100 {
		p.PageSize = 10
	}
}

// 计算偏移量
func (p *Param) Offset() int {
	return (p.Page - 1) * p.PageSize
}

// 通用分页查询函数
func Paginate(db *gorm.DB, param Param, dest interface{}) (Result, error) {
	var total int64
	var result Result

	// 处理分页参数
	param.Process()

	// 查询总数
	if err := db.Count(&total).Error; err != nil {
		return result, err
	}

	// 执行分页查询
	if err := db.Offset(param.Offset()).Limit(param.PageSize).Find(dest).Error; err != nil {
		return result, err
	}

	// 计算总页数
	totalPage := int((total + int64(param.PageSize) - 1) / int64(param.PageSize))

	// 组装返回结果
	result = Result{
		Page:      param.Page,
		PageSize:  param.PageSize,
		Total:     total,
		TotalPage: totalPage,
		Items:     dest,
	}

	return result, nil
}
