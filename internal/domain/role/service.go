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

func (service *RoleService) InsertSampleData() {

	tableExist := service.repository.db.Migrator().HasTable(&Role{})

	if !tableExist {
		service.repository.MigrateTable()

		roleNames := []string{"USER", "ADMIN"}

		for _, roleName := range roleNames {
			role := NewRole(roleName)
			service.repository.Create(role)
		}
	}

}
