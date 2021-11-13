package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

type jwtCustomClaims struct {
	Nombre     string `json:"Nombre"`
	Contraseña string `json:"Contraseña"`
	jwt.StandardClaims
}

func login(c echo.Context) error {

	jt := new(jwtCustomClaims)
	if err := c.Bind(jt); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Throws unauthorized error
	if jt.Nombre != "Salvador" || jt.Contraseña != "PIA" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		"Salvador Castro",
		"PIA",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("gobhgb76/&Jngnghn"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
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
	var r reservacion
	db.Where("IDcliente = ?", id).Delete(&r)
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
	var r reservacion
	db.Where("IDviaje = ?", id).Delete(&r)
	db.Where("idviaje = ? ", id).Delete(&v)
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
	var cl clientes
	clien := db.Where("idclientes = ?", r.IDcliente).Find(&cl)
	if clien.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, "NO EXISTE CLIENTE")
	}
	var v viaje
	via := db.Where("idviaje = ? AND Num_Personas >= ?", r.IDviaje, r.Personas).Find(&v)
	if via.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, "NO EXISTE VIAJE")
	} else {
		db.Model(&v).Where("idviaje = ?", r.IDviaje).Update("Num_Personas", v.Num_Personas-r.Personas)
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
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200", "http://localhost:3100"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	r := e.Group("/restricted")

	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("gobhgb76/&Jngnghn"),
	}

	r.Use(middleware.JWTWithConfig(config))
	// Crear
	e.POST("/login", login)
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

	// Borrar
	e.DELETE("Cliente/:id", deleteCliente)
	e.DELETE("/Viaje/:id", deleteViaje)
	e.DELETE("/Reservacion/:id", deleteReservacion)

	e.Logger.Fatal(e.Start("localhost:1323"))
}
