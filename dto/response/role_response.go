package response

type RoleResponse struct {
	Id          uint   `json:"id"`
	Role        string `json:"role"`
	Description string `json:"description"`
}
