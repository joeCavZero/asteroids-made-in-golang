package sliceSys

import "project/src/entities"

func RemoveShoot(s []entities.Shoot, i int) []entities.Shoot {
	return append(s[:i], s[i+1:]...)
}
