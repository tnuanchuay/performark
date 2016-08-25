package model

type (
	Form	struct{
		Action	string
		Method	string
		Input	[]Input
	}

	Input	struct{
		Name	string
	}
)
