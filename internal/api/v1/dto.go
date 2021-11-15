package v1

type UserDTO struct {
	ID       int    `json:"id,string,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	IsActive bool   `json:"is_active,string,omitempty"`
	Group    int    `json:"group,string,omitempty"`
}

type GroupDTO struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	IsActive bool   `json:"is_active,string,omitempty"`
}

type UsersDTO struct {
	Users []UserDTO
}
