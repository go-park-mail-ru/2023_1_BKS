package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_BKS/internal/post/domain"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type PostPostgressRepository struct {
	posts *sql.DB
}

func (t PostPostgressRepository) GetId(ctx context.Context,
	id uuid.UUID) (*domain.Post, int, error) {

	var result domain.Post

	err := t.posts.QueryRow(`SELECT userid, title, description,
	views, price, close, tags, images, time FROM posts WHERE id = $1
	LIMIT 1`, id).Scan(result.UserID, result.Title, result.Description,
		result.Views, result.Price, result.Status, result.Category,
		pq.Array(result.PathImages), result.Time)

	if err == sql.ErrNoRows {
		return nil, http.StatusNotFound, err
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	// Удалить от сюда
	*result.Views = *result.Views + 1
	_, err = t.posts.Exec("update posts set views = $1 where id = $2",
		*result.Views, id)
	//До сюда
	return &result, http.StatusOK, nil
}

func (t PostPostgressRepository) GetMiniPost(ctx context.Context,
	par domain.Parameters) ([]domain.Post, int, error) {

	if *par.Offset <= 0 {
		return nil, http.StatusBadRequest, errors.New("The page number cannot be less than zero")
	}

	var rows *sql.Rows
	var err error
	//SELECT * FROM EMPLOYEE WHERE name = 'Clark' and (dept IS NOT NULL) or name = 'Dave' and (dept IS  NULL);
	if *par.Sort == "New" || par.Sort == nil {
		rows, err = t.posts.Query(`SELECT id, userid, title, description,
		price,  images, time FROM posts WHERE tags = $1 and close = $2 and 
		ORDER BY time LIMIT $3 OFFSET $4`, par.Category, par.Offset, par.Limit,
			par.Status, par.UserId)
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var posts []domain.Post

	for rows.Next() {
		p := domain.Post{}
		err = rows.Scan(&p.Id, &p.UserID, &p.Title, &p.Description,
			&p.Price, pq.Array(&p.PathImages), &p.Time)
		if err != nil {
			// Сделать возврат ошибки
			fmt.Println(err)
			continue
		}
		posts = append(posts, p)
	}
	// Подумать нд именованием ошибки
	if err = rows.Err(); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	// Подумать нд именованием ошибки
	if err = rows.Close(); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return posts, http.StatusOK, nil
}

func (t PostPostgressRepository) GetFavorite(ctx context.Context,
	userId uuid.UUID) ([]uuid.UUID, int, error) {

	rows, err := t.posts.Query(`SELECT idpost FROM favorite
	WHERE userid = $1`, userId)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var postId []uuid.UUID

	for rows.Next() {
		p := uuid.UUID{}
		err = rows.Scan(pq.Array(&postId))
		if err != nil {
			fmt.Println(err)
			continue
		}
		q := uuid.UUID{}
		if p != q {
			postId = append(postId, p)
		}
	}

	// Подумать нд именованием ошибки
	if err = rows.Err(); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	// Подумать нд именованием ошибки
	if err = rows.Close(); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return postId, http.StatusOK, nil
}

func (t PostPostgressRepository) GetByCart(ctx context.Context,
	postId []uuid.UUID) ([]domain.Post, int, error) {
	var posts []domain.Post

	for _, idp := range postId {
		fmt.Println(idp)
		row := t.posts.QueryRow(`SELECT id, userid, title, description,
		price,  tags, images, time FROM posts WHERE close = false and id = $1`, idp)
		fmt.Println(row)
		p := domain.Post{}
		err := row.Scan(&p.Id, &p.UserID, &p.Title, &p.Description,
			&p.Price, &p.Status, pq.Array(&p.PathImages), &p.Time)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(p)
		posts = append(posts, p)
	}
	return posts, http.StatusOK, nil
}

func (t PostPostgressRepository) Search(ctx context.Context,
	search string) ([]uuid.UUID, int, error) {

	_, err := t.posts.Exec(`UPDATE posts SET fts = setweight(to_tsvector(title), 'A')
	|| setweight(to_tsvector(description), 'B')`)

	rows, err := t.posts.Query(`Select id FROM posts WHERE fts @@ to_tsquery($1)`, search)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var posts []uuid.UUID

	for rows.Next() {
		p := uuid.UUID{}
		err = rows.Scan(&p)
		if err != nil {
			fmt.Println(err)
			continue
		}
		posts = append(posts, p)
	}

	// Подумать нд именованием ошибки
	if err = rows.Err(); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	// Подумать нд именованием ошибки
	if err = rows.Close(); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return posts, http.StatusOK, nil
}

func (t *PostPostgressRepository) Create(ctx context.Context,
	post domain.Post) (int, error) {

	_, err := t.posts.Exec(`insert into posts (id, userid, title,
		description, price, close, tags, images, time, views)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		post.Id, post.UserID, post.Title, post.Description, post.Price,
		post.Status, post.Description, pq.Array(post.PathImages), post.Time, 1)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (t *PostPostgressRepository) Update(ctx context.Context,
	post domain.Post) (int, error) {

	_, err := t.posts.Exec(`update posts set  title = $1, description = $2,
	price = $3,  tags = $4, images = $5  where id = $6 and userid = $7`,
		post.Title, post.Description, post.Price,
		post.Description, pq.Array(post.PathImages), post.Id, post.UserID)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (t *PostPostgressRepository) Delete(ctx context.Context,
	id uuid.UUID) (int, error) {

	_, err := t.posts.Exec(`delete from posts where id = $1`, id)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (t *PostPostgressRepository) AddFavorite(ctx context.Context, userId uuid.UUID,
	postId uuid.UUID) (int, error) {

	pid, status, err := t.GetFavorite(ctx, userId)

	if status != http.StatusOK {
		return status, err
	}

	pid = append(pid, postId)

	_, err = t.posts.Exec("insert into favorite (idpost, userid) values ($1, $2)",
		pq.Array(pid), userId)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (t *PostPostgressRepository) RemoveFavorite(ctx context.Context, userId uuid.UUID,
	postId uuid.UUID) (int, error) {

	pid, status, err := t.GetFavorite(ctx, userId)

	if status != http.StatusOK {
		return status, err
	}

	var pidClear []uuid.UUID
	for _, fpid := range pid {
		if fpid != postId {
			pidClear = append(pidClear, fpid)
		}
	}

	_, err = t.posts.Exec("update favorite set idpost = $1 where UserId = $2",
		pq.Array(pidClear), userId)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
