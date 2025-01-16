package elastic

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type ElasticLayer struct {
	transport *http.Transport
	client    *elasticsearch.Client
}

func NewElasticLayer(client *elasticsearch.Client, transport *http.Transport) *ElasticLayer {
	return &ElasticLayer{
		client:    client,
		transport: transport,
	}
}

func NewESClient() (*elasticsearch.Client, *http.Transport) {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	transport := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 5 * time.Second,
	}

	addrs := os.Getenv("elastic_addrs")
	username := os.Getenv("elastic_username")
	password := os.Getenv("elastic_password")

	addrList := strings.Split(addrs, ",")

	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: addrList,
		Username:  username,
		Password:  password,
		Transport: transport,
	})
	if err != nil {
		log.Fatal(err)
	}

	return client, transport
}

func (el *ElasticLayer) Close() {
	if el.transport != nil {
		el.transport.CloseIdleConnections()
	}
}
