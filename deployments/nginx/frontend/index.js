const express = require('express');
const body = require('body-parser');
const cookie = require('cookie-parser');
const morgan = require('morgan');
const uuid = require('uuid').v4;
const path = require('path');
const fs = require('fs');

const app = express();

app.use(morgan('dev'));
app.use(express.static(path.resolve(__dirname, 'static')));
app.use(body.json());
app.use(cookie());

const users = require('./static/jsonData/users.json');
const anns = require('./static/jsonData/anns.json');

/** session identificators */
const ids = {};

app.post('/api/user', (req, res) => {
    const { password, login, username } = req.body;
    if (!username || login.length < 4) {
        return res.status(400).json({ error: 'Логин не менее 4 символов', errorFill: 'username' });
    }
    if (!password || !password.match(/^\S{4,}$/)) {
        return res.status(400).json({ error: 'Минимальная длина пароля 4 символа', errorFill: 'password' });
    }
    if (!email) {
        return res.status(400).json({ error: 'Невалидные данные пользователя', errorFill: 'email' });
    }
    if (users[email]) {
        return res.status(400).json({ error: 'Пользователь уже существует', errorFill: 'username' });
    }

    const id = uuid();
    const user = {
        ...req.body, purchs: [], anns: [], avatar: '/default.jpg',
    };

    ids[id] = email;
    users.push(user);

    res.cookie('appuniq', id, { expires: new Date(Date.now() + 1000 * 60 * 10) });
    return res.status(201).json({ id });
});

app.post('/api/login', (req, res) => {
    const { password, email } = req.body;
    if (!password || !email) {
        return res.status(400).json({ error: 'Не указан E-Mail или пароль' });
    }

    const user = users.find((user) => user.email === email);
    if (!user || user.password !== password) {
        return res.status(400).json({ error: 'Неверный E-Mail и/или пароль' });
    }

    const id = uuid();
    ids[id] = email;

    res.cookie('appuniq', id, { expires: new Date(Date.now() + 1000 * 60 * 10) });
    return res.status(200).json({ id });
});

app.post('/api/logout', (req, res) => {
    const id = req.cookies.appuniq;
    const emailSession = ids[id];
    const user = users.find((user) => user.email === emailSession);

    if (!emailSession || !user) {
        return res.status(401).end();
    }

    res.clearCookie('appuniq');
    return res.end();
});

app.get('/api/board', (req, res) => {
    const id = req.cookies.appuniq;
    const emailSession = ids[id];

    const user = users.find((user) => user.email === emailSession);

    if (!emailSession || !user) {
        return res.json(anns);
    } else {
        const result = anns.filter((_, i) => user.anns.includes(i));
        return res.json(result);
    }
});

app.get('/api/board/:id', (req, res) => {
    const annId = req.params.id;
    res.json(anns[annId]);
});

app.get('/api/sellers/:id', (req, res) => {
    const sellerId = req.params.id;
    res.json(users[sellerId]);
});

app.get('/seller', (req, res) => {
    const id = req.cookies.appuniq;
    const emailSession = ids[id];
    const user = users.find((user) => user.email === emailSession);

    if (!emailSession || !user) {
        return res.status(401).json({error: 'Пользователь не найден'});
    }

    return res.json({
        id: user.id,
        login: user.login,
        email: user.email,
        avatar: user.avatar,
    });
});

// get seller of announcement
app.get('/api/getseller/:id', (req, res) => {
    let sellerId = -1;
    for (let User in users) {
        if (users[User].anns.find(elem => elem == req.params.id)) {
            sellerId = users[User].id;
            break;
        };
    }
    // console.log("getseller output: ", sellerId);
    res.json({id: users[sellerId].id});
});

app.get('/api/bucket', (req, res) => {
    const id = req.cookies.appuniq;
    const emailSession = ids[id];
    const user = users.find((user) => user.email === emailSession);

    if (!emailSession || !user) {
        return res.status(401).json({error: 'Пользователь не найден'});
    } else {
        const result = anns.filter((_, i) => user.purchs.includes(i));
        return res.json(result);
    }
})

app.get('/api/user', (req, res) => {
    const id = req.cookies.appuniq;
    const emailSession = ids[id];
    const user = users.find((user) => user.email === emailSession);

    if (!emailSession || !user) {
        return res.status(401).json({error: 'Пользователь не найден'});
    }

    return res.json({ ...user });
});

app.get('/api/anns/:id', (req, res) => {
    const id = req.params.id;
    const user = users[id];

    if (!user) {
        return res.status(401).json({error: 'Пользователь не найден'});
    }

    const result = anns.filter((_, i) => user.anns.includes(i));
    return res.json(result);
});

app.get('/api/me/anns', (req, res) => {
    const id = req.cookies.appuniq;
    const emailSession = ids[id];
    const user = users.find((user) => user.email === emailSession);

    if (!emailSession || !user) {
        return res.status(401).json({error: 'Пользователь не найден'});
    }

    const result = anns.filter((_, i) => user.anns.includes(i));
    return res.json(result);
});

app.get('/api/me/purchs', (req, res) => {
    const id = req.cookies.appuniq;
    const emailSession = ids[id];
    const user = users.find((user) => user.email === emailSession);

    if (!emailSession || !user) {
        return res.status(401).json({error: 'Пользователь не найден'});
    }

    const result = anns.filter((_, i) => user.purchs.includes(i));
    return res.json(result);
});

/** port to listen */
const port = process.env.PORT || 8080;

app.listen(port, () => {
    console.log(`Server listening port ${port}`);
});
