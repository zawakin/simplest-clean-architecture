package main

import "fmt"

// Presenter は、レスポンスを出力するインターフェイスです.
// Clean Architecture では、Presenter という名前です。
type Presenter interface {
	Show(res *Response)
}

// presenterImpl は、Presenterインターフェイスの実装です.
type presenterImpl struct {
}

// NewPresenter は、Presenterインターフェイスを実装した PresenterImpl を返します.
func NewPresenter() Presenter {
	return &presenterImpl{}
}

func (presenter *presenterImpl) Show(res *Response) {
	// 役割：レスポンスを出力ビュー用のレスポンス(ViewResponse)に変換し、なんかする.
	vres := &ViewResponse{
		Reply: fmt.Sprintf("=== %s ===", res.Reply),
	}
	fmt.Println(vres)
}
