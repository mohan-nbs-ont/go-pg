# Connect to a postgresql database

```
go mod init go_pg
go mod tidy

go build
./go_pg
```

* create database tables

```
DROP SEQUENCE books_sequence;
DROP TABLE books;
CREATE TABLE books (id SERIAL PRIMARY KEY, title VARCHAR(100) NOT NULL, primary_author VARCHAR(100) NULL);
CREATE SEQUENCE books_sequence start 1 increment 1;
```
* use env variable

```
DATABASE_URL=postgres://user:password@127.0.0.1:5432/books_db
```
