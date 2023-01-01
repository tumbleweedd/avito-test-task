# avito-test-task
Тестовое задание на позицию стажёра в командку Backend (Avito): https://github.com/avito-tech/adv-backend-trainee-assignment

# ОПИСАНИЕ:
Cервис для хранения и подачи объявлений. Объявления хранятся в базе данных. Сервис предоставляет API, работающее поверх HTTP в формате JSON.
Данный проект является тренировочным, задумывался для оттачивания навыков написания веб-сервисов на Go.

# Описание конечных точек:

### Advertisement
| Метод         | URL                 |Описание|
| ------------- |:-------------------------------------:| -------------------------------: |
| POST          |/api/advertisement/                    |Создать объявление                |
| GET           |/api/advertisement/?limit=n&offset=m   |Получить все объявления           |
| GET           |/api/advertisement/:id                 |Получить объявление по его id     |
| PUT           |/api/advertisement/:id                 |Обновить объявление по его id     |
| DELETE        |/api/advertisement/:id                 |Удалить объявление по его id      |

### Images
| Метод         | URL                 |Описание|
| ------------- |:-------------------------------------:| -------------------------------------: |
| POST          |/api/advertisement/:id/image/          |Добавить изображение по id объявления   |
| GET           |/api/advertisement/:id/image/          |Получить все изображения для объявления |
| GET           |/api/advertisement/:advId/image/:imgId |Получить изображение                    |
| DELETE        |/api/advertisement/:advId/image/:imgId |Удалить изображение                     |
