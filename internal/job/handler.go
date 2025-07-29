package job 

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type Handler struct{
	service *Service
}

func NewHandler (s *Service)*Handler{
	return &Handler{s}
}

func (h *Handler)RegisterRoutes(r *gin.RouterGroup){
	r.POST("/jobs",h.CreateJob)
	r.GET("/jobs", h.GetJobs)
	r.DELETE("/jobs/:id", h.DeleteJob)
}

func (h *Handler) CreateJob(c *gin.Context){
	var input CreateJobRequest
	if err :=c.ShouldBindJSON(&input); err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID := c.MustGet("userID").(string)
	job := &Job{
		Title:       input.Title,
		Description: input.Description,
		Company:     input.Company,
		UserID:      userID,
	}

		if err := h.service.CreateJob(job); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create job"})
		return
	}

	c.JSON(http.StatusCreated, job)
}

func (h *Handler) GetJobs(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	jobs, err := h.service.GetJobsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch jobs"})
		return
	}

	c.JSON(http.StatusOK, jobs)
}

func (h *Handler) DeleteJob(c *gin.Context) {
	jobIDStr := c.Param("id")
	jobID, err := uuid.Parse(jobIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	userID := c.MustGet("userID").(string)

	if err := h.service.DeleteJob(jobID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete job"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job deleted"})
}