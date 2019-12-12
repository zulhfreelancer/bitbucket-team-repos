package main

import "time"

type apiTokenResponse struct {
	AccessToken  string `json:"access_token"`
	Scopes       string `json:"scopes"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

type repositoriesResponse struct {
	Next   string `json:"next"`
	Values []struct {
		Scm     string `json:"scm"`
		Website string `json:"website"`
		HasWiki bool   `json:"has_wiki"`
		UUID    string `json:"uuid"`
		Links   struct {
			Watchers struct {
				Href string `json:"href"`
			} `json:"watchers"`
			Branches struct {
				Href string `json:"href"`
			} `json:"branches"`
			Tags struct {
				Href string `json:"href"`
			} `json:"tags"`
			Commits struct {
				Href string `json:"href"`
			} `json:"commits"`
			Clone []struct {
				Href string `json:"href"`
				Name string `json:"name"`
			} `json:"clone"`
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Source struct {
				Href string `json:"href"`
			} `json:"source"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
			Hooks struct {
				Href string `json:"href"`
			} `json:"hooks"`
			Forks struct {
				Href string `json:"href"`
			} `json:"forks"`
			Downloads struct {
				Href string `json:"href"`
			} `json:"downloads"`
			Pullrequests struct {
				Href string `json:"href"`
			} `json:"pullrequests"`
		} `json:"links"`
		ForkPolicy string `json:"fork_policy"`
		Name       string `json:"name"`
		Project    struct {
			Key   string `json:"key"`
			Type  string `json:"type"`
			UUID  string `json:"uuid"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"links"`
			Name string `json:"name"`
		} `json:"project"`
		Language   string    `json:"language"`
		CreatedOn  time.Time `json:"created_on"`
		Mainbranch struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"mainbranch"`
		FullName  string `json:"full_name"`
		HasIssues bool   `json:"has_issues"`
		Owner     struct {
			Username    string `json:"username"`
			DisplayName string `json:"display_name"`
			Type        string `json:"type"`
			UUID        string `json:"uuid"`
			Links       struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"links"`
		} `json:"owner"`
		UpdatedOn   time.Time `json:"updated_on"`
		Size        int       `json:"size"`
		Type        string    `json:"type"`
		Slug        string    `json:"slug"`
		IsPrivate   bool      `json:"is_private"`
		Description string    `json:"description"`
	} `json:"values"`
	Pagelen  int    `json:"pagelen"`
	Size     int    `json:"size"`
	Page     int    `json:"page"`
	Previous string `json:"previous"`
}
