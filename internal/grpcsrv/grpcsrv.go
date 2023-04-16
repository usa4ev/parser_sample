package grpcsrv

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/usa4ev/parser_sample/internal/grpcsrv/protoparser"
	"github.com/usa4ev/parser_sample/internal/handlers"
	"github.com/usa4ev/parser_sample/internal/validation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	protoparser.ParserServer

	gs       *grpc.Server
	endpoint string // url where the server starts
}

func New(endpoint string) *server {
	srv := server{}
	srv.endpoint = endpoint

	return &srv
}

func (srv *server) ListenAndServe() error {
	listen, err := net.Listen("tcp", srv.endpoint)
	if err != nil {
		return err
	}

	gs := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	srv.gs = gs
	protoparser.RegisterParserServer(gs, srv)

	if err := gs.Serve(listen); err != nil {
		return err
	}

	return nil
}

// GetCompany returns company data if can successfully scrap it from rusprofile
func (srv *server) GetCompany(ctx context.Context, in *protoparser.GetCompanyRequest) (*protoparser.GetCompanyResponse, error) {
	res := protoparser.GetCompanyResponse{}

	if err := validation.ValidateINN(in.Inn); err != nil {
		return nil, fmt.Errorf("invalid inn &v: %v", in.Inn, err)
	}

	url := fmt.Sprintf("https://www.rusprofile.ru/search?query=%s", in.Inn)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:62.0) Gecko/20100101 Firefox/62.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Set("Connection", "keep-alive")

	webres, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get web data: %v", err)
	}

	if webres.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code from website: %v", webres.StatusCode)
	}

	defer webres.Body.Close()

	company, err := handlers.ScrapCompany(webres.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to scrap data: %v", err)
	}

	res.Ceo = company.CEO
	res.Inn = company.INN
	res.Kpp = company.KPP
	res.Name = company.Name

	return &res, nil
}
