# 作業📝:golang-migrate

- 作業ディレクトリ: ./supporting

- docker-compose up -d > データベース立ち上げ
- docker-compose down > 環境を全消去
- docker exec -it local-db bash
  - -it ${container_name} > 対象のコンテ内のBashにアクセスする
- root@${conatiner_name}プロンプトにてpsql -h localhost -d ausu -U admin
- ${databasename}プロンプトにて\qでデータベースプロンプトから切断
- database url "postgres://${username}:${password}@localhost:5432/${databasename}?sslmode=disable"
- migrate -database="postgres://admin:password@localhost:5432/ausu?sslmode=disable" -path=migrations/ up
- migrate -database="postgres://admin:password@localhost:5432/ausu?sslmode=disable" -path=migrations/ up ${version}

# 作業📝: sqlboiler

-　パッケージをインストールする
- 設定ファイルを準備する
- sqlboilerでコード生成をする
- governmentsテーブルに地方公共団体データをインソートする(地方公共団体){団体コード、名称}
- governmentsテーブルに存在するデータを削除する
- governmentsテーブルに存在するデータを編集する
