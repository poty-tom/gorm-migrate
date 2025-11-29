#!/usr/bin/env bash
set -euo pipefail

# Start background waiter + migrate
/usr/local/bin/wait-and-migrate.sh &

# Delegate to official entrypoint (PID 1)
if [[ $# -gt 0 ]]; then
  exec /usr/local/bin/docker-entrypoint.sh "$@"
else
  exec /usr/local/bin/docker-entrypoint.sh mariadbd
fi


