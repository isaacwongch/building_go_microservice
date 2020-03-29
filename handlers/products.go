package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/isaacwongch/building_go_microservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		p.l.Println("Handling PUT request")
		//need to handle yourself
		rp := regexp.MustCompile(`/([0-9]+)`)
		g := rp.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
		}

		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
		}

		idString := g[0][1]
		id, _ := strconv.Atoi(idString)

		p.l.Println("got id", id)

		p.updateProduct(id, rw, r)
		return
	}

	// catch
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, h *http.Request) {
	p.l.Println("Handling GET request")

	lp := data.GetProducts()
	//Marshall vs Encoder (no need to allocate memory when writing to io writer directly)
	//d, err := json.Marshal(lp)
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to do marshal", http.StatusInternalServerError)
	}

	//rw.Write(d)
}

func (p *Products) addProduct(rw http.ResponseWriter, h *http.Request) {
	p.l.Println("Handling POST request")

	prod := &data.Product{}

	// go hasn't read everything from the request, read that progressively, that's why a reader here
	err := prod.FromJSON(h.Body)

	if err != nil {
		http.Error(rw, "Unable to do marshal", http.StatusBadRequest)
	}

	p.l.Printf("PROD %#v", prod)

	data.AddProduct(prod)

}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, h *http.Request) {
	p.l.Println("Handling PUT request")

	prod := &data.Product{}

	// go hasn't read everything from the request, read that progressively, that's why a reader here
	err := prod.FromJSON(h.Body)

	if err != nil {
		http.Error(rw, "Unable to do marshal", http.StatusBadRequest)
	}

	p.l.Printf("PROD %#v", prod)

	err = data.UpdateProduct(id, prod)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
}
