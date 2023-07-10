package unsplash

type Photo struct {
	ID          string `json:"id"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Description string `json:"description"`
	URLs        struct {
		Thumb   string `json:"thumb"`
		Small   string `json:"small"`
		Regular string `json:"regular"`
		Full    string `json:"full"`
		Raw     string `json:"raw"`
	} `json:"urls"`
	CurrentUserCollections []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"current_user_collections"`
}

type SearchResponse struct {
	Total      int     `json:"total"`
	TotalPages int     `json:"total_pages"`
	Results    []Photo `json:"results"`
}
