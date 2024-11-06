package application

type SingleResp[T any] struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    T      `json:"data"`
}

type MultiResp[T any] struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    []*T   `json:"data"`
}

type PageResp[T any] struct {
	Success  bool   `json:"success"`
	Msg      string `json:"msg"`
	Data     []*T   `json:"data"`
	Total    int    `json:"total"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

func SingleRespOf[T any](data T, msg string) SingleResp[T] {
	return SingleResp[T]{
		Success: true,
		Msg:     msg,
		Data:    data,
	}
}

func SingleRespOk[T any]() SingleResp[T] {
	return SingleResp[T]{
		Success: true,
	}
}

func SingleRespFail[T any](msg string) SingleResp[T] {
	var emptyData T
	return SingleResp[T]{
		Success: false,
		Msg:     msg,
		Data:    emptyData,
	}
}

func MultiRespOf[T any](data []*T, msg string) MultiResp[T] {
	return MultiResp[T]{
		Success: true,
		Msg:     msg,
		Data:    data,
	}
}

func MultiRespFail[T any](msg string) MultiResp[T] {
	return MultiResp[T]{
		Success: false,
		Msg:     msg,
		Data:    nil,
	}
}

func PageRespOf[T any](total int, page int, pageSize int, data []*T, msg string) PageResp[T] {
	return PageResp[T]{
		Success:  true,
		Msg:      msg,
		Data:     data,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
}

func PageRespFail[T any](msg string) PageResp[T] {
	return PageResp[T]{
		Success: false,
		Msg:     msg,
		Data:    nil,
		Total:   0,
	}
}

func (resp SingleResp[T]) ConvertToAnyResponse() *SingleResp[any] {
	return &SingleResp[any]{
		Data:    resp.Data,
		Msg:     resp.Msg,
		Success: resp.Success,
	}
}

func (resp MultiResp[T]) ConvertToAnyResponse() *MultiResp[any] {
	var data []*any
	for _, d := range resp.Data {
		var anyData any = d
		data = append(data, &anyData)
	}
	return &MultiResp[any]{
		Data:    data,
		Msg:     resp.Msg,
		Success: resp.Success,
	}
}

func (resp PageResp[T]) ConvertToAnyResponse() *PageResp[any] {
	var data []*any
	for _, d := range resp.Data {
		var anyData any = d
		data = append(data, &anyData)
	}
	return &PageResp[any]{
		Data:     data,
		Msg:      resp.Msg,
		Success:  resp.Success,
		Total:    resp.Total,
		Page:     resp.Page,
		PageSize: resp.PageSize,
	}
}
