**Api событий**

Запускается на localhost:8080\n
База sqlite. По-умолчанию драйвер - memory.\n 
Можно создать файл `*.db` и указать его в `EventRepository.NewEventRepository`.
Отдаёт JSON. В случае ошибки отдаёт `{"error": "message"}` с http-кодом 400 (Bad Request) или 500 (Internal Error).

Методы:
GET /event - получение событий в промежутке дат
Запрос:
`GET localhost:8080/event?start=2020-01-01 00:00&end=2020-04-01 12:00`
Ответ:
`[
    {
        "Id": 1,
        "Start": "2020-01-02 14:00",
        "End": "2020-01-03 14:00",
        "Duration": "24h0m0s",
        "Name": "test123",
        "Description": "qweqwe123"
    },
    {
        "Id": 3,
        "Start": "2020-01-01 00:00",
        "End": "2020-01-01 14:00",
        "Duration": "14h0m0s",
        "Name": "test1",
        "Description": "qweqwe"
    }
]`

PUT /event - добавление новых событий
Тело запроса:
`[
    {
        "Id": 1,
        "Name": "test1",
        "Start": "2020-01-01 00:00",
        "End": "2020-01-01 14:00",
        "Description": "qweqwe"
    },
    {
        "Id": 2,
        "Name": "test2",
        "Start": "2020-01-02 14:00",
        "End": "2020-01-03 14:00",
        "Description": "qweqwe2"
    }
]`
Ответ: Пустое тело ответа с HTTP-кодом 200

POST /event - обновление событий
Тело запроса:
`[
    {
        "Id": 1,
        "Name": "test123",
        "Start": "2020-01-02 14:00",
        "End": "2020-01-03 14:00",
        "Description": "qweqwe123"
    },
    {
        "Id": 2,
        "Name": "test456",
        "Start": "2020-01-04 12:00",
        "End": "2020-01-04 14:00",
        "Description": "sdfsdf"
    }
]`
Ответ: Пустое тело ответа с HTTP-кодом 200

DELETE /event - удаление событий по id
Тело запроса:
`[1, 2]`
Ответ: Пустое тело ответа с HTTP-кодом 200