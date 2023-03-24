package jira

type Fields struct {
	Issuetype struct {
		Self        string `json:"self"`
		ID          string `json:"id"`
		Description string `json:"description"`
		IconURL     string `json:"iconUrl"`
		Name        string `json:"name"`
		Subtask     bool   `json:"subtask"`
		AvatarID    int    `json:"avatarId"`
	} `json:"issuetype"`
	Timespent any `json:"timespent"`
	Project   struct {
		Self           string `json:"self"`
		ID             string `json:"id"`
		Key            string `json:"key"`
		Name           string `json:"name"`
		ProjectTypeKey string `json:"projectTypeKey"`
		AvatarUrls     struct {
			Four8X48  string `json:"48x48"`
			Two4X24   string `json:"24x24"`
			One6X16   string `json:"16x16"`
			Three2X32 string `json:"32x32"`
		} `json:"avatarUrls"`
	} `json:"project"`
	FixVersions        []any  `json:"fixVersions"`
	Customfield10110   any    `json:"customfield_10110"`
	Customfield10111   any    `json:"customfield_10111"`
	Aggregatetimespent any    `json:"aggregatetimespent"`
	Resolution         any    `json:"resolution"`
	Customfield10105   any    `json:"customfield_10105"`
	Customfield10106   string `json:"customfield_10106"`
	Customfield10108   any    `json:"customfield_10108"`
	Customfield10109   any    `json:"customfield_10109"`
	Resolutiondate     any    `json:"resolutiondate"`
	Workratio          int    `json:"workratio"`
	LastViewed         string `json:"lastViewed"`
	Watches            struct {
		Self       string `json:"self"`
		WatchCount int    `json:"watchCount"`
		IsWatching bool   `json:"isWatching"`
	} `json:"watches"`
	Created  string `json:"created"`
	Priority struct {
		Self    string `json:"self"`
		IconURL string `json:"iconUrl"`
		Name    string `json:"name"`
		ID      string `json:"id"`
	} `json:"priority"`
	Customfield10100              any      `json:"customfield_10100"`
	Customfield10101              any      `json:"customfield_10101"`
	Labels                        []string `json:"labels"`
	Timeestimate                  any      `json:"timeestimate"`
	Aggregatetimeoriginalestimate any      `json:"aggregatetimeoriginalestimate"`
	Versions                      []any    `json:"versions"`
	Issuelinks                    []any    `json:"issuelinks"`
	Assignee                      struct {
		Self         string `json:"self"`
		Name         string `json:"name"`
		Key          string `json:"key"`
		EmailAddress string `json:"emailAddress"`
		AvatarUrls   struct {
			Four8X48  string `json:"48x48"`
			Two4X24   string `json:"24x24"`
			One6X16   string `json:"16x16"`
			Three2X32 string `json:"32x32"`
		} `json:"avatarUrls"`
		DisplayName string `json:"displayName"`
		Active      bool   `json:"active"`
		TimeZone    string `json:"timeZone"`
	} `json:"assignee"`
	Updated string `json:"updated"`
	Status  struct {
		Self           string `json:"self"`
		Description    string `json:"description"`
		IconURL        string `json:"iconUrl"`
		Name           string `json:"name"`
		ID             string `json:"id"`
		StatusCategory struct {
			Self      string `json:"self"`
			ID        int    `json:"id"`
			Key       string `json:"key"`
			ColorName string `json:"colorName"`
			Name      string `json:"name"`
		} `json:"statusCategory"`
	} `json:"status"`
	Components           []any  `json:"components"`
	Timeoriginalestimate any    `json:"timeoriginalestimate"`
	Description          string `json:"description"`
	Timetracking         struct {
	} `json:"timetracking"`
	Archiveddate          any    `json:"archiveddate"`
	Attachment            []any  `json:"attachment"`
	Aggregatetimeestimate any    `json:"aggregatetimeestimate"`
	Summary               string `json:"summary"`
	Creator               struct {
		Self         string `json:"self"`
		Name         string `json:"name"`
		Key          string `json:"key"`
		EmailAddress string `json:"emailAddress"`
		AvatarUrls   struct {
			Four8X48  string `json:"48x48"`
			Two4X24   string `json:"24x24"`
			One6X16   string `json:"16x16"`
			Three2X32 string `json:"32x32"`
		} `json:"avatarUrls"`
		DisplayName string `json:"displayName"`
		Active      bool   `json:"active"`
		TimeZone    string `json:"timeZone"`
	} `json:"creator"`
	Subtasks []any `json:"subtasks"`
	Reporter struct {
		Self         string `json:"self"`
		Name         string `json:"name"`
		Key          string `json:"key"`
		EmailAddress string `json:"emailAddress"`
		AvatarUrls   struct {
			Four8X48  string `json:"48x48"`
			Two4X24   string `json:"24x24"`
			One6X16   string `json:"16x16"`
			Three2X32 string `json:"32x32"`
		} `json:"avatarUrls"`
		DisplayName string `json:"displayName"`
		Active      bool   `json:"active"`
		TimeZone    string `json:"timeZone"`
	} `json:"reporter"`
	Customfield10000  string `json:"customfield_10000"`
	Aggregateprogress struct {
		Progress int `json:"progress"`
		Total    int `json:"total"`
	} `json:"aggregateprogress"`
	Environment any    `json:"environment"`
	Duedate     string `json:"duedate"`
	Progress    struct {
		Progress int `json:"progress"`
		Total    int `json:"total"`
	} `json:"progress"`
	Comment struct {
		Comments   []any `json:"comments"`
		MaxResults int   `json:"maxResults"`
		Total      int   `json:"total"`
		StartAt    int   `json:"startAt"`
	} `json:"comment"`
	Votes struct {
		Self     string `json:"self"`
		Votes    int    `json:"votes"`
		HasVoted bool   `json:"hasVoted"`
	} `json:"votes"`
	Worklog struct {
		StartAt    int   `json:"startAt"`
		MaxResults int   `json:"maxResults"`
		Total      int   `json:"total"`
		Worklogs   []any `json:"worklogs"`
	} `json:"worklog"`
	Archivedby any `json:"archivedby"`
}

type Issue struct {
	Expand string `json:"expand"`
	ID     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`
	Fields Fields `json:"fields"`
}

type ProjectElement struct {
	Expand     string `json:"expand"`
	Self       string `json:"self"`
	ID         string `json:"id"`
	Key        string `json:"key"`
	Name       string `json:"name"`
	AvatarUrls struct {
		Four8X48  string `json:"48x48"`
		Two4X24   string `json:"24x24"`
		One6X16   string `json:"16x16"`
		Three2X32 string `json:"32x32"`
	} `json:"avatarUrls"`
	ProjectTypeKey string `json:"projectTypeKey"`
	Archived       bool   `json:"archived"`
}

type AvatarUrls struct {
	Four8X48  string `json:"48x48"`
	Two4X24   string `json:"24x24"`
	One6X16   string `json:"16x16"`
	Three2X32 string `json:"32x32"`
}

type Project struct {
	Expand      string `json:"expand"`
	Self        string `json:"self"`
	ID          string `json:"id"`
	Key         string `json:"key"`
	Description string `json:"description"`
	Lead        struct {
		Self       string `json:"self"`
		Key        string `json:"key"`
		Name       string `json:"name"`
		AvatarUrls struct {
			Four8X48  string `json:"48x48"`
			Two4X24   string `json:"24x24"`
			One6X16   string `json:"16x16"`
			Three2X32 string `json:"32x32"`
		} `json:"avatarUrls"`
		DisplayName string `json:"displayName"`
		Active      bool   `json:"active"`
	} `json:"lead"`
	Components []any `json:"components"`
	IssueTypes []struct {
		Self        string `json:"self"`
		ID          string `json:"id"`
		Description string `json:"description"`
		IconURL     string `json:"iconUrl"`
		Name        string `json:"name"`
		Subtask     bool   `json:"subtask"`
		AvatarID    int    `json:"avatarId,omitempty"`
	} `json:"issueTypes"`
	AssigneeType string `json:"assigneeType"`
	Versions     []any  `json:"versions"`
	Name         string `json:"name"`
	Roles        struct {
		Administrators string `json:"Administrators"`
	} `json:"roles"`
	AvatarUrls     AvatarUrls `json:"avatarUrls"`
	ProjectTypeKey string     `json:"projectTypeKey"`
	Archived       bool       `json:"archived"`
}

type Workflow struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	Steps            int    `json:"steps"`
	Default          bool   `json:"default"`
	LastModifiedDate string `json:"lastModifiedDate,omitempty"`
	LastModifiedUser string `json:"lastModifiedUser,omitempty"`
}
