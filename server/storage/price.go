package storage

type Val struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (s *Store) Get_price() ([]Val, error) {
	var valArr []Val
	rows, err := s.DB.Query("select id, name, price from val")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		val := Val{}
		err := rows.Scan(&val.Id, &val.Name, &val.Price)
		
		if err != nil {
			//fmt.Println(err)
			continue
		}
		valArr = append(valArr, val)
	}
	
	return valArr, nil
}