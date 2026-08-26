package service

import "context"

type Page struct {
	Limit  int
	Offset int
	Total  int
}
type PageResult[T any] struct {
	Items []T
	Page  Page
}

func NormalizePage(limit, offset int) Page {
	if limit < 1 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}
	return Page{Limit: limit, Offset: offset}
}
func (p Page) Next() Page    { return Page{Limit: p.Limit, Offset: p.Offset + p.Limit, Total: p.Total} }
func (p Page) HasNext() bool { return p.Offset+p.Limit < p.Total }
func Paginate[T any](ctx context.Context, items []T, limit, offset int) PageResult[T] {
	p := NormalizePage(limit, offset)
	if p.Offset > len(items) {
		p.Offset = len(items)
	}
	end := p.Offset + p.Limit
	if end > len(items) {
		end = len(items)
	}
	p.Total = len(items)
	result := items[p.Offset:end]
	return PageResult[T]{Items: result, Page: p}
}
