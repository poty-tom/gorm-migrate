# gorm-migrate
MariaDBコンテナに対して`gorm.AutoMigrate`による自動生成を行うプログラム

## 動作
```bash
# .env
cp .env.template .env

# イメージビルド
docker build -t gorm-migrate:latest -f docker/mariadb/Dockerfile .

# コンテナ起動
docker compose up -d