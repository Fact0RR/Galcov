package storage

type PublicUser struct{
	Id int `json:"id"`
	Login    string `json:"login"`
}

func (s *Store) Get_people() ([]PublicUser, error) {
	var accArr []PublicUser
	rows, err := s.DB.Query("select id,login from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		acc := PublicUser{}
		err := rows.Scan(&acc.Id, &acc.Login)
		
		if err != nil {
			//fmt.Println(err)
			continue
		}
		accArr = append(accArr, acc)
	}
	
	return accArr, nil
}