# gtfs-gorm
ORM for gtfsdb-goedition

テストデータの出典先
http://opendata.pref.toyama.jp/pages/gtfs_jp.htm

## サブ機能 gtfsdb-go
gtfsをdbに格納する機能があります．  
gtfsファイルまたは，URLを指定し実行します．

### 使用方法
### Step1 build
以下のコマンドでビルドします．  
`go build ./cmd/gtfsdb-go`  

### Step2 run
Step1でビルドしたアプリケーションを以下のコマンドで実行します．  
`./gtfsdb-go`
### オプション
| オプション     | 説明                                                              | 例                                               | 
| ------------- |-----------------------------------------------------------------| ------------------------------------------------ | 
| --file        | 条件付き必須: gtfsのzipファイルをしています.  '--url'オプションをを使用する場合は必要ありません．      | --file hoge.zip                                  | 
| --url         | 条件付き必須: gtfsのurlを指定します．'--file'オプションを使用する場合は必要ありません．            | --url https://hoge.com/foo/bar                   | 
| --dsn         | オプション: PostgreSQL データベースに接続するための DSN（データソース名）を指定します.            | --dsn postgres://hoge:fuga@localhost:5432/foobar | 
| --schema      | オプション: データをインポートするデータベースのスキーマを指定します.指定がない場合は，publicスキーマが指定されます． | --schema fuga                                    | 
| --recal       | オプション: 幾何データや関連する空間データの再計算を行います．例)shape_dist_traveled           | --recal                                          | 
| --shapesex    | オプション: shapes_exテーブルの追加生成を行います.                                 | --shapesex                                       | 
| --shapesdetail | オプション: shapes_detailテーブルの追加生成を行います.                             | --shapesdetail                                   | 
| --geom        | オプション(有効を推奨): ジオメトリデータ処理を有効にし,PostGISの幾何計算を含むことができます.           | --geom                                           | 

### 使用例
`./gtfsdb-go --file hoge.zip --schema public --dsn postgres://hoge:fuga@localhost:5432/foobar`  
`./gtfsdb-go --url https://hoge.com/foo/bar --dsn postgres://hoge:fuga@localhost:5432/foobar --recal --shapesex --shapesdetail --geom`  
