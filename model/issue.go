package model

type Issue struct {
	Id   int
	Name string
}

var issues = []Issue{
	{Id:1, Name:"I need a help"},
	{Id:2, Name:"I need another help"},
}

type e string

func (s e) Error() string {
	return string(s)
}

func FindIssue(id int) (Issue, error) {
	for _, issue := range(issues) {
		if issue.Id == id {
			return issue,nil
		}
	}
	var e e
	e = "error"
	return Issue{},e
}

func GetIssues() []Issue {
	return issues;
}
