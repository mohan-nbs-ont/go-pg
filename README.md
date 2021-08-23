# My Project

mv go_pg.go main.go
go mod init go_pg # the name of the executable will be tdbill
go mod tidy        # will download deps

go build

git status
git add .
git status # executable should not be in list
git commit -S
git push origin master

.gitignore
```
*
!.gitignore
!*.go
!*.mod
```
