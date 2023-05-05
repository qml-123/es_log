package model

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
)

var esClient *elasticsearch.Client

func init() {
	var err error
	caBundle, err := ioutil.ReadFile("/home/elasticsearch/elasticsearch-8.7.0/config/certs/http_ca.crt")
	if err != nil {
		panic(err)
	}

	customCaBundle := "custom_ca_bundle.pem"
	err = ioutil.WriteFile(customCaBundle, caBundle, 0644)
	if err != nil {
		panic(err)
	}

	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(caBundle)

	tlsConfig := &tls.Config{
		RootCAs:            rootCAs,
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	esClient, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"https://localhost:9200"},
		Username:  "elastic",
		Password:  "FOWrYfQbfnRa1_WMepPk",
		Transport: transport,
	})
	if err != nil {
		panic(err)
	}
}

type QueryBuilder struct {
	PageNum   int32
	PageSize  int32
	Must      map[string][]string
	MustNot   map[string][]string
	Should    map[string][]string
	ShouldOr  bool
	TimeRange struct {
		Field string
		GTE   string
		LTE   string
	}
	Sort map[string]string
}
