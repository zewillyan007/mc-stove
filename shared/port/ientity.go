package port

type IEntity interface {
	IsValid() error
	GetId() int64
	SetId(id int64)
}
