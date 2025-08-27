package applications

type ApplicationRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}