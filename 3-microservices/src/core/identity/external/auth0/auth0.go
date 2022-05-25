package auth0

import (
	"identity/utils"

	"gopkg.in/auth0.v5/management"
)

func getManagement() (*management.Management, error) {
	domain := utils.GetWithDefault("AUTH0_DOMAIN", "")
	clientId := utils.GetWithDefault("AUTH0_MANAGEMENT_CLIENT_ID", "")
	clientSecret := utils.GetWithDefault("AUTH0_MANAGEMENT_CLIENT_SECRET", "")

	m, err := management.New(domain, management.WithClientCredentials(clientId, clientSecret))

	if err != nil {
		return nil, err
	}

	return m, nil
}

func AssignRolesToUser(authOID string, role *management.Role) error {
	m, err := getManagement()

	if err != nil {
		return err
	}

	err = m.User.AssignRoles(authOID, []*management.Role{role})

	if err != nil {
		return err
	}

	return nil
}
