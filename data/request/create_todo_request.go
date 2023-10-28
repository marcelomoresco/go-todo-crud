package request

type CreateTodoRequest struct {
	Name string `validate"required, min=1,max=10" json:"name"`
	Description string `validate"required, min=1,max=200" json:"description"`
	Done bool `json:"done"`
}