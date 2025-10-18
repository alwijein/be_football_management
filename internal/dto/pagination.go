package dto

// Pagination represents pagination parameters
type Pagination struct {
	Page       int   `json:"page" form:"page"`
	Limit      int   `json:"limit" form:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// GetOffset returns offset for database query
func (p *Pagination) GetOffset() int {
	return (p.Page - 1) * p.Limit
}

// GetLimit returns limit, with default value
func (p *Pagination) GetLimit() int {
	if p.Limit <= 0 {
		p.Limit = 10
	}
	if p.Limit > 100 {
		p.Limit = 100
	}
	return p.Limit
}

// GetPage returns page, with default value
func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

// CalculateTotalPages calculates total pages
func (p *Pagination) CalculateTotalPages() {
	if p.Total > 0 && p.Limit > 0 {
		p.TotalPages = int((p.Total + int64(p.Limit) - 1) / int64(p.Limit))
	}
}
