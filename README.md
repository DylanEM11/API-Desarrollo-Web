# API-Desarrollo-Web
## Nombres : Cesar Salvador Castro Mata  
## Dylan Eduardo Escobedo Martínez   
## Luis Jesus Garcia Almanza

## Introducción:
Una pequeña agencia de viajes operaba y organizaba todo su repertorio de viajes, lista de clientes, hoteles asociados, y todo lo relacionado a ellos mediante su pagina de Facebook, y creando grupos de conversaciones por Messenger, recibía sus pagos mediante transferencias, lo que significaba una causa de desconfianza de parte de los clientes y se traducía a su vez en perdidas de ganancias, por lo que se optó por la utilización de una API que redujera y automatizara todo el trabajo que realizaban ellos mismos, organizando mucho más fácilmente así, su lista de clientes, empleados, viajes, y hoteles asociados, así como la utilización de distintos métodos de pago mas seguros.

## Propuesta técnica:
### •	Entidades: 
-Tienda
-Empleado
-Admin
-Cliente
-Vuelo
-Hotel
-Cuarto

### •	Atributos:

-Tienda:
  IDTienda
  Telefono
  Direccion
  Correo

-Empleado:
 IDEmpleado
 NombreEmpleado
 Telefono
 Direccion
 Correo
 IDTienda

-Admin:
 IDAdmin
 IDEmpleado

-Cliente:
 IDCliente
 NombreCliente
 Telefono
 Direccion
 Correo
 IDVuelo
 IDHotel
 IDEmpleado

-Vuelo:
 IDVuelo
 NombreAvion
 LugarPartida
 LugarLlegada
 FechaSalida
 NumeroPasajeros

-Hotel:
 IDHotel
 NombreHotel
 CiudadHotel
 CapacidadHotel
 IDCuarto

-Cuarto:
 IDCuarto
 NumeroCuarto
 PersonasXCuarto
 
## •	Funcionalidades:
Nuestra API contiene una función que no viene dentro de la base de datos, que sería la función de cobro, buscando que sea un método de pago en línea, seguro y eficaz.

## •	Base de datos:
En nuestra API, decidimos utilizar el MySQL como base de datos, MySQL es un sistema de administración de bases de datos (Database Management System, DBMS) para bases de datos relacionales. Así, MySQL no es más que una aplicación que permite gestionar archivos llamados de bases de datos.

![Diagrama](https://user-images.githubusercontent.com/89227433/130306074-693c8a73-026e-4bab-a3c6-5267220d91e3.PNG)![WhatsApp Image 2021-08-20 at 8 58 14 PM](https://user-images.githubusercontent.com/89227433/130308029-30df696d-fd25-4e57-a415-402c2d17dfca.jpeg)


