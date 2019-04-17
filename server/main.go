package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
)

type Todo struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos []Todo

var todoType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Todo",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"userId": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"completed": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			/* Get (read) single todo by id
			   http://localhost:8080/graphql?query={todo(id:1){title,userId,completed}}
			*/
			"todo": &graphql.Field{
				Type:        todoType,
				Description: "Get todo by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						// Find todo
						for _, todo := range todos {
							if int(todo.ID) == id {
								return todo, nil
							}
						}
					}
					return nil, nil
				},
			},
			/* Get (read) todo list
			   http://localhost:8080/graphql?query={list{id,userId,title,completed}}
			*/
			"list": &graphql.Field{
				Type:        graphql.NewList(todoType),
				Description: "Get todo list",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return todos, nil
				},
			},
		},
	})

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		/* Create new todo item
		http://localhost:8080/graphql?query=mutation+_{create(name:"Inca Kola",info:"Inca Kola is a soft drink that was created in Peru in 1935 by British immigrant Joseph Robinson Lindley using lemon verbena (wiki)",price:1.99){id,name,info,price}}
		*/
		"create": &graphql.Field{
			Type:        todoType,
			Description: "Create new todo",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"info": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"price": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				rand.Seed(time.Now().UnixNano())
				todo := Todo{
					ID:        int64(rand.Intn(100000)), // generate random ID
					Title:     params.Args["title"].(string),
					UserID:    params.Args["userId"].(int64),
					Completed: params.Args["completed"].(bool),
				}
				todos = append(todos, todo)
				return todo, nil
			},
		},

		/* Update todo by id
		   http://localhost:8080/graphql?query=mutation+_{update(id:1,userId:5){id,title,userId,completed}}
		*/
		"update": &graphql.Field{
			Type:        todoType,
			Description: "Update todo by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"userId": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"title": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"completed": &graphql.ArgumentConfig{
					Type: graphql.Boolean,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(int)
				title, titleOk := params.Args["title"].(string)
				completed, completedOk := params.Args["completed"].(bool)
				userID, userIDOk := params.Args["userId"].(int64)
				// info, infoOk := params.Args["info"].(string)
				// price, priceOk := params.Args["price"].(float64)
				todo := Todo{}
				for i, t := range todos {
					if int64(id) == t.ID {
						if titleOk {
							todos[i].Title = title
						}
						if completedOk {
							todos[i].Completed = completed
						}
						if userIDOk {
							todos[i].UserID = userID
						}
						todo = todos[i]
						break
					}
				}
				return todo, nil
			},
		},

		/* Delete todo by id
		   http://localhost:8080/graphql?query=mutation+_{delete(id:1){id,title,userId,completed}}
		*/
		"delete": &graphql.Field{
			Type:        todoType,
			Description: "Delete todo by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(int)
				todo := Todo{}
				for i, p := range todos {
					if int64(id) == p.ID {
						todo = todos[i]
						// Remove from todo list
						todos = append(todos[:i], todos[i+1:]...)
					}
				}

				return todo, nil
			},
		},
	},
})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func initTodosData(p *[]Todo) {
	todo1 := Todo{ID: 1, Title: "Chicha Morada", UserID: 1, Completed: true}
	todo2 := Todo{ID: 2, Title: "Chicha de jora", UserID: 2, Completed: false}
	todo3 := Todo{ID: 3, Title: "Pisco", UserID: 3, Completed: true}
	*p = append(*p, todo1, todo2, todo3)
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

func handleTodo(w http.ResponseWriter, r *http.Request) {
	result := executeQuery(r.URL.Query().Get("query"), schema)
	json.NewEncoder(w).Encode(result)
}

func main() {
	// Primary data initialization
	initTodosData(&todos)

	todoHandler := http.HandlerFunc(handleTodo)

	http.HandleFunc("/graphql", indexHandler(todoHandler))

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
