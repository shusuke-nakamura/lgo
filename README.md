# lgo

# heyのインストール
go install github.com/rakyll/hey@latest

# リンターのshadow
go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest

module-root/
|---- README.md
|---- cmd/
|  |---- cmd-tools/
|    |---- ...
|    |---- main.go
|  |---- web-apps/
|    |---- ...
|    |---- main.go
|---- go.mod
|---- go.sum
|---- pkg/
|  |---- customer/
|    |---- customer1.go
|    |---- customer2.go
|  |---- inventory/
|    |---- inventory1.go
|    |---- inventory2.go
|    |---- inventory3.go
|  |---- package3/
|    |---- package3.go
|---- main.go

◇ モジュール
・利用可能なバージョンを確認するコマンド
【コマンド】
go list
(例)
go list -m -versions github.com/learning-go-book/simpletax

・依存関係にあるモジュールのバージョン変更
【コマンド】
go get 
(例)
go get github.com/learning-go-book/simpletax@v1.0.0

◇ pkg.go.devサイト
https://pkg.go.dev/