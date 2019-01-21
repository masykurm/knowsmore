# Knowsmore
Simple Autocomplete Engine using go ft Elasticsearch

## Preparation
Please follow these steps to prepare elasticsearch and golang:

* Make sure you have golang in your device. I use `go 1.10.1 linux/amd64`

you can open this url to install [golang 1.10](https://www.admfactory.com/how-to-install-golang-1-10-on-ubuntu/)

* Install your elastic node, you can install in your local or just using docker

follow these steps to install in your [ubuntu](https://www.elastic.co/guide/en/beats/libbeat/5.6/elasticsearch-installation.html)

follow these steps to install in your [docker](https://www.elastic.co/guide/en/elasticsearch/reference/5.6/docker.html)

* Create mapping for your elastic index

```json
{
    "mappings": {
        "_doc" : {
            "properties" : {
                "suggest" : {
                    "type" : "completion"
                },
                "name" : {
                    "type": "keyword",
                    "copy_to": "suggest"
                },
                "url" : {
                    "type": "keyword"
                }
            }
        }
    }
}
```

you can use my dummy data
```json
{
    "member": [
      {
        "name": "Ruben Panjaitan",
        "url": "https://www.facebook.com/search/str/ruben panjaitan/keywords_search",
        "suggest": "Ruben Panjaitan"
      },
      {
        "name": "Roberto Tambunan",
        "url": "https://www.facebook.com/search/str/roberto tambunan/keywords_search",
        "suggest": "Roberto Tambunan"
      },
      {
        "name": "Ivandra Oktovan",
        "url": "https://www.facebook.com/search/str/ivandra oktovan/keywords_search",
        "suggest": "Ivandra Oktovan"
      },
      {
        "name": "Iman Situmorang",
        "url": "https://www.facebook.com/search/str/iman situmorang/keywords_search",
        "suggest": "Iman Situmorang"
      },
      {
        "name": "Septian Siahaan",
        "url": "https://www.facebook.com/search/str/septian siahaan/keywords_search",
        "suggest": "Septian Siahaan"
      },
      {
        "name": "Setia Simaremare",
        "url": "https://www.facebook.com/search/str/setia simaremare/keywords_search",
        "suggest": "Setia Simaremare"
      },
      {
        "name": "Raymond Sitepu",
        "url": "https://www.facebook.com/search/str/raymond sitepu/keywords_search",
        "suggest": "Raymond Sitepu"
      },
      {
        "name": "Witri Manurung",
        "url": "https://www.facebook.com/search/str/witri manurung/keywords_search",
        "suggest": "Witri Manurung"
      }
    ]
  }
```


## Installation
* Clone the project 
```bash
git clone https://github.com/robertotambunan/knowsmore.git
```

* Restore dependencies using go get
```bash
go get -v
```

* Run Project
```bash
cd cmd/knowsmore && go build && ./knowsmore
```


## Usage
Hit your go api to `POST localhost:9002/autocomplete` with example payload body:
```json
{
	"keyword": "r"
}
```
