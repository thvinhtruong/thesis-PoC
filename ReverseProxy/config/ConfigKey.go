package config

type ConfigKey struct {
	key string
}

var MAIN_SERVICE_HOST ConfigKey = ConfigKey{key: "MainServiceHost"}

var MAIN_SERVICE_PORT ConfigKey = ConfigKey{key: "MainServicePort"}

var USER_SERVICE_HOST ConfigKey = ConfigKey{key: "UserServiceHost"}

var USER_SERVICE_PORT ConfigKey = ConfigKey{key: "UserServicePort"}

var STUDY_SERVICE_HOST ConfigKey = ConfigKey{key: "StudyServiceHost"}

var STUDY_SERVICE_PORT ConfigKey = ConfigKey{key: "StudyServicePort"}

var HMAC_KEY ConfigKey = ConfigKey{key: "HMAC_KEY"}
