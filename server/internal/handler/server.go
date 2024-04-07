package handler

import (
	"bytes"
	"context"
	"io"
	"service-pdf-compose/pkg/composer"
	"service-pdf-compose/pkg/composer/pdfcompose"
)

type PdfComposeServiceClient struct {
	pdfcompose.UnimplementedPdfComposeServiceServer
}

func (p *PdfComposeServiceClient) Send(ctx context.Context, params *pdfcompose.FormData) (*pdfcompose.File, error) {
	files := []io.ReadCloser{io.NopCloser(bytes.NewReader(params.Upfile1))}
	if params.Upfile2 != nil {
		files = append(files, io.NopCloser(bytes.NewReader(params.Upfile2)))
	}
	if params.Upfile3 != nil {
		files = append(files, io.NopCloser(bytes.NewReader(params.Upfile3)))
	}

	result, err := composer.ComposeFromFiles(files)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(result)
	if err != nil {
		return nil, err
	}

	return &pdfcompose.File{File: body}, nil
}
