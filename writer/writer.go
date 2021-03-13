package writer

import (
	"database/sql"
	"db-loader/proto"
)

const (
	connectionInfoTemplate = ""
	insertQueryTemplate    = ""
	selectQueryTemplate    = ""
	updateQueryTemplate    = ""
	deleteQueryTemplate    = ""
)

type DBWriter struct {
	replica *sql.DB
}

func connectionStringBuild(dbi *proto.DBinfo) string {
	return ""
}

func (dbw *DBWriter) New(dbi *proto.DBinfo) error {
	instance, err := sql.Open("postgres", connectionStringBuild(dbi))
	if err != nil {
		return err
	}
	dbw.replica = instance
}

func (dbw *DBWriter) Insert(insInfo *proto.InsertInfo) error {
	return nil
}

func (dbw *DBWriter) Select(selectInfo *proto.SelectInfo) (*proto.SelectInfo, error) {
	return nil, nil
}

func (dbw *DBWriter) Update(updInfo *proto.UpdateInfo) error {
	return nil
}

func (dbw *DBWriter) Delete(delInfo *proto.DeleteInfo) error {
	return nil
}
