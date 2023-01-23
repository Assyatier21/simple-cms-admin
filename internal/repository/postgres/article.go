package postgres

import (
	m "cms-admin/models"
	"context"
)

func (r *repository) GetArticles(ctx context.Context) ([]m.ResArticle, error) {
	return []m.ResArticle{}, nil
}
func (r *repository) GetArticleDetails(ctx context.Context, id int) (m.ResArticle, error) {
	return m.ResArticle{}, nil
}
func (r *repository) InsertArticle(ctx context.Context, article m.Article) (m.ResArticle, error) {
	return m.ResArticle{}, nil
}
func (r *repository) UpdateArticle(ctx context.Context, article m.Article) (m.ResArticle, error) {
	return m.ResArticle{}, nil
}
func (r *repository) DeleteArticle(ctx context.Context, id int) error {
	return nil
}
