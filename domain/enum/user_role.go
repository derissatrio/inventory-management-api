package enum

type UserRole string

const (
	RoleAdmin    UserRole = "admin"
	RoleEmployee UserRole = "employee"
)

func (r UserRole) IsValid() bool {
	switch r {
	case RoleAdmin, RoleEmployee:
		return true
	default:
		return false
	}
}