## GO 

### Project create 

```bash

PS D:\sdev\222dis\gventas\app> go mod init github.com/222dis/gventas
go: creating new go.mod: module github.com/222dis/gventas
PS D:\sdev\222dis\gventas\app>
```
*github.com/222dis/gventas debe ser único


```bash
Tools environment: GOPATH=C:\Users\Lab_Redes\go
Installing 7 tools at C:\Users\Lab_Redes\go\bin in module mode.
  gotests
  gomodifytags
  impl
  goplay
  dlv
  staticcheck
  gopls


```

### Runing local  

```bash
PS D:\sdev\222dis\gventas\app>go install github.com/cosmtrek/air@latest

PS D:\sdev\222dis\gventas\app>air

```
### Runing docker  

```bash
PS D:\sdev\222dis\gventas>docker-compose up --build -d

PS D:\sdev\222dis\gventas>docker exec -it gventas_app bash
root@a5b44fbef5b1:/go/src/app# air

```

ToDo ...




