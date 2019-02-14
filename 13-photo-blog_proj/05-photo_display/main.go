package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getcookie(w, req)
	//process form
	if req.Method == http.MethodPost {
		// open
		f, h, err := req.FormFile("newfile")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		//create sha for file name
		ext := strings.Split(h.Filename, ".")[1]
		fh := sha1.New()
		io.Copy(fh, f)
		fname := fmt.Sprintf("%x", fh.Sum(nil)) + "." + ext

		// store on server
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		dst, err := os.Create(filepath.Join(wd, "public", "pics", fname))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		//copy
		f.Seek(0, 0)
		io.Copy(dst, f)

		//add to cookie
		c = associate(w, c, fname)
	}

	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs[1:])
}

func getcookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

		http.SetCookie(w, c)
	}

	return c

}

func associate(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, fname) {
		s += "|" + fname
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}
