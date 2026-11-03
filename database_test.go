package golang_db

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test(t *testing.T) {

}

func TestConnect(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
