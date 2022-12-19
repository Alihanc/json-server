package helper

import (
	model "json-server/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	// we will query the data based on the ID
	id := c.Param("id")
	// and the data model to execute new data
	var updatedBunker model.Bunker
	// based on our dummy data we will loop through each and every data
	for i, singleBunker := range model.InitialData {
		// if the ID match with our reference ID
		if singleBunker.ID == id {
			// we will update all these parameters
			singleBunker.ID = updatedBunker.ID
			singleBunker.NAME = updatedBunker.NAME
			singleBunker.CATEGORY = updatedBunker.CATEGORY
			singleBunker.PLAQUE = updatedBunker.PLAQUE
			// after that we will bind this new DATA to OLD data
			if err := c.BindJSON(&updatedBunker); err != nil {
				// if there is an error we will panic
				println("Panic: Error Inserting NEW DATA")
				return
			}
			// than we will append all the data as JSON
			model.InitialData = append(model.InitialData[:i], updatedBunker)
			// it will give us response
			c.IndentedJSON(http.StatusCreated, updatedBunker)
		}
		// println("Updated DATA: ", singleBunker)
	}
}
