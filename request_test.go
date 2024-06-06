package requests

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	// Get Query parameters
	query := r.URL.Query()
	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(j)
}

func PostFormHandler(w http.ResponseWriter, r *http.Request) {
	// Get Form data
	form := r.Form
	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(j)
}

func PostJSONHandler(w http.ResponseWriter, r *http.Request) {
	// Get JSON data
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(j)
}

func PutJSONHandler(w http.ResponseWriter, r *http.Request) {
	// Get JSON data
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(j)
}

func DeleteJSONHandler(w http.ResponseWriter, r *http.Request) {
	// Get JSON data
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(j)
}

func server(listenAddr string) error {
	// Create GET API receive Query parameters and return JSON response
	http.HandleFunc("/api/get", GetHandler)
	// Create POST API receive Form data and return JSON response
	http.HandleFunc("/api/post-form", PostFormHandler)
	// Create POST API receive JSON data and return JSON response
	http.HandleFunc("/api/post-json", PostJSONHandler)
	// Create PUT API receive JSON data and return JSON response
	http.HandleFunc("/api/put-json", PutJSONHandler)
	// Create DELETE API receive JSON data and return JSON response
	http.HandleFunc("/api/delete-json", DeleteJSONHandler)

	return http.ListenAndServe(listenAddr, nil)
}

func TestRequest(t *testing.T) {
	go server(":8080")
	// Wait for server to start
	time.Sleep(1 * time.Second)
	// Test GET API
	t.Run("Test GET API", func(t *testing.T) {
		resp, err := Requests("get", "http://localhost:8080/api/get")
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()
	})
	t.Run("Test GET Query API", func(t *testing.T) {
		query := Params{"name": "john", "age": "30"}
		resp, err := Requests("GET", "http://localhost:8080/api/get", query)
		if err != nil {
			t.Error(err)
		}
		var data map[string]interface{}
		if err := resp.JSON(&data); err != nil {
			t.Error(err)
		}
		t.Log(data)
	})
	// Test POST Form API
	t.Run("Test POST Form API", func(t *testing.T) {
		form := Form{"name": "john", "age": "30"}
		resp, err := Requests(POST, "http://localhost:8080/api/post-form", form)
		if err != nil {
			t.Error(err)
		}
		var data map[string]interface{}
		if err := resp.JSON(&data); err != nil {
			t.Error(err)
		}
		t.Log(data)
	})
	// Test POST JSON API
	t.Run("Test POST JSON API", func(t *testing.T) {
		data := struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{
			Name: "john",
			Age:  30,
		}
		resp, err := Requests("POST", "http://localhost:8080/api/post-json", data)
		if err != nil {
			t.Error(err)
		}
		var result map[string]interface{}
		if err := resp.JSON(&result); err != nil {
			t.Error(err)
		}
		t.Log(result)
	})
	// Test PUT JSON API
	t.Run("Test PUT JSON API", func(t *testing.T) {
		data := map[string]interface{}{"name": "john", "age": 30}
		resp, err := Requests("PUT", "http://localhost:8080/api/put-json", data)
		if err != nil {
			t.Error(err)
		}
		var result map[string]interface{}
		if err := resp.JSON(&result); err != nil {
			t.Error(err)
		}
		t.Log(result)
	})
	// Test DELETE JSON API
	t.Run("Test DELETE JSON API", func(t *testing.T) {
		data := map[string]interface{}{"name": "john", "age": 30}
		resp, err := Requests("DELETE", "http://localhost:8080/api/delete-json", data)
		if err != nil {
			t.Error(err)
		}
		var result map[string]interface{}
		if err := resp.JSON(&result); err != nil {
			t.Error(err)
		}
		t.Log(result)
	})
}
