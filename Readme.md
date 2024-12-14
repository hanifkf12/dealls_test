# Technical Test - Backend Engineer

## Project Structure
1. Use fiber as routing packcage.
2. Use sqlx to connect and operation database.
3. Use layer Repository, UseCase, Routing
4. Database use MySQL
5. Implement jwt for authorization.

## Instruction

1. Replace configuration on .env files
    ```
   NAME=test
    PORT=9000
    DB_HOST=localhost
    DB_PORT=3306
    DB_USERNAME=root
    DB_PASSWORD=hanifkf12
    DB_NAME=dealls_test
   ```
2. Running migration use command
    ```
   make run-migration
   ```
3. Running service use command
    ```
    make run-http
   ```