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

func CreateWorkSpace(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(string)
	if !ok {
		utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userObjID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		utils.SendError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var workspaceReq models.Workspace
	if err := utils.ParseBody(r, &workspaceReq); err != nil {
		utils.SendError(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	if workspaceReq.Name == "" {
		utils.SendError(w, "Workspace name is required", http.StatusBadRequest)
		return
	}

	newWorkSpace := models.Workspace{
		UserID:    userObjID,
		Name:      workspaceReq.Name,
		Tasks:     []primitive.ObjectID{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	workspaceCollection := config.GetCollection("workspaces")
	result, err := workspaceCollection.InsertOne(r.Context(), newWorkSpace)
	if err != nil {
		utils.SendError(w, "Failed to create workspace", http.StatusInternalServerError)
		return
	}

	userCollection := config.GetCollection("users")
	_, err = userCollection.UpdateOne(
		r.Context(),
		bson.M{"_id": userObjID},
		bson.M{"$push": bson.M{"workspaces": result.InsertedID}},
	)
	if err != nil {
		utils.SendError(w, "Failed to link workspace to user", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"id":        result.InsertedID,
		"name":      newWorkSpace.Name,
		"userId":    newWorkSpace.UserID,
		"createdAt": newWorkSpace.CreatedAt,
	}

	utils.SendResponse(w, response, http.StatusCreated)
}