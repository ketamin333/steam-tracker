package game_repository

func (r *GameRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM games WHERE id = $1`, id)

	return err
}
