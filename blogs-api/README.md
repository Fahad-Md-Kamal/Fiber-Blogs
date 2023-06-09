# Go-Fiber Project

Go is a programming language that's commonly used for building web applications and microservices, Fiber is a web framework for Go that's fast, flexible, and easy to use, and GORM is an ORM library for Go that simplifies database interactions. Combining Fiber and GORM can provide a powerful toolset for building web applications with Go. Moreover, JWT is the modern authentication and authorization mechanism that we are going to use in our project.

Here I'm going to develop a complete API with fiber + GORM + JWT to demonstrate the complete project architecture.

The complete code could is available at:  
[https://github.com/Fahad-Md-Kamal/Fiber-Blogs](https://github.com/Fahad-Md-Kamal/Fiber-Blogs)

## Create and Run Fiber-Go Project

Start go project with the command:

```bash
go mod init github.com/fahad-md-kamal/fiber-blogs
```

> This will start a go project with the file `go.mod`.

```bash
module github.com/fahad-md-kamal/fiber-blogs

go 1.20
```

> All of our dependencies lists will be stored here. This is similar to node project’s `pacakge.json` or `requirements.txt` file of Python projects.

Now let's install Fiber with the command.

```bash
go get github.com/gofiber/fiber/v2
```

**N.B:** This will update `go.mod` file with the dependencies of fiber’s dependencies as follows:

```go
module github.com/fahad-md-kamal/fiber-blogs

go 1.20

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/gofiber/fiber/v2 v2.44.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/klauspost/compress v1.16.3 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/philhofer/fwd v1.1.2 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/savsgio/dictpool v0.0.0-20221023140959-7bf2e61cea94 // indirect
	github.com/savsgio/gotils v0.0.0-20230208104028-c358bd845dee // indirect
	github.com/tinylib/msgp v1.1.8 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.45.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
)
```

> This will also generate a file called `go.sum`. But do not have any concerns with this file since this file is automatically managed by the go package manager on `go.mod` files modification.

Now we are ready to start our Fiber-Go project.  
Let's add a file to our project’s root directory as `main.go`

```go
package main

import "fmt"

func main() {
	fmt.Println("What a day to start fiber-go")
}
```

**Go project’s entrypoint is always `main()` function which resides in `package main`**  
Now let's start our go fiber server. Create a file called **server/server.go.** Now add the following code to the server.go file.

```go
package server

import (
	"github.com/gofiber/fiber/v2"
)

func SetupAndListen() {
	app := fiber.New()
	app.Listen(":3000")
}
```

> Here we are telling our server to create an app instance of fiber and listen to port **3000**.

Now update the main.go file with the following line.

```go
package main

import (
	"github.com/fahad-md-kamal/fiber-blogs/server"
)

func main() {
	server.SetupAndListen()
}
```

**N.B**. If you are using vscode’s go plugin then after writing the package name which in this case is **server** you will see suggestions to be auto-imported to the project.  
Now open the terminal and type

```bash
go run main.go
```

> This will start your go project and start listening to the port **3000**. You will see this something similar to the following to your terminal.

![Screenshot 2023-04-25 at 12.56.39 PM.png](https://user-images.githubusercontent.com/34704464/235286857-858a7ee1-7c40-444e-b0c9-e931e9c241ab.png)

Congratulations !!

We have started the Fiber Go server.

Code for this far could be found at:  
[https://github.com/Fahad-Md-Kamal/Fiber-Blogs/tree/part-1/project-setup](https://github.com/Fahad-Md-Kamal/Fiber-Blogs/tree/part-1/project-setup)

<br>
<br>
<br>


## Create API Endpoint
---


Now let's create an API endpoint With Fiber:

Create a folder called users (***With the intention to modularize the project***)  
Now create another file called `users/userControllers.go` and add an API handler function to it.

```go
package controllers

import "github.com/gofiber/fiber/v2"

func GetUsersListHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Yeee!!!, Fiber Project has started",
	})
}
```

Now register this route to the `users/routes.go` routes list as

```go
package users

import (
	"github.com/fahad-md-kamal/fiber-blogs/users/controllers"
	"github.com/gofiber/fiber/v2"
)

func UsersRouts(app *fiber.App) {
	router := app.Group("users")

	router.Get("/", controllers.GetUsersListHandler)
}
```

Finally include this Users Module route to Our main app on `server.go` as follows.

```go
package server

import (
	"github.com/fahad-md-kamal/fiber-blogs/users"
	"github.com/gofiber/fiber/v2"
)

func SetupAndListen() {
	app := fiber.New()

	users.UsersRouts(app)
	app.Listen(":3000")
}
```

Now restart the project and hit the API [http://localhost:3000/users](http://localhost:3000/users) and you will see the following response:

![Run Project Success](https://user-images.githubusercontent.com/34704464/235286896-a44e7b20-c994-4a87-8703-5aa69f22a338.png)

Congratulations !!

Our Fiber API receives API requests from clients. returns response.

Now our project’s structure should look as follows:


```md
├── LICENSE
├── README.md
├── go.mod
├── go.sum
├── main.go
├── server
│ └── server.go
├── users
│ ├── controllers
│ │ └── userControllers.go
│ └── routes.go
```

The code could be found here:

[https://github.com/Fahad-Md-Kamal/Fiber-Blogs/tree/part-1/project-setup](https://github.com/Fahad-Md-Kamal/Fiber-Blogs/tree/part-1/project-setup)

<br>
<br>
<br>

## Connect to database (Postgres) using GORM
---


Inorder to connect Postgres Database we need to install the GORM package and GORM’s postgres database driver:

```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

We cannot push our credentials publicly, therefore, we also need some additional packages to load environment variables.

```bash
go get github.com/joho/godotenv
go get github.com/mitchellh/mapstructure
```

> Here, `godotenv` will load environment (env) variables and `mapstructure` will be mapping those env variables into go structs.

Read Environment variables and return those as go struct from `configs/envVars.go`

```go
package configs

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
)

type EnvConfig struct {
	ServingPort  string `mapstructure:"SERVING_PORT"`
	DbHost       string `mapstructure:"DB_HOST"`
	DbPort       string `mapstructure:"DB_PORT"`
	DbName       string `mapstructure:"DB_NAME"`
	DbUser       string `mapstructure:"DB_USER"`
	DbPassword   string `mapstructure:"DB_PASSWORD"`
	SecretKey    string `mapstructure:"SECRET_KEY"`
	JwtSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

var ENVs EnvConfig

func LoadEnvs() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	envVars := make(map[string]string)
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		envVars[pair[0]] = pair[1]
	}

	// var cfg EnvConfig
	err = mapstructure.Decode(envVars, &ENVs)
	if err != nil {
		return fmt.Errorf("error decoding env vars: %w", err)
	}

	return nil
}
```

> Here we have declared a struct `EnvConfig` and mapping environment variables according to `EnvConfig` Struct through the `LoadEnvs` function.
>
> > Note that, here LoadEnvs tries to load environments from `.env` file.

***N.B: Add a file to the project’s root directory named `.env` and add the variables that you have declared on `EnvConfig`.***

```
SERVING_PORT=:8000
DB_HOST=localhost
DB_PORT=5432
DB_NAME=blog_db
DB_USER=postgres
DB_PASSWORD=postgres
SECRET_KEY=123456789
JWT_SECRET_KEY=2
```

Now let's create a database connection.  
Create a file `database/dbSetup.go`

```go
package database

import (
	"fmt"

	"github.com/fahad-md-kamal/fiber-blogs/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConfig() error {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Dhaka",
		configs.ENVs.DbHost, configs.ENVs.DbUser, configs.ENVs.DbPassword, configs.ENVs.DbName, configs.ENVs.DbPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
```

> Here we are creating a Global `var DB *gorm.DB` variable that could be accessed from all over the project to interact with the database.
> 
> The `DbConfig` function loads environment configs and generates interpolated environment values for the Database connection to be executed.

Now in order to load environment variables before creating a database connection need to execute `LoadEnvs()` function from `main.go` so that it loads environment configs before creating a database connection.

```go
package main

import (
	"github.com/fahad-md-kamal/fiber-blogs/configs"
	"github.com/fahad-md-kamal/fiber-blogs/database"
	"github.com/fahad-md-kamal/fiber-blogs/server"
)

func main() {
	if err := configs.LoadEnvs(); err != nil {
		panic(err.Error())
	}
	if err := database.DbConfig(); err != nil {
		panic(err.Error())
	}
	server.SetupAndListen()
}
```

Create a GORM struct for Users as follows:

```go
package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Username    string `gorm:"unique;not null" json:"username"`
	Email       string `gorm:"unique;not null" json:"email"`
	Password    string `gorm:"not null" json:"password"`
	IsSuperuser bool   `gorm:"default=false;not null" json:"is_superuser"`
	IsActive    bool   `gorm:"default=true;not null" json:"is_active"`
}
```

> **N.B:** Since we are adding `gorm.Model`, GORM will automatically add `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` fields to the struct and database table.

Update `database/dbSetup.go` file to auto-migrate changes to the database on system start.

```go
package database

import (
	"fmt"

	"github.com/fahad-md-kamal/fiber-blogs/configs"
	usermodels "github.com/fahad-md-kamal/fiber-blogs/users/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate() {
	DB.AutoMigrate(&usermodels.Users{})
}

var DB *gorm.DB

func DbConfig() error {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Dhaka",
		configs.ENVs.DbHost, configs.ENVs.DbUser, configs.ENVs.DbPassword, configs.ENVs.DbName, configs.ENVs.DbPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	Migrate() // Calling this function to apply changes to the database
	return nil
}
```

> Here we have added Migrate function that will be called from DbConfig function. Since we want to apply all our changes to the database during the server start, therefore, we are calling the `Migrate()` function from `DbConfig()` function which is called from `main.go` before starting the fiber server.
> 
> In Migrate function we are passing our Users struct to create the database table based on the gorm struct.
>
>> **N.B:** All the structs that we are going to generate, will be added to `DB.AutoMigrate(&usermodels.Users{})` comma separated.  

Besides, since our user model is in the Models package of the `Users module`, therefore, we are importing the package as:

```go
package migrations

import (
	...
	usermodels "github.com/fahad-md-kamal/fiber-blogs/users/models"
	...
)
```

That’s it now run the project and check the database. It will apply all the changes to the database right before the server starts.

### Congratulations !!

We have connected the database and created User table with GORM.

*Folder Architecture should look like this.*

```md
├── LICENSE
├── README.md
├── configs
│ └── envVars.go
├── database
│ └── dbSetup.go
├── go.mod
├── go.sum
├── main.go
├── server
│ └── server.go
├── users
│ ├── controllers
│ │ └── userControllers.go
│ ├── models
│ │ └── users.go
│ └── routes.go
```

***The code could be found here:***

[https://github.com/Fahad-Md-Kamal/Fiber-Blogs/tree/part-2/database-connection](https://github.com/Fahad-Md-Kamal/Fiber-Blogs/tree/part-2/database-connection)

<br>
<br>
<br>

## CRUD
---

Before going any further, we need to configure our development server to auto reload after any change we made to our codebase, incase of avoiding manual server restart. 
Here I've used [air](https://github.com/cosmtrek/air) to auto reload my developement server.  
You can follow the instruction from the following url:
[https://github.com/cosmtrek/air](https://github.com/cosmtrek/air).

**I've followed install.sh but you can choose something else.**

```bash
# binary will be $(go env GOPATH)/bin/air
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# or install it into ./bin/
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

air -v
```

`air`'s config file on Root diractory `.air.toml`

```txt
# .air.toml

root = "."
tmp_dir = "tmp"
build_dir = "tmp/build"

[[runners]]
  name = "Fiber"
  path = "."
  args = ["./tmp/build/main"]
  env = {}

[runners.log]
  mode = "console"
  prefix = "Fiber"
  color = true

```

Instead of running server with `go run main.go` now run the server with the command

```bash
air
```

This should now run the server and detect any changes to any file and auto reload the server.  
Now Lets start developing our API Endpoints

<br>
<br>
<br>

## CRUD - C: Create

I have seperated database migration machanisams to a seperate package named migrations as follows:

```go
package migrations

import (
	"github.com/fahad-md-kamal/fiber-blogs/database"
	usermodels "github.com/fahad-md-kamal/fiber-blogs/users/models"
)

func MigrateChanges() {
	database.DB.AutoMigrate(
		&usermodels.Users{},
	)
}
```
> This will help us to avoid cyclic import and seperate database configurations from project's configurations.

`dbSetup.go` file is modified to

```go
package database

import (
	"fmt"

	"github.com/fahad-md-kamal/fiber-blogs/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConfig() error {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Dhaka",
		configs.ENVs.DbHost, configs.ENVs.DbUser, configs.ENVs.DbPassword, configs.ENVs.DbName, configs.ENVs.DbPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
```

> I am planning to seperate model validation and model representation layer seperated from application and database layer. Since, I belive, application layer's responsibility is to implement business logic and database layer's responsibility is to communicate data.
> 
> Dto's for validation and representation layer.
> 
> models for database layer.
>
> Controller / Utils / helpers for business layer.

Add a file `users/dtos/userDtos.go` that will be used to validate user's request and return API response.

```go
package dtos

type UserCreateDto struct {
	Username string `json:"username" validate:"required,min=4,max=50"`
	Email    string `json:"email" validate:"required,email,min=8,max=100"`
	Password string `json:"password" validate:"required,min=6"`
}
```

And the Response dto

```go
type UserResponseDto struct {
	Id          uint      `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	IsActive    bool      `json:"is_active"`
	IsSuperuser bool      `json:"is_superuser"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (udto *UserResponseDto) ParseToResponseDto(user *models.Users) {
	udto.Id = user.ID
	udto.Username = user.Username
	udto.Email = user.Email
	udto.IsSuperuser = user.IsSuperuser
	udto.IsActive = user.IsActive
	udto.CreatedAt = user.CreatedAt
	udto.UpdatedAt = user.UpdatedAt
}

```

> This will help us to avoid returing user's password or similar type secure credentials.

Now we want to validate the data before creating the user. Therefore, we are going to use a package called validator from the go.  
Install the package:

```bash
go get github.com/go-playground/validator/v10
```

Now let's create an utility package that could be used globally for any model that we want to validate.  
Add a file `utils/validateStructs.go`.

```go
package utils

import "github.com/go-playground/validator"

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(inputStruct interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(inputStruct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

```

> Here we are creating ErrorResponse struct to generate all errors as error list.
> 
> In `ValidateStruct()` function we are passing our struct. Then this will check each fields of the struct using it's validat rules. It will show errors list and will return it.

Now we are going to use it on `userDtos.go` as:

```go
package dtos

...

func (data *UserCreateDto) ValidateUserCreateDto() ([]*utils.ErrorResponse, bool) {
	errors := utils.ValidateStruct(data)
	return errors, len(errors) == 0
}
```

Now add two functions to Users model on `users.go`

```go
package models

import (
	"fmt"

	"github.com/fahad-md-kamal/fiber-blogs/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username    string `gorm:"unique;not null" json:"username"`
	Email       string `gorm:"unique;not null" json:"email"`
	Password    string `gorm:"not null" json:"password"`
	IsSuperuser bool   `gorm:"default=false;not null" json:"is_superuser"`
	IsActive    bool   `gorm:"default=true;not null" json:"is_active"`
}

func (u *Users) ValidateUserExists() (string, bool) {
	var user Users
	result := database.DB.Where("username = ? OR email = ?", u.Username, u.Email).First(&user)
	return fmt.Sprintf("User exists with username: %s OR email: %s", u.Username, u.Email), result.RowsAffected > 0
}

func (u *Users) Save() error {
	if u.ID == 0 {
		if result := database.DB.Create(&u); result.Error != nil {
			return result.Error
		}
	} else {
		if result := database.DB.Save(&u); result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func (u *Users) GeneratePasswordHash() (error, bool) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err, false
	}
	u.Password = string(hashedPassword)
	return nil, true
}
```

> `ValidateUserExists()` will check if user exists with username or email.
>
> `Save()` will create object if there is no Id otherwise will save it
>
> `GeneratePasswordHash()` will generate password hash before saving it.

Now we will update our `AddUserHandler()` for creating users as:

```go
func AddUserHandler(c *fiber.Ctx) error {

	var userCreateDto dtos.UserCreateDto

	if err := c.BodyParser(&userCreateDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if errors, ok := userCreateDto.ValidateUserCreateDto(); !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors})
	}

	UserToCreate := userCreateDto.ParseFromDto()
	if err, ok := UserToCreate.GeneratePasswordHash(); !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if message, ok := UserToCreate.ValidateUserExists(); ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": message})
	}

	if err := UserToCreate.Save(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	responseDto := new(dtos.UserResponseDto)
	responseDto.ParseToResponseDto(UserToCreate)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": responseDto,
	})
}
```

Let's Register this handler function to the routes as:

```go
package users

import (
	"github.com/fahad-md-kamal/fiber-blogs/users/controllers"
	"github.com/gofiber/fiber/v2"
)

func UsersRouts(app *fiber.App) {
	router := app.Group("users")

	router.Post("/", controllers.AddUserHandler)
}
```


Let's request from Postman to create a user.

**_It show error if any fields were missing._**
![Missing Field Error](https://user-images.githubusercontent.com/34704464/235286588-3c50ffcd-68c6-48ef-b853-4f869579db94.png)

> It will create user with hashed password. On successfully creating user, it will return the user as `UserResponseDto`

Our Current Folder stracture should be now as:

```md
├── LICENSE
├── README.md
├── configs
│ └── envVars.go
├── database
│ └── dbSetup.go
├── example.env
├── go.mod
├── go.sum
├── main.go
├── migrations
│ └── migrations.go
├── server
│ └── server.go
├── users
│ ├── controllers
│ │ └── userControllers.go
│ ├── dtos
│ │ └── userDtos.go
│ ├── models
│ │ └── users.go
│ └── routes.go
│ │── utils
│ │ │── pagination.go
│ │ └── validateStructs.go
```

Code for this far could be found here:  
[https://github.com/Fahad-Md-Kamal/Fiber-Blogs/tree/part-3/api/user-create](https://github.com/Fahad-Md-Kamal/Fiber-Blogs/tree/part-3/api/user-create)

<br>
<br>
<br>

## CRUD - R: Read (List)

Lets create our User's List API Handler

```go
func GetUsersListHandler(c *fiber.Ctx) error {

	// Parse pagination parameters
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset := (page - 1) * limit

	// Get Users List
	users, totalCount, err := models.GetUsersList(limit, offset)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	// Convert User's list into response Dtos
	userDtos := dtos.ParseUsersListToResponseDto(&users)

	// Get Paginated Response
	pagination := utils.Paginate(int(totalCount), limit, page, userDtos)
	return c.JSON(pagination)
}
```

> Here we have added page, limit and offset for paginated reponse of user's list.
> 
> We have added a function `GetUsersList()` to our model Users since we are interacting with database from our models package.
>
> We are passing `limit` and `offset` to it as parameters and receiving `users list`, `totalCount` and `error` from it.

```go
// users/models/users.go

func GetUsersList(limit, offset int) ([]Users, int64, error) {
	var users []Users
	var totalCount int64

	if err := database.DB.Model(Users{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	if err := database.DB.Model(Users{}).Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}

```

Then we are Parsing `models.Users` list into User's Response Dto list in order to hide some fields from users.

```go
func ParseUsersListToResponseDto(users *[]models.Users) []UserResponseDto {
	usersList := []UserResponseDto{}
	for _, user := range *users {
		usersList = append(usersList, UserResponseDto{
			Id:          user.ID,
			Username:    user.Username,
			Email:       user.Email,
			IsActive:    user.IsActive,
			IsSuperuser: user.IsSuperuser,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
		})
	}
	return usersList
}

```

Finally preparing the paginated Response to send to the user.

```go
type Pagination struct {
	TotalCount  int64       `json:"total_count"`
	Limit       int         `json:"limit"`
	CurrentPage int         `json:"current_page"`
	TotalPages  int         `json:"total_pages"`
	HasNextPage bool        `json:"has_next_page"`
	HasPrevPage bool        `json:"has_prev_page"`
	NextPage    int         `json:"next_page"`
	PrevPage    int         `json:"prev_page"`
	Data        interface{} `json:"data"`
}

func Paginate(totalCount, limit, currentPage int, data interface{}) *Pagination {
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))
	hasNextPage := currentPage < totalPages
	hasPrevPage := currentPage > 1
	nextPage := currentPage + 1
	prevPage := currentPage - 1

	return &Pagination{
		TotalCount:  int64(totalCount),
		Limit:       limit,
		CurrentPage: currentPage,
		TotalPages:  totalPages,
		HasNextPage: hasNextPage,
		HasPrevPage: hasPrevPage,
		NextPage:    nextPage,
		PrevPage:    prevPage,
		Data:        data,
	}
}
```

Register it to our routes list:
```go
package users

import (
	"github.com/fahad-md-kamal/fiber-blogs/users/controllers"
	"github.com/gofiber/fiber/v2"
)

func UsersRouts(app *fiber.App) {
	router := app.Group("users")

	...
	router.Get("/", controllers.GetUsersListHandler)
}

```


Now run the application and see the paginated response.

```json
{
  "total_count": 2,
  "limit": 10,
  "current_page": 1,
  "total_pages": 1,
  "has_next_page": false,
  "has_prev_page": false,
  "next_page": 2,
  "prev_page": 0,
  "data": [
    {
      "id": 1,
      "username": "fahad",
      "email": "fahadmdkamal@gmail.com",
      "is_active": false,
      "is_superuser": false,
      "created_at": "2023-04-26T00:35:07.392797+06:00",
      "updated_at": "2023-04-26T00:35:07.392797+06:00"
    },
    {
      "id": 2,
      "username": "fahadmdkamal",
      "email": "faahad.hossain@gmail.com",
      "is_active": false,
      "is_superuser": false,
      "created_at": "2023-04-26T00:40:31.994202+06:00",
      "updated_at": "2023-04-26T00:40:31.994202+06:00"
    }
  ]
}
```

> That's it about `Read (List)` API.

Code for this far could be found here:
[https://github.com/Fahad-Md-Kamal/Fiber-Blogs/tree/part-3/api/paginated-list](https://github.com/Fahad-Md-Kamal/Fiber-Blogs/tree/part-3/api/paginated-list)

<br>
<br>
<br>

## CRUD - R : Read (Detail)
---

First we need to create another function named `GetUserDetailHandler` at `users/models/users.go` file since we are following a principle of communicating to database from models.

```go
func GetUserById(userId uint) (*Users, error) {
	var user Users
	result := database.DB.First(&user, userId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
```

> This will take userId as parameter and return `User` and `error` as return type.

Now lets create a parsing function to parse DB user into `DtoUser` function that will take a DB user model and convert it into a `DtoUser`.

```go
func ParseUserToResponseDto(user *models.Users) *UserResponseDto {
	userDto := UserResponseDto{
		Id:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		IsActive:    user.IsActive,
		IsSuperuser: user.IsSuperuser,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
	return &userDto
}
```

> Here we are passing DB User and maping each field to a `UserResponseDto` model.

Now Lets create a `UserDetailHandler` function.

```go
func GetUserDetailHandler(c *fiber.Ctx) error {
	userId, err := strconv.ParseUint(c.Params("id"), 10, 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Invalid User Id",
		})
	}

	user, err := models.GetUserById(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to get user",
		})
	}

	dtoUser := dtos.ParseUserToResponseDto(user)
	return c.JSON(fiber.Map{
		"data": &dtoUser,
	})
}
```

First of all, we are converting `id` into a `uint` type (**_e.g. We could work with string. But for standered practice, it's better we convert it here_**)

> If there is any error we immidiately show the actual error as error field and message as our custom error field.

Then we are getting user from the database using the `GetUserById()` function that we have already created in our `users/models/users.go` file as:

```go
	user, err := models.GetUserById(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to get user",
		})
	}
```

> If there is any error, it will return custom error message and system error message with a status of 404 not found.

Finally, it will convert the database user into Dto User in order to hide fields that shouldn't be seen by end users.

```go
	dtoUser := dtos.ParseUserToResponseDto(user)
	return c.JSON(fiber.Map{
		"data": &dtoUser,
	})
```

Now create an API route in our `users/routes.go`:

```go
router.Get("/:id", controllers.GetUserDetailHandler)
```

> Here we are providing `:id` as dynamic value into our url that will be parsed into uint into our handler finction.

Let's hit the endpoint `{{url}}/users/1` and test our api:

![User Detail API POST](https://user-images.githubusercontent.com/34704464/235286434-7f5e1648-a3d7-4cb2-8092-2f689aada587.png)

### Congratulations !!

We have developed the Detail API endpoint.

Code for this far could be found here:  
[https://github.com/Fahad-Md-Kamal/Fiber-Blogs/tree/part-3/api/user-details](https://github.com/Fahad-Md-Kamal/Fiber-Blogs/tree/part-3/api/user-details)

<br>
<br>
<br>

## CRUD - U : UPDATE (PUT)
---

In order to develop update endpoint first, lets create a `UpdateUser()` function into our`users/models/users.go` file that will be responsible for communicating with the database.

```go
func (userToUpdate *Users) UpdateUser(updateDto interface{}, omitFields ...string) (*Users, error) {
	if result := database.DB.Model(userToUpdate).Omit(omitFields...).Updates(updateDto); result.Error != nil {
		return nil, result.Error
	}
	return userToUpdate, nil
}
```

> Here we are passing an interface that will be passed from the handler to update the user.
>
> We are also accepting optional parameters that should be passed to be ignored on update fields.
>
> > **_e.g. This is because we want to restric general users to update certain fields such IsSuperuser, IsActive, Username etc._**

We need to update our `ValidateUserExists()` function that we have developed for creating user and create another function `ValidateUserExistsWithEmailOrUsername` to check user exists with the `username` and `email`.

```go
type UserCheckParams struct {
	UserId   uint
	Username string
	Email    string
}

func ValidateUserExistsWithEmailOrUsername(params UserCheckParams) (string, bool) {
	var count int64
	query := database.DB.Model(&Users{}).Where("username = ? OR email = ?", params.Username, params.Email)
	if params.UserId > 0 {
		query = query.Not("id = ?", params.UserId)
	}
	err := query.Count(&count).Error
	if err != nil {
		return err.Error(), true
	}
	return "User exists with the given attribute(s)", count > 0
}

func (u *Users) ValidateUserExists() (string, bool) {
	userParams := UserCheckParams{
		UserId:   u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
	return ValidateUserExistsWithEmailOrUsername(userParams)
}
```

> Here we have created a `UserCheckParams` struct to control the parameters from single point.

Now lets add an update `UpdateUserDto` that will be responsible for validating user update fields

```go
type UserUpdateDto struct {
	Email       string `json:"email" validate:"omitempty,email,min=8,max=100"`
	IsSuperuser *bool  `json:"is_superuser" validate:"omitempty"`
	IsActive    *bool  `json:"is_active" validate:"omitempty"`
}

func (data *UserUpdateDto) ValidateUserUpdateDto() ([]*utils.ErrorResponse, bool) {
	errors := utils.ValidateStruct(data)
	return errors, len(errors) == 0
}
```

> Here we added `validate:"omitempty"` to `UserUpdateDto` struct to make fields optional for validator to validate.
>
> We are also taking `IsSupseruser` or `IsActive` boolean pointer so that, if user passes a value it shouldn't be null otherwise it will be a null.

**Handler Function**

```go

func UpdateUserHandler(c *fiber.Ctx) error {
	userId, err := strconv.ParseUint(c.Params("id"), 10, 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Invalid User Id",
		})
	}

	var userUpdateDto dtos.UserUpdateDto
	if err := c.BodyParser(&userUpdateDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to parse provided data",
		})
	}

	if errors, ok := userUpdateDto.ValidateUserUpdateDto(); !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   errors,
			"message": "Invalid data to update",
		})
	}

	userToUpdate, err := models.GetUserById(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to get user",
		})
	}

	userCheckParams := models.UserCheckParams{
		UserId: userToUpdate.ID,
		Email:  userUpdateDto.Email,
	}
	msg, exists := models.ValidateUserExistsWithEmailOrUsername(userCheckParams)

	if exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   msg,
			"message": msg,
		})
	}

	updatedUser, err := userToUpdate.UpdateUser(&userUpdateDto)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to update user",
		})
	}

	dtoUser := dtos.ParseUserToResponseDto(updatedUser)
	return c.JSON(fiber.Map{
		"data": &dtoUser,
	})
}
```

> Here we are taking user id from the request and parsing it to `uint` and Getting user with the given ID.
>
> Parsing provided data to `UserUpdateDto` and validating fields.
>
> Later we are geting user from the database
>
> Validating User with the given email exists.
>
> Finally we are updating user and returning the udpated user data by persing it `intoUserToResponseDto`

Now register this handler to `users/routes.go` with for put request:

```go

func UsersRouts(app *fiber.App) {
	router := app.Group("users")

	...
	router.Put("/:id", controllers.UpdateUserHandler)
}
```

Now hit the api : `{{url}}/users/1`

![User UPDATE API](https://user-images.githubusercontent.com/34704464/235298500-e5cba53e-6d5d-4727-80c5-a366a5e0c226.png)

### Congratulations !!!

We have created an update API.

Our Folder stracture's will be same as before.

Code for this far could be found at:  
[https://github.com/Fahad-Md-Kamal/Fiber-Blogs/tree/part-3/api/user-update](https://github.com/Fahad-Md-Kamal/Fiber-Blogs/tree/part-3/api/user-update)

<br>
<br>
<br>

## CRUD - D : DELETE (Delete)

---


`users/models/users.go`

```go
func (u *Users) DeleteUser() error {
	if result := database.DB.Delete(&u); result.Error != nil {
		return result.Error
	}
	return nil
}
```

In this function we are simply deleting user from the database. (**_e.g. Soft delete_**)

**Delete User Handler**

```go
func DeleteUserHandler(c *fiber.Ctx) error {
	userId, err := strconv.ParseUint(c.Params("id"), 10, 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Invalid User Id",
		})
	}

	userToDelete, err := models.GetUserById(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to get user",
		})
	}

	if err := userToDelete.DeleteUser(); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Couldn't delete user",
		})
	}
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}
```

> Here we are reusing most of the topic from `GetUserDetailHandler` but instead of returning users we are using the DeleteUser() function that we have developed early and on success delete we are returning `204` status without any error message.

Now Lets register this Handler to our routes

```go
func UsersRouts(app *fiber.App) {
	router := app.Group("users")
	...
	router.Delete("/:id", controllers.DeleteUserHandler)
}
```

Congratulations !!
We have successfully developed the Delete API.

After hitting `{{url}}/users/<user_id_to_delete>` **_e.g: `{{url}}/users/2`_** we can successfully delete a user.

![User Delete API](https://user-images.githubusercontent.com/34704464/235302936-70d17000-860f-44ef-8755-54c9465ab8dd.png)

After this far, folder stracture should look like this

```
├── LICENSE
├── README.md
├── configs
│ └── envVars.go
├── database
│ └── dbSetup.go
├── example.env
├── go.mod
├── go.sum
├── main.go
├── migrations
│ └── migrations.go
├── server
│ └── server.go
├── users
│ ├── controllers
│ │ └── userControllers.go
│ ├── dtos
│ │ └── userDtos.go
│ ├── models
│ │ └── users.go
│ └── routes.go
└── utils
├── pagination.go
└── validateStructs.go
```

<br>
<br>
<br>

## *N.B. I have update the project to show log instead of returning system error to users*

See in code:  
[https://github.com/Fahad-Md-Kamal/Fiber-Blogs/tree/part-4/logger](https://github.com/Fahad-Md-Kamal/Fiber-Blogs/tree/part-4/logger)

<br>
<br>
<br>

## Let's Create User Login API's Including JWT token

---

<br>


First we will install Golang-jwt package for jwt token

```bash
go get github.com/golang-jwt/jwt
```

`configs/envVars.go`

```go
type EnvConfig struct {
	...
	JwtSecretKey  string `mapstructure:"JWT_SECRET_KEY"`
	TokenLifeTime string `mapstructure:"TOKEN_LIFETIME"`
}
```

> Here I have added to environment variables that will be used for generating JWT token

Update `ValidateUserExistsWithEmailOrUsername` to return 3 data instead of only two. This will help use reuse the function from several points.

```go
func ValidateUserExistsWithEmailOrUsername(params UserCheckParams) (*Users, string, bool) {
	....
	return &dbUser, msg, exists
}
```

Added several model Dto structs for login Request, login Response, token claim as well as added `ParseToLoginResponseDto` function.

```go

type TokenClaimPayload struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type LoginRequestDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (data *LoginRequestDto) ValidateLoginRequestDto() ([]*utils.ErrorResponse, bool) {
	errors := utils.ValidateStruct(data)
	return errors, len(errors) == 0
}

type LoginResponseDto struct {
	UserID      uint   `json:"userId"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Token       string `json:"token"`
	IsSuperuser bool   `json:"isSuperUser"`
}

func ParseToLoginResponseDto(token string, u *models.Users) *LoginResponseDto {
	loginResponseDto := LoginResponseDto{
		UserID:   u.ID,
		Username: u.Username,
		Email:    u.Email,
		Token:    token,
	}

	return &loginResponseDto
}
```

Added `ValidatePasswordHash()` at `users/models/users.go` to validate password hash

```go
func (user *Users) ValidatePasswordHash(password string) (string, bool) {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Printf("comparing token got error %s", err.Error())
		return "Invalid credentials", false
	}
	return "", true
}
```

Added `GenerateJwtToken()` that will be generating jwt tokens.

```go

func GenerateJwtToken(user *models.Users) (string, bool) {
	tokenLifetime, err := strconv.ParseInt(configs.ENVs.TokenLifeTime, 10, 0)
	if err != nil {
		log.Printf("Failed to read token lifetime environment variable: %s", err.Error())
		return "Failed to read token lifetime environment variable", false
	}
	claims := &dtos.TokenClaimPayload{
		ID:       user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(tokenLifetime) * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(configs.ENVs.JwtSecretKey))
	if err != nil {
		log.Printf("Internal server error while attaching signe to the token: %s", err.Error())
		return "Error for generating signed token", false
	}

	return tokenString, true
}
```

> Here token claim fields are added along with token lifetime duration.
>
> Later I'm signing the token with our JwtSecretKey

Now I'am adding a loginHandler function on `users/controllers/authControllers.go` file

```go

func LoginHandler(c *fiber.Ctx) error {
	loginRequestData := dtos.LoginRequestDto{}
	if err := c.BodyParser(&loginRequestData); err != nil {
		log.Printf("Error parsing Login Request: %s | Error: %s", c.Params("id"), err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Login Data"})
	}

	if errors, ok := loginRequestData.ValidateLoginRequestDto(); !ok {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": errors})
	}

	userCheckParams := models.UserCheckParams{
		Username: loginRequestData.Username,
		Email:    loginRequestData.Username,
	}
	dbUser, _, exists := models.ValidateUserExistsWithEmailOrUsername(userCheckParams)
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	tokenString, success := helpers.GenerateJwtToken(dbUser, loginRequestData.Password)
	if !success {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Faild to create auth token"})
	}
	userResponseDto := dtos.ParseToLoginResponseDto(tokenString, dbUser)
	return c.JSON(&userResponseDto)
}
```

> First I'm parsing user data from request to `LoginUserRequestDto` and validating it with the function `loginRequestData.ValidateLoginRequestDto()`.
>
> Then I'm validating and fetching user from database if user exists.
>
> Later I'm generating jwt token and parsing the user to `LoginUserResponseDto` before returning it.

Finally, I've created an additional router named `unProtectedRoute` since that will be open for anyone to login into the system.

```go
func UsersRouts(app *fiber.App) {
	router := app.Group("users")
	...

	unProtectedRoute := app.Group("")
	unProtectedRoute.Post("/login", controllers.LoginHandler)
}

```

That's it, our API now Allows us to login into the system.

![Login API Response](https://user-images.githubusercontent.com/34704464/235322125-e69257b6-db80-4c0b-bfce-6551819fd8d3.png)

<br>
<br>
<br>

## Now let's develope logout api
---

First we need to create a `BlacklistedToken` model to store tokens into the database with one functions to Check if token is balcklisted or not. And other to store the token into the database as blacklisted.

```go
type BlacklistedTokens struct {
	gorm.Model
	Token string `gorm:"uniqueIndex" json:"token"`
}

func (t *BlacklistedTokens) IsTokenBlacklisted() bool {
	var count int64
	if err := database.DB.Where("token = ?", t.Token).Find(&t).Count(&count).Error; err == nil {
		return count > 0
	}
	return count > 0
}

func (t *BlacklistedTokens) BlacklistToken() bool {
	if err := database.DB.Model(&t).Create(&t).Error; err != nil {
		log.Printf("Error making token as blacklisted %s", err.Error())
		return false
	}
	return true
}

```

`users/controllers/authControllers.go`

```go
func LogoutHandler(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"message": "Missing Authorization header",
		})
	}

	authToken := strings.Split(authHeader, " ")[1]

	// Invalidate the token by adding it to the blacklist
	blacklistedToken := models.BlacklistedTokens{Token: authToken}

	if ok := blacklistedToken.IsTokenBlacklisted(); ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token or token expired",
		})
	}

	if ok := blacklistedToken.BlacklistToken(); !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Couldn't blacklist token",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User logged out",
	})
}
```

Here we are parsing token from request header and validating if token is already blacklisted or not. Finally we are making token as blacklisted.

Now register this token to the routers and hit the api. It will logout users successfully.

![user logout success](https://user-images.githubusercontent.com/34704464/235322440-04daf536-07e5-41a8-9c91-be46ae2e5129.png)

**_N.B. We will move the validation part of token if it's blacklisted or not into our middleware function later since that will be responsible for validating every requests. Once we develop that._**

Congratulations !!

Our Login & Logout API with JWT token (serverless access) complete.

<br>
<br>
<br>

## Middleware
---

Let's create middleware function to protect our routes from unauthorized users.  
Create a file called `middlewares/jwtMiddleware.go` add a middleware function as follows:

```go
package middlewares

import (

	...
	userdtos "github.com/fahad-md-kamal/fiber-blogs/users/dtos"
	usermodels "github.com/fahad-md-kamal/fiber-blogs/users/models"
	...

)

func JwtMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing authorization header"})
		}
		tokenString := authHeader[len("Bearer "):]

		blacklistedToken := usermodels.BlacklistedTokens{Token: tokenString}
		if blacklistedToken.IsTokenBlacklisted() {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token or token expired",
			})
		}

		token, err := jwt.ParseWithClaims(tokenString, &userdtos.TokenClaimPayload{},
			func(token *jwt.Token) (interface{}, error) {
				// Get the signing key from your authentication server or config file
				signingKey := []byte(configs.ENVs.JwtSecretKey)
				return signingKey, nil
			})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid JWT signature"})
			}
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid JWT token"})
		}

		// Check if token is valid
		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid JWT token"})
		}

		// Extract claims from JWT token
		claims, ok := token.Claims.(*userdtos.TokenClaimPayload)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid JWT token"})
		}

		var user usermodels.Users
		user.ID = claims.ID
		user.Username = claims.Username
		user.IsSuperuser = claims.IsSuperuser
		user.Email = claims.Email
		c.Locals("user", user)
		return c.Next()
	}
}
```

> Here we are getting Authorization value from request header with `authHeader := c.Get("Authorization")`

```go
	tokenString := authHeader[len("Bearer "):]

	blacklistedToken := usermodels.BlacklistedTokens{Token: tokenString}
	if blacklistedToken.IsTokenBlacklisted() {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token or token expired",
		})
	}
```

> Here we are checking token if it has already been used and logged out which we are metioning.

```go
	token, err := jwt.ParseWithClaims(tokenString, &userdtos.TokenClaimPayload{},
		func(token *jwt.Token) (interface{}, error) {
			// Get the signing key from your authentication server or config file
			signingKey := []byte(configs.ENVs.JwtSecretKey)
			return signingKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid JWT signature"})
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid JWT token"})
	}

	// Check if token is valid
	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid JWT token"})
	}
```

> Here we are chekcking token signature
>
> And whether token is valid or not.

Finally, we are populating token payload to a user and passing it to fiber context so that we can access user from our api's and moving to our next middleware.

```go
	var user usermodels.Users
	user.ID = claims.ID
	user.Username = claims.Username
	user.IsSuperuser = claims.IsSuperuser
	user.Email = claims.Email
	c.Locals("user", user)
	return c.Next()
```

**Now we can use it to our route as:**

```go
func UsersRouts(app *fiber.App) {
	router := app.Group("users", middlewares.JwtMiddleware())

	router.Post("/", controllers.AddUserHandler)
	router.Get("/", controllers.GetUsersListHandler)
	router.Get("/:id", controllers.GetUserDetailHandler)
	router.Put("/:id", controllers.UpdateUserHandler)
	router.Delete("/:id", controllers.DeleteUserHandler)
	...
}
```

> Here all the routes that are registerd with `router` are using the `JwtMiddleware()` and they will require authentication token in the header.

**_N.B. Let's remove the validation code from logout function since it is being done by our JwtMiddleware now._**

```go
	if ok := blacklistedToken.IsTokenBlacklisted(); ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token or token expired",
		})
	}
```

**Congratulations !!**

We have secured our api's from un authenticated users.

![Secured API Route](https://user-images.githubusercontent.com/34704464/235442646-bec55b8b-70f4-4bfa-b533-f7073aa901f4.png)

**Our Current project stracture should be as**

```md
├── LICENSE
├── README.md
├── configs
│ └── envVars.go
├── database
│ └── dbSetup.go
├── example.env
├── go.mod
├── go.sum
├── main.go
├── middlewares
│ └── jwtMiddlewars.go
├── migrations
│ └── migrations.go
├── server
│ └── server.go
├── users
│ ├── controllers
│ │ ├── authControllers.go
│ │ └── userControllers.go
│ ├── dtos
│ │ ├── authDtos.go
│ │ └── userDtos.go
│ ├── helpers
│ │ └── jwtTokeHelpers.go
│ ├── models
│ │ ├── auth.go
│ │ └── users.go
│ └── routes.go
└── utils
  ├── pagination.go
  └── validateStructs.go
```

Code for this far could be found here:  
[https://github.com/Fahad-Md-Kamal/Fiber-Blogs/tree/part-4/middleware/jwt-auth](https://github.com/Fahad-Md-Kamal/Fiber-Blogs/tree/part-4/middleware/jwt-auth)
