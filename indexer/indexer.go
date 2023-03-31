package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"zse/dataProcessing"
	"zse/types"

	"github.com/joho/godotenv"
	"github.com/pkg/profile"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	files, err := dataProcessing.ReadFilesPath()
	if err != nil {
		fmt.Println("error al leer archivos de la carpeta maildir: ", err)
	}

	lineData := &types.DataEnron{
		Index:   "messi",
		Records: []types.Record{},
	}

	//dataStr es un objeto del struct Data el contiene varios atributos de tipo string, dicho objeto dataStr se agregara varias veces como un elemento del slice Records el cual es un atributo del struct DataEnron
	var dataStr types.Data
	var count int = 0
	for _, file := range files {

		//lines es un slice de string que contiene como elementos todos los renglones de un archivo, osea cada elemento es un renglon de un archivo
		lines := dataProcessing.ReadFilesLines(file)

		//esta funcion toma el objeto dataStr e inicializa cada uno de sus atributos con ciertos elementos del slice de string lines, despues retorna el objeto dataStr ya inicializado el cual se asigna a la variable dataString. este proceso se repitevarias veces porque esta dentro de un for
		dataString := dataProcessing.LinesToStruct(lines, dataStr)

		//se agrega un objeto Data al slice []Record
		lineData.Records = append(lineData.Records, types.Record{
			Document: types.Data{
				MessageID:               dataString.MessageID,
				Date:                    dataString.Date,
				From:                    dataString.From,
				To:                      dataString.To,
				Subject:                 dataString.Subject,
				Cc:                      dataString.Cc,
				MimeVersion:             dataString.MimeVersion,
				ContentType:             dataString.ContentType,
				ContentTransferEncoding: dataString.ContentTransferEncoding,
				XFrom:                   dataString.XFrom,
				XTo:                     dataString.XTo,
				Xcc:                     dataString.Xcc,
				Xbcc:                    dataString.Xbcc,
				XFolder:                 dataString.XFolder,
				XOrigin:                 dataString.XOrigin,
				XFileName:               dataString.XFileName,
				Body:                    dataString.Body,
			},
		})
		count++
		fmt.Println("file appended", count)
	}

	jsonData := dataProcessing.StructToJson(*lineData)

	//solo es necesario crer un cliente para hacer una o varias peticiones
	clientHttp := &http.Client{
		Timeout: time.Second * 10,
	}
	fmt.Println("indexing data")
	dataProcessing.IndexJson(jsonData, clientHttp)

}
