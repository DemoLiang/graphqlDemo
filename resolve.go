package main

import (
	"github.com/graphql-go/graphql"
	"github.com/segmentio/ksuid"
	"log"
)

var gUser User

//查询用户
func QueryUserResolve(p graphql.ResolveParams) (interface{}, error) {
	log.Printf("call QueryUserResolve")
	name := p.Args["name"].(string)
	if gUser.UserInfo.Name == name {
		return &gUser, nil
	}
	phone := p.Args["phone"].(string)
	if gUser.UserInfo.Phone == phone {
		return &gUser, nil
	}

	return &User{
		UserInfo: UserBasicInfo{
			Name:  "default",
			Phone: "88888",
			Sex:   Male,
		},
	}, nil
}

//查询用户列表
func QueryUserBasicResolve(p graphql.ResolveParams) (interface{}, error) {
	log.Printf("call QueryUserBasicResolve")
	return &gUser.UserInfo, nil
}

//查询用户的朋友
func articlesResolve(p graphql.ResolveParams) (interface{}, error) {
	log.Printf("call articlesResolve")
	return gUser.Articles, nil
}

//修改用户信息
func MutationUserResolve(p graphql.ResolveParams) (interface{}, error) {
	log.Printf("call MutationUserResolve")

	if name, ok := p.Args["name"].(string); ok {
		gUser.UserInfo.Name = name
	}
	if phone, ok := p.Args["phone"].(string); ok {
		gUser.UserInfo.Phone = phone
	}
	if age, ok := p.Args["age"].(int); ok {
		gUser.UserInfo.Age = age
	}
	if sex, ok := p.Args["sex"].(Sex); ok {
		gUser.UserInfo.Sex = sex
	}
	gUser.UserInfo.Id = ksuid.New().String()
	return &gUser, nil
}

//创建文章
func MutationAddArticleResolve(p graphql.ResolveParams) (interface{}, error) {
	log.Printf("call MutationAddArticleResolve")
	var article Article
	if title, ok := p.Args["title"].(string); ok {
		article.Title = title
	}
	if detail, ok := p.Args["detail"].(string); ok {
		article.Detail = detail
	}
	article.Id = ksuid.New().String()
	gUser.Articles = append(gUser.Articles, article)
	return &article, nil
}
