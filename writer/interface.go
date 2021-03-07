package writer

import "db-loader/decls"

type IWriter interface {
	Write(inData *decls.Incident) error
}
