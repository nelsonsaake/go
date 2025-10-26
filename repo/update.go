package repo

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func isZero(v any) bool {
	return reflect.Zero(reflect.TypeOf(v)) == v
}

func updateOldRecord[T any](old *T, updates any) (*T, error) {

	var result = old

	die := func(f string, a ...any) (*T, error) {
		return nil, fmt.Errorf(f, a...)
	}

	oldMap, err := toMap(old)
	if err != nil {
		return die("error casting old record to map: %w", err)
	}

	updatesMap, err := toMap(updates)
	if err != nil {
		return die("error casting updates to map record: %w", err)
	}

	for key, value := range updatesMap {
		if isZero(value) {
			continue
		}
		oldMap[key] = value
	}

	bs, err := json.Marshal(oldMap)
	if err != nil {
		return die("error marshalling: %w", err)
	}

	err = json.Unmarshal(bs, result)
	if err != nil {
		return die("error unmarshalling %q: %w", getRscName[T](), err)
	}

	return result, nil
}

func (r *Repo[T]) Update(id string, data any) (*T, error) {

	var (
		old = new(T)
	)

	die := func(f string, a ...any) (*T, error) {
		return nil, fmt.Errorf(f, a...)
	}

	rscName := getRscName[T]()

	err := r.DB.Where("id = ?", id).First(old).Error
	if err != nil {
		return die("error finding %q with id %s: %v", rscName, id, err)
	}

	result, err := updateOldRecord(old, data)
	if err != nil {
		return die("error casting %q", rscName)
	}

	err = r.DB.Save(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func Update[T any](id string, data any) (*T, error) {
	return Do[T]().Update(id, data)
}
