package db


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
		"phone": "varchar(20)",
		"email": "varchar(300)",
		"name": "varchar(50)",
		"hash": "varchar(128)",
		"share_location": "varchar(10)",							// all | friends | {phone} | no
		"online": "int(11)" })

	CreateTableIfNotExists("user_code", map[string]string{
		"id": "int(11) not null auto_increment",
		"user_id": "int(11)",
		"code": "varchar(10)" })

	CreateTableIfNotExists("friends", map[string]string{
		"id": "int(11) not null auto_increment",
		"user1_id": "int(11)",
		"user2_id": "int(11)" })
}