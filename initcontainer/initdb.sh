#! /bin/sh

PGHOST=$POSTGRES_SERVER
PGPORT='5432'
PGDATABASE='postgres'
PGUSER=$POSTGRES_USER
PGPASSWORD=$POSTGRES_PASSWORD

while true
    do
    response=$(psql --csv<<EOF |grep -v pg_is_in_recovery
\pset tuples_only off
\pset footer off
SELECT pg_is_in_recovery();
EOF
    )
    echo "$response" | grep -q "f"
    if [[ $? -eq 0 ]]; then
        break
    else
        echo "DB isn't ready - trying again"
    fi
    sleep 1
done


psql<<EOF
\pset tuples_only off;
\pset footer off;
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    user_name VARCHAR(100)  PRIMARY KEY,
    user_email VARCHAR(100) NOT NULL,
    post_title VARCHAR(100) NOT NULL,
    post_content VARCHAR(500) NOT NULL
);
INSERT INTO users VALUES 
('Bob Joe','bob@nowhere.com','Birds and Geology','Myrspoven rocks!'),
('Alice Adam','alice@nowhere.com','URGENT','Can I eat all the cookies?');
commit;
EOF