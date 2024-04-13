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
	ipp := items_per_page
	//Adding each complete slice to the array 
	for it := 0; it < int(len(colection)/ipp); it++ {
		subArray := colection[(it * ipp):(ipp + ipp*it)]
		r.selfColection = append(r.selfColection, subArray)
	}
	//adding the last array
	mod := len(colection) % ipp
	if mod > 0 {
		subArray := colection[(len(colection) - mod):]
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
