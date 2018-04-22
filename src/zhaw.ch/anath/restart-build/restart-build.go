package main

import (
	"flag"

	"zhaw.ch/anath/travis"
)

func main() {
	repoSlugPtr := flag.String("repository", "", "GitHub repository name, e.g. AnathPKI/demo")
	branchNamePtr := flag.String("branch", "master", "Branch name")

	flag.Parse()

	lastBuildID := travis.LastBuildIDForBranch(*repoSlugPtr, *branchNamePtr)
	travis.RestartBuild(lastBuildID)
}
