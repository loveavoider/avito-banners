Команды:

make build - собирает приложение

make up - запускает приложение, и оно начинае слушать порт 8080

Интеграционный тест находится в файле internal/api/banner/get_integration_test.go

Апи как в swagger

GET /banner - получение всех баннеров

GET /user_banner - получение контента баннера

POST /banner - создать баннер

PATCH /banner/{id} - обновить баннер

DELETE /banner/{id} - удалить баннер

Постман <br/>
[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 128px; height: 32px;">](https://app.getpostman.com/run-collection/26184379-8f9601f9-5a8a-4b76-97c7-982a142fe760?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D26184379-8f9601f9-5a8a-4b76-97c7-982a142fe760%26entityType%3Dcollection%26workspaceId%3Da042dd5e-c99d-468d-bc8d-7a671f156387)


TODO

- [ ] Fix Errors Wrapping
- [ ] Logger
- [ ] Tests
- [ ] Change DI
- [ ] GitHub Actions