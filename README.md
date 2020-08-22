# 最もシンプルなクリーンアーキテクチャ

クリーンアーキテクチャの核となる概念を説明するために、非常に簡単なアプリを作りました。

# リクエストの流れ

1. 起動しているアプリに、なんらかの入力が入ってきます。
1. まずはその入力を、入力ビュー用リクエストモデル( `ViewRequest` )に詰め込み、コントローラ( `Controller` )に渡します。
1. コントローラは、次の二つをします。
    1. `ViewRequest` を処理しやすいリクエストモデル( `Request` )に変換します。
    1. リクエストモデル `Request` をユースケース ( `Usecase` )に渡して、処理を投げます。
1. ユースケースは、その実装である `Interactor` に従って、次のような処理を行います。
    1. リクエスト `Request` になんらかの処理を加え、レスポンス( `Response` )を作ります。
    1. 得られたレスポンスをビューに返すために、プレゼンター( `Presenter` )を呼び出します。
1. プレゼンターは、その実装に従って、次のような処理を行います。
    1. 出力ビュー用レスポンスモデル( `ViewResponse` )に変換を行います。
    1. ビューにレスポンスを返すことを通知します。