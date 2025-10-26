package repo

func (r *Repo[T]) GetRandomID() (string, error) {

	var id string

	err := r.DB.Select("id").Order("RAND()").Limit(1).Scan(&id).Error

	return id, err
}

func GetRandomID[T any]() (string, error) {

	return Do[T]().GetRandomID()
}
