package main

import (
	"fmt"
	"indexerZinc/dataProcessing"
	"net/http"
	"sync"
	"time"

	"github.com/pkg/profile"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(517_424)
	defer profile.Start(profile.CPUProfile).Stop()
	files, err := dataProcessing.ReadFilesPath()
	if err != nil {
		fmt.Println("error al leer archivos de la carpeta maildir: ", err)
	}
	//se crea un cliente http que se usa en todas las peticiones
	clientHttp := &http.Client{
		Timeout: time.Second * 10,
	}

	for _, file := range files {
		lines := dataProcessing.ReadFilesLines(file) //lines es un slice de string donde cada elemento es un renglon del archivo
		jsonData := dataProcessing.LinesToJson(lines)
		fmt.Println("indexing: ", file)
		go dataProcessing.IndexJson(jsonData, clientHttp, wg)
	}
	wg.Wait()
}
