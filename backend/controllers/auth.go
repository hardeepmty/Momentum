package controllers

import (
	"back/config"
	"back/models"
	"back/utils"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hashedPassword),
		Workspaces: []primitive.ObjectID{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result ,err := collectionUsers.InsertOne(r.Context(),newUser)
	if err != nil {
		utils.SendError(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	userID := result.InsertedID.(primitive.ObjectID)


	//creating a default workspace for the user as "My Tasks"
	collectionWorkspace := config.GetCollection("workspaces")
	defaultWorkspace := models.Workspace{
		UserID: result.InsertedID.(primitive.ObjectID),
		Name: "My Tasks",
		Tasks: []primitive.ObjectID{},
		CreatedAt: time.Now(),
		UpdatedAt:time.Now(),
	}

	workspaceResult , err := collectionWorkspace.InsertOne(r.Context(),defaultWorkspace)
	if err != nil {
		utils.SendError(w, "Failed to create deafult workspace", http.StatusInternalServerError)
		return 
	}

	workspaceID := workspaceResult.InsertedID.(primitive.ObjectID)

_, err = collectionUsers.UpdateByID(
		r.Context(),
		userID,
		bson.M{"$push": bson.M{"workspaces": workspaceID}},
	)
	if err != nil {
		utils.SendError(w, "Failed to link workspace to user", http.StatusInternalServerError)
		return
	}


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
			"id":        userID,
			"username":  newUser.Username,
			"email":     newUser.Email,
			"workspaces": []primitive.ObjectID{workspaceID},
		},
	}, http.StatusCreated)
}



func Login(w http.ResponseWriter, r *http.Request) {
	var req models.AuthRequest
	if err := utils.ParseBody(r, &req) ; err !=nil {
		utils.SendError(w, "invalid req", http.StatusBadRequest)
		return 
	}

	//find the user 
	collection := config.GetCollection("users")
	var user models.User
	err := collection.FindOne(r.Context(),map[string]interface{}{"email":req.Email}).Decode(&user)
	if err != nil {
		utils.SendError(w, "Invalid creds", http.StatusUnauthorized)
		return 
	}

	//compare the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(req.Password)); err != nil {
		utils.SendError(w, "invalid creds", http.StatusUnauthorized)
		return
	} 

	//generate the token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"exp": time.Now().Add(time.Hour*24).Unix() ,
	})

	tokenString, err := token.SignedString([]byte("hardeep"))

	if err != nil {
		utils.SendError(w, "cannot gen token", http.StatusInternalServerError)
		return 
	}

	//send the user as response
		utils.SendResponse(w, map[string]interface{}{
		"token": tokenString,
		"user": map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	}, http.StatusOK)
}