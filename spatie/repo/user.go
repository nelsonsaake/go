package repo

type User interface {
	GetID() string
}

type Scope struct {
	User
	Roles       []string
	Permissions []string
}

func (s *Spatie) Scope(user User) (*Scope, error) {

	// Fetch roles and permissions in one query using UNION ALL
	type resultRow struct {
		Type  string
		Value string
	}
	var rows []resultRow

	query := `
		SELECT 'role' AS type, r.name AS value
		FROM user_roles ur
		JOIN roles r ON ur.role_id = r.id
		WHERE ur.user_id = ?
		UNION ALL
		SELECT 'permission' AS type, p.name AS value
		FROM user_permissions up
		JOIN permissions p ON up.permission_id = p.id
		WHERE up.user_id = ?
	`
	if err := s.db.Raw(query, user.GetID(), user.GetID()).Scan(&rows).Error; err != nil {
		return nil, err
	}

	var (
		roles       = make([]string, 0)
		permissions = make([]string, 0)
	)
	for _, row := range rows {
		switch row.Type {
		case "role":
			roles = append(roles, row.Value)
		case "permission":
			permissions = append(permissions, row.Value)
		}
	}

	return &Scope{
		User:        user,
		Roles:       roles,
		Permissions: permissions,
	}, nil
}
