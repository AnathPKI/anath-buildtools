package travis

import (
	"testing"
)

func TestComposeBuildsURL(t *testing.T) {
	actual := composeBuildsURL("test/slug", "release/1.0.0").String()
	expected := TravisAPIBaseURL + "/repo/test%2Fslug/builds?branch.name=release%2F1.0.0&limit=1&sort_by=started_at%3Adesc"
	if actual != expected {
		t.Errorf("'%s' does not match expected '%s'", actual, expected)
	}
}

func TestComposeRequestsURL(t *testing.T) {
	actual := composeRequestsURL("test/slug").String()
	expected := TravisAPIBaseURL + "/repo/test%2Fslug/requests"
	if actual != expected {
		t.Errorf("'%s' does not match expected '%s'", actual, expected)
	}
}
