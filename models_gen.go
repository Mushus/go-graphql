// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package go_graphql

type Category struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Documents []*Document `json:"documents"`
}

type Document struct {
	ID         string      `json:"id"`
	Title      string      `json:"title"`
	Body       string      `json:"body"`
	Categories []*Category `json:"categories"`
}

type NewCategory struct {
	Name string `json:"name"`
}

type NewDocumument struct {
	Title      string   `json:"title"`
	Body       string   `json:"body"`
	Categories []string `json:"categories"`
}

type NewWebSite struct {
	Name string `json:"name"`
}

type WebSite struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Documents []*Document `json:"documents"`
}
