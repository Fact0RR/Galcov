package storage

func (s *Store) TransferToSelf(id_user int, id_val_old int, id_val_new int, count float64)  error {
	_, err := s.DB.Exec("call t2s($1,$2,$3,$4)",id_user,id_val_old,id_val_new,count)
	if err != nil {
		return err
	}
	
	return nil
}

func (s *Store) TransferToUser(id_old int, id_new int, id_val int, count float64)  error {
	_, err := s.DB.Exec("call t2u($1,$2,$3,$4)",id_old,id_new,id_val,count)
	if err != nil {
		return err
	}
	
	return nil
}