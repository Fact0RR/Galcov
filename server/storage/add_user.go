package storage

func (s *Store) Add_user(login, password string) error {
	var id int
	err := s.DB.QueryRow("select add_user($1,$2)",login,password).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}