package database

func (u *User) GetID() int {
	return int(u.ID)
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetCreatedAt() string {
	return u.CreatedAt
}

func (u *User) GetUpdatedAt() string {
	return u.UpdatedAt
}

func (u *User) GetApiKey() string {
	return u.ApiKey
}
