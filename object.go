package firepower

import (
	"context"
)

// ObjectService handles communication with the object related methods of the
type ObjectService service

// Object represents a FirePower object
type Object struct {
	Links  Links  `json:"links"`
	Items  []Item `json:"items"`
	Paging Paging `json:"paging"`
}

type Links struct {
	Self string `json:"self"`
}

type Paging struct {
	Count  int `json:"count"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Pages  int `json:"pages"`
}

type Item struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	Links       struct {
		Parent string `json:"parent"`
		Self   string `json:"self"`
	} `json:"links"`
	Metadata struct {
		Domain struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"domain"`
		IPType   string `json:"ipType"`
		LastUser struct {
			Name string `json:"name"`
		} `json:"lastUser"`
		ParentType string `json:"parentType"`
		ReadOnly   struct {
			Reason string `json:"reason"`
			State  bool   `json:"state"`
		} `json:"readOnly"`
		Timestamp int `json:"timestamp"`
	} `json:"metadata"`
	Name        string `json:"name"`
	Overridable bool   `json:"overridable"`
	Type        string `json:"type"`
	Value       string `json:"value"`
}

// ObjectListOptions specifies the optional parameters to the IssuesService.List
// and IssuesService.ListByOrg methods.
type ObjectListOptions struct {
	// Expanded specifies which issues to list. Possible values are: assigned,
	// created, mentioned, subscribed, all. Default is "assigned".
	Expanded bool `url:"expanded,true"`
}

func (s *ObjectService) listObjects(ctx context.Context, u string, opt *ObjectListOptions) (*Object, *Response, error) {
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var objects *Object
	resp, err := s.client.Do(ctx, req, &objects)
	if err != nil {
		return nil, resp, err
	}

	return objects, resp, nil
}

func (s *ObjectService) getObject(ctx context.Context, u string, opt *ObjectListOptions) (*Item, *Response, error) {
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var item *Item
	resp, err := s.client.Do(ctx, req, &item)
	if err != nil {
		return nil, resp, err
	}

	return item, resp, nil
}

//InsertObject Insert
func (s *ObjectService) InsertObject(ctx context.Context, u string, i Item, opt *ObjectListOptions) (*Response, error) {
	req, err := s.client.NewRequest("POST", u, i)
	if err != nil {
		return nil, err
	}

	var item *Item
	resp, err := s.client.Do(ctx, req, &item)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ListHosts the Objects for the authenticated user.
func (s *ObjectService) ListHosts(ctx context.Context, opt *ObjectListOptions) (*Object, *Response, error) {
	return s.listObjects(ctx, "object/hosts", opt)
}

// GetHost the Objects for the authenticated user.
func (s *ObjectService) GetHost(ctx context.Context, hostid string, opt *ObjectListOptions) (*Item, *Response, error) {
	return s.getObject(ctx, "object/hosts/"+hostid, opt)
}

// InsertHost the Objects for the authenticated user.
func (s *ObjectService) InsertHost(ctx context.Context, i Item) (*Response, error) {
	return s.InsertObject(ctx, "object/hosts/", i, nil)
}
