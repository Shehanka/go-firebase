package main

import (
	"firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	sa := option.WithCredintialsFile("./ServiceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)

	client, err := app.Firestore(context)
	if err != nil {
		log.Fatalln(err)
	}
	quote := getQuote()
	log.Print(quote)
	result, err := client.Collection("sampleData").Doc("inspiration").Set(context.Background(), quote)
	if err != nil{
		log.Fatalln(err)
	}
	log.Print(result)
	defer client.Close()
}

func getQuote() *Quote {
	/* returns Quote
	{

	}
	*/
}