#!/bin/bash
#
# Create bintray-deploy.json as expected by .travis.yml.
#
# Call this script from the repository root

TRAVIS_TAG=${TRAVIS_TAG:-tag expected}

if ! echo "${TRAVIS_TAG}" | grep -E "^v[0-9]+\.[0-9]+\.[0-9]+\$" >/dev/null 2>&1
then
    echo 'Tag does not match "^v[0-9]+.[0-9]+.[0-9]+$". Assuming non-release tag and doing nothing.'
    # Don't make the job fail. Maybe it's a legit non-release tag
    exit 0
fi

SEMANTIC_VERSION=${TRAVIS_TAG#v}

cat > bintray-deploy.json <<EOF
{
    "package": {
        "name": "go-binaries",
        "repo": "anathpki/anath-buildtools",
        "desc": "I was pushed completely automatically",
        "website_url": "https://github.com/AnathPKI/anath-buildtools",
        "vcs_url": "https://github.com/AnathPKI/anath-buildtools.git",
        "licenses": ["BSD 2-Clause"]
    },
    "version": {
        "name": "${SEMANTIC_VERSION}",
        "vcs_tag": "${TRAVIS_TAG}",
        "gpgSign": false
    },

    "files":
        [
        {"includePattern": "\./(restart-build|trigger-build)", "uploadPattern": "${TRAVIS_TAG}/\$1"}
        ],
    "publish": true
}
EOF