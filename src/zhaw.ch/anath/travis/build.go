package travis

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type buildsJSONDoc struct {
	Builds []struct {
		Id int64
	}
}

type buildRequest struct {
	Request struct {
		Branch  string
		Message string
	}
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

/*
TriggerBuild triggers a build of the given branch and repository
*/
func TriggerBuild(repoSlug, branchName, message string) {
	requestsURL := composeRequestsURL(repoSlug)
	travisRequest := newTravisRequest(http.MethodPost, requestsURL)

	log.Printf("Trigger build of branch '%s' in repo '%s': %s'", branchName, repoSlug, requestsURL)

	var postBodyStructure = buildRequest{}
	postBodyStructure.Request.Branch = branchName
	postBodyStructure.Request.Message = message

	postBody, err := json.Marshal(postBodyStructure)
	if err != nil {
		log.Fatal(err)
	}

	travisRequest.Body = nopCloser{bytes.NewBuffer(postBody)}

	client := &http.Client{}
	resp, err := client.Do(travisRequest)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		log.Fatalf("Triggering build resulted in HTTP Status Code %d", resp.StatusCode)
	}
}

/*
RestartBuild restarts the build with the given ID
*/
func RestartBuild(buildID string) {
	restartURL, err := url.Parse(TravisAPIBaseURL + "/build/" + buildID + "/restart")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Restart build %s: %s", buildID, restartURL)

	rebuildRequest := newTravisRequest(http.MethodPost, restartURL)

	client := &http.Client{}
	resp, err := client.Do(rebuildRequest)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		log.Fatalf("Restart of build %s resulted in HTTP Status Code %d", buildID, resp.StatusCode)
	}
}

func newTravisRequest(method string, apiURL *url.URL) *http.Request {
	// We will set the URL further down
	req, err := http.NewRequest(method, "", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.URL = apiURL
	req.Header.Add(TravisAPIHeaderName, TravisAPIVersion3)
	req.Header.Add("Authorization", "token "+AuthenticationToken())
	req.Header.Add("Content-Type", ApplicationJSONMime)
	req.Header.Add("Accept", ApplicationJSONMime)

	return req
}

/*
LastBuildIDForBranch gets the ID of the last build for a given repository slug and branch.
*/
func LastBuildIDForBranch(repoSlug, branchName string) string {
	apiURL := composeBuildsURL(repoSlug, branchName)
	log.Printf("Request '%s'", apiURL)

	buildIDRequest := newTravisRequest(http.MethodGet, apiURL)

	client := &http.Client{}
	resp, err := client.Do(buildIDRequest)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Received HTTP Status Code %d from Travis", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var b buildsJSONDoc

	err = json.Unmarshal(body, &b)
	if err != nil {
		log.Fatal(err)
	}

	return strconv.FormatInt(b.Builds[0].Id, 10)
}

func composeBuildsURL(repoSlug, branchName string) *url.URL {
	var travisURL *url.URL
	travisURL, err := url.Parse(TravisAPIBaseURL)
	if err != nil {
		log.Fatal(err)
	}

	travisURL.Path = "/repo/" + repoSlug + "/builds"
	travisURL.RawPath = "/repo/" + url.PathEscape(repoSlug) + "/builds"

	parameters := url.Values{}
	parameters.Add(TravisAPILimitQueryParameter, "1")
	parameters.Add(TravisAPIBranchQueryParameter, branchName)
	parameters.Add(TravisAPISortByQueryParameter, "started_at:desc")
	travisURL.RawQuery = parameters.Encode()

	return travisURL
}

func composeRequestsURL(repoSlug string) *url.URL {
	var travisURL *url.URL
	travisURL, err := url.Parse(TravisAPIBaseURL)
	if err != nil {
		log.Fatal(err)
	}

	travisURL.Path = "/repo/" + repoSlug + "/requests"
	travisURL.RawPath = "/repo/" + url.PathEscape(repoSlug) + "/requests"
	return travisURL
}
