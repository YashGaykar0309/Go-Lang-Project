docker run --name postgres-db \
  -e POSTGRES_USER=admin \
  -e POSTGRES_PASSWORD=admin123 \
  -e POSTGRES_DB=mydb \
  -p 5432:5432 \
  -d postgres:15
