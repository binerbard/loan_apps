package setting

import (
	utils_setting "loan_apps/scripts/setting"
)

// Config : Struct for setting
type Config struct {
	AppService              string
	AppHost                 string
	AppPort                 string
	APIPrefix               string
	APIVersion              string
	APIDirectory            string
	DBUser                  string
	DBPassword              string
	DBDatabase              string
	DBHost                  string
	DBPort                  string
	DBDriver                string
	StorageAccessKey            string
	StorageSecretKey        string
	StorageHost            string
	StoragePort            string
	StorageBucket          string
	ExpiredToken           int
	TokenLength            int
	JwtSecretKey           string
	JwtExpiredPerSession   int
	JwtRefreshExpiredPerDay int
	JwtAudience            string
	JwtIssuer              string
}



// GetConfig : Config function
func GetConfig() Config {
	utils_setting.LoadEnv()

	return Config{
		AppService:              utils_setting.GetEnv("APP_SERVICE", ""),
		AppHost:                 utils_setting.GetEnv("API_HOST", ""),
		AppPort:                 utils_setting.GetEnv("API_PORT", ""),
		APIPrefix:               utils_setting.GetEnv("API_PREFIX", ""),
		APIVersion:              utils_setting.GetEnv("API_VERSION", ""),
		APIDirectory:            utils_setting.GetEnv("API_PATH", ""),
		DBUser:                  utils_setting.GetEnv("DB_USER", ""),
		DBPassword:              utils_setting.GetEnv("DB_PASSWORD", ""),
		DBDatabase:              utils_setting.GetEnv("DB_DATABASE", ""),
		DBHost:                  utils_setting.GetEnv("DB_HOST", ""),
		DBPort:                  utils_setting.GetEnv("DB_PORT", ""),
		DBDriver:                utils_setting.GetEnv("DB_DRIVER", ""),
		StorageAccessKey:             utils_setting.GetEnv("STORAGE_ACCESSKEY", ""),
		StorageSecretKey:          utils_setting.GetEnv("STORAGE_SECRETKEY", ""),
		StorageHost:             utils_setting.GetEnv("STORAGE_HOST", ""),
		StoragePort:             utils_setting.GetEnv("STORAGE_PORT", ""),
		StorageBucket:           utils_setting.GetEnv("STORAGE_BUCKET", ""),
		ExpiredToken:            utils_setting.CastInt(utils_setting.GetEnv("EXPIRED_TOKEN", "")),
		TokenLength:             utils_setting.CastInt(utils_setting.GetEnv("TOKEN_LENGTH", "")),
		JwtSecretKey:            utils_setting.GetEnv("JWT_SECRET_KEY", ""),
		JwtExpiredPerSession:    utils_setting.CastInt(utils_setting.GetEnv("JWT_EXPIRED_PER_SESSION", "")),
		JwtRefreshExpiredPerDay: utils_setting.CastInt(utils_setting.GetEnv("JWT_REFRESH_EXPIRED_PER_DAY", "")),
		JwtAudience:             utils_setting.GetEnv("JWT_AUDIENCE", ""),
		JwtIssuer:               utils_setting.GetEnv("JWT_ISSUER", ""),
	}
}

// AppService : App service
var AppService = GetConfig().AppService

// APIPathPrefix : API path
var APIPathPrefix = "/" + GetConfig().APIPrefix + "/" + GetConfig().APIVersion + "/" + GetConfig().APIDirectory

// APILisen : API listen
var APILisen = GetConfig().AppHost + ":" + GetConfig().AppPort

// DatabaseURL : Database URL
var DatabaseURL = GetConfig().DBUser + ":" + GetConfig().DBPassword + "@tcp(" + GetConfig().DBHost + ":" + GetConfig().DBPort + ")/" + GetConfig().DBDatabase + "?parseTime=true"

// DatabaseDNS : Database DNS
var DatabaseDNS = "host=" + GetConfig().DBHost + " user=" + GetConfig().DBUser + " dbname=" + GetConfig().DBDatabase + " password=" + GetConfig().DBPassword + " port=" + GetConfig().DBPort + " sslmode=disable"

var StoragePathPrefix = GetConfig().StorageHost + ":" + GetConfig().StoragePort

var StorageAccessKey = GetConfig().StorageAccessKey

var StorageSecretKey = GetConfig().StorageSecretKey

var StorageBucket = GetConfig().StorageBucket

// ExpiredToken : Expired token
var ExpiredToken = GetConfig().ExpiredToken

// TokenLength : Token length
var TokenLength = GetConfig().TokenLength

// JwtSecretKey is the secret key for JWT token
var JwtSecretKey = []byte(GetConfig().JwtSecretKey)

// JwtExpiredPerSession is the expired time of JWT token per session
var JwtExpiredPerSession = GetConfig().JwtExpiredPerSession

// JwtRefreshExpiredPerDay is the expired time of JWT refresh token per day
var JwtRefreshExpiredPerDay = GetConfig().JwtRefreshExpiredPerDay

// JwtAudience is the audience of JWT token
var JwtAudience = GetConfig().JwtAudience

// JwtIssuer is the issuer of JWT token
var JwtIssuer = GetConfig().JwtIssuer
