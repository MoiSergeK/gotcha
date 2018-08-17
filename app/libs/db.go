package libs

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var dbConnection = initDBConn()
func DBConn() *sql.DB { return dbConnection }

func initDBConn() *sql.DB {
	dbconf := DB()
	db, err := sql.Open(dbconf["driver"], dbconf["user"] + ":" + dbconf["pass"] + "@/" + dbconf["db"])
	
	if err != nil { panic(err) }
	
	return db
}

func Insert(table string, values map[string]string) {
	dbconf := DB()

	vals_declared := "("
	vals := "("

	for key, value := range values{
		vals_declared += key + ","
		vals += "'" + value + "',"
	}

	vals_declared = vals_declared[0: len(vals_declared)-1]
	vals = vals[0: len(vals)-1]

	vals_declared += ") values"
	vals += ")"

	dbConnection.Exec("insert into " + dbconf["db"] + "." + table + vals_declared + vals)
}

func Select(what string, from string, where string) []map[string]string {
	dbconf := DB()

	if where != "" { where = " where " + where }

	rows, err := dbConnection.Query("select " + what + " from " + dbconf["db"] + "." + from + where)
	
	defer rows.Close()

	if err != nil { panic(err) }

	var (
		id string
		coords string
		address string
	)

	var places []map[string]string

    for rows.Next(){
		err := rows.Scan(&id, &coords, &address)

		if err != nil{ panic(err) }
		
        places = append(places, map[string]string{"id": id, "coords": coords, "address": address})
	}

	return places
}

func Delete(from string, where string) {
	dbconf := DB()

	dbConnection.Exec("delete from " + dbconf["db"] + "." + from + " where " + where)
}

func CreateTableIfNotExists(tableName string) {
	
}