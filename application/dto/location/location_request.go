package location

type CreateLocationRequest struct {
	Name        string `json:"name" binding:"required"`
	Area        string `json:"area" binding:"required"`
	Description string `json:"description"`
	Capacity    int    `json:"capacity" binding:"min=0"`
}

type UpdateLocationRequest struct {
	Name        string `json:"name,omitempty"`
	Area        string `json:"area,omitempty"`
	Description string `json:"description,omitempty"`
	Capacity    *int   `json:"capacity,omitempty" binding:"omitempty,min=0"`
}