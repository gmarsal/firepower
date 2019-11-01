package firepower

import (
	"context"
)

// ObjectService handles communication with the object related methods of the
// GitHub API docs: https://developer.github.com/v3/enterprise/
type ObjectService service

// Object represents a GitHub issue on a repository.
//
// Note: As far as the GitHub API is concerned, every pull request is an issue,
// but not every issue is a pull request. Some endpoints, events, and webhooks
// may also return pull requests via this struct. If PullRequestLinks is nil,
// this is an issue, and if PullRequestLinks is not nil, this is a pull request.
// The IsPullRequest helper method can be used to check that.
type Object struct {
	Links struct {
		Self   string `json:"self"`
		Parent string `json:"parent"`
	} `json:"links"`
	Type        string `json:"type"`
	Value       string `json:"value"`
	Overridable bool   `json:"overridable"`
	Description string `json:"description"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Metadata    struct {
		Timestamp int64 `json:"timestamp"`
		LastUser  struct {
			Name string `json:"name"`
		} `json:"lastUser"`
		Domain struct {
			Name string `json:"name"`
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"domain"`
		IPType     string `json:"ipType"`
		ParentType string `json:"parentType"`
	} `json:"metadata"`
}

// IssueRequest represents a request to create/edit an issue.
// It is separate from Issue above because otherwise Labels
// and Assignee fail to serialize to the correct JSON.
type IssueRequest struct {
	Title     *string   `json:"title,omitempty"`
	Body      *string   `json:"body,omitempty"`
	Labels    *[]string `json:"labels,omitempty"`
	Assignee  *string   `json:"assignee,omitempty"`
	State     *string   `json:"state,omitempty"`
	Milestone *int      `json:"milestone,omitempty"`
	Assignees *[]string `json:"assignees,omitempty"`
}

// IssueListOptions specifies the optional parameters to the IssuesService.List
// and IssuesService.ListByOrg methods.
type ObjectListOptions struct {
	// Filter specifies which issues to list. Possible values are: assigned,
	// created, mentioned, subscribed, all. Default is "assigned".
	Filter string `url:"limit,omitempty"`

	ListOptions
}

// List the Objects for the authenticated user. If all is true, list issues
func (s *ObjectService) List(ctx context.Context, all bool, opt *ObjectListOptions) ([]*Object, *Response, error) {
	var u string
	if all {
		u = "issues"
	} else {
		u = "user/issues"
	}
	return s.listObjects(ctx, u, opt)
}

func (s *ObjectService) listObjects(ctx context.Context, u string, opt *ObjectListOptions) ([]*Object, *Response, error) {
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var objects []*Object
	resp, err := s.client.Do(ctx, req, &objects)
	if err != nil {
		return nil, resp, err
	}

	return objects, resp, nil
}

type ObjectRequest struct {
	Type        string `json:"type"`
	Value       string `json:"value"`
	Overridable bool   `json:"overridable"`
	Description string `json:"description"`
	ID          string `json:"id"`
	Name        string `json:"name"`
}
