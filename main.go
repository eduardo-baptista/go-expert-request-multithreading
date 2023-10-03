package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please provide a CEP as argument")
		return
	}

	cep := os.Args[1]
	if len(cep) != 9 {
		fmt.Println("invalid CEP, please provide a valid CEP on format 00000-000")
		return
	}

	chAPICEP := make(chan []byte)
	ctxAPICEP, cancelAPICEP := context.WithCancel(context.Background())
	chViaCEP := make(chan []byte)
	ctxViaCEP, cancelViaCEP := context.WithCancel(context.Background())

	go requestAPICEP(ctxAPICEP, cep, chAPICEP)
	go requestViaCEP(ctxViaCEP, cep, chViaCEP)

	select {
	case resp := <-chAPICEP:
		fmt.Printf("Response from APICEP: %s\n", resp)
		cancelViaCEP()
	case resp := <-chViaCEP:
		fmt.Printf("Response from ViaCEP: %s\n", resp)
		cancelAPICEP()
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
		cancelAPICEP()
		cancelViaCEP()
	}
}

func requestAPICEP(ctx context.Context, cep string, ch chan<- []byte) {
	url := fmt.Sprintf("https://cdn.apicep.com/file/apicep/%s.json", cep)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("APICEP: Error: %s\n", err.Error())
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("APICEP: Error: %s\n", err.Error())
		return
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("APICEP: Error: %s\n", err.Error())
		return
	}

	ch <- b
}

func requestViaCEP(ctx context.Context, cep string, ch chan<- []byte) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("ViaCEP: Error: %s\n", err.Error())
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("ViaCEP: Error: %s\n", err.Error())
		return
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("ViaCEP: Error: %s\n", err.Error())
		return
	}

	ch <- b
}
