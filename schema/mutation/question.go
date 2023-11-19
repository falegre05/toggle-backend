package mutation

import (
	"github.com/graphql-go/graphql"
	"github.com/togglhire/backend-homework/resolvers"
	"github.com/togglhire/backend-homework/schema/output"
	"github.com/togglhire/backend-homework/security"
)

var AddQuestion = &graphql.Field{
	Type: output.Question,
	Args: graphql.FieldConfigArgument{
		"body": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: security.Check(security.PermissionsUser, func(p graphql.ResolveParams) (interface{}, error) {
		return resolvers.AddQuestion(p.Context, p.Args)
	}),
}

var UpdateQuestion = &graphql.Field{
	Type: output.Question,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"body": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: security.Check(security.PermissionsUser, func(p graphql.ResolveParams) (interface{}, error) {
		return resolvers.UpdateQuestion(p.Context, p.Args)
	}),
}

var DeleteQuestion = &graphql.Field{
	Type: graphql.Boolean,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: security.Check(security.PermissionsUser, func(p graphql.ResolveParams) (interface{}, error) {
		return resolvers.DeleteQuestion(p.Context, p.Args)
	}),
}
