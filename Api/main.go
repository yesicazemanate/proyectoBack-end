package main

import (
	"context"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type Alimento struct {
	ID string `json:"id"`
	Nombre string `json:"nombre"`	
	Description string `json:"description"`
}
var alimentos = []Alimento{}

func getAlimento(a *gin.Context){
	a.IndentedJSON(http.StatusOK, alimentos)
}
func postAlimento(a *gin.Context){
	var newAlimento Alimento
	if err:= a.BindJSON(&newAlimento); err!=nil{
		return
	}
	alimentos= append(alimentos, newAlimento )
	a.IndentedJSON(http.StatusCreated, alimentos)

}
func getAlimentoId(a *gin.Context){
id := a.Param("id")
for _, c:= range alimentos{
	if c.ID == id {
		a.IndentedJSON(http.StatusOK, c)
		return
	}
}
a.IndentedJSON(http.StatusNotFound, gin.H{"message": "data not found"})
}
 func deleteAlimento(a * gin.Context){
	id := a.Param("id")
	for i, c := range alimentos{
if c.ID == id {
	alimentos = append(alimentos[:i], alimentos[i+1:]...)
a.JSON(http.StatusOK, gin.H{"message": "alimento eliminado correctamente"})
return 
	}
}
a.JSON(http.StatusOK, gin.H{"message": "alimento no encontrado"})
 }

func main() {
	ctx:= context.TODO()
	alimentOption:= options.Client().ApplyURI("mongodb://localhost:27017")
	aliment, error := mongo.Connect(ctx, alimentOption)
	if error != nil{
		panic(error)
	}
	error= aliment.Ping(ctx, nil)
	if error !=nil{
		log.Fatal(error)
	}
	log.Println("conectado a MongoDB")


	router := gin.Default() 
	 //es la primera peticion que ya se le hizo al servicio 
	router.GET("/alimentos", getAlimento)
	 router.POST("/alimentos", postAlimento)
	 router.GET("/alimentos/:id", getAlimentoId)
	 router.DELETE("/alimentos/:id", deleteAlimento)
	 //esta es para correr el servidor
	router.Run("localhost:8080")
	

}