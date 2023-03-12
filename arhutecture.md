# Entity

```
type User struct {
	ID				int
	FirstName		string
	SecondName		string
	Email			string
	Password		string \\  В зашифрованном виде
	CardNumber		string \\  В зашифрованном виде
	CardCvv			string \\  В зашифрованном виде
	Image			image.RGBA
}
```
Часть с номером карт не уверен нужна ли учитывая, что у нас не маркет плейс, а сайт с объявлениями, и возможно стоит добавить адрес не уверен.


```
type Post struct {
    ID          int
    Title		string
	Description	string
	Price		string
	Date		data.Time
	Tags		[]string
    Image		image.RGBA
}
```
Тут содержатся структуры используемые во всех, частях системы.

Если будем использовать лайки/дизлайки, то необходимо использовать следующую структуры:

```
type Like struct {
    ID          int
    User_       *User
    Post_       *Post
}

type Dislike struct {
    ID          int
    User_       *User
    Post_       *Post
}
```

# Repository

## Базы данных
```
type UserRepo struct {
    User entity.User
}

func (user *UserRepo) Get() entity.User {
    return user.User
}

func (user *UserRepo) Set(db DataBase) error

type PostRepo struct {
    Post entity.Post
}

func (post *PostRepo) Get() entity.Post

func (post *PostRepo) Set(db DataBase) error

type ValueRepo[T any] struct {
    value T
}

func (value *ValueRepo) Get[T any]() T

func (value *ValueRepo) Set[T any](value T) error


type Entity[T any] interface {
    Get() T
    Set(DataBase, int) error
}

type DataBase[T any] interface {
    Get() T
    Set(Entity) error
}

func GetDataBase[T any] (entity Entity[T], db DataBase[T], id int) error {
	// Обработка ошибок
    entity.Set(db, id)
}

func SetDataBase[T any] (entity Entity[T], db DataBase[T]) error {
	// Обработка ошибок
    db.Set(entity)
}

```
В синтаксисе не до конца уверен.

```
type MyConstraint interface {
 int | int8 | int16 | int32 | int64
}
```
Чисто теоритически для большей безопасности можно использовать вместо any, что то на подобии.


### Основной принцип работы

Через два основных интерфейса реализуется основная логика работы с базой данных. В процессе написания понял, что осмысленность данного решения будет только при опредлении типа данных во время компиляции(что то на подобии такого GetDataBase[value.TypeOf()]).
В данный момент не до конца уверен, что это возможно(исходя из шаблонов плюсов).

Но в целом это не отменяет того факта, что архитектура будет полезна. Без интерфейсов все будет решаться методами set и get в рамках структуры. Вся реализация рассмотрена при условии идентичности названия полей и структуры с названиями столбцов и таблицы в БД. В таком случае с помощью пакета Reflection. Во первых он позволяет итерировать по структуре, а во вторых получать названия полей и соответсвенно самой структуры.
Пример работы: При set в БД в первую очередь ищется таблица с названием структуры.
Далее идёт валидация на не заполненность данных в строке с данным id. Далее идёт итерирование по structure и запись в БД.
Get подобным образом.
