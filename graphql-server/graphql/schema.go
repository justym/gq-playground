package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/justym/gq-playground/graphql-server/model"
)

var (
	players []model.Player
)

func init() {
	players = model.NewPlayers()
}

func NewSchema() (*graphql.Schema, error) {
	fields := graphql.Fields{
		"players": &graphql.Field{
			Type:        graphql.NewList(playerObject),
			Description: "All players",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return players, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, err
	}

	return &schema, nil
}
