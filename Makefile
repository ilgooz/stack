run:
	go build -ldflags "-X main.buildstamp=`date -u '+%Y-%m-%dT%I:%M:%S'` -X main.githash=`git rev-parse HEAD`" && ./stack

win32:
	env GOOS=windows GOARCH=386 go build

win64:
	env GOOS=windows GOARCH=amd64 go build
