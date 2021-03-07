package writer

import "db-loader/decls"

type DBWriter struct{}

func (dbw *DBWriter) producer(msgs chan<- decls.Incident) {}
func (dbw *DBWriter) consumer(query chan<- string)        {}

func (dbw *DBWriter) Write(inData *decls.Incident) error {
	go dbw.producer(inData)
	go dbw.consumer("")
	return nil
}
