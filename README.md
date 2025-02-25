# gtfs-gorm
ORM for gtfsdb-goedition

テストデータの出典先  
jp: http://opendata.pref.toyama.jp/pages/gtfs_jp.htm. 
  
schedule: https://mobilitydatabase.org/feeds?gtfs=true&official=true

## サブ機能 gtfsdb-go
gtfsをdbに格納する機能があります．  
gtfsファイルまたは，URLを指定し実行します．

### 使用方法
### Step1 build
以下のコマンドでビルドします．

`go build ./cmd/gtfsdb-go`  書き換える

### Step2 run
Step1でビルドしたアプリケーションを以下のコマンドで実行します．  
`./gtfsdb-go`書き換える
### オプション
| オプション          | 説明                                                              | 例                                                        | 
|----------------|-----------------------------------------------------------------|----------------------------------------------------------| 
| --type         | 必須: gtfsの種類を選択します。（jp　か　schedule）                               | --type jp                                                | 
| --file         | 条件付き必須: gtfsのzipファイルをしています.  '--url'オプションをを使用する場合は必要ありません．      | --file filename.zip                                      | 
| --url          | 条件付き必須: gtfsのurlを指定します．'--file'オプションを使用する場合は必要ありません．            | --url https://hoge.com/foo/bar                           | 
| --dsn          | オプション: PostgreSQL データベースに接続するための DSN（データソース名）を指定します.            | --dsn postgres://username:password@localhost:5432/dbname | 
| --schema       | オプション: データをインポートするデータベースのスキーマを指定します.指定がない場合は，publicスキーマが指定されます． | --schema fuga                                            | 
| --recal        | オプション: 幾何データや関連する空間データの再計算を行います．例)shape_dist_traveled           | --recal                                                  | 
| --shapesex     | オプション: shapes_exテーブルの追加生成を行います.                                 | --shapesex                                               | 
| --shapesdetail | オプション: shapes_detailテーブルの追加生成を行います.                             | --shapesdetail                                           | 
| --geom         | オプション(有効を推奨): ジオメトリデータ処理を有効にし,PostGISの幾何計算を含むことができます.           | --geom                                                   | 

### 使用例
`./gtfsdb-go --type jp --file hoge.zip --schema public --dsn postgres://username:password@localhost:5432/dbname`  
`./gtfsdb-go --type schedule --url https://hoge.com/foo/bar --dsn postgres://username:password@localhost:5432/dbname --recal --shapesex --shapesdetail --geom`  
