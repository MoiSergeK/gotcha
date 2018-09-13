package db


func SyncDB() {
	CreateTableIfNotExists("self_places", map[string]string{
		"id": "int(11) not null auto_increment", 
		"user_id": "int(11) not null", 
		"address": "varchar(300) not null",
		"lat": "varchar(16) not null",
		"lng": "varchar(16) not null" })

	CreateTableIfNotExists("self_routes", map[string]string{
		"id": "int(11) not null auto_increment", 
		"user_id": "int(11) not null",
		"name": "varchar(300)",
		"route": "text",
		"share": "varchar(10)" })											// all | friends | {id} | no

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
		"share": "varchar(10)",												// all | friends | {id} | no
		"online": "int(11)" })

	CreateTableIfNotExists("user_code", map[string]string{
		"id": "int(11) not null auto_increment",
		"user_id": "int(11)",
		"code": "varchar(10)" })

	CreateTableIfNotExists("friends", map[string]string{
		"id": "int(11) not null auto_increment",
		"user1_id": "int(11)",
		"user2_id": "int(11)" })

	CreateTableIfNotExists("location_requests", map[string]string{
		"id": "int(11) not null auto_increment",
		"from_user_id": "int(11)",
		"to_user_id": "int(11)" })

	CreateTableIfNotExists("location_responses", map[string]string{
		"id": "int(11) not null auto_increment",
		"from_user_id": "int(11)",
		"to_user_id": "int(11)",
		"lat": "varchar(16) not null",
		"lng": "varchar(16) not null" })
}