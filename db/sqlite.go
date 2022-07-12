package db

type PagedResult[D interface{}] struct {
	Data   []D   `json:"data"`
	PrevId int64 `json:"prevId"`
	NextId int64 `json:"nextId"`
}
type PageOptions struct {
	LastId int64
	Total  int64
}

func (p *PageOptions) PrevId() int64 {
	prevId := p.LastId - p.Total
	if prevId <= 0 {
		prevId = 0
	}
	return prevId
}
func (p *PageOptions) NextId(total int64) int64 {
	if total <= 0 {
		return p.LastId
	}
	return p.LastId + p.Total
}

func DefaultPageOptions(lastId int64, total int64) PageOptions {
	if lastId <= 0 {
		lastId = 0
	}

	if total <= 0 {
		total = 0
	}
	if total >= 100 {
		total = 100
	}
	return PageOptions{
		LastId: lastId,
		Total:  total,
	}
}
