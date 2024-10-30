package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	webhook "example/user/webhook_proxy/webhook"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type ConsoleReceiver struct {
	Key     string `json:"key"`
	Version string `json:"version"`
}

func sendConsoleWebhook(user string, key string, version string) {
	if key == "...-...-...-..." {
		return
	}

	url := ""

	title := "User opened Devtools"
	productImage := "https://i.postimg.cc/vB3MDK2s/t-pfp.png"
	date := time.Now().Format(time.RFC3339)
	hook := webhook.Create(productImage, title, date, version)

	hook.AddField("User", user, false)
	hook.AddField("Key", key, false)

	for {
		webhookReq, err := hook.Send(url)
		if err != nil {
			fmt.Println(err.Error())
		}

		if webhookReq.StatusCode == 204 { //204 status is successful webhook post
			fmt.Println("Webhook sent")
			break
		} else {
			fmt.Println("Webhook failed")
			fmt.Println(webhookReq.StatusCode)
			retryHeader := webhookReq.Header.Get("Retry-After")
			fmt.Println(retryHeader)
			retry, _ := strconv.Atoi(retryHeader)
			time.Sleep(time.Second * time.Duration(retry))
		}
	}
}

type Hyper struct {
	Status string `json:"status"`
	Key    string `json:"key"`
	User   User   `json:"user"`
}

type User struct {
	Discord Discord `json:"discord"`
}

type Discord struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
}

func validate(key string) (Hyper, error) {
	bearer := "..."
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.hyper.co/v6/licenses/"+key, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+bearer)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var data Hyper
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err.Error())
	}

	return data, err
}

func console(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var receiver ConsoleReceiver
	_ = json.NewDecoder(r.Body).Decode(&receiver)
	license, err := validate(receiver.Key)
	if license.Status != "active" || err != nil {
		w.WriteHeader(http.StatusNotFound)
		resp := make(map[string]string)
		resp["message"] = "License Not Found"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}
	json.NewEncoder(w).Encode(receiver)
	sendConsoleWebhook(license.User.Discord.Username+"#"+license.User.Discord.Discriminator, license.Key, receiver.Version)
}

type ProductReceiver struct {
	Product      string `json:"product"`
	Store        string `json:"store"`
	Size         string `json:"size"`
	ProductURL   string `json:"product_url"`
	ProductImage string `json:"product_image"`
	Pid          string `json:"pid"`
	Date         string `json:"date"`
	Mode         string `json:"mode"`
	User         string `json:"user"`
	Key          string `json:"key"`
	Version      string `json:"version"`
}

func sendPublicWebhook(product string, size string, store string, productURL string, productImage string, pid string, date string, mode string, version string) {
	url := "https://discord.com/api/webhooks/..."

	title := "A storm has come! :cloud_lightning:"
	hook := webhook.Create(productImage, title, date, version)
	hook.AddField("Store", "||"+store+"||", false)
	hook.AddField("Product", product, true)
	hook.AddField("Size", size, true)
	hook.AddField("PID", "||"+pid+"||", true)
	hook.AddField("Mode", "||"+mode+"||", true)
	hook.AddField("Useful Links", productURL, true)

	for {
		webhookReq, err := hook.Send(url)
		if err != nil {
			fmt.Println(err.Error())
		}

		if webhookReq.StatusCode == 204 { //204 status is successful webhook post
			fmt.Println("Webhook sent")
			break
		} else {
			fmt.Println("Webhook failed")
			fmt.Println(webhookReq.StatusCode)
			retryHeader := webhookReq.Header.Get("Retry-After")
			fmt.Println(retryHeader)
			retry, _ := strconv.Atoi(retryHeader)
			time.Sleep(time.Second * time.Duration(retry))
		}
	}
}

func sendSecretWebhook(product string, size string, store string, productURL string, productImage string, pid string, date string, mode string, user string, key string, version string) {
	url := "https://discord.com/api/webhooks/..."

	title := "A storm has come! :cloud_lightning:"
	hook := webhook.Create(productImage, title, date, version)
	hook.AddField("Store", "||"+store+"||", false)
	hook.AddField("Product", product, true)
	hook.AddField("Size", size, true)
	hook.AddField("PID", "||"+pid+"||", true)
	hook.AddField("Mode", "||"+mode+"||", true)
	hook.AddField("User", "||"+user+"||", true)
	hook.AddField("Key", "||"+key+"||", true)
	hook.AddField("Useful Links", productURL, true)

	for {
		webhookReq, err := hook.Send(url)
		if err != nil {
			fmt.Println(err.Error())
		}

		if webhookReq.StatusCode == 204 { //204 status is successful webhook post
			fmt.Println("Webhook sent")
			break
		} else {
			fmt.Println("Webhook failed")
			fmt.Println(webhookReq.StatusCode)
			retryHeader := webhookReq.Header.Get("Retry-After")
			fmt.Println(retryHeader)
			retry, _ := strconv.Atoi(retryHeader)
			time.Sleep(time.Second * time.Duration(retry))
		}
	}
}

func product(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var receiver ProductReceiver
	_ = json.NewDecoder(r.Body).Decode(&receiver)
	license, err := validate(receiver.Key)
	if license.Status != "active" || err != nil {
		w.WriteHeader(http.StatusNotFound)
		resp := make(map[string]string)
		resp["message"] = "License Not Found"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}
	json.NewEncoder(w).Encode(receiver)
	sendPublicWebhook(receiver.Product, receiver.Size, receiver.Store, receiver.ProductURL, receiver.ProductImage, receiver.Pid, receiver.Date, receiver.Mode, receiver.Version)
	sendSecretWebhook(receiver.Product, receiver.Size, receiver.Store, receiver.ProductURL, receiver.ProductImage, receiver.Pid, receiver.Date, receiver.Mode, license.User.Discord.Username+"#"+license.User.Discord.Discriminator, receiver.Key, receiver.Version)
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// handle cors
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},    // All origins
		AllowedMethods: []string{"POST"}, // Allowing only get, just an example
	})
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/console", console).Methods("POST")
	myRouter.HandleFunc("/product", product).Methods("POST")
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":8000", c.Handler(myRouter)))
}

func main() {
	fmt.Println("Volt API - Online")

	handleRequests()
}
