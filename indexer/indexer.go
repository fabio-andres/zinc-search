package main

import (
	"fmt"
	"indexerZinc/dataProcessing"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	files, err := dataProcessing.ReadFilesPath()
	if err != nil {
		fmt.Println("error al leer archivos de la carpeta maildir: ", err)
	}

	for _, file := range files {
		lines := dataProcessing.ReadFilesLines(file) //lines es un slice de string donde cada elemento es un renglon del archivo
		jsonData := dataProcessing.LinesToJson(lines)
		fmt.Println("indexing: ", file)
		dataProcessing.IndexJson(jsonData)
	}

}
