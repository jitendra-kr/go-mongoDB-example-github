package internal

import (
	"context"
	"log"
	"net/http"

	httpresponse "github.com/go-mongoDB-example-github/httpResponse"
	"github.com/go-mongoDB-example-github/models"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Name  string
	Email string
}

var userCollection = "user"

func (s *Server) createUser(c echo.Context) error {

	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return httpresponse.InternalServerError(c)
	}

	insertResult, err := s.Store.Collection(userCollection).InsertOne(context.TODO(), user)
	if err != nil {
		return httpresponse.InternalServerError(c)
	}
	return c.JSON(http.StatusCreated, models.CreateUserResp{
		Id: insertResult.InsertedID.(primitive.ObjectID).Hex(),
	})
}

func (s *Server) getUserByID(c echo.Context) error {
	var result models.User
	userID := c.Param("id")

	_id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Printf("failed to convert ObjectIDFromHex")
		return httpresponse.InternalServerError(c)
	}

	filter := bson.M{"_id": _id}
	err = s.Store.Collection(userCollection).FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return httpresponse.RecordNotFound(c)
		}
		log.Printf("failed to read record from DB")
		return httpresponse.InternalServerError(c)
	}
	return c.JSON(http.StatusOK, result)
}

func (s *Server) deleteUserByID(c echo.Context) error {

	userID := c.Param("id")

	_id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Printf("failed to convert ObjectIDFromHex")
		return httpresponse.InternalServerError(c)
	}

	filter := bson.M{"_id": _id}
	result, err := s.Store.Collection(userCollection).DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Printf("failed to delete record from DB")
		return httpresponse.InternalServerError(c)
	}

	if result.DeletedCount == 0 {
		return httpresponse.BadRequest(c, "Invalid userID")

	}
	return c.JSON(http.StatusNoContent, result)
}
