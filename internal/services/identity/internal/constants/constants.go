package constants

type Gender string
type Status string

const (
	MALE   string = "MALE"
	FEMALE string = "FEMALE"
	OTHER  string = "OTHER"
)

const (
	ACTIVE   string = "ACTIVE"
	INACTIVE string = "INACTIVE"
	DELETED  string = "DELETED"
)

func GetGender(s string) string {
	switch s {
	case string(MALE):
		return string(MALE)
	case string(FEMALE):
		return string(FEMALE)
	default:
		return string(OTHER)
	}
}

func GetStatus(s string) string {
	switch s {
	case string(ACTIVE):
		return string(ACTIVE)
	case string(DELETED):
		return string(DELETED)
	default:
		return string(INACTIVE)
	}
}
