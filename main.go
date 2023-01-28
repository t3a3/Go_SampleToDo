package main

import (
	"ToDo_app/app/controllers"
	"ToDo_app/app/models"
	"fmt"
)

func main() {
	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/

	fmt.Println(models.Db)

	controllers.StartMainServer()

	// user, _ := models.GetUserByEmail("test@example.com")
	// fmt.Println(user)

	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(session)

	// valid, _ := session.CheckSession()
	// fmt.Println(valid)
	//userの作成

	// u := &models.User{}
	// u.Name = "test3"
	// u.Email = "test3@exmaple.com"
	// u.PassWord = "testtest3"
	// fmt.Println(u)

	// u.CreatedUser()

	//userの取得
	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Name = "Test2"
	// u.Email = "Test2@example.com"
	// u.UpdateUser()

	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()

	// user, _ := models.GetUser(2)
	// user.CreatedToDo("FirstTodo")

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	//user, _ := models.GetUser(3)
	//user.CreatedToDo("ThirdTodo")

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// user2, _ := models.GetUser(2)
	// todos, _ := user2.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	//t, _ := models.GetTodo(3)
	// t.Content = "Update Todo"
	// t.UpdateTodo()
	// t.DeleteTodo()
}
