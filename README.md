# Connect to a postgresql database

```
mv go_pg.go main.go
go mod init go_pg
go mod tidy

go build
./go_pg
```

## git

* create .gitignore

```
*
!.gitignore
!*.go
!*.mod
!*.sum
!README.md
```

* initialize, add, push

```
git init
git add .gitignore
git add .
git status
git commit
git remote add origin git@github.com:mohan-nbs-ont/go_pg.git
gh auth login
gh repo create go_pg --private
git push --set-upstream origin master
git push
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
