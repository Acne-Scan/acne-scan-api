package repository

func (articleRepository *ArticleRepositoryImpl) Delete(id string) error {
	_, err := articleRepository.DB.Exec("delete from article where article_id=?", id)
	if err != nil {
		return err
	}

	return nil
}
