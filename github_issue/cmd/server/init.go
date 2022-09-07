package server

func Run() {
	c := NewClient()

	c.CreateIssue()
}
