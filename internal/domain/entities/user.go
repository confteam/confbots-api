package entities

type Role string

const (
	RoleMember     Role = "MEMBER"
	RoleAdmin      Role = "ADMIN"
	RoleSuperAdmin Role = "SUPERADMIN"
	RoleBanned     Role = "BANNED"
)

type User struct {
	ID   int
	TgId int64
}

type UserChannel struct {
	ID        int
	UserID    int
	ChannelID int
	role      Role
	anonimity bool
}
