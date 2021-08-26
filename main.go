package main

import (
	"fmt"
	"math"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"strconv"
)

func main (){
	r := chi.NewRouter()
  r.Use(middleware.Logger)
  r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("ingrese un número entre 1 y un millón"))
  })
  r.Route("/{num}", func(r chi.Router) {
	  r.Get("/", func(w http.ResponseWriter, r *http.Request){
	  numero := chi.URLParam(r, "num")
	  
	  fmt.Println(numero)
		strN, _ := strconv.Atoi(numero)
		num := validar(strN)
		if num == 0{
			http.Error(w, http.StatusText(406), 406)
			return
		}
	  w.Write([]byte(fmt.Sprintf("Respuesta: %d\n", num)))

  	})
  })
  

  http.ListenAndServe(":3000", r)
}

//itera el numero hasta que encuentrar el menor numero primo y palindromo
func validar(num int) (int){
	sum:= num
	if num >= 1 && num <=1000000{
		for  {
			if esPrimo(sum){
				if esPalindromo(sum){
					return sum
				}
			}
			sum++
		}
			
	}
	return 0 
}

//busca el numero primo
func esPrimo(num int) bool{
	
	if (num == 2){
		return true
	}
	if(math.Mod(float64(num), 2) == 0){
		return false;
	}

	for i:=3; i<num; i+=2{
		if math.Mod(float64(num), float64(i)) == 0{
			return false
		}
	}
	return true
}

//determina si el nro es primo
func esPalindromo (num int) bool{
	num2 := num
	reverso:=0
	for {
		if num > 0 {
			reverso = reverso*10 + num%10
			num = num /10
		} else{
			break
		}
		
	}
	
	if num2 == reverso { return true }; return false
	
}
