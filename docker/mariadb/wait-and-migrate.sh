#!/usr/bin/env bash
set -euo pipefail

host="${MARIADB_HOST:-127.0.0.1}"
port="${MARIADB_PORT:-3306}"
user="${MARIADB_USER:-}"
pass="${MARIADB_PASSWORD:-}"

if [[ -z "$user" ]]; then
  user="root"
  pass="${MARIADB_ROOT_PASSWORD:-$pass}"
fi

echo "Waiting for MariaDB to be ready at ${host}:${port} as ${user} ..."
for i in $(seq 1 600); do
  if mariadb-admin ping -h "${host}" -P "${port}" -u "${user}" -p"${pass}" --silent; then
    echo "MariaDB is ready."
    break
  fi
  sleep 1
done

echo "Running AutoMigrate ..."
/usr/local/bin/migrate
echo "AutoMigrate finished."


