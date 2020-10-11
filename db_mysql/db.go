package db_mysql

import "database/sql"
var Db *sql.DB
func ConnectDB()  {
	//config := beego.AppConfig
	//dbDriver := config.String("db_driverName")
	//dbUser := config.String("db_user")
	//dbPassword := config.String("db_password")
	//dbIp := config.String("db_ip")
	//dbName := config.String("db_name")
	//fmt.Println(dbDriver,dbUser,dbPassword,dbName)
	//
	//connUrl := dbUser +":" + dbPassword + "@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	//
	//db, err := sql.Open(dbDriver,connUrl)
	//
	//if err != nil {
	//	panic("数据库连接错误，请检查配置")
	//}
    //     Db = db
}
