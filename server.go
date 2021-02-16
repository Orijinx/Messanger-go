package main
import "net"
import "fmt"
import "bufio"
// import "strings" // требуется только ниже для обработки примера
func main() {
  fmt.Println("Launching server...")

  // Устанавливаем прослушивание порта
  ln, err := net.Listen("tcp", ":8081")

  if err !=nil{
    fmt.Println(err)
    return
  }
 

  // Открываем порт
  conn, err := ln.Accept()
  fmt.Println("Connected to device")
  if err !=nil{
    fmt.Println(err)
    return
  }
  // Запускаем цикл
  for {
    
    
    // Будем прослушивать все сообщения разделенные \n
    message, _ := bufio.NewReader(conn).ReadString('\n')
    // Распечатываем полученое сообщение
    fmt.Print("Message Received:", string(message))
    // // Процесс выборки для полученной строки
    // newmessage := strings.ToUpper(message)
    // // Отправить новую строку обратно клиенту
    // conn.Write([]byte(newmessage + "\n"))
  }
}