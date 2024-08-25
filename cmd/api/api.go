package api 

import "database/sql"

type ApiServer struct {
	addr string 
	db  *sql.DB
}

