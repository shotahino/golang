PKG := main.go

all:
	GOOS=linux GOARCH=arm go build -v ${PKG}

old:
	GOOS=linux GOARCH=arm GOARM=5 go build -v ${PKG}

