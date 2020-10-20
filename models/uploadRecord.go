package models

import (
	""
)

type UploadRecord struct {
	Id  int
	FileName  string
	FileCert string
	FileTitle string
	CertTime int
	Phone string
}

func (u UploadRecord) SaveRecord()(int64,error){
	rs , err:=
}
