### golang-migrateのインストール

Mac:   
brew install golang-migrate   
Linux:   
curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -   
echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list   
apt-get update   
apt-get install -y migrate
centos:   
https://qiita.com/libra_lt/items/548fcbdfbcf992cba4ed   
   
### Migration実行
migrate -database="mysql://myapp:admin@tcp(127.0.0.1:3306)/myapp" -path=db/migrate up   
   
### Migrationファイル作成
migrate create -ext sql -dir ~/myappTest/db/migrate create_user_table   
   
### Dirty version ~~のエラーが出た時
上手く起動できなかったSQLがあった時、そいつのバージョンが残ったままで邪魔しているのでmigrate forceを使用する   
   
migrate -database="mysql://myapp:admin@tcp(127.0.0.1:3306)/myapp" -path=db/migrate force 20220707100531   
   
### 困った時は
履歴自体はgolang-migrateが自動でschema_migrationsというテーブルが管理している。そのレコードを消したり操作すれば状態を変えられる。   
   
公式：https://github.com/golang-migrate/migrate/blob/master/GETTING_STARTED.md#forcing-your-database-version   
その他：   
https://asapoon.com/golang/4002/golang-migrate/   
https://qiita.com/juchilian/items/0bfed79cc1229deb4c62   
https://qiita.com/yutaiii/items/b2579bdf989d6d8f3c11   