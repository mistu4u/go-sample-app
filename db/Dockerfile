FROM amacneil/dbmate

WORKDIR /scripts

COPY migrations ./db/migrations
COPY migrate.sh .

RUN chmod +x ./migrate.sh

ENTRYPOINT ["sh", "migrate.sh"]