package logger


type Category string
type SubCategory string
type ExtraKey string

const (
	General Category = "General"
	Postgres Category = "Postgres"
	Redis Category = "Redis"
	Validation Category = "Validation"
	RequestResponse Category = "RequestResponse"
	IO Category = "IO"
)

const (
		// General
		Startup         SubCategory = "Startup"
		ExternalService SubCategory = "ExternalService"

		// Postgres
		Migration SubCategory = "Migration"
		Select    SubCategory = "Select"
		Rollback  SubCategory = "Rollback"
		Update    SubCategory = "Update"
		Delete    SubCategory = "Delete"
		Insert    SubCategory = "Insert"
		Connect   SubCategory = "Connect"

		// Internal
		Api	SubCategory = "Api"
	)

const (
		AppName      ExtraKey = "AppName"
		LoggerName   ExtraKey = "Logger"
		ClientIp     ExtraKey = "ClientIp"
		HostIp       ExtraKey = "HostIp"
		Method       ExtraKey = "Method"
		StatusCode   ExtraKey = "StatusCode"
		BodySize     ExtraKey = "BodySize"
		Path         ExtraKey = "Path"
		Latency      ExtraKey = "Latency"
		RequestBody  ExtraKey = "RequestBody"
		ResponseBody ExtraKey = "ResponseBody"
		ErrorMessage ExtraKey = "ErrorMessage"
)