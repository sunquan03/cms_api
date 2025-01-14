package elastic

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

type ElsaticLayer struct {
	client *elasticsearch.Client
}

func NewElsaticLayer(client *elasticsearch.Client) *ElsaticLayer {
	return &ElsaticLayer{client: client}
}

func NewESClient() *elasticsearch.Client {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	addrs := os.Getenv("elastic_addrs")
	username := os.Getenv("elastic_username")
	password := os.Getenv("elastic_password")

	addrList := strings.Split(addrs, ",")

	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: addrList,
		Username:  username,
		Password:  password,
	})
	if err != nil {
		log.Fatal(err)
	}

	return client
}
