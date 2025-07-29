package job

type CreateJobRequest struct{
	Title string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
}