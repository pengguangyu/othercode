package conf

// 是否整个数据库的表全量生成
const ModelAllTables = false
const ModelTablesList = "lian_topnews_feed_recommend,lian_user_detail"

// model保存路径
const ModelPath = "model/"

// 是否覆盖已存在model
const ModelReplace = true

// 数据库驱动
const DriverName = "mysql"

type DbConf struct {
	Host   string
	Port   string
	User   string
	Pwd    string
	DbName string
}

// 数据库链接配置
var MasterDbConfig = DbConf{
	Host:   "172.16.40.211",
	Port:   "3306",
	User:   "root",
	Pwd:    "123456",
	DbName: "mytestdb",
}
