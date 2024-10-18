#! /bin/sh

psql -U $POSTGRES_USER -h $POSTGRES_SERVER -d postgres -W $POSTGRES_PASSWORD<<EOF
DROP TABLE IF EXIST users;
CREATE TABLE users (
    user_name VARCHAR(100)  PRIMARY KEY,
    user_email VARCHAR(100) NOT NULL,
    post_title VARCHAR(100) NOT NULL,
    post_content VARCHAR(500) NOT NULL
);
INSERT INTO users VALUES 
('Bob Joe','bob@nowhere.com','Birds and Geology','Myrspoven rocks!')
('Alice Adam','alice@nowhere.com','URGENT','Can I eat all the cookies?');
EOF