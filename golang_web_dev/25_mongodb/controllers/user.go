package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/j1mmyson/mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	client *mongo.Client
}

func NewUserController(c *mongo.Client) *UserController {
	return &UserController{c}
}

func (uc UserController) GetUser(c *gin.Context) {
	doc := uc.client.Database("test").Collection("users").FindOne(context.TODO(), bson.M{})

	var u models.User
	doc.Decode(&u)

	c.JSON(http.StatusOK, u)
}

func (uc UserController) CreateUser(c *gin.Context) {

	var u models.User

	if err := c.ShouldBindJSON(&u); err != nil {

	}
	uc.client.Database("test").Collection("users").InsertOne(context.Background(), u)

	fmt.Println(u)

	c.JSON(http.StatusOK, u)
}

func (uc UserController) DeleteUser(c *gin.Context) {

	var u models.User

	res, err := uc.client.Database("test").Collection("users").DeleteOne(context.Background(), u)

	if err != nil {

	}
	fmt.Println(res)

	c.JSON(http.StatusOK, gin.H{"message": "delete user"})
}
