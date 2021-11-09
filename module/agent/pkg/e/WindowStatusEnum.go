package e
type WindowStatus int
const (
	INIT WindowStatus =iota
	OPENING
	RUNNING
	CLOSE
	ABNORMAL
)
func ( w WindowStatus)String()string{
	return [...]string{"Init","Running","Opening","Close","Abnormal"}[w]
}
