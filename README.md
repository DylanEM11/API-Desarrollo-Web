# API-Desarrollo-Web
## Nombres :
## Cesar Salvador Castro Mata      1794668
## Dylan Eduardo Escobedo Martínez 1866199
## Luis Jesus Garcia Almanza       1869209

## Introducción:
Una pequeña agencia de viajes operaba y organizaba todo su repertorio de viajes, lista de clientes, hoteles asociados, y todo lo relacionado a ellos mediante su pagina de Facebook, y creando grupos de conversaciones por Messenger, recibía sus pagos mediante transferencias, lo que significaba una causa de desconfianza de parte de los clientes y se traducía a su vez en perdidas de ganancias, por lo que se optó por la utilización de una API que redujera y automatizara todo el trabajo que realizaban ellos mismos, organizando mucho más fácilmente así, su lista de clientes, empleados, viajes, y hoteles asociados, así como la utilización de distintos métodos de pago mas seguros.

## Propuesta técnica:
### •	Entidades: 
-Cliente
-Viaje
-Reservacion

### •	Atributos:

-Cliente:
  IDclientes
  Nombre     
  Direccion  
  Correo     
  Telefono  
  
-Viaje:
  IDviaje     
  Origen        
  Destino       
  Aerolinea     
  Hotel         
  Num_Personas  
  Fecha_ida    
  Fecha_regreso
  
-Reservacion:
  IDReservacion
  Personas      
  IDcliente    
  IDviaje       
 
## •	Funcionalidades:
Nuestra API contiene una función que no viene dentro de la base de datos, que sería la función de cobro, buscando que sea un método de pago en línea, seguro y eficaz.

## •	Base de datos:
En nuestra API, decidimos utilizar el MySQL como base de datos, MySQL es un sistema de administración de bases de datos (Database Management System, DBMS) para bases de datos relacionales. Así, MySQL no es más que una aplicación que permite gestionar archivos llamados de bases de datos.

![LOL](https://user-images.githubusercontent.com/89227433/141645498-eb3fbcf3-7b62-4ad2-9019-3063155f72a5.PNG)



