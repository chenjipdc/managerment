package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"io/ioutil"
)

const (
	username = "xxx@xxx"
	password = "xxx"
	user = "xxx"
	owner = user
)

var ignores = []string{
	"chenjipdc.github.io",
	"18.06-linalg-notes",
	"algorithms",
	"config-consul",
	"CrashLog",
	"golang-notes",
	"learn-regex",
}

var client *github.Client

func init()  {
	tp := github.BasicAuthTransport{
		Username:username,
		Password:password}
	client = github.NewClient(tp.Client())
}

func main() {
	// 删除repo
	for _, repo := range repos() {
		name := *repo.Name
		keep := false
		for _, ignore := range ignores {
			if name == ignore {
				keep = true
				break
			}
		}
		if !keep {
			deleteRepo(name)
		}
	}
}

// 获取repositories
func repos() []*github.Repository {
	ctx := context.Background()

	page := 1
	perPage := 50
	rps := new([]*github.Repository)

	for {
		var repositories []*github.Repository
		repositories, _, err := client.Repositories.List(ctx, user, &github.RepositoryListOptions{
			ListOptions: github.ListOptions{
				Page:page,
				PerPage:perPage}})
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fmt.Println("page:", page)
		if len(repositories) == 0 {
			break
		}
		page++
		*rps = append(*rps, repositories...)
	}
	return *rps
}

// 删除repo
func deleteRepo(repo string) {
	ctx := context.Background()
	response, err := client.Repositories.Delete(ctx, owner, repo)
	if err != nil {
		fmt.Println(err)
	}
	b, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(b))
}


// 获取start
func stars() []*github.StarredRepository {
	ctx := context.Background()

	page := 1
	perPage := 50
	stRepos := new([]*github.StarredRepository)

	for {
		starredRepos, _, err := client.Activity.ListStarred(ctx, owner, &github.ActivityListStarredOptions{
			ListOptions: github.ListOptions{
				Page: page,
				PerPage: perPage}})
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("page:", page)
		if len(starredRepos) == 0 {
			break
		}
		page++
		*stRepos = append(*stRepos, starredRepos...)
	}
	return *stRepos
}

// unstart
func unstar(repo string)  {
	ctx := context.Background()
	_, err := client.Activity.Unstar(ctx, owner, repo)
	if err != nil {
		fmt.Println(err)
	}
}

// start
func star(repo string)  {
	ctx := context.Background()
	_, err := client.Activity.Star(ctx, owner, repo)
	if err != nil {
		 fmt.Println(err)
	}
}