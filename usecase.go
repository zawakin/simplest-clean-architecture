package main

// Usecase は、単一リクエストを処理するインターフェイスです.
type Usecase interface {
	Handle(req *Request)
}

// interactor は、Usecaseインターフェイスの実装です.
type interactor struct {
	presenter Presenter
}

// NewInteractorUsecase は、Usecaseインターフェイスを実装した interactor を返します.
func NewInteractorUsecase(presenter Presenter) Usecase {
	return &interactor{
		presenter: presenter,
	}
}

func (in *interactor) Handle(req *Request) {
	// 役割：リクエスト(Request)をレスポンス(Response)に変換し、レスポンスを出力する.
	res := &Response{
		Reply: req.Message + " " + req.Message,
	}
	in.presenter.Show(res)
}
