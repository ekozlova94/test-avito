# test-avito

## Решение
Диаграмма работы сервиса:

![Диаграмма](https://raw.githubusercontent.com/ekozlova94/test-avito/master/diagram.png)

Фрагменты кода, решающие конкретные задачи:
* Подписка на изменение цены - https://github.com/ekozlova94/test-avito/blob/b75fb3eeecc95fb193b0e2d35dde1429bfe26fc1/internal/app/test/server/subscription.go#L50
* Отслеживание изменений цены - https://github.com/ekozlova94/test-avito/blob/b75fb3eeecc95fb193b0e2d35dde1429bfe26fc1/internal/app/test/server/subscription.go#L113
* Отправка уведомления на почту - https://github.com/ekozlova94/test-avito/blob/b75fb3eeecc95fb193b0e2d35dde1429bfe26fc1/internal/app/test/sender/prodsender/sender-impl.go#L22
* Работа с БД - https://github.com/ekozlova94/test-avito/tree/master/internal/app/test/store

## Требования к оборудования для запуска
* Git
* Docker
* Docker Compose

## Инструкция по запуску проекта
1. Hеобходимо склонировать репозиторий
2. Запустить сервис из папки с проектом с помощью команды: docker compose up

### Для подписки на изменение цены нужно выполнить GET запрос, например:
curl --location --request GET 'http://localhost:7000/api/v1/subscription?link=https://www.avito.ru/irbit/kollektsionirovanie/banknota_1953852134&email=test@gmail.com' ,

где link - это ссылка на объявление, email - адрес электронной почты (куда прислать уведомление)

