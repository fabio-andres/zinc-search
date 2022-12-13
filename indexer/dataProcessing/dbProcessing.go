package dataProcessing

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"zse/types"
)

func ReadFilesPath() ([]string, error) {
	mainFolder := "enronDB/maildir"
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

	//os.Open() abre el archivo que se le pase como parametro y retorna el archivo ya abierto. esto es necesario para poder acceder al contenido de un archivo
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	//bufio.NewScanner() crea un objeto scanner inicializando su atributo r de tipo io.Reader con el archivo de texto abierto que se le pase como parametro
	fileScanner := bufio.NewScanner(readFile)

	//Split(bufio.ScanLines) añade un token a cada renglon del archivo para que el metodo Scan() solo detecte renglones
	fileScanner.Split(bufio.ScanLines)

	//Scan() escanea el archivo tal como se indico con Split(bufio.ScanLines), osea Scan() escanea renglon por renglon del archivo
	for fileScanner.Scan() {
		//Text() convierte un renglon en string
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines //retorna un slice donde cada elemento es un renglon del archivo
}

func LinesToStruct(lines []string, dataStr types.Data) types.Data {
	//line representa un renglon de un archivo representado por lines
	for _, line := range lines {
		switch 0 {
		//si en el line (renglon) del archivo se halla la palabra Message-ID retornar la posicion del renglon donde inicia esa palabra
		case strings.Index(line, "Message-ID:"):
			dataStr.MessageID = line[11:len(line)]
		case strings.Index(line, "Date:"):
			dataStr.Date = line[5:len(line)]
		case strings.Index(line, "From:"):
			dataStr.From = line[5:len(line)]
		case strings.Index(line, "To:"):
			dataStr.To = line[3:len(line)]
		case strings.Index(line, "Subject:"):
			dataStr.Subject = line[8:len(line)]
		case strings.Index(line, "Cc:"):
			dataStr.Cc = line[3:len(line)]
		case strings.Index(line, "Mime-Version:"):
			dataStr.MimeVersion = line[13:len(line)]
		case strings.Index(line, "Content-Type:"):
			dataStr.ContentType = line[13:len(line)]
		case strings.Index(line, "Content-Transfer-Encoding:"):
			dataStr.ContentTransferEncoding = line[26:len(line)]
		case strings.Index(line, "X-From:"):
			dataStr.XFrom = line[7:len(line)]
		case strings.Index(line, "X-To:"):
			dataStr.XTo = line[5:len(line)]
		case strings.Index(line, "X-cc:"):
			dataStr.Xcc = line[5:len(line)]
		case strings.Index(line, "X-bcc:"):
			dataStr.Xbcc = line[6:len(line)]
		case strings.Index(line, "X-Folder:"):
			dataStr.XFolder = line[9:len(line)]
		case strings.Index(line, "X-Origin:"):
			dataStr.XOrigin = line[9:len(line)]
		case strings.Index(line, "X-FileName:"):
			dataStr.XFileName = line[11:len(line)]
		default:
			//el line (renglon) se concatena con el anterior line
			dataStr.Body += line
		}
	}
	return dataStr
}

func StructToJson(lineData types.DataEnron) []byte {
	jsonData, err := json.Marshal(lineData) //convierte el struct en un json
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	return jsonData
}

func IndexJson(jsonData []byte, clientHttp *http.Client) {
	user := "admin"
	password := os.Getenv("ZINC_SEARCH_PASSWORD")
	encodeCredentials := base64.StdEncoding.EncodeToString([]byte(user + ":" + password))

	url := "http://localhost:4080/api/_bulkv2"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creando petición: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+encodeCredentials)
	resp, err := clientHttp.Do(req)
	if err != nil {
		log.Fatalf("Error haciendo petición: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("request successful", resp.Status)
	} else {
		fmt.Println(resp.Status)
	}

	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error leyendo respuesta: %v", err)
	}

	fmt.Println(string(bodyResp))
}
