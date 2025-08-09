package db

func (r *Repository) GetUserRoleByID(roleId int64) (string, error) {
	var role string
	err := r.db.QueryRowContext(context.Background(), "SELECT role FROM user_roles WHERE id = $1", roleId).Scan(&role)
	if err != nil {
		logger.Error("Failed to get user role by ID: ", err)
		return "", err
	}
	return role, nil
}
