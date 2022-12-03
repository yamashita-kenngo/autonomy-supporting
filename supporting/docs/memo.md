# 作業📝

- 作業ディレクトリ: ./supporting

- docker-compose up -d > データベース立ち上げ
- docker-compose down > 環境を全消去
- docker exec -it local-db bash
  - -it ${container_name} > 対象のコンテ内のBashにアクセスする
- root@${conatiner_name}プロンプトにてpsql - localhost -d ausu -U admin
- ${databasename}プロンプトにて\qでデータベースプロンプトから切断
