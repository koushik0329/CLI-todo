package main

func main() {
	todo := Todos{}

	storage := Newstorage[Todos]("todoslist.json")
	storage.Load(&todo)

	cmdf := Newcmd()
	cmdf.Execute(&todo)

	storage.Save(todo)
}
