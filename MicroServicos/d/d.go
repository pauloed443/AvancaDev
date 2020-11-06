package main

import (
	"fmt"
	"net/http"
)

type Coupon struct {
	Code string
}

type Coupons struct {
	Coupon []Coupon
}

var coupons Coupons

func main() {
	coupon := Coupon{
		Code: "abc",
	}

	coupons.Coupon = append(coupons.Coupon, coupon)

	http.HandleFunc("/", validate)
	http.ListenAndServe(":9093", nil)
}

func validate(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")

	for _, item := range coupons.Coupon {
		if coupon == item.Code {
			fmt.Fprintf(w, "Valid")
		}
	}
	fmt.Fprintf(w, "Invalid")
}
