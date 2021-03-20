package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/qri-io/jsonschema"
	"github.com/servicesys/GoHeadlessCMS/pkg/model"
	"github.com/servicesys/GoHeadlessCMS/pkg/storage"
	"time"
)

func main() {

	fmt.Println("TESTANDO")

	contentStorage := storage.NewContentStoragePostgres()

	//testCategory(contentStorage)
	//testType(contentStorage)
	testContent(contentStorage)
	//testSchema(contentStorage)
	//testSearch(contentStorage)

}

func testContent(contentStorage storage.ContentStorage) {

	fmt.Println("--------------------")
	textoJSon := `{ "titulo" : "Este e um texto de exemplo -NOVO - ola" ,
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
		//Name: "EXEMPLO_2",
		Content:    []byte(textoJSon),
		CreatedOn:  time.Time{},
		ModifiedOn: time.Time{},
		Status:     "AP",
	}

	errorCreate := contentStorage.CreateContent(content)

	fmt.Println(content.Status)
	fmt.Println(errorCreate)
	fmt.Println("--------------------")

	errorDelete := contentStorage.DeleteContent("76d24f2b-299b-43c9-ab8a-bae420366f61")
	fmt.Println(errorDelete)

	content.Type = model.Type{
		Cod:         "TYPE",
		Metadata:    []byte(SCHEMA),
		Description: "",
	}
	content.Uuid = "68bfe91b-6dd0-46fb-8748-cdfa7d1a81be"
	content.Status = "RM"
	content.Name = "EXEMPLO_RM"
	errorUpdate := contentStorage.UpdateContent(content)
	fmt.Println(errorUpdate)

	mcontent, errorGet := contentStorage.GetContent("68bfe91b-6dd0-46fb-8748-cdfa7d1a81be")
	fmt.Println(errorGet)
	fmt.Println(mcontent.Category, mcontent.Name, content.Type.Cod)

	contents, err := contentStorage.GetAllContentByCategory("TEXTO_EXEMPLO")

	fmt.Println(err)
	//fmt.Println(contents[10].Type.Cod)
	if err == nil {
		fmt.Println("\nget category-----------")
		for _, s := range contents {

			fmt.Println(s.Type.Cod, s.Name)

		}
	}

	mcontentName, errorName := contentStorage.GetContentByName("EXEMPLO_RM")
	fmt.Println(errorName)
	fmt.Println(mcontentName.Category, mcontentName.Name, mcontentName.Type.Cod)

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
		Cod:         "TYPE",
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
		"TEXTO_BLOG",
		"CATEGORIA DE TEXTO",
	})

	fmt.Println(error)

	listCategorys, errCategory := contentStorage.GetAllCategory()

	fmt.Println(listCategorys)
	fmt.Println(errCategory)

	errorUpdate := contentStorage.UpdateCategory(model.Category{
		"TEXTO_BLOG",
		"CATEGORIA DE TEXTO BLOG",
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

	contents, error := contentStorage.SearchContent([]string{"titulo", "texto"}, []string{"ola", "ola"})

	if error == nil {
		fmt.Println("\nSearch-----------")
		for _, s := range contents {

			fmt.Println(s.Type.Cod, s.Name, s.Category.Cod)

		}
	}
	//fmt.Println(contents)
	fmt.Println(error)
	fmt.Println(len(contents))
}
