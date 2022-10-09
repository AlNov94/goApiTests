before run:

go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite

run:

go test "${GOPATH}goApiTests\tests" -parallel 10