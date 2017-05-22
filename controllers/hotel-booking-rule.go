package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"GoApi/common"
	"GoApi/entities"

	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func HotelBookingRuleIndex(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	session, err := mgo.Dial(common.MongoAddress)
	if err != nil {
		fmt.Fprint(w, err)
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("hotels")

	result := []entities.Hotel{}
	err = c.Find(nil).All(&result)
	if err != nil {
		log.Fatal(err)
		fmt.Fprint(w, err)
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		fmt.Fprint(w, err)
		panic(err)
	}

	fmt.Fprint(w)
}
func HotelBookingRuleFind(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	var ID int
	var err error
	if ID, err = strconv.Atoi(vars["id"]); err != nil {
		fmt.Fprint(w, err)
		panic(err)
	}

	session, err := mgo.Dial(common.MongoAddress)
	if err != nil {
		fmt.Fprint(w, err)
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("hotels")

	result := entities.Hotel{}
	err = c.Find(bson.M{"id": ID}).One(&result)
	if err != nil {
		//log.Fatal(err)
		fmt.Fprint(w, err)
		return
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		fmt.Fprint(w, err)
		panic(err)
	}

	fmt.Fprint(w)
}
func HotelBookingRuleAdd(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var hotel entities.Hotel
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		fmt.Fprint(w, err)
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		fmt.Fprint(w, err)
		panic(err)
	}
	bodyStr := string(body)
	fmt.Print(bodyStr)
	if err := json.Unmarshal(body, &hotel); err != nil {
		w.WriteHeader(422) // unprocessable entity

		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Fprint(w, err)
			panic(err)
		}
	}

	session, err := mgo.Dial(common.MongoAddress)
	if err != nil {
		fmt.Fprint(w, err)
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("hotels")
	err = c.Insert(hotel)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, "{\"success\":\"true\"}")
}
