package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "127.0.0.1"
	DB_PORT     = "5433"
	DB_USER     = "postgres"
	DB_PASSWORD = "days"
	DB_NAME     = "go_graphql_db"
)

type Author struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Word struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func indexHandler(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	}
}

func main() {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)

	defer db.Close()

	authorType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Author",
		Description: "An author",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The identifier of the author.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if author, ok := p.Source.(*Author); ok {
						return author.ID, nil
					}

					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the author.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if author, ok := p.Source.(*Author); ok {
						return author.Name, nil
					}

					return nil, nil
				},
			},
			"created_at": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The created_at date of the author.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if author, ok := p.Source.(*Author); ok {
						return author.CreatedAt, nil
					}

					return nil, nil
				},
			},
		},
	})

	wordType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Word",
		Description: "A Word",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The identifier of the word.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if word, ok := p.Source.(*Word); ok {
						return word.ID, nil
					}

					return nil, nil
				},
			},
			"content": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The content of the word.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if word, ok := p.Source.(*Word); ok {
						return word.Content, nil
					}

					return nil, nil
				},
			},
			"created_at": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The created_at date of the word.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if word, ok := p.Source.(*Word); ok {
						return word.CreatedAt, nil
					}

					return nil, nil
				},
			},
			"author": &graphql.Field{
				Type: authorType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if word, ok := p.Source.(*Word); ok {
						author := &Author{}
						err = db.QueryRow("select id, name from words where id = $1", word.AuthorID).Scan(&author.ID, &author.Name)
						checkErr(err)

						return author, nil
					}

					return nil, nil
				},
			},
		},
	})

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"author": &graphql.Field{
				Type:        authorType,
				Description: "Get an author.",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)

					author := &Author{}
					err = db.QueryRow("select id, name, from authors where id = $1", id).Scan(&author.ID, &author.Name)
					checkErr(err)

					return author, nil
				},
			},
			"authors": &graphql.Field{
				Type:        graphql.NewList(authorType),
				Description: "List of authors.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					rows, err := db.Query("SELECT id, name FROM authors")
					checkErr(err)
					var authors []*Author

					for rows.Next() {
						author := &Author{}

						err = rows.Scan(&author.ID, &author.Name)
						checkErr(err)
						authors = append(authors, author)
					}

					return authors, nil
				},
			},
			"word": &graphql.Field{
				Type:        wordType,
				Description: "Get a word.",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)

					word := &Word{}
					err = db.QueryRow("select id, content, author_id from words where id = $1", id).Scan(&word.ID, &word.Content, &word.AuthorID)
					checkErr(err)

					return word, nil
				},
			},
			"words": &graphql.Field{
				Type:        graphql.NewList(wordType),
				Description: "List of words.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					rows, err := db.Query("SELECT id, content, author_id FROM words")
					checkErr(err)
					var words []*Word

					for rows.Next() {
						word := &Word{}

						err = rows.Scan(&word.ID, &word.Content, &word.AuthorID)
						checkErr(err)
						words = append(words, word)
					}

					return words, nil
				},
			},
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			// Author
			"createAuthor": &graphql.Field{
				Type:        authorType,
				Description: "Create new author",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					name, _ := params.Args["name"].(string)
					createdAt := time.Now()

					var lastInsertId int
					err = db.QueryRow("INSERT INTO authors(name, created_at) VALUES($1, $2) returning id;", name, createdAt).Scan(&lastInsertId)
					checkErr(err)

					newAuthor := &Author{
						ID:        lastInsertId,
						Name:      name,
						CreatedAt: createdAt,
					}

					return newAuthor, nil
				},
			},
			"updateAuthor": &graphql.Field{
				Type:        authorType,
				Description: "Update an author",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)
					name, _ := params.Args["name"].(string)

					stmt, err := db.Prepare("UPDATE authors SET name = $1 WHERE id = $2")
					checkErr(err)

					_, err2 := stmt.Exec(name, id)
					checkErr(err2)

					newAuthor := &Author{
						ID:   id,
						Name: name,
					}

					return newAuthor, nil
				},
			},
			"deleteAuthor": &graphql.Field{
				Type:        authorType,
				Description: "Delete an author",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)

					stmt, err := db.Prepare("DELETE FROM authors WHERE id = $1")
					checkErr(err)

					_, err2 := stmt.Exec(id)
					checkErr(err2)

					return nil, nil
				},
			},
			// Word
			"createWord": &graphql.Field{
				Type:        wordType,
				Description: "Create new word",
				Args: graphql.FieldConfigArgument{
					"content": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"author_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					content, _ := params.Args["content"].(string)
					authorId, _ := params.Args["author_id"].(int)
					createdAt := time.Now()

					var lastInsertId int
					err = db.QueryRow("INSERT INTO words(content, author_id, created_at) VALUES($1, $2, $3) returning id;", content, authorId, createdAt).Scan(&lastInsertId)
					checkErr(err)

					newWord := &Word{
						ID:        lastInsertId,
						Content:   content,
						AuthorID:  authorId,
						CreatedAt: createdAt,
					}

					return newWord, nil
				},
			},
			"updateWord": &graphql.Field{
				Type:        wordType,
				Description: "Update a word",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"content": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"author_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)
					content, _ := params.Args["content"].(string)
					authorId, _ := params.Args["author_id"].(int)

					stmt, err := db.Prepare("UPDATE words SET content = $1, author_id = $2 WHERE id = $3")
					checkErr(err)

					_, err2 := stmt.Exec(content, authorId, id)
					checkErr(err2)

					newWord := &Word{
						ID:       id,
						Content:  content,
						AuthorID: authorId,
					}

					return newWord, nil
				},
			},
			"deleteWord": &graphql.Field{
				Type:        wordType,
				Description: "Delete a word",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)

					stmt, err := db.Prepare("DELETE FROM words WHERE id = $1")
					checkErr(err)

					_, err2 := stmt.Exec(id)
					checkErr(err2)

					return nil, nil
				},
			},
		},
	})
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// serve HTTP
	// http.Handle("/graphql", h)
	http.HandleFunc("/graphql", indexHandler(h))

	http.ListenAndServe(":8080", nil)
}
