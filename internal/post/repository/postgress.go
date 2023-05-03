package repository

import (
	"context"
	"database/sql"
	"fmt"
	"post/domain"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type PostPostgressRepository struct {
	posts *sql.DB
}

func (t PostPostgressRepository) GetId(ctx context.Context, id uuid.UUID) (domain.Post, error) {
	var (
		userID      uuid.UUID
		title       string
		description string
		price       string
		close       bool
		views       int
		time        time.Time
		tag         string
		pathImages  []string
	)

	row := t.posts.QueryRow("SELECT userid, title, description, views, price, close, tags, images, time FROM posts WHERE id = $1 LIMIT 1", id)
	err := row.Scan(&userID, &title, &description, &views, &price, &close, &tag, pq.Array(&pathImages), &time)
	if err != nil {
		return domain.Post{}, err
	}
	_, err = t.posts.Exec("update posts set views = $1 where id = $2", views+1, id)

	return domain.Post{
		Close:      close,
		Id:         id,
		UserID:     userID,
		Title:      title,
		Views:      views + 1,
		Desciption: description,
		Price:      price,
		Tags:       tag,
		PathImages: pathImages,
		Time:       time,
	}, err
}

func (t PostPostgressRepository) GetByTag(ctx context.Context, tag string, number int) ([]domain.Post, error) {
	if number <= 0 {
		return []domain.Post{}, nil // Ошибка страницы
	}

	rows, err := t.posts.Query(`SELECT id, userid, title, description,
	price,  images, time FROM posts WHERE tags = $1 and close = false ORDER BY time LIMIT 10 OFFSET $2`, tag, (number-1)*10)
	if err != nil {
		return []domain.Post{}, err
	}

	var posts []domain.Post

	for rows.Next() {
		p := domain.Post{}
		err = rows.Scan(&p.Id, &p.UserID, &p.Title, &p.Desciption,
			&p.Price, pq.Array(&p.PathImages), &p.Time)
		if err != nil {
			fmt.Println(err)
			continue
		}
		posts = append(posts, p)
	}

	return posts, err
}

func (t PostPostgressRepository) GetByUserIdOpen(ctx context.Context, idUser uuid.UUID, number int) ([]domain.Post, error) {
	if number <= 0 {
		return []domain.Post{}, nil // Ошибка страницы
	}
	rows, err := t.posts.Query(`SELECT id, userid, title, description, price, 
	 tags, images, time FROM posts WHERE userid = $1 and close = false ORDER BY time LIMIT 10 OFFSET $2`, idUser, (number-1)*10)
	if err != nil {
		return []domain.Post{}, err
	}

	var posts []domain.Post

	for rows.Next() {
		p := domain.Post{}
		err := rows.Scan(&p.Id, &p.UserID, &p.Title, &p.Desciption,
			&p.Price, &p.Tags, pq.Array(&p.PathImages), &p.Time)
		if err != nil {
			fmt.Println(err)
			continue
		}
		posts = append(posts, p)
	}

	return posts, err
}

func (t PostPostgressRepository) GetByUserIdClose(ctx context.Context, idUser uuid.UUID, number int) ([]domain.Post, error) {
	if number <= 0 {
		return []domain.Post{}, nil // Ошибка страницы
	}
	rows, err := t.posts.Query(`SELECT id, userid, title, description, price, 
	 tags, images, time FROM posts WHERE userid = $1 and close = true ORDER BY time LIMIT 10 OFFSET $2`, idUser, (number-1)*10)
	if err != nil {
		return []domain.Post{}, err
	}

	var posts []domain.Post

	for rows.Next() {
		p := domain.Post{}
		err := rows.Scan(&p.Id, &p.UserID, &p.Title, &p.Desciption,
			&p.Price, &p.Tags, pq.Array(&p.PathImages), &p.Time)
		if err != nil {
			fmt.Println(err)
			continue
		}
		posts = append(posts, p)
	}

	return posts, err
}

func (t PostPostgressRepository) GetFavorite(ctx context.Context, userId uuid.UUID) ([]uuid.UUID, error) {
	rows, err := t.posts.Query(`SELECT idpost FROM favorite WHERE userid = $1`, userId)
	if err != nil {
		return []uuid.UUID{}, err
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

	if err != nil {
		return []uuid.UUID{}, err
	}

	return postId, err
}

func (t PostPostgressRepository) GetByArray(ctx context.Context, postId []uuid.UUID) ([]domain.Post, error) {
	var posts []domain.Post

	for _, idp := range postId {
		fmt.Println(idp)
		row := t.posts.QueryRow(`SELECT id, userid, title, description,
		price,  tags, images, time FROM posts WHERE close = false and id = $1`, idp)
		fmt.Println(row)
		p := domain.Post{}
		err := row.Scan(&p.Id, &p.UserID, &p.Title, &p.Desciption,
			&p.Price, &p.Tags, pq.Array(&p.PathImages), &p.Time)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(p)
		posts = append(posts, p)
	}
	return posts, nil
}

func (t PostPostgressRepository) GetSortNew(ctx context.Context, number int) ([]domain.Post, error) {
	if number <= 0 {
		return []domain.Post{}, nil // Ошибка страницы
	}
	rows, err := t.posts.Query(`SELECT id, userid, title, description,
	price,  tags, images, time FROM posts WHERE close = false ORDER BY time LIMIT 10 OFFSET $1`, (number-1)*10)
	if err != nil {
		return []domain.Post{}, err
	}

	var posts []domain.Post

	for rows.Next() {
		p := domain.Post{}
		err = rows.Scan(&p.Id, &p.UserID, &p.Title, &p.Desciption,
			&p.Price, &p.Tags, pq.Array(&p.PathImages), &p.Time)
		if err != nil {
			fmt.Println(err)
			continue
		}
		posts = append(posts, p)
	}

	return posts, err
}

func (t PostPostgressRepository) Search(ctx context.Context, search string) ([]domain.Post, error) {
	rows, err := t.posts.Query(`Select id FROM posts WHERE ts_rank(to_tsvector(description), plainto_tsquery($1)) ORDER BY ts_rank(to_tsvector(description), plainto_tsquery($1)`, search)
	if err != nil {
		return []domain.Post{}, err
	}

	var posts []domain.Post

	for rows.Next() {
		p := domain.Post{}
		err = rows.Scan(&p.Id, &p.UserID, &p.Title, &p.Desciption,
			&p.Price, &p.Tags, pq.Array(&p.PathImages), &p.Time)
		if err != nil {
			fmt.Println(err)
			continue
		}
		posts = append(posts, p)
	}

	return posts, err
}

func (t *PostPostgressRepository) Create(ctx context.Context, post domain.Post) error {
	_, err := t.posts.Exec("insert into posts (id, userid, title, description, price, close, tags, images, time, views) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		post.Id, post.UserID, post.Title, post.Desciption, post.Price, post.Close,
		post.Tags, pq.Array(post.PathImages), post.Time, 0)
	return err
}

func (t *PostPostgressRepository) Update(ctx context.Context, post domain.Post) error {
	_, err := t.posts.Exec("update posts set  userid = $1, title = $2, description = $3, price = $4,  tags = $5, images = $6  where id = $7",
		post.UserID, post.Title, post.Desciption, post.Price,
		post.Tags, pq.Array(post.PathImages), post.Id)
	return err
}

func (t *PostPostgressRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := t.posts.Exec("delete from posts where id = $1", id)
	return err
}

func (t *PostPostgressRepository) Close(ctx context.Context, id uuid.UUID) error {
	_, err := t.posts.Exec("update posts set close = true where id = $1", id)
	return err
}

func (t *PostPostgressRepository) AddFavorite(ctx context.Context, userId uuid.UUID, postId uuid.UUID) error {
	pid, _ := t.GetFavorite(ctx, userId)

	pid = append(pid, postId)
	fmt.Println(pid)
	_, err := t.posts.Exec("insert into favorite (idpost, userid) values ($1, $2)",
		pq.Array(pid), userId)
	return err
}

func (t *PostPostgressRepository) RemoveFavorite(ctx context.Context, userId uuid.UUID, postId uuid.UUID) error {
	pid, _ := t.GetFavorite(ctx, userId)
	var pidClear []uuid.UUID
	for _, fpid := range pid {
		if fpid != postId {
			pidClear = append(pidClear, fpid)
		}
	}

	_, err := t.posts.Exec("update favorite set idpost = $1 where UserId = $2", pq.Array(pidClear), userId)
	return err
}
