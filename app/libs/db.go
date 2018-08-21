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

func SyncDB() {
	CreateTableIfNotExists("places", map[string]string{
		"id": "int(11) not null auto_increment", 
		"address": "varchar(300) not null",
		"lat": "varchar(16) not null",
		"lng": "varchar(16) not null" })

	CreateTableIfNotExists("common_places", map[string]string{
		"id": "int(11) not null auto_increment", 
		"name": "varchar(300) not null",
		"lat": "varchar(16) not null",
		"lng": "varchar(16) not null" })

	CreateTableIfNotExists("users", map[string]string{
		"id": "int(11) not null auto_increment", 
		"phone": "varchar(20) not null" })
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

	Println("insert into " + dbconf["db"] + "." + table + vals_declared + vals)

	dbConnection.Exec("insert into " + dbconf["db"] + "." + table + vals_declared + vals)
}

func Select(what string, from string, where string) []map[string]string {
	dbconf := DB()

	if where != "" { where = " where " + where }

	rows, err := dbConnection.Query("select " + what + " from " + dbconf["db"] + "." + from + where)
	
	defer rows.Close()

	if err != nil { panic(err) }

	var places []map[string]string

	if from == "places" {
		var (
			id string
			address string
			lat string
			lng string
		)
	
		for rows.Next(){
			err := rows.Scan(&id, &address, &lat, &lng)
	
			if err != nil{ panic(err) }
			
			places = append(places, map[string]string{"id": id, "lat": lat, "lng": lng, "address": address})
		}
	} else if from == "common_places" {
		var (
			id string
			name string
			lat string
			lng string
		)
	
		for rows.Next(){
			err := rows.Scan(&id, &name, &lat, &lng)
	
			if err != nil{ panic(err) }
			
			places = append(places, map[string]string{"id": id, "name": name, "lat": lat, "lng": lng})
		}
	}

	return places
}

func Exists(table string, where string) int {
	dbconf := DB()

	if where != "" { where = " where " + where }

	rows, err := dbConnection.Query("select count(*) as count from" + dbconf["db"] + "." + table + " where " + where)

	if err != nil{ panic(err) }

	defer rows.Close()

	var count int

	var recs map[string]int

    for rows.Next(){
		err := rows.Scan(&count)

		if err != nil{ panic(err) }
		
        recs = map[string]int{"count": count}
	}

	var exists = 0

	if recs["count"] > 0 {exists = 1}

	return exists
}

func Delete(from string, where string) {
	dbconf := DB()

	dbConnection.Exec("delete from " + dbconf["db"] + "." + from + " where " + where)
}

func CreateTableIfNotExists(table string, fields map[string]string) {
	createQuery := ""

	for key, value := range fields {
		createQuery += key + " " + value + ", "
	}

	createQuery += "PRIMARY KEY (id)"

	dbConnection.Exec("create table if not exists gotcha." + table + " (" + createQuery + ") ")
}