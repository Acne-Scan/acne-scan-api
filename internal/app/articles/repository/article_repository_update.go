package repository

import (
	"fmt"
	"strings"
	"time"
)

func (articleRepository *ArticleRepositoryImpl) Update(name,description, image string, id string, updatedAt time.Time) error {
	  // Base query
    query := "UPDATE article SET "
    updates := []string{}
    values := []interface{}{}

    // Add fields to update if they are not nil
    if image != "" {
        updates = append(updates, "image = ?")
        values = append(values, image)
    }

    if name != "" {
        updates = append(updates, "name = ?")
        values = append(values, name)
    }

    if description != "" {
        updates = append(updates, "description = ?")
        values = append(values, description)
    }

    // updated_at should always be updated
    updates = append(updates, "updated_at = ?")
    values = append(values, updatedAt)

    // Add the WHERE clause
    query += strings.Join(updates, ", ") + " WHERE article_id = ?"
    values = append(values, id)

    // Execute the query
    _, err := articleRepository.DB.Exec(query, values...)
    if err != nil {
        return fmt.Errorf("failed to update article: %s", err.Error())
    }

    return nil

}
