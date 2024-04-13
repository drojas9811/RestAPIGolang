package helpers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idStr)
	}
	return id, nil
}

type itemByValue []string
type PaginationHelper struct {
	selfColection []itemByValue
}

// the receiving collecction is like the
//following: (['a','b','c','d','e','f'], 4)
//In golang it means slide of elements

func (r *PaginationHelper) Init(colection []string, items_per_page int) {
	//[0,1,2,3,4,5,6,7,8,9] -> 3 => [[0,1,2],[3,4,5],[6,7,8],[9]]
	//while the array has more than 0 elements, iteration is going.
	//iteration in the original array every x items_per_page
	//size=10 and items_per_page=3 --> 3 and extra 1
	// index+(iteration)*(items_per_page)
	// (0,1,2,3,4,5,6,7,8,9)+(0,1,2)*3
	// [initial:end]=[0:2][3:5][6:8][9::]
	//[(iteration)*(items_per_page):((iteration+1)*(items_per_page))-1]
	//[(0)*(3):(1)*(3)]= [0:3]
	//[(1)*(3):(2)*(3)]= [3:6]
	//[(2)*(3):(3)*(3)]= [6:9]
	//((iteration+1)*(items_per_page))-1
	//((x+1)(y))-1
	//xy+y-1////////////
	//(xy+y)-1
	//(y(x+1))-1
	//[(it*ipp):(ipp+ipp*it-1)]///////
	ipp := items_per_page
	for it := 0; it < int(len(colection)/items_per_page); it++ {
		subArray := colection[(it * ipp):(ipp + ipp*it)] //0:2//3:5//6:8//9:11
		r.selfColection = append(r.selfColection, subArray)
	}
	// len(colection)%items_per_page = 1
	// if mod>0 => [len(colection)-mod::]
	//[10-1::]->[9::]
	//[11-2::]->[9::]
	//[12-3::]->[9::]
	mod:=len(colection)%items_per_page
	if mod >0{
		subArray:= colection[(len(colection)-mod):]
		r.selfColection = append(r.selfColection, subArray)
	}

}
func (r *PaginationHelper) Page_count() int {
	return 5
}
func (r *PaginationHelper) Item_count() int {
	return 0
}
func (r *PaginationHelper) Page_item_count(page_index int) int {
	return 0
}
func (r *PaginationHelper) Page_index(item_index int) int {
	return 0
}
