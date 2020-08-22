package main

// Request は、アプリで扱いやすい形式のリクエストのモデルです. Usecase によって処理されます。
type Request struct {
	Message string
}

// Response は、アプリで扱いやすい形式のレスポンスのモデルです. これは、Presenter によって ViewResponse に変換されます.
type Response struct {
	Reply string
}

// ViewRequest は、このアプリケーションに外部から入力されたリクエストのモデルです.
type ViewRequest struct {
	Message string
}

// ViewResponse は、ビューでそのまま表示できるように変換されたレスポンスのモデルです.
type ViewResponse struct {
	Reply string
}
