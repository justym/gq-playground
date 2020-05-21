package graphql

import "github.com/graphql-go/graphql"

var playerObject = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Player",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"highScore": &graphql.Field{
				Type: graphql.String,
			},
			"isOnline": &graphql.Field{
				Type: graphql.Boolean,
			},
			"location": &graphql.Field{
				Type: graphql.String,
			},
			"levelsUnlocked": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)
