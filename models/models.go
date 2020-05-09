package models

// Model Model
type Model interface {
	GetID() interface{}
	GetCreateAt() interface{}
	GetUpdateAt() interface{}
}
