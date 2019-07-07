# Go GraphQL

## 作業手順

### 新規作成

リポジトリを作成

```bash
mkdir go-graphql
go mod init github.com/Mushus/go-graphql
```

`schema.graphql` ファイルを作成

```bash
go run github.com/99designs/gqlgen init
```

以下のファイルが生成される

```plain
.
├── generated.go
├── gqlgen.yml
├── models_gen.go
├── resolver.go   基本的にここにコードを書く
└── server
    └── server.go 実行エントリーポイント
```

新しくエンドポイントが増えた場合は再度生成すれば良い

```bash
go run github.com/99designs/gqlgen
```

NOTE: 再度生成するときにエラーが出たため使用していない部分を一旦コメントアウトした

```gqlgen.yaml
# directives:
#   deprecated:
#     skip_runtime: true
#   include:
#     skip_runtime: true
#   skip:
#     skip_runtime: true
```

# Resolver の作成

エントリーポイントになるやつ。  
ルートの `resolver.go` は自動生成のたびに中が書き換わるのでこちらを書き換えたくない。  
幸い resolver 自体は単なる Interface なのでコピペしてコードを用意した。
