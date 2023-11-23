package web

import (
	"net/http"

	"lamoda_task/internal/store"

	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
)

const (
	PostCreateStoreURL     = "/create-store"
	PostCreateProductURL   = "/create-product"
	DeleteProductURL       = "/delete-product/"
	PostReserveProductsURL = "/reserve-products"
	PostReleaseProductsURL = "/release-products"
	GetRemainProductsURL   = "/remain-products"
)

type postCreateStoreRequest struct {
	Payload *struct {
		Name        store.Name      `json:"name"`
		IsAvailable store.Available `json:"isAvailable"`
	} `in:"body=json"`
}

type postCreateStoreResponse struct {
	Id store.Id `json:"id"`
}

type postCreateProductRequest struct {
	Payload *struct {
		Name     store.Name     `json:"name"`
		Size     store.Size     `json:"size"`
		Code     store.Code     `json:"code"`
		Quantity store.Quantity `json:"quantity"`
		StoreID  store.Id       `json:"store_id"`
	} `in:"body=json"`
}

type postCreateProductResponse struct {
	Id store.Id `json:"id"`
}

type deleteProductRequest struct {
	Id store.Id `in:"path=id"`
}

type messageResponse struct {
	Message string `json:"message"`
}

type postProductCodesRequest struct {
	Payload *struct {
		Codes []store.Code `json:"codes"`
	} `in:"body=json"`
}

type getRemainProductsRequest struct {
	Id store.Id `in:"query=id;required"`
}

type getRemainProductsResponse struct {
	Products []store.Product `json:"products"`
}

func InitRoutes(httpRouter *chi.Mux, sm *store.StoreManagement) {
	httpin.UseGochiURLParam("path", chi.URLParam)

	httpRouter.Post(PostCreateStoreURL, func(w http.ResponseWriter, r *http.Request) {
		postCreateStore(w, r, sm.CreateStore)
	})

	httpRouter.Post(PostCreateProductURL, func(w http.ResponseWriter, r *http.Request) {
		postCreateProduct(w, r, sm.CreateProduct)
	})

	httpRouter.Delete(DeleteProductURL+"{id}/", func(w http.ResponseWriter, r *http.Request) {
		deleteProduct(w, r, sm.DeleteProduct)
	})

	httpRouter.Post(PostReserveProductsURL, func(w http.ResponseWriter, r *http.Request) {
		postReserveProducts(w, r, sm.ReserveProducts)
	})

	httpRouter.Post(PostReleaseProductsURL, func(w http.ResponseWriter, r *http.Request) {
		postReleaseProducts(w, r, sm.ReleaseProducts)
	})

	httpRouter.Get(GetRemainProductsURL, func(w http.ResponseWriter, r *http.Request) {
		getRemainProducts(w, r, sm.GetRemainingProducts)
	})
}
