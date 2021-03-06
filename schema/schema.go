package schema

import (
	"io/ioutil"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/scottjr632/graphql-chat/schema/resolver"
)

const defaultQqlFileName = "schema.gql"

func getGqlSchema(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", nil
	}
	return string(data), nil
}

func mustGetQqlSchema(filePath string) string {
	schema, err := getGqlSchema(filePath)
	if err != nil {
		panic(err)
	}
	return schema
}

// New returns a new graphql schema
func New() (*graphql.Schema, error) {
	schema := mustGetQqlSchema(defaultQqlFileName)
	s, err := graphql.ParseSchema(schema, resolver.New())
	return s, err
}
