package storage

type Acc struct{
	Id_val int `json:"id_val"`
	Name string `json:"name"`
	Price float64 `json:"count"`

}

func (s *Store) Get_accaount(id int) ([]Acc, error) {
	var accArr []Acc
	rows, err := s.DB.Query("select val_id,name,count from acc_val av join val as v on v.id = av.val_id where login_id = $1",id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		acc := Acc{}
		err := rows.Scan(&acc.Id_val, &acc.Name, &acc.Price)
		
		if err != nil {
			//fmt.Println(err)
			continue
		}
		accArr = append(accArr, acc)
	}
	
	return accArr, nil
}