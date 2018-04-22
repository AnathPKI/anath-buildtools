/*
Constants used in travis package
*/

package travis

import (
	"log"
	"os"
)

/*
Travis Constants
*/
const (
	TravisAPIBaseURL              string = "https://api.travis-ci.org"
	TravisAPIHeaderName           string = "Travis-API-Version"
	TravisAPIVersion3             string = "3"
	TravisAPIBranchQueryParameter string = "branch.name"
	TravisAPISortByQueryParameter string = "sort_by"
	TravisAPILimitQueryParameter  string = "limit"
	AuthenticationTokenEnvVar     string = "TRAVIS_TOKEN"
	ApplicationJSONMime           string = "application/json"
)

/*
AuthenticationToken returns the Travis Authentication token from the environment variable $TRAVIS_TOKEN
*/
func AuthenticationToken() string {
	value, ok := os.LookupEnv(AuthenticationTokenEnvVar)
	if !ok {
		log.Fatalf("'%s' environment variable not set", AuthenticationTokenEnvVar)
	}

	return value
}
