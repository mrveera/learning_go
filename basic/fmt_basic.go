package main

import "fmt"

type Programmer struct {
  Name string
  Age uint
  Fav_lang string
  Sleeping_hrs uint
}

func (p Programmer) String() string {
  return fmt.Sprintf("\n I am %s , %d years old , As a Programmer I like %s languages.Everyday I sleep for %d hrs",p.Name,p.Age,p.Fav_lang,p.Sleeping_hrs)
}

func main() {
  //printing variables
  var pens=2
  fmt.Printf("I have %d pens", pens)

  //playing with String()
  veera := Programmer{
    Name:"veera",
    Age:18,
    Fav_lang:"all",
    Sleeping_hrs:3,
  }

  fmt.Println(veera)
  fmt.Printf("using printf %+v",veera)
}
