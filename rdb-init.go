package main

import "database/sql"

func main() {
	rdb, _ := sql.Open("sqlite3", "./test.db")
	defer rdb.Close()
	rdb.Exec("CREATE TABLE users (user_id integer primary key,username varchar(64) not null,password varchar(64) not null,privilege integer not null)")
}
