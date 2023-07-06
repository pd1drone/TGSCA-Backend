package testing

import (
	"fmt"
	"testing"
	"tgsca/database"
	"tgsca/rest"
)

func TestLoginStudent(t *testing.T) {
	tgscadb, err := rest.New()
	if err != nil {
		fmt.Println(err)
	}

	Success, userID, err := database.LoginStudent(tgscadb.TGSCAdb, "123", "e16b2ab8d12314bf4efbd6203906ea6c", "09/10/1997")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Success)
	fmt.Println(userID)

	Success1, userID1, err := database.LoginStudent(tgscadb.TGSCAdb, "123", "wrongpass", "09/10/1997")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Success1)
	fmt.Println(userID1)

}

func TestLoginAdmin(t *testing.T) {
	tgscadb, err := rest.New()
	if err != nil {
		fmt.Println(err)
	}

	Success, err := database.LoginAdmin(tgscadb.TGSCAdb, "admin", "0192023a7bbd73250516f069df18b500")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Success)

	Success1, err := database.LoginAdmin(tgscadb.TGSCAdb, "admin", "wrongpass")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Success1)
}

func TestGetApproveDetails(t *testing.T) {
	tgscadb, err := rest.New()
	if err != nil {
		fmt.Println(err)
	}

	Data, err := database.GetStudentEnrollmentStatus(tgscadb.TGSCAdb)
	if err != nil {
		fmt.Println(err)
	}

	for _, d := range Data {
		fmt.Println(d)
	}
}
