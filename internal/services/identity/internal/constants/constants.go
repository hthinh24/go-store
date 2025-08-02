package constants

type Gender string
type UserStatus string

const (
	MALE   Gender = "MALE"
	FEMALE Gender = "FEMALE"
	OTHER  Gender = "OTHER"
)

const (
	ACTIVE   UserStatus = "ACTIVE"
	INACTIVE UserStatus = "INACTIVE"
	DELETED  UserStatus = "DELETED"
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
