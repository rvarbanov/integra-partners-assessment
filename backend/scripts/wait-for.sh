#!/bin/sh
# wait-for.sh

set -e

host="$1"
shift
cmd="$@"

until nc -z "$host" 5432; do
  echo "Waiting for database to be ready... ($host:5432)"
  sleep 1
done

>&2 echo "Database is up - executing command"
exec $cmd