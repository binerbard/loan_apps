package setting

import (
	utils_setting "loan_apps/scripts/setting"
)

// Config : Struct for setting
type Config struct {
	AppService              string
	AppHost                 string
	AppPort                 string
	GRPCHost                string
	GRPCPort                string
	APIPrefix               string
	APIVersion              string
	APIDirectory            string
	DBUser                  string
	DBPassword              string
	DBDatabase              string
	DBHost                  string
	DBPort                  string
	DBDriver                string
	DOCUser                 string
	DOCPassword             string
	DOCDatabase             string
	DCHost                  string
	DOCPort                 string
	DOCDriver               string
	ExpiredToken            int
	TokenLength             int
	JwtSecretKey            string
	JwtExpiredPerSession    int
	JwtRefreshExpiredPerDay int
	JwtAudience             string
	JwtIssuer               string
	LimitAccountDevices     int
	smtpHost                string
	smtpPort                int
	smtpEmail               string
	smtpPassword            string
	smtpEmailSender         string
	smtpNameSender          string
}


// GetConfig : Config function
func GetConfig() Config {
	utils_setting.LoadEnv()

	return Config{
		AppService:              utils_setting.GetEnv("APP_SERVICE", ""),
		AppHost:                 utils_setting.GetEnv("API_HOST", ""),
		AppPort:                 utils_setting.GetEnv("API_PORT", ""),
		APIPrefix:               utils_setting.GetEnv("API_PREFIX", ""),
		GRPCHost:                utils_setting.GetEnv("GRPC_HOST", ""),
		GRPCPort:                utils_setting.GetEnv("GRPC_PORT", ""),
		APIVersion:              utils_setting.GetEnv("API_VERSION", ""),
		APIDirectory:            utils_setting.GetEnv("API_PATH", ""),
		DBUser:                  utils_setting.GetEnv("DB_USER", ""),
		DBPassword:              utils_setting.GetEnv("DB_PASSWORD", ""),
		DBDatabase:              utils_setting.GetEnv("DB_DATABASE", ""),
		DBHost:                  utils_setting.GetEnv("DB_HOST", ""),
		DBPort:                  utils_setting.GetEnv("DB_PORT", ""),
		DBDriver:                utils_setting.GetEnv("DB_DRIVER", ""),
		DOCUser:                 utils_setting.GetEnv("DOC_USER", ""),
		DOCPassword:             utils_setting.GetEnv("DOC_PASSWORD", ""),
		DOCDatabase:             utils_setting.GetEnv("DOC_DATABASE", ""),
		DCHost:                  utils_setting.GetEnv("DOC_HOST", ""),
		DOCPort:                 utils_setting.GetEnv("DOC_PORT", ""),
		DOCDriver:               utils_setting.GetEnv("DOC_DRIVER", ""),
		ExpiredToken:            utils_setting.CastInt(utils_setting.GetEnv("EXPIRED_TOKEN", "")),
		TokenLength:             utils_setting.CastInt(utils_setting.GetEnv("TOKEN_LENGTH", "")),
		JwtSecretKey:            utils_setting.GetEnv("JWT_SECRET_KEY", ""),
		JwtExpiredPerSession:    utils_setting.CastInt(utils_setting.GetEnv("JWT_EXPIRED_PER_SESSION", "")),
		JwtRefreshExpiredPerDay: utils_setting.CastInt(utils_setting.GetEnv("JWT_REFRESH_EXPIRED_PER_DAY", "")),
		JwtAudience:             utils_setting.GetEnv("JWT_AUDIENCE", ""),
		JwtIssuer:               utils_setting.GetEnv("JWT_ISSUER", ""),
		LimitAccountDevices:     utils_setting.CastInt(utils_setting.GetEnv("LIMIT_ACCOUNT_DEVICES", "")),
		smtpHost:                utils_setting.GetEnv("SMTP_HOST", ""),
		smtpPort:                utils_setting.CastInt(utils_setting.GetEnv("SMTP_PORT", "")),
		smtpEmail:               utils_setting.GetEnv("SMTP_EMAIL", ""),
		smtpPassword:            utils_setting.GetEnv("SMTP_PASSWORD", ""),
		smtpEmailSender:         utils_setting.GetEnv("SMTP_EMAIL_SENDER", ""),
		smtpNameSender:          utils_setting.GetEnv("SMTP_NAME_SENDER", ""),
	}
}

// AppService : App service
var AppService = GetConfig().AppService

// APIPathPrefix : API path
var APIPathPrefix = "/" + GetConfig().APIPrefix + "/" + GetConfig().APIVersion + "/" + GetConfig().APIDirectory

// APILisen : API listen
var APILisen = GetConfig().AppHost + ":" + GetConfig().AppPort

// DatabaseURL : Database URL
var DatabaseURL = GetConfig().DBDriver + "://" + GetConfig().DBUser + ":" + GetConfig().DBPassword + "@" + GetConfig().DBHost + ":" + GetConfig().DBPort + "/" + GetConfig().DBDatabase

// DocumentURL : Document URL
var DocumentURL = GetConfig().DOCDriver + "://" + GetConfig().DOCUser + ":" + GetConfig().DOCPassword + "@" + GetConfig().DCHost + ":" + GetConfig().DOCPort + "/"

// DatabaseDNS : Database DNS
var DatabaseDNS = "host=" + GetConfig().DBHost + " user=" + GetConfig().DBUser + " dbname=" + GetConfig().DBDatabase + " password=" + GetConfig().DBPassword + " port=" + GetConfig().DBPort + " sslmode=disable"

// DocumentDNS : Document DNS
var DocumentDNS = "host=" + GetConfig().DCHost + " user=" + GetConfig().DOCUser + " dbname=" + GetConfig().DOCDatabase + " password=" + GetConfig().DOCPassword + " port=" + GetConfig().DOCPort

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

// DOCDatabase is the name of document database
var DOCDatabase = GetConfig().DOCDatabase

// GRPCHost is the host of gRPC server
var GRPCHost = GetConfig().GRPCHost

// GRPCPort is the port of gRPC server
var GRPCPort = GetConfig().GRPCPort

// GRPCUri is the uri of gRPC server
var GRPCUri = GRPCHost + ":" + GRPCPort

// LimitAccountDevices is the limit of account devices
var LimitAccountDevices = GetConfig().LimitAccountDevices

// SMTPHost is the host of SMTP
var SMTPHost = GetConfig().smtpHost

// SMTPPort is the port of SMTP
var SMTPPort = GetConfig().smtpPort

// SMTPEmail is the email of SMTP
var SMTPEmail = GetConfig().smtpEmail

// SMTPPassword is the password of SMTP
var SMTPPassword = GetConfig().smtpPassword

// SMTPEmailSender is the email sender of SMTP
var SMTPEmailSender = GetConfig().smtpEmailSender

// SMTPNameSender is the name sender of SMTP
var SMTPNameSender = GetConfig().smtpNameSender