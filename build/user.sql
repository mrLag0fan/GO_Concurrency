CREATE DATABASE "go_user_db";
\c "go_user_db"
DROP TABLE IF EXISTS "go_user";
CREATE TABLE "go_user" (
    name VARCHAR NOT NULL,
    comment VARCHAR NOT NULL
);
