# 作業📝:golang-migrate

- 作業ディレクトリ: ./supporting

- docker-compose up -d > データベース立ち上げ
- docker-compose down > 環境を全消去
- docker exec -it local-db bash
  - -it ${container_name} > 対象のコンテ内のBashにアクセスする
- root@${conatiner_name}プロンプトにてpsql -h localhost -d ausu -U admin
- ${databasename}プロンプトにて\qでデータベースプロンプトから切断
- migrate create -ext sql -dir migrations -seq create_${tablename}_table
- database url "postgres://${username}:${password}@localhost:5432/${databasename}?sslmode=disable"
- migrate -database="postgres://admin:password@localhost:5432/ausu?sslmode=disable" -path=migrations/ up ${version}

# 作業📝:sqlboiler
- sqlboilerのインストール
- sqlboiler postgresql driverのインストール
- go.modにsqlboiler/v4の依存を追加する
- go.modにnull/v8の依存を追加する
- sqlboilerコマンドがcommand not foundになる場合はPATHが通っているか確認するとよい
- sqlboiler.tomlはcオプションを付けないと既存環境変数と衝突するのでつける
- 設定ファイルにsslmodeの項目を追加しないとSSL is not enabledのエラーが出るためつける
- sqlboiler psql -c sqlboiler.toml  --output apps/supporting-api/infra/postgres/models