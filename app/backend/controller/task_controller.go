package controller

import (
	"context"
	"net/http"

	"github.com/MouslyCode/three-of-a-perfect-pair/backend/database"
	"github.com/MouslyCode/three-of-a-perfect-pair/backend/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateTask(ctx *gin.Context) {
	var input model.Request
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	task := model.Task{
		Title:     input.Title,
		Completed: false,
	}

	response, err := database.TaskCollection.InsertOne(context.TODO(), task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Task Created Succesfully",
		"Data":    response,
	})

}

func UpdateTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := bson.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Id"})
		return
	}

	var input model.Request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": bson.M{
			"title": input.Title,
		},
	}

	result, err := database.TaskCollection.UpdateOne(
		context.TODO(),
		filter,
		update,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task Updated",
		"data":    result,
	})

}

func DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := bson.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Id"})
		return
	}

	filter := bson.M{"_id": id}

	result, err := database.TaskCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task Deleted",
	})
}

func GetTasks(ctx *gin.Context) {
	cursor, err := database.TaskCollection.Find(ctx.Request.Context(), bson.D{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx.Request.Context())

	var tasks []model.Task
	if err = cursor.All(ctx.Request.Context(), &tasks); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tasks)

}

func CompletedTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := bson.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Id"})
		return
	}

	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": bson.M{
			"completed": true,
		},
	}

	result, err := database.TaskCollection.UpdateOne(
		context.TODO(),
		filter,
		update,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task Completed",
		"data":    result,
	})
}
