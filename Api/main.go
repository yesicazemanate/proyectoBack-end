package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)


 var mongoAliment *mongo.Client

func init(){
	if err := conectMongo(); err != nil{
		log.Fatal("no se pudo conectar a MongoDB")
	}
	println("conexion establecida")
}
func main() {
	router := gin.Default() 
	 //es la primera peticion que ya se le hizo al servicio 
	router.GET("/alimentos", getAlimento)
	router.POST("/alimentos/aggregate", postAlimento)
	router.GET("/alimentos/:id", getAlimentoId)
	router.DELETE("/alimentos/:id", deleteAlimento)
	 //esta es para correr el servidor
	router.Run("localhost:8080")
	

}
func conectMongo() error{
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	ctx:= context.TODO()
	alimentOption:= options.Client().ApplyURI("mongodb://localhost:27017").SetServerAPIOptions(serverAPI)
	aliment, error := mongo.Connect(ctx, alimentOption)
	if error != nil{
		panic(error)
	}
	error = aliment.Ping(context.TODO(),nil)
	mongoAliment=aliment
	return error
 }
 
func getAlimento(a *gin.Context){
	cursor, error := mongoAliment.Database("Restaurante").Collection("Alimento").Find(context.TODO(), bson.D{{}})
	if error!=nil{
		a.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}
	var alimentos []bson.M
	if error = cursor.All(context.TODO(), &alimentos); error !=nil{
		a.JSON(http.StatusInternalServerError, gin.H{"error": error.Error})
		return
	}
	a.JSON(http.StatusOK, alimentos)
}
func postAlimento(a *gin.Context){
	var aliment map[string]interface{}
	if error := a.ShouldBindJSON(&aliment); error != nil{
		a.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		if _, error := mongoAliment.Database("Restaurante").Collection("Alimento").Aggregate(context.TODO(), aliment); error != nil {
			a.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
			return
		}
		a.JSON(http.StatusOK, aliment)
	}
}
func getAlimentoId(a *gin.Context){
idstr := a.Param("id")
id , error := primitive.ObjectIDFromHex(idstr)
if error != nil{
	a.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	return
}
var alimento bson.M
error = mongoAliment.Database("Restaurante").Collection("Alimento").FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&alimento)
if error != nil{
	a.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
	return
} 
a.JSON(http.StatusOK, alimento)
}
 func deleteAlimento(a * gin.Context){
	idstr := a.Param("id")
id , error := primitive.ObjectIDFromHex(idstr)
if error != nil{
	a.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	return
}
result ,error := mongoAliment.Database("Restaurante").Collection("Alimento").DeleteOne(context.TODO(), bson.D{{"_id", id}})
if error != nil{
	a.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
	return
} 
if result.DeletedCount==0{
	a.JSON(http.StatusNotFound, gin.H{"message":"alimento no encontrado"})
return
}
a.JSON(http.StatusOK, gin.H{"message": "alimento borrado exitosamente"})
 }