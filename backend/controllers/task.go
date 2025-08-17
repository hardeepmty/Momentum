package controllers

import (
	"back/config"
	"back/models"
	"back/utils"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	userId , ok := r.Context().Value("userId").(string)
	if !ok {
		utils.SendError(w,"invalid user id", http.StatusBadRequest)
		return 
	}

	//convert to object id
	userObjId , err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		utils.SendError(w, "invalid user id", http.StatusBadRequest)
		return 
	}

	//parsing the task req
	var taskReq models.Task
	if err := utils.ParseBody(r, &taskReq) ;err != nil {
		utils.SendError(w, "invalid task request", http.StatusBadRequest)
		return 
	}

	//validating the req fields
	if taskReq.Title == "" {
		utils.SendError(w, "title is req", http.StatusBadRequest)
		return
	}

	if taskReq.Priority == ""{
		//set it to med
		taskReq.Priority= "medium"
	}

	//validating the priority rule
	validPriorities := map[string]bool{"low": true, "medium": true, "high": true}
	if !validPriorities[taskReq.Priority] {
		utils.SendError(w, "Invalid priority value", http.StatusBadRequest)
		return
	}

	newTask := models.Task{
		UserID:      userObjId,
		WorkspaceID: taskReq.WorkspaceID,
		Title:       taskReq.Title,
		Description: taskReq.Description,
		Tag:         taskReq.Tag,
		Completed:   false, 
		Priority:    taskReq.Priority,
		DueDate:     taskReq.DueDate,
		Subtasks:    taskReq.Subtasks,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	taskCollection := config.GetCollection("tasks")
	result , err := taskCollection.InsertOne(r.Context(),newTask)
	if err != nil {
		utils.SendError(w, "failed tp create task", http.StatusInternalServerError)
		return 
	}

	//if task has workspace id we need to add it to workspace
	if taskReq.WorkspaceID != nil {
		workspaceCollection := config.GetCollection("workspaces")
		_, err := workspaceCollection.UpdateOne(
			r.Context(),
			bson.M{"_id":taskReq.WorkspaceID , "userId":userObjId},
			bson.M{"$push":bson.M{"tasks":result.InsertedID}},
		)
		if err != nil {
			utils.SendError(w, "Failed to link task to workspace", http.StatusInternalServerError)
			return
		}
	}

	response := map[string]interface{}{
		"id":          result.InsertedID,
		"title":       newTask.Title,
		"description": newTask.Description,
		"tag":         newTask.Tag,
		"completed":   newTask.Completed,
		"priority":    newTask.Priority,
		"dueDate":     newTask.DueDate,
		"workspaceId": newTask.WorkspaceID,
		"createdAt":   newTask.CreatedAt,
	}

	utils.SendResponse(w, response, http.StatusCreated)
}


func GetUserTasks (w http.ResponseWriter, r *http.Request) {
	userId , ok := r.Context().Value("userId").(string)
	if !ok {
		utils.SendError(w,"invalid user id", http.StatusBadRequest)
		return 
	}

	userObjId , err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		utils.SendError(w, "invalid user id", http.StatusBadRequest)
		return 
	}

	collection := config.GetCollection("tasks")
	cursor , err := collection.Find(r.Context(), bson.M{"userId":userObjId})
	if err != nil {
		utils.SendError(w, "unable to get tasks", http.StatusInternalServerError)
		return 
	}

	var tasks []models.Task
	if err := cursor.All(r.Context(), &tasks) ; err != nil {
		utils.SendError(w, "failed to fetch tasks", http.StatusInternalServerError)
		return 
	}

	utils.SendResponse(w, tasks, http.StatusOK)

}