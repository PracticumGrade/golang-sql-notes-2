## Пререквизиты

1. Создаем базу данных `notes`:
```bash
CREATE DATABASE notes;
```

2. Создаем нового пользователя с именем `notes` и паролем `notes_pass`:
```bash
CREATE USER notes WITH ENCRYPTED PASSWORD 'notes_pass';
```

3. Предоставляем все привилегии на базу данных `notes` пользователю `notes`:
```bash
GRANT ALL PRIVILEGES ON DATABASE notes TO notes;
```

4. Переключаем на БД `notes`:
```bash
\c notes
```

5. Создаем таблицу `notes`:
```bash
CREATE TABLE notes (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);
```

6. Предоставляем все привилегии на таблицу `notes` пользователю `notes`:
```bash
GRANT ALL PRIVILEGES ON TABLE notes TO notes;
```
