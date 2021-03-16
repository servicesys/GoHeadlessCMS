package storage

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/qri-io/jsonschema"
	uuid "github.com/satori/go.uuid"
	"github.com/servicesys/GoHeadlessCMS/pkg/model"
)

type ContentStoragePostgres struct {
	dbConnection *pgx.Conn
	//jsonValidator schema.JsonSchemaValidator
}

func NewContentStoragePostgres() ContentStorage {

	connection := Connect()
	contentStoragePostgres := &ContentStoragePostgres{
		dbConnection: connection,
		//jsonValidator: schema.JsonSchemaValidatorQri{},
	}

	return contentStoragePostgres
}

func (contentStorage *ContentStoragePostgres) SearchContent(strSearch string) []model.Content {
	panic("implement me")
}

func (contentStorage *ContentStoragePostgres) GetContent(uuid string) (model.Content, error) {

	strQuery :=
		`SELECT     uuid,
					value,
	                content_status, 
					created_on,
					modified_on,
                    ct.cod,
					ct.description,
					ct.metadata,
	                cc.cod,
					cc.description 
					FROM headless_cms.content_value cv
					INNER  JOIN headless_cms.content_type ct ON (cv.content_type_cod =ct.cod)
                    INNER  JOIN headless_cms.content_category cc  ON (cv.content_category_cod =cc.cod)
				    WHERE uuid=$1;`
	rows, errQuery := contentStorage.dbConnection.Query(context.Background(), strQuery, uuid)
	content := model.Content{}

	defer rows.Close()
	if errQuery != nil {
		return content, errQuery
	}
	for rows.Next() {
		content.Type = model.Type{}
		content.Category = model.Category{}
		rows.Scan(&content.Uuid,
			&content.Content,
			&content.Status,
			&content.CreatedOn,
			&content.ModifiedOn,
			&content.Type.Cod,
			&content.Type.Description,
			&content.Type.Metadata,
			&content.Category.Cod,
			&content.Category.Description)
	}
	return content, nil
}

func (contentStorage *ContentStoragePostgres) GetAllContentByCategory(categoryCod string) ([]model.Content, error) {

	strQuery :=
		`SELECT     uuid,
					value,
	                content_status, 
					created_on,
					modified_on,
                    ct.cod,
					ct.description,
					ct.metadata,
	                cc.cod,
					cc.description 
					FROM headless_cms.content_value cv
					INNER  JOIN headless_cms.content_type ct ON (cv.content_type_cod =ct.cod)
                    INNER  JOIN headless_cms.content_category cc  ON (cv.content_category_cod =cc.cod)
				    WHERE   cc.cod=$1;`
	rows, errQuery := contentStorage.dbConnection.Query(context.Background(), strQuery, categoryCod)

	contents := make([]model.Content, 0)

	defer rows.Close()
	if errQuery != nil {
		return contents, errQuery
	}
	for rows.Next() {
		content := model.Content{}
		content.Type = model.Type{}
		content.Category = model.Category{}
		errScan := rows.Scan(&content.Uuid,
			&content.Content,
			&content.Status,
			&content.CreatedOn,
			&content.ModifiedOn,
			&content.Type.Cod,
			&content.Type.Description,
			&content.Type.Metadata,
			&content.Category.Cod,
			&content.Category.Description)
		if errScan != nil {
			return contents, errScan
		}
		contents = append(contents, content)
	}
	return contents, nil
}

func validateContent(content model.Content) (model.Content, error) {
	//validate schema
	if !json.Valid(content.Type.Metadata) {
		return content, errors.New("SCHEMA INVALID")
	}
	if !json.Valid(content.Content) {
		return content, errors.New("Content INVALID")
	}

	/* @TODO VALIDAR O CONTEUDO COMPATIVEL COM O TYPO
	ctx := context.TODO()
	valid, erroStr := content.Type.Validator(ctx, content.Content)
	if !valid {
		return content, errors.New(erroStr[0])
	}*/
	return model.Content{}, nil
}

func (contentStorage *ContentStoragePostgres) CreateContent(content model.Content) error {

	_, errValidate := validateContent(content)
	if errValidate != nil {
		return errValidate
	}
	queryInsert := `INSERT INTO headless_cms.content_value(uuid, content_category_cod, 
                    content_type_cod, value, content_status , created_on, modified_on) 
                    VALUES($1, $2 ,$3, $4 , $5 ,  now()  ,  now() )`

	content.Uuid = uuid.NewV4().String()
	err := doExecute(contentStorage.dbConnection, queryInsert, content.Uuid, content.Category.Cod,
		content.Type.Cod, content.Content, content.Status)
	return err
}

func (contentStorage *ContentStoragePostgres) UpdateContent(content model.Content) error {

	_, errValidate := validateContent(content)
	if errValidate != nil {
		return errValidate
	}
	queryUpdate := `UPDATE headless_cms.content_value
	SET content_category_cod=$1 , content_type_cod=$2, value=$3,  modified_on=now() , content_status=$4
	WHERE uuid=$5`
	err := doExecute(contentStorage.dbConnection, queryUpdate, content.Category.Cod, content.Type.Cod, content.Content, content.Status, content.Uuid)
	return err

}

func (contentStorage *ContentStoragePostgres) DeleteContent(uuid string) error {

	queryDelete := `DELETE FROM headless_cms.content_value WHERE uuid=$1;`
	err := doExecute(contentStorage.dbConnection, queryDelete, uuid)
	return err
}

func (contentStorage *ContentStoragePostgres) CreateCategory(category model.Category) error {
	query := `
		INSERT INTO headless_cms.content_category (cod, description) VALUES($1, $2)`

	err := doExecute(contentStorage.dbConnection, query, category.Cod, category.Description)
	return err
}

func (contentStorage *ContentStoragePostgres) GetAllCategory() ([]model.Category, error) {

	strQuery := `SELECT cod, description FROM headless_cms.content_category;`

	rows, errQuery := contentStorage.dbConnection.Query(context.Background(), strQuery)
	var categorys []model.Category
	if errQuery != nil {
		return categorys, errQuery
	}

	for rows.Next() {

		category := model.Category{}
		err := rows.Scan(
			&category.Cod,
			&category.Description)

		if err != nil {
			return categorys, err
		}
		categorys = append(categorys, category)
	}

	defer rows.Close()

	return categorys, nil
}

func (contentStorage *ContentStoragePostgres) UpdateCategory(category model.Category) error {

	query := `UPDATE headless_cms.content_category SET description = $1 WHERE cod=$2`
	err := doExecute(contentStorage.dbConnection, query, category.Description, category.Cod)
	return err
}

func (contentStorage *ContentStoragePostgres) CreateType(mtype model.Type) error {

	//validate schema
	if err := json.Unmarshal(mtype.Metadata, &jsonschema.Schema{}); err != nil {
		return err
	}
	queryInsert := `INSERT INTO headless_cms.content_type(cod, metadata, description) VALUES($1, $2, $3);`
	err := doExecute(contentStorage.dbConnection, queryInsert, mtype.Cod, mtype.Metadata, mtype.Description)
	return err
}

func (contentStorage *ContentStoragePostgres) UpdateType(mtype model.Type) error {

	if err := json.Unmarshal(mtype.Metadata, &jsonschema.Schema{}); err != nil {
		return err
	}
	queryInsert := `UPDATE headless_cms.content_type SET metadata=$1, description=$2 WHERE cod=$3;`
	err := doExecute(contentStorage.dbConnection, queryInsert, mtype.Metadata, mtype.Description, mtype.Cod)
	return err
}

func (contentStorage *ContentStoragePostgres) GetType(cod string) (model.Type, error) {

	strQuery := `SELECT cod, metadata, description FROM headless_cms.content_type WHERE cod=$1;`
	rows, errQuery := contentStorage.dbConnection.Query(context.Background(), strQuery, cod)
	mtype := model.Type{}

	defer rows.Close()
	if errQuery != nil {
		return mtype, errQuery
	}
	for rows.Next() {
		rows.Scan(&mtype.Cod, &mtype.Metadata, &mtype.Description)
	}
	return mtype, nil
}
