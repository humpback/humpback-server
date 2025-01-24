package enum

type LengthLimit struct {
	Min int `json:"Min"`
	Max int `json:"Max"`
}

var (
	LimitUsername    = LengthLimit{Min: 2, Max: 100}
	LimitTeamName    = LengthLimit{Min: 2, Max: 100}
	LimitEmail       = LengthLimit{Min: 0, Max: 200}
	LimitPassword    = LengthLimit{Min: 8, Max: 20}
	LimitPhone       = LengthLimit{Min: 0, Max: 11}
	LimitDescription = LengthLimit{Min: 0, Max: 500}
	LimitConfigName  = LengthLimit{Min: 1, Max: 100}
	LimitConfigValue = LengthLimit{Min: 1, Max: 10000}
)

var RuleLength = map[string]LengthLimit{
	"Username":    LimitUsername,
	"TeamName":    LimitTeamName,
	"Email":       LimitEmail,
	"Password":    LimitPassword,
	"Phone":       LimitPhone,
	"Description": LimitDescription,
	"ConfigName":  LimitConfigName,
	"ConfigValue": LimitConfigValue,
}

var (
	RegularEmail = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	RegularPhone = `^\d+$`
)

var RuleFormat = map[string]string{
	"Email": RegularEmail,
	"Phone": RegularPhone,
}
