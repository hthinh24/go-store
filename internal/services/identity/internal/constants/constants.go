package constants

type Gender string
type UserStatus string

const (
	GenderMale   Gender = "MALE"
	GenderFemale Gender = "FEMALE"
	GenderOther  Gender = "OTHER"
)

const (
	UserStatusActive   UserStatus = "ACTIVE"
	UserStatusInactive UserStatus = "INACTIVE"
	UserStatusDeleted  UserStatus = "DELETED"
)

func GetGender(s string) string {
	switch s {
	case string(GenderMale):
		return string(GenderMale)
	case string(GenderFemale):
		return string(GenderFemale)
	default:
		return string(GenderOther)
	}
}

func GetStatus(s string) string {
	switch s {
	case string(UserStatusActive):
		return string(UserStatusActive)
	case string(UserStatusDeleted):
		return string(UserStatusDeleted)
	default:
		return string(UserStatusInactive)
	}
}
