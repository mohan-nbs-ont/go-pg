# Project - `go_pg`

```
mv go_pg.go main.go

# the name of the executable will be go_pg
go mod init go_pg

# will download deps
go mod tidy

go build
```

push changes to github

```
git status
git add .
# executable should not be in list
git status
git commit -S
git push origin master
```

.gitignore

```
*
!.gitignore
!*.go
!*.mod
```
