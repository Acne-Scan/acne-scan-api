package service

import "fmt"

func (articleService *ArticleServiceImpl) Delete(id string) error {

	ifExist, err := articleService.ArticleRepository.GetById(id)
	if err != nil {
		return fmt.Errorf("failed to check is article already exist:%s", err.Error())
	}

	if ifExist == nil {
		return fmt.Errorf("article not found")
	}

	err = articleService.ArticleRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
