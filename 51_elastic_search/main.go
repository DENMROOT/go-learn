package main

import (
	"encoding/json"
	common "go-learn/51_elastic_search/common/config"
	"go-learn/51_elastic_search/config"
	"go-learn/51_elastic_search/controllers"
	"go-learn/51_elastic_search/domain"
	"go-learn/51_elastic_search/interactors"
	"go-learn/51_elastic_search/repositories"
	"os"
	"strconv"

	"github.com/labstack/gommon/log"
)

var _configuration config.Configuration

func init() {
	if err := common.LoadConfigFromFile("config", &_configuration); err != nil {
		log.Error("Error loading config: ", err)
		os.Exit(0)
	}
}

func main() {
	log.Infof("JOB START")

	//controller := controllers.Controller {
	//	Interactor: interactors.Interactor{
	//		Repository: repositories.Repository {
	//			Configuration: _configuration,
	//		},
	//	},
	//}

	repository := repositories.NewRepository(_configuration)
	interactor := interactors.NewInteractor(repository)
	controller := controllers.Controller{
		Interactor: *interactor,
	}

	controller.CreateIndex("dummy")

	bulkInsert := [10]domain.Something{}
	type Teste struct {
		Fruit    string `json:"fruit"`
		Category string `json:"category"`
	}

	for i := 0; i < 10; i++ {
		teste := Teste{Fruit: "banana", Category: "natureza"}
		bytes, _ := json.Marshal(teste)
		bulkInsert[i] = domain.Something{
			TYPE: "Teste " + strconv.Itoa(i),
			ID:   strconv.Itoa(i),
			DATA: bytes,
		}
	}

	controller.Insert(bulkInsert[:])

	log.Infof("JOB END")
}
