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

/// Seccion LOGIN

func login(c echo.Context) error {

	jt := new(jwtCustomClaims)
	if err := c.Bind(jt); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Verificacion del usuario
	if jt.Nombre != "Salvador" || jt.Contraseña != "PIA" {
		return echo.ErrUnauthorized
	}

	// Costumizacion de cliams
	claims := &jwtCustomClaims{
		"Salvador Castro",
		"PIA",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Crear token con claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Genera un token y lo envia
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
	//Crea un cliente y recibe el json
	cl := new(clientes)
	if err := c.Bind(cl); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	/// Conexion a base de datos
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}
	// Crea al cliente en la base de datos
	db.Create(&cl)

	// returna el los datos recibidos
	return c.JSON(http.StatusOK, cl)
}

func getCliente(c echo.Context) error {
	//Traemos la estructura
	var cl []clientes
	//Conexion de base de datos
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}
	// Encuentra los clientes registrados en la base de datos
	db.Find(&cl)
	// Regresa el valor encontrado
	return c.JSON(http.StatusOK, cl)
}

func updateCliente(c echo.Context) error {
	// recibe el id del url
	id, _ := strconv.Atoi(c.Param("id"))
	//Conexion de base de datos
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLO CONEXION A BDD")
	}
	//Crea un cliente y busca con el id al cliente con la base de datos
	cl := new(clientes)
	db.First(&cl, id)
	if errs := c.Bind(cl); errs != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//Guarda los datos modificados del cliente
	db.Where("idclientes = ? ", id).Save(&cl)
	//Regresa el valor
	return c.JSON(http.StatusOK, cl)

}

func deleteCliente(c echo.Context) error {
	// recibe el id del url
	id, _ := strconv.Atoi(c.Param("id"))
	//Conexion de base de datos
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLO CONEXION A BDD")
	}
	//Traemos la estructura
	var cl clientes
	var r reservacion
	// buscamos los clientes y sus reservaciones para eliminarlos
	db.Where("IDcliente = ?", id).Delete(&r)
	db.Where("idclientes = ? ", id).Delete(&cl)
	return c.NoContent(http.StatusNoContent)
}

//////// SECCION VIAJES

func CrearViaje(c echo.Context) error {
	//Crea un viaje y recibe el json
	v := new(viaje)
	if err := c.Bind(v); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//Conexion de base de datos
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}
	//Crea el viaje en la base de datos
	db.Create(&v)

	//regresamos los datos dados
	return c.JSON(http.StatusOK, v)
}

func getViaje(c echo.Context) error {
	//traemos la estructura
	var v []viaje
	//conexion de base de datos
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}
	//encontramos los viajes
	db.Find(&v)
	//regresamos los datos resultantes
	return c.JSON(http.StatusOK, v)
}

func updateViaje(c echo.Context) error {
	//recibimos el id del url
	id, _ := strconv.Atoi(c.Param("id"))
	//conexion a BD
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLO CONEXION A BDD")
	}
	//Creamos un nuevo viaje y buscamos por el id en la BD
	v := new(viaje)
	db.First(&v, id)
	if errs := c.Bind(v); errs != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//Salvamos los datos cambiados en la BD
	db.Where("idviaje = ? ", id).Save(&v)
	return c.JSON(http.StatusOK, v)

}

func deleteViaje(c echo.Context) error {
	//recibimos el id del url
	id, _ := strconv.Atoi(c.Param("id"))
	//conexion BD
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLO CONEXION A BDD")
	}
	//Traemos las estructuras
	var v viaje
	var r reservacion
	// buscamos los viajes y sus reservaciones para eliminarlos
	db.Where("IDviaje = ?", id).Delete(&r)
	db.Where("idviaje = ? ", id).Delete(&v)
	return c.NoContent(http.StatusNoContent)
}

/////// SECCION RESERVACION

func CrearReservacion(c echo.Context) error {
	//Crea un cliente y recibe el json
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
	r.POST("/Viaje", CrearViaje)
	r.POST("/Reservacion", CrearReservacion)

	// Ver
	r.GET("/Cliente", getCliente)
	r.GET("/Viaje", getViaje)
	r.GET("/Reservacion/:id", getReservacionID)

	//Actualizar
	r.PUT("/Cliente/:id", updateCliente)
	r.PUT("/Viaje/:id", updateViaje)

	// Borrar
	r.DELETE("Cliente/:id", deleteCliente)
	r.DELETE("/Viaje/:id", deleteViaje)
	r.DELETE("/Reservacion/:id", deleteReservacion)

	e.Logger.Fatal(e.Start("localhost:1323"))
}
