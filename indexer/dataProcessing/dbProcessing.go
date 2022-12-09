package dataProcessing

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func ReadFilesPath() ([]string, error) {
	mainFolder := "enronDB/maildirComplete"
	files := []string{}
	//el parametro path string es la ruta de cada archivo o carpeta dentro de la carpeta mainFolder que se paso como paraemtro. cada vez que el metodo filepath.WalkDir() encuentre una carpeta o archivo dentro de la ruta que esta en el parametro mainFolder entonces ejecutara el callback func(path string, info fs.DirEntry, err error)
	err := filepath.WalkDir(mainFolder, func(path string, info fs.DirEntry, err error) error {
		//si info no es un directorio es porque es un archivo entonces agreguelo al slice files
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func ReadFilesLines(file string) []string {
	var fileLines []string

	readFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines //retorna un slice donde cada elemento es un renglon del archivo
}

func LinesToJson(lines []string) []byte {
	mapData := make(map[string]string)

	for _, line := range lines {
		switch 0 {
		case strings.Index(line, "Message-ID:"):
			mapData["Message-ID"] = line[11:len(line)]
		case strings.Index(line, "Date:"):
			mapData["Date"] = line[5:len(line)]
		case strings.Index(line, "From:"):
			mapData["From"] = line[5:len(line)]
		case strings.Index(line, "To:"):
			mapData["To"] = line[3:len(line)]
		case strings.Index(line, "Subject:"):
			mapData["Subject"] = line[8:len(line)]
		case strings.Index(line, "Cc:"):
			mapData["Cc"] = line[3:len(line)]
		case strings.Index(line, "Mime-Version:"):
			mapData["Mime-Version"] = line[13:len(line)]
		case strings.Index(line, "Content-Type:"):
			mapData["Content-Type"] = line[13:len(line)]
		case strings.Index(line, "Content-Transfer-Encoding:"):
			mapData["Content-Transfer-Encoding"] = line[26:len(line)]
		case strings.Index(line, "X-From:"):
			mapData["X-From"] = line[7:len(line)]
		case strings.Index(line, "X-To:"):
			mapData["X-To"] = line[5:len(line)]
		case strings.Index(line, "X-cc:"):
			mapData["X-cc"] = line[5:len(line)]
		case strings.Index(line, "X-bcc:"):
			mapData["X-bcc"] = line[6:len(line)]
		case strings.Index(line, "X-Folder:"):
			mapData["X-Folder"] = line[9:len(line)]
		case strings.Index(line, "X-Origin:"):
			mapData["X-Origin"] = line[9:len(line)]
		case strings.Index(line, "X-FileName:"):
			mapData["X-FileName"] = line[11:len(line)]
		default:
			mapData["Body"] += line
		}
	}

	jsonData, err := json.Marshal(mapData) //convierte el map en un json
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	return jsonData
}

func IndexJson(jsonData []byte, clientHttp *http.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	user := "admin"
	password := "Complexpass#123"
	encodeCredentials := base64.StdEncoding.EncodeToString([]byte(user + ":" + password))

	url := "http://localhost:4080/api/enron1/_doc"

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creando petición: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+encodeCredentials)
	_, err = clientHttp.Do(req)
	if err != nil {

		log.Fatalf("Error haciendo petición: %v", err)
	}

}
