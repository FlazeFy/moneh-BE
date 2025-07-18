package admin

import "moneh/models"

// Admin Interface
type AdminService interface {
	GetAllAdminContact() ([]models.UserContact, error)
}

// Admin Struct
type adminService struct {
	adminRepo AdminRepository
}

// Admin Constructor
func NewAdminService(adminRepo AdminRepository) AdminService {
	return &adminService{
		adminRepo: adminRepo,
	}
}

func (s *adminService) GetAllAdminContact() ([]models.UserContact, error) {
	// Repo : Get All Admin Audit
	admin, err := s.adminRepo.FindAllAdminContact()
	if err != nil {
		return nil, err
	}

	return admin, nil
}
