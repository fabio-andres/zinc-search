package main

import (
	"fmt"
	"net/http"
	"time"
	"zse/dataProcessing"
	"zse/types"

	"github.com/pkg/profile"
)

func main() {
	//wg := new(sync.WaitGroup)
	//wg.Add(517_424)
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
	for _, file := range files {
		lines := dataProcessing.ReadFilesLines(file) //lines es un slice de string donde cada elemento es un renglon del archivo
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
		fmt.Println("line appended")
	}

	jsonData := dataProcessing.StructToJson(*lineData)

	fmt.Println(string(jsonData))

	clientHttp := &http.Client{
		Timeout: time.Second * 10,
	}
	dataProcessing.IndexJson(jsonData, clientHttp)

}
