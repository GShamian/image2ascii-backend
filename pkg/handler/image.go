package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qeesung/image2ascii/convert"
)

var defaultOptions = convert.Options{
	Ratio:           1,
	FixedWidth:      -1,
	FixedHeight:     -1,
	FitScreen:       true,
	Colored:         false,
	Reversed:        false,
	StretchedScreen: false,
}

func (h *Handler) createASCII(c *gin.Context) {
	file, err := c.FormFile("myFile")
	if err != nil {
		fmt.Println("dasdadas")
		log.Fatal(err)
	}
	log.Println(file.Filename)

	err = c.SaveUploadedFile(file, "storage/"+file.Filename)
	if err != nil {
		log.Fatal(err)
	}
	converter := convert.NewImageConverter()
	resp := converter.ImageFile2ASCIIString("storage/"+file.Filename, &defaultOptions)
	c.String(http.StatusOK, resp)
}
