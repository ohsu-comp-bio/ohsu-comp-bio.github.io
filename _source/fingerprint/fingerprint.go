package main

import (
  "context"
  "fmt"
  "github.com/google/go-github/github"
)

func main() {
  ctx := context.Background()
  cli := github.NewClient(nil)

  org, _, _ := cli.Organizations.Get(ctx, "ohsu-comp-bio")

  opt := &github.RepositoryListByOrgOptions{
    ListOptions: github.ListOptions{PerPage: 100},
  }

  var allRepos []*github.Repository
  for {
    repos, resp, _ := cli.Repositories.ListByOrg(ctx, "ohsu-comp-bio", opt)
    allRepos = append(allRepos, repos...)
    if resp.NextPage == 0 {
      break
    }
    opt.ListOptions.Page = resp.NextPage
  }


  for _, r := range allRepos {
    fmt.Println(*r.Name)
  }
}
