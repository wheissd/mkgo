package main

import (
	"encoding/xml"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/go-faster/errors"
)

type githubReleases struct {
	Entry []struct {
		Author struct {
			Name string `xml:"name"`
		} `xml:"author"`
		Content struct {
			Type     string `xml:"type,attr"`
			CharData string `xml:",chardata"`
		} `xml:"content"`
		ID   string `xml:"id"`
		Link struct {
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
		Thumbnail struct {
			Height int    `xml:"height,attr"`
			Url    string `xml:"url,attr"`
			Width  int    `xml:"width,attr"`
		} `xml:"thumbnail"`
		Title   string    `xml:"title"`
		Updated time.Time `xml:"updated"`
	} `xml:"entry"`
	ID   string `xml:"id"`
	Link []struct {
		Href string `xml:"href,attr"`
		Rel  string `xml:"rel,attr"`
		Type string `xml:"type,attr"`
	} `xml:"link"`
	Title   string    `xml:"title"`
	Updated time.Time `xml:"updated"`
}

func (cmd *cmd) checkNewGithubVersionAvailable(currentVersion string, project string) error {
	availableVersionsRawResponse, err := cmd.depCheckClient.Get(fmt.Sprintf("https://github.com/%s/releases.atom", project))
	if err != nil {
		return errors.Wrap(err, "http.Get atlas available versions")
	}

	d := xml.NewDecoder(availableVersionsRawResponse.Body)
	var releases githubReleases
	err = d.Decode(&releases)
	if err != nil {
		return errors.Wrap(err, "xml.Decode")
	}

	var latestVersion string
	for _, release := range releases.Entry {
		if latestVersion == "" {
			latestVersion = release.Title
			continue
		}

		if release.Title > latestVersion {
			latestVersion = release.Title
		}
	}

	if currentVersion < latestVersion {
		cmd.logRed("there is new atlas version available")
	}

	return nil
}

func (cmd *cmd) checkVersion(dep dependency) error {
	vOutput, err := cmd.runCmd(dep.getVersionCmd)
	if err != nil {
		return errors.Wrap(err, "cmd.runCmd(dep.getVersionCmd)")
	}

	if !strings.HasPrefix(string(vOutput), dep.getVersionPrefix) {
		return errors.New(fmt.Sprintf("%s not found", dep.name))
	}

	vSplitted := strings.Split(strings.TrimPrefix(string(vOutput), dep.getVersionPrefix), "\n")
	currentVersion := vSplitted[0]

	err = cmd.checkNewGithubVersionAvailable(currentVersion, dep.project)
	if err != nil {
		return errors.Wrap(err, "cmd.checkNewGithubVersionAvailable")
	}

	return nil
}

type dependency struct {
	name               string
	installCmd         string
	checkHasNewVersion func() error
	getVersionCmd      string
	getVersionPrefix   string
	project            string
}

func (cmd *cmd) deps() []dependency {
	return []dependency{
		{
			name:             "ogen",
			project:          "ogen-go/ogen",
			getVersionCmd:    "ogen -version",
			getVersionPrefix: "ogen version ",
			installCmd:       "go install github.com/ogen-go/ogen/cmd/ogen@latest",
		},
		{
			name:             "atlas",
			project:          "ariga/atlas",
			getVersionCmd:    "atlas version",
			getVersionPrefix: "atlas version ",
			installCmd:       "curl -sSf https://atlasgo.sh | ATLAS_VERSION=v0.19.3-cfa638c-canary sh",
		},
	}
}

func (cmd *cmd) checkDependencies(info *runInfo) error {
	// check dependency updates once a week
	if info.LastDepCheck.After(time.Now().Add(-time.Hour * 24 * 7)) {
		return nil
	}
	info.LastDepCheck = time.Now()
	for _, dep := range cmd.deps() {
		if dep.checkHasNewVersion != nil {
			if err := dep.checkHasNewVersion(); err != nil {
				cmd.logger.Error("dep.checkHasNewVersion()", slog.Any("error", err))
			}
		}
		if err := cmd.checkVersion(dep); err != nil {
			cmd.logger.Error("cmd.checkVersion(dep)", slog.Any("error", err))
		}
	}

	return nil
}
