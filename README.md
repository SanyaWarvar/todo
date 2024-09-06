# TodoAPP

base api url = https://todo-vice.onrender.com
## docs



POST **/auth/sign-in**

Принимает json формата: {"username": string, "password": string}

Возвращает: 201 и id созданного пользователя

POST **/auth/sign-up**

Принимает json формата: {"username": string, "password": string}

Возвращает: 200 и json формата {"token": "access_token"}

**ВСЕ ЗАПРОСЫ СЕКЦИИ /api ТРЕБУЮТ header с bearer токеном (полученний на /auth/sign-in) "Authorization": "access_token"**



POST **/api/lists**

Принимает json формата: {"title": string, "description": string}. "description" является необязательным полем.

Создает новый todo список.

Возвращает: 201 и json формата {"list_id": int}

GET **/api/lists**

Возвращает все todo списки, которые создал пользователь

Возвращает: 200 и json формата {"data": [
{"id": int, "title": string, "description": string}]
}

GET **/api/lists/:id**

:id целое число - айди списка, который принадлежит пользователю

Возвращает todo список

Возвращает: 200 и json формата {"id": int, "title": string, "description": string}

GET **/api/lists/:id/items**

:id целое число - айди списка, который принадлежит пользователю

Возвращает все записи тасков, принадлежащих этому списку

Возвращает: 200 и json формата {"data": null | [ "id": int, "title": string, "description": string, "done": bool ] }

POST **/api/lists/:id/items**

:id целое число - айди списка, который принадлежит пользователю

Принимает: json формата {"title": string, "description", "done": bool}. "description", "done" являются необязательными полями.

Добавляет таску в список

Возвращает: 201 и json формата {"id": int }

GET **/api/items/:id**

:id целое число - айди таски, которая принадлежит пользователю. Список указывать не надо.

Принимает: json формата {"title": string, "description", "done": bool}. "description", "done" являются необязательными полями.

Добавляет таску в список

Возвращает: 200 и json формата {"id": int }

PUT **/api/items/:id**

:id целое число - айди таски, которая принадлежит пользователю. Список указывать не надо.

Принимает: json формата {"title": string, "description", "done": bool}. Ни одно из полей не является обязательным, но необходимо, чтобы было хотя бы одно из полей. Менятся будут только переданные поля.

Например если передан только "title", то поля "done" и "description" не изменятся

Изменяет данные таски

Возвращает: 200 и json формата {"details": "Success"}

DELETE **/api/items/:id**

:id целое число - айди таски, которая принадлежит пользователю. Список указывать не надо.

Удаляет таску

Возвращает: 200 и json формата {"details": "Success"}
