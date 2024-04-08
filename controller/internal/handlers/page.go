package handlers

import (
	"context"
	"controller-service/pkg/composer/pdfcompose"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
)

type Handler struct {
	HttpClient *http.Client
}

func (h *Handler) Send(w http.ResponseWriter, r *http.Request) {
	cwt, _ := context.WithTimeout(context.Background(), time.Second*5)
	conn, err := grpc.DialContext(cwt, "pdf-compose-service:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	uc := pdfcompose.NewPdfComposeServiceClient(conn)

	files := &pdfcompose.FormData{
		Upfile1: h.getFile(r, "upfile1"),
	}

	if bytes := h.getFile(r, "upfile2"); bytes != nil {
		files.Upfile2 = bytes
	}
	if bytes := h.getFile(r, "upfile3"); bytes != nil {
		files.Upfile3 = bytes
	}

	resp, err := uc.Send(cwt, files)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Disposition", "attachment; filename=result.pdf")
	w.Write(resp.GetFile())

	return
}

func (h *Handler) getFile(r *http.Request, param string) []byte {
	file, _, err := r.FormFile(param)
	if err != nil {
		if errors.Is(err, http.ErrMissingFile) {
			return nil
		}
		log.Fatal(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func (h *Handler) Web(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/templates/form.html")
	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, nil)
}
