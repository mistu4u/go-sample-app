#!/bin/sh
CMD=$1

URL=$DATABASE_URL
if [ -z "$URL" ]
then
  URL="$DB_DRIVER://$DB_USER:$DB_PASS@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=$SSL_MODE"
fi

export DATABASE_URL=$URL

dbmate wait
if [ -z $CMD ] || [ $CMD != "rollback" ]
then
    echo "migrating to newest schema"
    dbmate migrate
else
    echo "rollback $ROLLBACK"
    dbmate down
fi