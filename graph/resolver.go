package graph

import "github.com/sasakitimaru/git_gql/graph/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	srv services.Services
}
