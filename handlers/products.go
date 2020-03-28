package handlers

import (
	"log"
	"net/http"

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

	// catch
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	//Marshall vs Encoder (no need to allocate memory when writing to io writer directly)
	//d, err := json.Marshal(lp)
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to do marshal", http.StatusInternalServerError)
	}

	//rw.Write(d)
}
