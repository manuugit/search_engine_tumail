package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/pprof"
	"strings"

	env "git.com/searchEngineTumail"
)

type Datamail struct {
	MessageID      string
	Date           string
	UserName       string
	SenderEmail    string
	RecipientEmail string
	Subject        string
	CategoryFolder string
	Body           string
}

// map to match file and struct keys
var fileKeys = map[string]string{
	"messageID":      "Message-ID",
	"date":           "Date",
	"senderEmail":    "From",
	"recipientEmail": "To",
	"subject":        "Subject",
	"userName":       "X-Origin",
	"categoryFolder": "X-Folder",
	"body":           "body",
}
var mails []Datamail
var indexJson string = `{ "index" : { "_index" : "mails" } }`
var USER_PASSWORD_DB = env.USERDB + ":" + env.PASSWORDDB

// this function will be executed for each file
// Save file info in a Datamail struct and append it to mails list
func getMailFromFile(path string, info os.DirEntry, err error) error {
	if err != nil {
		log.Fatal("Error al acceder a:", path, err)
		return err
	}

	if !info.IsDir() {
		// map aux
		data := make(map[string]string)
		// lectura del archivo
		content, err := os.ReadFile(path)
		if err != nil {
			log.Fatal("Error leyendo archivo", err)
			return err
		}
		// se pasa cada línea del archivo a una lista
		lines := strings.Split(string(content), "\n")
		body := ""

		// se obtiene el cuerpo del mensaje, concatenando las últimas líneas
		if len(lines) > 15 {
			body = strings.Join(lines[15:len(lines)], "\n")
			// se guarda en un map
			data["body"] = body
			// líneas del archivo por indexar, sin el body
			lines = lines[:15]
		}

		// se guardan los atributos de interés en el map
		for _, line := range lines {
			atribute := strings.SplitN(line, ":", 2)
			if len(atribute) == 2 {
				key := strings.TrimSpace(atribute[0])
				value := strings.TrimSpace(atribute[1])
				data[key] = value
			}
		}
		// creamos un struct con la info guardada en el map y lo agregamos a la lista
		mail := Datamail{MessageID: data[fileKeys["messageID"]], Date: data[fileKeys["date"]], SenderEmail: data[fileKeys["senderEmail"]],
			RecipientEmail: data[fileKeys["recipientEmail"]], Subject: data[fileKeys["subject"]], UserName: data[fileKeys["userName"]],
			CategoryFolder: data[fileKeys["categoryFolder"]], Body: data[fileKeys["body"]]}
		// se valida que no sea un archivo vacío
		if !(mail.MessageID == "" && mail.Date == "" && mail.SenderEmail == "" && mail.RecipientEmail ==
			"" && mail.Subject == "" && mail.UserName == "" && mail.CategoryFolder == "" && mail.Body == "") {
			mails = append(mails, mail)
		}
	}
	return nil
}

// this function create a ndjson file from mails list. This file is used to bulk ingestion
func createNdjson(list []Datamail) {
	file, err := os.Create("mails.ndjson")
	if err != nil {
		log.Fatal("Error creando archivo", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, element := range list {
		// se escribe el json con el index
		_, err := writer.WriteString(indexJson + "\n")
		// se escribe el registro
		err = json.NewEncoder(writer).Encode(element)
		if err != nil {
			log.Fatal("Error al escribir en ndjson ", err)
		}
		writer.Flush()
	}
}

func main() {
	// inicia profiling de CPU
	f, err := os.Create("indexer.prof")
	if err != nil {
		log.Println("error generando perfil de cpu")
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// se pasa el nombre de la carpeta por línea de comandos
	if len(os.Args) == 2 {
		dirName := os.Args[1]
		path := "../../../" + dirName
		log.Println("procesando...")

		err := filepath.WalkDir(path, getMailFromFile)
		if err != nil {
			log.Fatal("Error recorriendo subdirectorios:", err)
		}

		createNdjson(mails)
		// con el archivo ndjson generado, se puede hacer la carga masiva en zincsearch usando el comando curl
		// curl http://localhost:4080/api/_bulk -i -u userName:password  --data-binary "@mails.ndjson"
		cmd := exec.Command("curl", "http://localhost:4080/api/_bulk", "-i", "-u", USER_PASSWORD_DB, "--data-binary", "@mails.ndjson")
		err = cmd.Run()
		if err != nil {
			log.Fatal("Error al ejecutar el comando curl:", err)
		}

	} else {
		log.Fatal("Se debe proporcionar un parámetro: el nombre de la carpeta")
	}

	//time.Sleep(30 * time.Second)
}
