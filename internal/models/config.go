package models


type Config struct{
	Name string `json:"name" yaml:"name"`
	Role string `json:"role" yaml:"role"` 
	Stack string  `json:"stack" yaml:"stack"`
}