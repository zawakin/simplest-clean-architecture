package main

// Controller は ViewRequest から Request を作り、 Usecase にそれを渡します.
type Controller struct {
	usecase Usecase
}

// NewController は、新しいControllerインスタンスを返します.
func NewController(usecase Usecase) *Controller {
	return &Controller{
		usecase: usecase,
	}
}

// Control は、入力ビュー用の形式のリクエストをアプリで扱いやすいリクエストに変換し、Usecaseにハンドルさせます.
func (c *Controller) Control(vreq *ViewRequest) {
	req := &Request{Message: vreq.Message}
	c.usecase.Handle(req)
}
