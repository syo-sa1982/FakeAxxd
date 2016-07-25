package resources
import (
	"net/http"
	"github.com/syo-sa1982/GoNTAkun/model"
	"github.com/labstack/echo"
	"log"
)





func (resource *Resource) JoinEvent() echo.HandlerFunc {

	return func(c echo.Context) error {

		var (
			db = resource.DB
		)
		db = resource.SetDBConnection()
		defer db.Close()

		u := new(JoinRequest)

		if err := c.Bind(u); err != nil {
			return err
		}

		member := model.Member{
			EventID:u.EventId,
			Name:u.Name,
		}
		log.Println(member)

		db.Create(&member)

		responseApi := map[string]int{"id": member.ID}

		api := APIFormat{"success", 1, 0, responseApi}
		return c.JSON(http.StatusOK, &api)
	}
}

func (resource *Resource) CancelEvent() echo.HandlerFunc {

	return func(c echo.Context) error {

		var (
			db = resource.DB
			member = model.Member{}
		)
		db = resource.SetDBConnection()
		defer db.Close()

		responseApi := map[string]string{"id": c.Param("join_id")}

		db.Model(member).Where("id = ?",c.Param("join_id")).Delete(&member)
		log.Println(responseApi)

		api := APIFormat{"success", 1, 0, responseApi}
		return c.JSON(http.StatusOK, &api)
	}
}