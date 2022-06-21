before run:

go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite

run:

rm "${GOPATH}tests\allure-results"
go test "${GOPATH}goApiTests\tests"