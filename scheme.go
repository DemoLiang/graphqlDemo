package main

import "github.com/graphql-go/graphql"

var GSchema, err = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    Query,
		Mutation: Mutation,
	})

var Query = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Query",
		Description: "查询",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type:        GLUser,
				Description: "查询用户信息",
				//Args: graphql.FieldConfigArgument{
				//	"name": &graphql.ArgumentConfig{
				//		Type:        graphql.NewNonNull(graphql.String),
				//		Description: "用户姓名",
				//	},
				//	"phone": &graphql.ArgumentConfig{
				//		Type:        graphql.String,
				//		Description: "电话号码",
				//	},
				//},
				Resolve: QueryUserResolve,
			},
		}})

var Mutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Mutation",
		Description: "更新信息",
		Fields: graphql.Fields{
			"create_or_edit_user": &graphql.Field{
				Type:        GLUser,
				Description: "创建或者更新用户信息",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.String),
						Description: "名字",
					},
					"phone": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.String),
						Description: "电话",
					},
					"age": &graphql.ArgumentConfig{
						Type:        graphql.Int,
						Description: "电话",
					},
					"sex": &graphql.ArgumentConfig{
						Type:        GLUserSex,
						Description: "性别",
					},
				},
				Resolve: MutationUserResolve,
			},
			"add_article": &graphql.Field{
				Type:        GLArticle,
				Description: "创建文章",
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.String),
						Description: "标题",
					},
					"detail": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.String),
						Description: "详情",
					},
				},
				Resolve: MutationAddArticleResolve,
			},
		},
	})
