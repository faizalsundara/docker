set MYSQL_CONN_STRING=root@tcp(localhost:3306)/hello_world?parseTime=true
set PORT=8083
go run main.go lib.go model.go