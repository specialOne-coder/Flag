package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

func FirstApi() string {

	fmt.Println("*************************** FirstApi ***************************")
	
	var secret string
	for i := 3000; i <= 4001; i++ {
		address := fmt.Sprintf("34.77.36.161:%d", i)

		conn, err := net.Dial("tcp", address)
		if err != nil {
			if strings.Contains(err.Error(), "refused") {
				continue
			}
			fmt.Println("Erreur de connexion :", err)
			break
		}
		defer conn.Close()
		fmt.Println("Connexion établie avec succès à l'adresse :", address)
		resp, err := http.Get(fmt.Sprintf("http://%s", address))
		if err != nil {
			fmt.Println("Others thing in this port", err)
			break
		}
		defer resp.Body.Close()
		// Récupération de la réponse de la requête GET
		body, err := ioutil.ReadAll(resp.Body)
		secret = string(body)
		fmt.Println(secret)
	}
	return secret
}

func SecondApi() string {
	
	secret := FirstApi()
	fmt.Println("")
	fmt.Println("*************************** SecondApi ***************************")
	key := strings.TrimPrefix(secret, "The secret key is: ")
	url := Api + ":3941?secretKey=" + key
	resp, err := http.Post(url, "", nil)
	if err != nil {
		fmt.Println("Erreur lors de la requête POST :", err)
	}
	defer resp.Body.Close()

	// Lecture du corps de la réponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du corps de la réponse :", err)
	}
	// Affichage du corps décodé
	fmt.Println("File recupéré =>", string(body))
	return string(body)
}

func ThirdApi() string {
	SecondApi()
	fmt.Println("")
	fmt.Println("*************************** ThirdApi ***************************")
	address :=  Api + ":3610"
	url := address+"?finalKey=8116fdd3f12b6d7c4b136cbdaa3360a57eb4eb676ae63294450ee1f4f34b36f3"

	resp, err := http.Post(url,"",nil)
	if err != nil {
		fmt.Println("Erreur lors de la requête GET :", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("ThirdApi :", string(body))
	return string(body)

}

