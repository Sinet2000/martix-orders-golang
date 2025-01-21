package entities

type PagedResponse[T any] struct {
	Data        []T    `json:"db"`          // List of items (generic type)
	PageSize    int    `json:"pageSize"`    // Number of items per page
	PageIndex   int    `json:"pageIndex"`   // Current page index
	HasPrevPage bool   `json:"hasPrevPage"` // Indicates if a previous page exists
	HasNextPage bool   `json:"hasNextPage"` // Indicates if a next page exists
	Sorts       string `json:"sorts"`       // Sorting criteria (e.g., "name ASC, created_at DESC")
	Filters     string `json:"filters"`     // Applied filters (Like in sieve)
}

func NewPagedResponse[T any](data []T, totalCount int64, pageIndex, pageSize int, sorts string, filters string) *PagedResponse[T] {
	hasPrevPage := pageIndex > 1
	hasNextPage := int(totalCount) > pageIndex*pageSize

	return &PagedResponse[T]{
		Data:        data,
		PageSize:    pageSize,
		PageIndex:   pageIndex,
		HasPrevPage: hasPrevPage,
		HasNextPage: hasNextPage,
		Sorts:       sorts,
		Filters:     filters,
	}
}
