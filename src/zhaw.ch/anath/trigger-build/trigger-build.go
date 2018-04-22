package main

import (
	"flag"

	"zhaw.ch/anath/travis"
)

func main() {
	repoSlugPtr := flag.String("repository", "", "GitHub repository name, e.g. AnathPKI/demo")
	branchNamePtr := flag.String("branch", "master", "Branch name")
	messagePtr := flag.String("message", "Triggered via API", "Message to Travis")

	flag.Parse()

	travis.TriggerBuild(*repoSlugPtr, *branchNamePtr, *messagePtr)
}
