package controllers

import (
	"back/config"
	"back/models"
	"back/utils"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var req models.SignupRequest

	err := utils.ParseBody(r, &req) 
	if err != nil {
		utils.SendError(w, "Inavlid req format", http.StatusBadRequest)
		return 
	}

	collectionUsers := config.GetCollection("users") 
	var existingUsers models.User
	err = collectionUsers.FindOne(r.Context(), map[string]interface{}{"email":req.Email}).Decode(&existingUsers)
	if err == nil {
		utils.SendError(w, "user exits", http.StatusConflict)
		return 
	}

	//pwd hashing
	hashedPassword , err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.SendError(w, "failed to hash pwd", http.StatusInternalServerError)
		return 
	}

	//create the user 
	newUser := models.User{
		Username: req.Username,
		Email: req.Email,
		Password: string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result ,err := collectionUsers.InsertOne(r.Context(),newUser)
	if err != nil {
		utils.SendError(w, "Failed to create user", http.StatusInternalServerError)
		return
	}


	// //Creating a default workspace for the user as "My Tasks"
	// collectionWorkspace := config.GetCollection("workspaces")
	// defaultWorkspace := models.Workspace{
	// 	UserID: result.InsertedID.(primitive.ObjectID),
	// 	Name: "My Tasks",
	// 	Tasks: []primitive.ObjectID{},
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt:time.Now(),
	// }

	// workspaceResult , err := collectionWorkspace.InsertOne(r.Context(),defaultWorkspace)
	// if err != nil {
	// 	utils.SendError(w, "Failed to create deafult workspace", http.StatusInternalServerError)
	// 	return 
	// }

	// _, err = collectionUsers.UpdateByID(
	// 	r.Context(),
	// 	result.InsertedID,
	// 	bson.M{
	// 	"$push": bson.M{
	// 		"workspaces": []primitive.ObjectID{workspaceResult.InsertedID.(primitive.ObjectID)},
	// 	},
	// },
	// )

	// if err != nil {
	// 	utils.SendError(w, "failed to link the Workspace", http.StatusInternalServerError)
	// 	return
	// }

	//generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": result.InsertedID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte("hardeep")) 
	if err != nil {
		utils.SendError(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	//response of user creadeted
	utils.SendResponse(w, map[string]interface{}{
		"token": tokenString,
		"user": map[string]interface{}{
			"id":       result.InsertedID,
			"username": newUser.Username,
			"email":    newUser.Email,
			//"workspace": workspaceResult.InsertedID,
		},
	}, http.StatusCreated)
}



