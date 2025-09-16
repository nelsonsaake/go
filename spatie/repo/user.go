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
		GROUP BY r.name
		UNION ALL
		SELECT 'permission' AS type, p.name AS value
		FROM user_permissions up
		JOIN permissions p ON up.permission_id = p.id
		WHERE up.user_id = ?
		GROUP BY p.name
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

// ScopeMany returns scopes for multiple users
func (s *Spatie) ScopeMany(users ...User) (map[string]*Scope, error) {

	scopes := make(map[string]*Scope)
	if len(users) == 0 {
		return scopes, nil
	}

	// Build userID slice
	userIDs := make([]string, len(users))
	for i, u := range users {
		userIDs[i] = u.GetID()
	}

	// Query all roles and permissions for all users
	type resultRow struct {
		UserID string
		Type   string
		Value  string
	}
	var rows []resultRow

	query := `
		SELECT ur.user_id AS user_id, 'role' AS type, r.name AS value
		FROM user_roles ur
		JOIN roles r ON ur.role_id = r.id
		WHERE ur.user_id IN ?
		GROUP BY ur.user_id, r.name
		UNION ALL
		SELECT up.user_id AS user_id, 'permission' AS type, p.name AS value
		FROM user_permissions up
		JOIN permissions p ON up.permission_id = p.id
		WHERE up.user_id IN ?
		GROUP BY up.user_id, p.name
	`
	if err := s.db.Raw(query, userIDs, userIDs).Scan(&rows).Error; err != nil {
		return nil, err
	}

	// Build scopes map
	for _, u := range users {
		scopes[u.GetID()] = &Scope{
			User:        u,
			Roles:       []string{},
			Permissions: []string{},
		}
	}
	for _, row := range rows {
		scope := scopes[row.UserID]
		if scope == nil {
			continue
		}
		switch row.Type {
		case "role":
			scope.Roles = append(scope.Roles, row.Value)
		case "permission":
			scope.Permissions = append(scope.Permissions, row.Value)
		}
	}
	return scopes, nil
}
