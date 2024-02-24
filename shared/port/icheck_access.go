package port

type ICheckAccessService interface {
	CheckAccess(UserData interface{}, endPoint, method string) (bool, error)
}
