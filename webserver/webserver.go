package webserver

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"test/tagger"
	"time"
)

type Webserver struct {
	ServerPort        string
	MaxFileUploadSize int64
	Tagger            *tagger.Tagger
}

func NewServer(serverPort string, maxFileUploadSize int64, t *tagger.Tagger) *Webserver {
	return &Webserver{
		ServerPort:        serverPort,
		MaxFileUploadSize: maxFileUploadSize,
		Tagger:            t,
	}
}

func (ws *Webserver) Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ws.indexHandler)
	mux.HandleFunc("/upload", ws.uploadHandler)

	log.Println("FileHandler API listening in port " + ws.ServerPort)
	log.Fatal(http.ListenAndServe(":"+ws.ServerPort, mux))
}

func (ws *Webserver) indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}

func (ws *Webserver) uploadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, ws.MaxFileUploadSize)
	if err := r.ParseMultipartForm(ws.MaxFileUploadSize); err != nil {
		http.Error(w, "The uploaded file is too big. Please choose an file that's less than 1MB in size", http.StatusBadRequest)
		return
	}

	// The argument to FormFile must match the name attribute
	// of the file input on the frontend
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	// Create the uploads folder if it doesn't
	// already exist
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new file in the uploads directory
	dst, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Upload successful")
	score, err := ws.Tagger.Tag(dst.Name())

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Fprintf(w, "Probably: "+score)
}
