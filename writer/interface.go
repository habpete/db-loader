package writer

import "db-loader/decls"

type DBInfo struct{}

type IWriter interface {
	Write(inData *decls.Incident) error
}
