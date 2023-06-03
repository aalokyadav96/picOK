package main

import(
	"net/http"
	"fmt"
	
	"crypto/rand"
	"io/ioutil"
	"mime"
	"os"
	"path/filepath"
	
    "github.com/julienschmidt/httprouter"
	
)

const maxUploadSize = 2 * 1024 * 1024 // 2 mb
var uploadPath = "./uploads"
//var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("path", r.URL.Path)
	if r.Method == "GET" {
		if isLoggedIn(r) {
			res := "fytrdjvhtfudy"
			fmt.Println(res)
			tmpl.ExecuteTemplate(w, "index.html", res)
		} else {
			tmpl.ExecuteTemplate(w, "head.html", nil)
			tmpl.ExecuteTemplate(w, "nav.html", nil)
			tmpl.ExecuteTemplate(w, "nonloginhome.html", nil)
			tmpl.ExecuteTemplate(w, "footer.html", nil)
		}
	}
}
// isLoggedIn()
//				http.Redirect(w, r, "/register", http.StatusSeeOther)			


func NewPhotoGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("path", r.URL.Path)
	if r.Method == "GET" {
		if isLoggedIn(r) {
			tmpl.ExecuteTemplate(w, "head.html", nil)
			tmpl.ExecuteTemplate(w, "nav.html", nil)
			tmpl.ExecuteTemplate(w, "upload.html", nil)
			tmpl.ExecuteTemplate(w, "footer.html", nil)
		} else {
				http.Redirect(w, r, "/login", http.StatusSeeOther)			
		}
	}
}
func NewPhotoPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("path", r.URL.Path)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			fmt.Printf("Could not parse multipart form: %v\n", err)
			renderError(w, "CANT_PARSE_FORM", http.StatusInternalServerError)
			return
		}
var pics []string
		files := r.MultipartForm.File["imgfile"]
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		// Get and print out file size
		fileSize := fileHeader.Size
		fmt.Printf("File size (bytes): %v\n", fileSize)
		// validate file size
		if fileSize > maxUploadSize {
			renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
			return
		}
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}

		// check file type, detectcontenttype only needs the first 512 bytes
		detectedFileType := http.DetectContentType(fileBytes)
		switch detectedFileType {
		case "image/png", "image/jpg", "image/jpeg":
			break
		default:
			renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
			return
		}
		fileName := randToken(12)
		fileEndings, err := mime.ExtensionsByType(detectedFileType)
		if err != nil {
			renderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
			return
		}
	if fileEndings[0] == ".jfif" {fileEndings[0] = ".jpg"}
		newFileName := fileName + fileEndings[0]

		newPath := filepath.Join(uploadPath, newFileName)
		fmt.Printf("FileType: %s, File: %s\n", detectedFileType, newPath)

		// write file
		newFile, err := os.Create(newPath)
		if err != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		defer newFile.Close() // idempotent, okay to call twice
		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		

		pics = append(pics,"/files/"+newFileName)
		}
		
	fmt.Println(pics)
//	http.Redirect(w, r, "/", http.StatusSeeOther)			
	tmpl.ExecuteTemplate(w,"res.html", pics)

//		w.Write([]byte(fmt.Sprintf("SUCCESS - use /files/%v to access the file", newFileName)))
//			t, _ := template.ParseFiles("templates/returnUploads.html")
//			t, _ := template.ParseFiles("res.html")
//			t.Execute(w, "files/"+newFileName)
//			return
	}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}



func ShowPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		fmt.Println("path: ", r.URL.Path)
		
		postID := ps.ByName("postid")
		fmt.Println("postID: ", postID)
		
		tmpl.ExecuteTemplate(w, "head.html", nil)
		tmpl.ExecuteTemplate(w, "nav.html", nil)
		tmpl.ExecuteTemplate(w, "post.html", postID)
		tmpl.ExecuteTemplate(w, "footer.html", nil)
}

func DeletePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		fmt.Println("Method :", r.Method)
		
		fmt.Println("path: ", r.URL.Path)
		postID := ps.ByName("postid")
		fmt.Println("postID: ", postID)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}