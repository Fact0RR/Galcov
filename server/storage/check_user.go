package storage

func (s *Store) Check_user(login, password string) (int, error) {

	
	var id int
	err := s.DB.QueryRow("select check_user($1,$2)",login,password).Scan(&id)
	if err != nil {
		return 0, err
	}
	
	return id, nil
}