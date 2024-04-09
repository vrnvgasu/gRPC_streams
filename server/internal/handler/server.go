package handler

import (
	"bufio"
	"bytes"
	"io"
	"service-pdf-compose/pkg/composer"
	"service-pdf-compose/pkg/composer/pdfcompose"
)

type PdfComposeServiceServer struct {
	pdfcompose.UnimplementedPdfComposeServiceServer
}

func (c *PdfComposeServiceServer) Send(params *pdfcompose.FormData, stream pdfcompose.PdfComposeService_SendServer) error {
	files := []io.ReadCloser{io.NopCloser(bytes.NewReader(params.Upfile1))}
	if params.Upfile2 != nil {
		files = append(files, io.NopCloser(bytes.NewReader(params.Upfile2)))
	}
	if params.Upfile3 != nil {
		files = append(files, io.NopCloser(bytes.NewReader(params.Upfile3)))
	}

	result, err := composer.ComposeFromFiles(files)
	if err != nil {
		return err
	}

	r := bufio.NewReader(result)
	b := make([]byte, 3)
	for {
		n, err := r.Read(b)
		if err != nil {
			return err
		}

		if err := stream.Send(&pdfcompose.Chunk{Content: b[0:n]}); err != nil {
			return err
		}
	}
}
