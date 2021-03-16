package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/qri-io/jsonschema"
)

type Type struct {
	Cod         string `json:"cod"`
	Metadata    []byte `json:"metadata"`
	Description string `json:"description"`
}

func (mtype Type) Validator(ctx context.Context , jsonBytes []byte) (bool, []string) {


	var listErrors = make([]string, 0)
	rs := &jsonschema.Schema{}

	if err := json.Unmarshal(mtype.Metadata, rs); err != nil {
		listErrors = append(listErrors, err.Error())
		return false, listErrors
	}


	fmt.Println( string(jsonBytes))
	errs, err := rs.ValidateBytes(ctx, jsonBytes)
	if err != nil {
		listErrors = append(listErrors, err.Error())
		return false, listErrors
	}

	if len(errs) > 0 {
		for _, e := range errs {
			listErrors = append(listErrors, e.Error())
		}
		return false, listErrors
	}

	return true, nil
}

func (mtype Type) Validate() error {

	//validate schema
	if err := json.Unmarshal(mtype.Metadata, &jsonschema.Schema{}); err != nil {
		return err
	}
	return nil
}
