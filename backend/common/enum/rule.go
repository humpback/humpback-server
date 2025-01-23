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
)

var RuleLength = map[string]LengthLimit{
	"LimitUsername":    LimitUsername,
	"LimitTeamName":    LimitTeamName,
	"LimitEmail":       LimitEmail,
	"LimitPassword":    LimitPassword,
	"LimitPhone":       LimitPhone,
	"LimitDescription": LimitDescription,
}

var (
	RegularEmail = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	RegularPhone = `^\d+$`
)

var RuleFormat = map[string]string{
	"RegularEmail": RegularEmail,
	"RegularPhone": RegularPhone,
}
