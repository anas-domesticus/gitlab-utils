package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/xanzy/go-gitlab"
)

func main() {
	initialiseLogger()
	log.Info("Starting application...")
	appConfig := NewConfig()
	git, err := gitlab.NewClient(appConfig.Token)
	if err != nil {
		log.Fatal(err)
	}
	t := true
	f := false
	p, _, err := git.Projects.ListProjects(&gitlab.ListProjectsOptions{Archived: &f, Owned: &t})
	if err != nil {
		log.Fatal(fmt.Sprintf("Problem getting projects %s", err.Error()))
	}
	for _, v := range p {
		pl, _, err := git.Pipelines.ListProjectPipelines(v.ID, &gitlab.ListProjectPipelinesOptions{
			Status: gitlab.BuildState("pending"),
		})
		if err != nil {
			log.Fatal(fmt.Sprintf("Problem getting pipelines for project %s: %s", v.Name, err.Error()))
		}
		if len(pl) == 0 {
			log.Info(fmt.Sprintf("Project %s has no pending pipelines", v.Name))
		} else {
			log.Info(fmt.Sprintf("Project %s has pending pipelines", v.Name))
			for _, vpl := range pl {
				j, _, _ := git.Jobs.ListPipelineJobs(v.ID, vpl.ID, &gitlab.ListJobsOptions{})
				for _, vj := range j {
					log.Info(vj.Name, vj.Tag)
				}
			}
		}
	}
}
