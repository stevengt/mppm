package vst

type VSTComponentType int32

// https://www.reddit.com/r/edmproduction/comments/1peyan/could_someone_explain_to_me_the_difference/

const (
	Program VSTComponentType = iota
	Bank    VSTComponentType = iota
)
