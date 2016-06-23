package artifact

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const urlBase string = "https://search.maven.org/solrsearch/select?rows=20&wt=json&q="

func GetInfo(name string) []Artifact {
	res, err := http.Get(urlBase + name)
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	var artifacts Artifacts
	if err := json.Unmarshal(body, &artifacts); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return nil
	}

	return artifacts.Response.Docs
}

type Artifacts struct {
	ResponseHeader ResponseHeader
	Response       Response
	Spellcheck     Spellcheck
}

type ResponseHeader struct {
	Status int
	QTime  int
	Params Params
}

type Params struct {
	Spellcheck      string
	Fl              string
	Sort            string
	Indent          string
	Q               string
	Qf              string
	SpellcheckCount string `json:"spellcheck.count"`
	Wt              string
	Rows            string
	Version         string
	DefType         string
}

type Response struct {
	NumFound int
	Start    int
	Docs     []Artifact
}

type Artifact struct {
	Id            string
	G             string
	A             string
	LatestVersion string
	RepositoryId  string
	P             string
	TimeStamp     int
	VersionCount  int
	Text          []string
	Ec            []string
}

type Spellcheck struct {
	Suggestions []string
}
