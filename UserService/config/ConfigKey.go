package config

type ConfigKey struct {
	key string
}

var MAIN_SERVICE_HOST ConfigKey = ConfigKey{key: "MainServiceHost"}

var MAIN_SERVICE_PORT ConfigKey = ConfigKey{key: "MainServicePort"}

var USER_SERVICE_HOST ConfigKey = ConfigKey{key: "UserServiceHost"}

var USER_SERVICE_PORT ConfigKey = ConfigKey{key: "UserServicePort"}

var HMAC_KEY ConfigKey = ConfigKey{key: "HMAC_KEY"}

var MYSQL_USERNAME ConfigKey = ConfigKey{key: "MysqlUsername"}

var MYSQL_PASSWORD ConfigKey = ConfigKey{key: "MysqlPassword"}

var MYSQL_HOST ConfigKey = ConfigKey{key: "MysqlHost"}

var MYSQL_PORT ConfigKey = ConfigKey{key: "MysqlPort"}

var MYSQL_DATABASE ConfigKey = ConfigKey{key: "MysqlDatabase"}
