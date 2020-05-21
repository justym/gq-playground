package main

import (
	"github.com/justym/gq-playground/errors"
	"github.com/justym/gq-playground/model"
	"github.com/justym/gq-playground/service"
)

const endpoint = "https://api.github.com/graphql"

func main() {
	MITlicense := service.NewLisenseClient(endpoint, model.MITQuery)
	Apache2License := service.NewLisenseClient(endpoint, model.Apache2Query)
	err := MITlicense.Do()
	errors.Check(err, "FAILED")
	err = Apache2License.Do()
	errors.Check(err, "FAILED")
}
