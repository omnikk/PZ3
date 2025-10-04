    PZ3 Технологии индустриального программирования.  
    Студент: Выборнов О.А.  
    Группа: ЭФМО-02-25  

Как запустить:

1)cклонировать репозиторий:

    git clone -b main --single-branch https://github.com/omnikk/PZ3.git

2.1) Через .exe файл (tasks.exe)

2.2)Через cmd командой "go run ./cmd/server"

3)Использование переменной окружения PORT (по умолчанию 8080):
   
    $env:PORT="9090"
    .\tasks.exe run


Проверка работы
![Гифка с Gifius ru (3)](https://github.com/user-attachments/assets/c5be4f59-2825-40d4-aad8-a1d1260f19ab)


GET /health  
<img width="1052" height="372" alt="image" src="https://github.com/user-attachments/assets/615b9af4-9cbb-49d7-b534-a4e8c1309f82" />


GET /tasks  
<img width="1048" height="426" alt="image" src="https://github.com/user-attachments/assets/5f4d37ed-d5d6-4cf4-97cc-7f13b8b786f4" />


POST /tasks  
<img width="1041" height="444" alt="image" src="https://github.com/user-attachments/assets/5f4c860f-ead0-41b1-9f52-ae8f79fb7ab8" />


GET /tasks/{id}  
<img width="1051" height="408" alt="image" src="https://github.com/user-attachments/assets/59cce9b7-eab8-427f-a485-9ca5d9395429" />


CORS и preflight OPTIONS  
<img width="658" height="317" alt="image" src="https://github.com/user-attachments/assets/e7d99952-6f50-444d-87a2-acd0ab477dcc" />


CORS middleware  
<img width="1077" height="244" alt="image" src="https://github.com/user-attachments/assets/1b4cd5be-811f-4bf9-a020-ccd06801ff48" />


Проверка title на пустоту при создании задачи 
<img width="1196" height="472" alt="image" src="https://github.com/user-attachments/assets/76cf73d7-8fea-457b-bda5-9a44d35e8482" />
<img width="1079" height="501" alt="image" src="https://github.com/user-attachments/assets/c22608ef-f1b2-4fa1-9501-012655e22403" />


Запуск через PowerShell-скрипт tasks.ps1 или exe  
<img width="760" height="229" alt="image" src="https://github.com/user-attachments/assets/11aac0a8-e020-4db5-86be-764d80c3cd01" />


Cтруктура проекта:

    pz3-http/
      cmd/
        server/
          main.go
      internal/
        api/
          handlers.go
          middleware.go
          cors.go
          responses.go
        storage/
          memory.go
      tasks.ps1
      go.mod
      go.sum
      README.md

Описание структуры:

  cmd/myapp - точка входа в приложение  
  internal/app - основная логика приложения и обработчики
  internal/storage — хранилище задач
  server.exe — собранный бинарник  
  tasks.ps1 — PowerShell скрипт для запуска, сборки, тестов  
  go.mod - файл модуля Go с зависимостями  
  go.sum - контрольные суммы зависимостей  
