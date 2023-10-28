package response

type TodoResponse struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Done bool `json:"done"`
}