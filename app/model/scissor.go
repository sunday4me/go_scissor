package model


func GetAllScissors() ([]Scissor, error) {
	var scissors []Scissor

	tx := db.Find(&scissors)
	if tx.Error != nil {
		return []Scissor{}, tx.Error
	}

	return scissors, nil
}

func GetScissor(id uint64) (Scissor, error) {
	var scissor Scissor

	tx := db.Where("id = ?", id).First(&scissor)

	if tx.Error != nil {
		return Scissor{}, tx.Error
	}

	return scissor, nil
}

func CreateScissor(scissor Scissor) error {
	tx := db.Create(&scissor)
	return tx.Error
}

func UpdateScissor(scissor Scissor) error {

	tx := db.Save(&scissor)
	return tx.Error
}

func DeleteScissor(id uint64) error {

	tx := db.Unscoped().Delete(&Scissor{}, id)
	return tx.Error
}

func FindByScissorUrl(url string) (Scissor, error) {
	var scissor Scissor
	tx := db.Where("scissor = ?", url).First(&scissor)
	return scissor, tx.Error
}
