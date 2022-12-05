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
