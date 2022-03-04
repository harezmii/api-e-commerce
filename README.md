# Golang Fiber Api E-Commerce Boilerplate

# Application Stack
- <font color="cyan">Golang v1.17.5</font>
- <font color="cyan">Fiber v2.27.0</font>
- <font color="cyan">Minio latest</font>
- <font color="cyan">Validation v10 (Turkish Translate)</font>
- <font color="cyan">Hashicorp Vault  v1.3.1</font>
- <font color="cyan">Viper Env v1.10</font>
- <font color="cyan">SwagGo Swag v1.7.8</font>
- <font color="cyan">Fiber Swagger v1.0.3</font>
- <font color="cyan">Ent Entity v0.10.0</font>
- <font color="cyan">Zap Log v1.20.0 </font>
- <font color="cyan">Confluent kafka v1.8.2</font>
- <font color="cyan">Firebase Auth</font>
- <font color="cyan">Postgres pq v1.10.4 </font>
## Directory Structure
```
├───api - Api initialize
│   └───rest
├───cmd - Start App 
├───docs - Swagger Documentation
├───ent - Ent Orm 
│   ├───category
│   ├───comment
│   ├───enttest
│   ├───faq
│   ├───hook
│   ├───image
│   ├───message
│   ├───migrate
│   ├───predicate
│   ├───product
│   ├───profile
│   ├───runtime
│   ├───schema
│   ├───settings
│   └───user
├───internal - Main Application
│   ├───controller - Controller
│   │   ├───category
│   │   ├───comment
│   │   ├───faq
│   │   ├───image
│   │   ├───message
│   │   ├───product
│   │   ├───profile
│   │   └───user
│   ├───entity - Error,Succes Response And Seed
│   │   ├───dto
│   │   ├───response
│   │   └───seed
│   ├───handle - App Handlers
│   ├───infraStructure
│   │   ├───messageQueue
│   │   │   └───kafkaQueue
│   │   └───minio
│   ├───logs - Log Manage
│   ├───secret - Middleware And Secret Process
│   │   ├───hash
│   │   ├───jwtManage
│   │   ├───middleware
│   │   │   └───firabaseAuth
│   │   └───vault
│   ├───storage - Redis .etc Storage 
│   └───validate - Entity Validation
├───pkg - Config 
    └───config

```
## Swagger generates
![](https://user-images.githubusercontent.com/55887187/152854587-672fd01b-b16f-4705-b2a2-f175b283b5ce.png)