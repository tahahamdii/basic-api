package api 

import "database/sql"

type ApiServer struct {
	addr string 
	db  *sql.DB
}


func NewApiServer(addr string , db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db: db,
	}
}
 
func (s *ApiServer) Run() error {
	return nil
}