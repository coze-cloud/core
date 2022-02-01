package core

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var GraphQlModule = fx.Provide(
	fx.Annotated{
		Name:   "query",
		Target: newGraphQlQuery,
	},
	fx.Annotated{
		Name:   "mutation",
		Target: newGraphQlMutation,
	},
	newGraphQlSchema,
)

func newGraphQlQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Description: "The application's root query object",
		Fields: graphql.Fields{},
	})
}

func newGraphQlMutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Description: "The application's root mutation object",
		Fields: graphql.Fields{},
	})
}

type graphqlDependencies struct {
	fx.In

	Query *graphql.Object `name:"query"`
	Mutation *graphql.Object `name:"mutation"`
}

func newGraphQlSchema(dependencies graphqlDependencies, logger *logrus.Logger) *graphql.Schema {
	query := dependencies.Query
	mutation := dependencies.Query

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: func() *graphql.Object {
			if len(query.Fields()) <= 0 {
				return nil
			}
			return query
		}(),
		Mutation: func() *graphql.Object {
			if len(mutation.Fields()) <= 0 {
				return nil
			}
			return mutation
		}(),
	})
	if err != nil {
		logger.Error(err)
	}

	return &schema
}

func UseGraphQl(schema *graphql.Schema, server *echo.Echo, logger *logrus.Logger) {
	graphqlHandler := handler.New(&handler.Config{
		Schema: schema,
	})

	echoHandler := echo.WrapHandler(graphqlHandler)
	server.GET("/graphql", echoHandler)
	server.POST("/graphql", echoHandler)
}

func UseGraphQlWithPlayground(schema *graphql.Schema, server *echo.Echo, logger *logrus.Logger) {
	graphqlHandler := handler.New(&handler.Config{
		Schema: schema,
		Playground: true,
	})

	echoHandler := echo.WrapHandler(graphqlHandler)
	server.GET("/graphql", echoHandler)
	server.POST("/graphql", echoHandler)

	logger.Info("Playground can be accessed on route /graphql")
}


