package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/qri-io/jsonschema"
	"github.com/servicesys/GoHeadlessCMS/pkg/model"
	"github.com/servicesys/GoHeadlessCMS/pkg/storage"
	"go.uber.org/zap"
	"time"
)

func main() {

	fmt.Println("TESTANDO")

	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	sugar.Infow("failed to fetch URL",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("failed to fetch URL: %s", "http://example.com")

	contentStorage := storage.NewContentStoragePostgres()

	//testCategory(contentStorage)
	//testType(contentStorage)
	//testContent(contentStorage)
	//testSchema(contentStorage)
	testSearch(contentStorage)

}

func testContent(contentStorage storage.ContentStorage) {
	fmt.Println("--------------------")
	textoJSon := `{ "titulo" : "Este e um texto de exemplo -NOVO" ,
"texto" :  "Corpo do texto texto de exemplo"
		 }`

	SCHEMA := `{
     "$id": "https://qri.io/schema/",
    "$comment" : "sample comment",
    "title": "Person",
    "type": "object",
	"properties": {
		"titulo": {
			
			"title": "Titulo", 
			"type": "string",
			"default": "",
			"examples": [
				"Este e um texto de exemplo"
			],
			"pattern": "^.*$"
		},
		"texto": {
			
			"title": "Texto", 
			"type": "string",
			"default": "",
			"examples": [
				"<p>Este e  o corpot do texto texto de exemplo</p>"
			],
			"pattern": "^.*$"
		}
	},
	"required": [
		"titulo",
		"texto"
	]
}
`
	mtype := model.Type{
		Cod:         "TYPE",
		Metadata:    []byte(SCHEMA),
		Description: "TIPO",
	}

	content := model.Content{
		Uuid: "",
		Type: mtype,
		Category: model.Category{
			Cod:         "TEXTO_EXEMPLO",
			Description: "",
		},
		Content:    []byte(textoJSon),
		CreatedOn:  time.Time{},
		ModifiedOn: time.Time{},
		Status:     "AP",
	}

	errorCreate := contentStorage.CreateContent(content)

	fmt.Println(content.Status)
	fmt.Println(errorCreate)
	fmt.Println("--------------------")


		errorDelete := contentStorage.DeleteContent("513265f0-ae9e-4763-afe3-1a42805749ed")
		fmt.Println(errorDelete)

		content.Type = model.Type{
			Cod:         "TYPE2",
			Metadata:    []byte(SCHEMA),
			Description: "",
		}
		content.Uuid = "0d5a1aac-4489-4803-a356-9215ff8a3ea1"
		content.Status = "RM"
		errorUpdate := contentStorage.UpdateContent(content)
		fmt.Println(errorUpdate)

		mcontent, errorGet := contentStorage.GetContent("513265f0-ae9e-4763-afe3-1a42805749ed")
		fmt.Println(errorGet)
		fmt.Println(mcontent.Category)

		contents, err := contentStorage.GetAllContentByCategory("TEXTO_EXEMPLO")

		fmt.Println(err)
		//fmt.Println(contents[10].Type.Cod)
		if err == nil {
			fmt.Println("-----------")
			for _, s := range contents {

				fmt.Println(s.Type.Cod)

			}
		}

}

func testType(contentStorage storage.ContentStorage) {

	schemaTypeRaw := `{
	"definitions": {},
	"$schema": "http://json-schema.org/draft-07/schema#", 
	"$id": "https://example.com/object1615860509.json", 
	"title": "Root", 
	"type": "object",
	"required": [
		"titulo",
		"texto"
	],
	"properties": {
		"titulo": {
			"$id": "#root/titulo", 
			"title": "Titulo", 
			"type": "string",
			"default": "",
			"examples": [
				"Este e um texto de exemplo"
			],
			"pattern": "^.*$"
		},
		"texto": {
			"$id": "#root/texto", 
			"title": "Texto", 
			"type": "string",
			"default": "",
			"examples": [
				"<p>Este e  o corpot do texto texto de exemplo</p>"
			],
			"pattern": "^.*$"
		}
	}
}
`
	erroCreate := contentStorage.CreateType(model.Type{
		Cod:         "TYPE2",
		Metadata:    []byte(schemaTypeRaw),
		Description: "TIPO PARA DESCREVER TEXO E CORPO DO TEXTO",
	})

	fmt.Println(erroCreate)

	errorUpdate := contentStorage.UpdateType(model.Type{
		Cod:         "TYPE",
		Metadata:    []byte(schemaTypeRaw),
		Description: "TIPO PARA DESCREVER TEXO E CORPO DO TEXTO - ALTERADO",
	})
	fmt.Println(errorUpdate)

	mtype, errorGet := contentStorage.GetType("TYPE")
	fmt.Println(mtype.Description)
	fmt.Println(errorGet)

}
func testCategory(contentStorage storage.ContentStorage) {

	error := contentStorage.CreateCategory(model.Category{
		"TEXTO_EXEMPLO",
		"CATEGORIA DE TEXTO",
	})

	fmt.Println(error)

	listCategorys, errCategory := contentStorage.GetAllCategory()

	fmt.Println(listCategorys)
	fmt.Println(errCategory)

	errorUpdate := contentStorage.UpdateCategory(model.Category{
		"TEXTO_BLOG",
		"CATEGORIA DE TEXTO - ALTERADO",
	})

	fmt.Println(errorUpdate)
}

func testSchema(contentStorage storage.ContentStorage) {

	var schemaData = `{
    "$id": "https://qri.io/schema/",
    "$comment" : "sample comment",
    "title": "Person",
    "type": "object",
    "properties": {
        "firstName": {
            "type": "string"
        },
        "lastName": {
            "type": "string"
        },
        "age": {
            "description": "Age in years",
            "type": "integer",
            "minimum": 0
        },
        "friends": {
          "type" : "array",
          "items" : { "title" : "REFERENCE", "$ref" : "#" }
        }
    },
    "required": ["firstName", "lastName"]
  }`
	textoJSon := `{ "titulo" : "Este e um texto de exemplo -NOVO" ,
		"texto" :  "Corpo do texto texto de exemplo" }`

	jsonBytes := []byte(textoJSon)

	var listErrors = make([]string, 0)
	rs := &jsonschema.Schema{}

	if err := json.Unmarshal([]byte(schemaData), rs); err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonBytes))
	errs, err := rs.ValidateBytes(context.Background(), jsonBytes)
	if err != nil {
		listErrors = append(listErrors, err.Error())
		fmt.Println(err)
	}
	fmt.Println(errs)

}

func testSearch(contentStorage storage.ContentStorage) {

	contents := contentStorage.SearchContent("testo")

	fmt.Println(len(contents))
}
