package main

import "github.com/graphql-go/graphql"

type Sex uint

type UserBasicInfo struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
	Sex   Sex    `json:"sex"`
}

type Article struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type User struct {
	UserInfo UserBasicInfo `json:"user_info"`
	Articles []Article     `json:"articles"`
}

var GLUserBasicInfoConfig = graphql.ObjectConfig{
	Name:        "UserBasicInfo",
	Description: "用户基础信息",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: "ID",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "姓名",
		},
		"age": &graphql.Field{
			Type:        graphql.Int,
			Description: "年龄",
		},
		"phone": &graphql.Field{
			Type:        graphql.String,
			Description: "电话",
		},
		"sex": &graphql.Field{
			Type:        GLUserSex,
			Description: "性别",
		},
	},
}

var GLUserBasicInfo = graphql.NewObject(GLUserBasicInfoConfig)

var GLArticleConfig = graphql.ObjectConfig{
	Name:        "Article",
	Description: "文章信息",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: "id",
		},
		"title": &graphql.Field{
			Type:        graphql.String,
			Description: "标题",
		},
		"detail": &graphql.Field{
			Type:        graphql.String,
			Description: "文章详情",
		},
	},
}

var GLArticle = graphql.NewObject(GLArticleConfig)

var GLUserConfig = graphql.ObjectConfig{
	Name:        "User",
	Description: "用户信息",
	Fields: graphql.Fields{
		"user_info": &graphql.Field{
			Type:        GLUserBasicInfo,
			Description: "用户基础信息",
			//Resolve:     QueryUserBasicResolve,
		},
		"articles": &graphql.Field{
			Type:        graphql.NewList(GLArticle),
			Description: "文章列表",
			Resolve:     articlesResolve,
		},
	},
}

var GLUser = graphql.NewObject(GLUserConfig)

const (
	Unknown Sex = iota
	Male
	Female
)

var GLUserSex = graphql.NewEnum(graphql.EnumConfig{
	Name:        "UserSex",
	Description: "用户性别",
	Values: graphql.EnumValueConfigMap{
		"Other": &graphql.EnumValueConfig{
			Value:       Unknown,
			Description: "其他",
		},
		"Male": &graphql.EnumValueConfig{
			Value:       Male,
			Description: "男性",
		},
		"Female": &graphql.EnumValueConfig{
			Value:       Female,
			Description: "女性",
		},
	},
})
