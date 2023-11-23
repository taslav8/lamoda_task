package web

import (
	"log"
	"net/http"

	"lamoda_task/internal/store"

	"github.com/ggicci/httpin"
	"github.com/unrolled/render"
)

var rnr = render.New(render.Options{
	StreamingJSON: true,
})

func postCreateStore(w http.ResponseWriter, r *http.Request, createStore func(name store.Name, isAvailable store.Available) (store.Id, error)) {
	input := r.Context().Value(httpin.Input).(*postCreateStoreRequest)

	id, err := createStore(input.Payload.Name, input.Payload.IsAvailable)
	if err != nil {
		log.Printf("Cannot be store created: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = rnr.JSON(w, http.StatusOK, postCreateStoreResponse{Id: id}); err != nil {
		log.Printf("Cannot make HTTP response back: %v\n", err)
	}
}

func postCreateProduct(w http.ResponseWriter, r *http.Request, createProduct func(name store.Name, size store.Size, code store.Code, quantity store.Quantity, storeID store.Id) (store.Id, error)) {
	input := r.Context().Value(httpin.Input).(*postCreateProductRequest)

	id, err := createProduct(input.Payload.Name, input.Payload.Size, input.Payload.Code, input.Payload.Quantity, input.Payload.StoreID)
	if err != nil {
		log.Printf("Cannot be product created: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = rnr.JSON(w, http.StatusOK, postCreateProductResponse{Id: id}); err != nil {
		log.Printf("Cannot make HTTP response back: %v\n", err)
	}
}

func deleteProduct(w http.ResponseWriter, r *http.Request, deleteProduct func(id store.Id) error) {
	input := r.Context().Value(httpin.Input).(*deleteProductRequest)

	if err := deleteProduct(input.Id); err != nil {
		log.Printf("Cannot be product deleted: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := rnr.JSON(w, http.StatusOK, messageResponse{Message: "Deleted successfully"}); err != nil {
		log.Printf("Cannot make HTTP response back: %v\n", err)
	}
}

func postReserveProducts(w http.ResponseWriter, r *http.Request, ReserveProducts func(productCodes []store.Code) error) {
	input := r.Context().Value(httpin.Input).(*postProductCodesRequest)

	if err := ReserveProducts(input.Payload.Codes); err != nil {
		log.Printf("Cannot be product reserved: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := rnr.JSON(w, http.StatusOK, messageResponse{Message: "Reserved successfully"}); err != nil {
		log.Printf("Cannot make HTTP response back: %v\n", err)
	}
}

func postReleaseProducts(w http.ResponseWriter, r *http.Request, releaseProducts func(productCodes []store.Code) error) {
	input := r.Context().Value(httpin.Input).(*postProductCodesRequest)

	if err := releaseProducts(input.Payload.Codes); err != nil {
		log.Printf("Cannot be product released: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := rnr.JSON(w, http.StatusOK, messageResponse{Message: "Released successfully"}); err != nil {
		log.Printf("Cannot make HTTP response back: %v\n", err)
	}
}

func getRemainProducts(w http.ResponseWriter, r *http.Request, getRemainingProducts func(storeID store.Id) ([]store.Product, error)) {
	input := r.Context().Value(httpin.Input).(*getRemainProductsRequest)

	products, err := getRemainingProducts(input.Id)
	if err != nil {
		log.Printf("Cannot be product remained: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = rnr.JSON(w, http.StatusOK, getRemainProductsResponse{Products: products}); err != nil {
		log.Printf("Cannot make HTTP response back: %v\n", err)
	}
}
