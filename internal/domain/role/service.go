package role

type RoleService struct {
	repository RoleRepository
}

func NewRoleService(repository RoleRepository) *RoleService {
	service := &RoleService{
		repository: repository,
	}
	return service
}
