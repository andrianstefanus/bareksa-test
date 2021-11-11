package model

type Response struct {
	Status string
	Msg    string
}

type Err struct {
	Status int
	Msg    string
}
