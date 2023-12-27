# mock-alert-notifier
Bing AIに手取り足取りLayered ArichitectureでGoでいい感じに書くのを教えてもらいつつモックを作っていく

## prompt
```
golangのレイアードアーキテクチャによくあるservice、repositoryで以下の要件のシステムを作成するときのディレクトリ構成を教えて。

- ユーザ情報管理機能、メール送信機能、SMS送信機能が存在する
- ユーザ情報にユーザ名、メールアドレス、電話番号が記録されている
- メール送信機能は指定ユーザのメールアドレスにsendgrid apiを使用してメールを送信する
- SMS送信機能はtwilio apiを使用して指定ユーザの電話番号にSMSを送信する
```

## files
```bash
/myapp
├── main.go
├── go.mod
├── go.sum
└── pkg
    ├── user
    │   ├── handler.go
    │   ├── service.go
    │   └── repository.go
    ├── email
    │   ├── handler.go
    │   ├── service.go
    │   └── repository.go
    └── sms
        ├── handler.go
        ├── service.go
        └── repository.go
```

## concept
- main.go: アプリケーションのエントリーポイントです。
- pkg: ビジネスロジックを含むパッケージのディレクトリです。
- user, email, sms: 各機能に対応するパッケージです。それぞれには、handler.go（HTTPリクエストを処理）、service.go（ビジネスロジックを実装）、repository.go（データストレージとのインタラクションを担当）が含まれます
