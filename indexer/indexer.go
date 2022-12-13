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
	var dataStr types.Data
	var count int = 0
	for _, file := range files {

		//lines es un slice de string donde cada elemento es un renglon del archivo
		lines := dataProcessing.ReadFilesLines(file)
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
	clientHttp := &http.Client{
		Timeout: time.Second * 10,
	}
	fmt.Println("indexing data")
	dataProcessing.IndexJson(jsonData, clientHttp)

}
