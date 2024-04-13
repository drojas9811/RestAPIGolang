package helpers

import (
	"fmt"
	"testing"
)

func TestPagination(t *testing.T) {

	testCases := []struct {
		Name string
	}{
		{
			Name: "test1",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			fmt.Println("This is the testing for the testcase:", tc.Name, ".")
			var testedStructure PaginationHelper
			testedStructure.Init([]string{"a", "b", "c", "d", "e", "f","g","h","i","j"}, 3)
			pages := testedStructure.Page_count()              //should return 2
			items := testedStructure.Item_count()              //should return 6
			itemsByPage := testedStructure.Page_item_count(0)  //should return 4
			itemsByPage1 := testedStructure.Page_item_count(1) //last page -should =
			itemsByPage2 := testedStructure.Page_item_count(2) //should ==-1 since the page
			pageOfItem := testedStructure.Page_index(5)        //should ==1 (Zero based index)
			fmt.Println("Results from fuctions are:",
				pages,
				items,
				itemsByPage,
				itemsByPage1,
				itemsByPage2,
				pageOfItem,
			)
		})
	}
}
