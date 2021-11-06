package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type clientes struct {
	IDclientes int    `json:"ID"`
	Nombre     string `json:"Nombre"`
	Direccion  string `json:"Direccion"`
	Correo     string `json:"Correo"`
	Telefono   string `json:"Telefono"`
}

type reservacion struct {
	IDReservacion int `json:"ID"`
	Personas      int `json:"Personas"`
	IDcliente     int `json:"IDCliente"`
	IDviaje       int `json:"IDViaje"`
}

type viaje struct {
	IDviaje       int    `json:"ID"`
	Origen        string `json:"Origen"`
	Destino       string `json:"Destino"`
	Aerolinea     string `json:"Aerolinea"`
	Hotel         string `json:"Hotel"`
	Num_Personas  int    `json:"Num_Personas"`
	Fecha_ida     string `json:"Fecha_ida"`
	Fecha_regreso string `json:"Fecha_regreso"`
}

////// SECCION CLIENTE

func CrearCliente(c echo.Context) error {
	cl := new(clientes)
	if err := c.Bind(cl); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}
	db.Create(&cl)

	return c.JSON(http.StatusOK, cl)
}

func getCliente(c echo.Context) error {
	var cl []clientes
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}
	db.Find(&cl)
	return c.JSON(http.StatusOK, cl)
}

func updateCliente(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLO CONEXION A BDD")
	}
	cl := new(clientes)
	db.First(&cl, id)
	if errs := c.Bind(cl); errs != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db.Where("idclientes = ? ", id).Save(&cl)
	return c.JSON(http.StatusOK, cl)

}

func deleteCliente(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLO CONEXION A BDD")
	}
	var cl clientes
	db.Where("idclientes = ? ", id).Delete(&cl)
	return c.NoContent(http.StatusNoContent)
}

//////// SECCION VIAJES

func CrearViaje(c echo.Context) error {
	v := new(viaje)
	if err := c.Bind(v); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}
	db.Create(&v)

	return c.JSON(http.StatusOK, v)
}

func getViaje(c echo.Context) error {
	var v []viaje
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}
	db.Find(&v)
	return c.JSON(http.StatusOK, v)
}

func updateViaje(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLO CONEXION A BDD")
	}
	v := new(viaje)
	db.First(&v, id)
	if errs := c.Bind(v); errs != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db.Where("idviaje = ? ", id).Save(&v)
	return c.JSON(http.StatusOK, v)

}

func deleteViaje(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLO CONEXION A BDD")
	}
	var v viaje
	db.Where("idviajes = ? ", id).Delete(&v)
	return c.NoContent(http.StatusNoContent)
}

/////// SECCION RESERVACION

func CrearReservacion(c echo.Context) error {
	r := new(reservacion)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}
	db.Create(&r)

	return c.JSON(http.StatusOK, r)
}

func getReservacionID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}
	var r reservacion
	db.First(&r, id)
	return c.JSON(http.StatusOK, r)
}

func updateReservacion(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLO CONEXION A BDD")
	}
	r := new(reservacion)
	db.First(&r, id)
	if errs := c.Bind(r); errs != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db.Where("id_reservacion = ? ", id).Save(&r)
	return c.JSON(http.StatusOK, r)

}

func deleteReservacion(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLO CONEXION A BDD")
	}
	var r reservacion
	db.Where("id_reservacion = ? ", id).Delete(&r)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	// Crear
	e.POST("/Cliente", CrearCliente)
	e.POST("/Viaje", CrearViaje)
	e.POST("/Reservacion", CrearReservacion)

	// Ver
	e.GET("/Cliente", getCliente)
	e.GET("/Viaje", getViaje)
	e.GET("/Reservacion/:id", getReservacionID)

	//Actualizar
	e.PUT("/Cliente/:id", updateCliente)
	e.PUT("/Viaje/:id", updateViaje)
	e.PUT("/Reservacion/:id", updateReservacion)

	// Borrar
	e.DELETE("Cliente/:id", deleteCliente)
	e.DELETE("/Viaje/:id", deleteViaje)
	e.DELETE("/Reservacion/:id", deleteReservacion)

	e.Logger.Fatal(e.Start("localhost:1323"))
}
