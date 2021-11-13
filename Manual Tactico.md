#**MANUAL TACTICO DE LA AGENCIA DE VIAJES**

## **1) LOGIN**

### Primeramente para poder administrar nuestra agencia de viajes tenemos que logearnos para poder tener acceso a todas las funciones de nuestra agencia de viajes
###Entramos en la carpeta admin del postman y seleccionamos el metodo post de login y ingresamos el sig json con sus datos respectivos y le damos en send

### {
###    "Nombre": "Nombre Administrador",
###    "Contraseña": "Contraseña Administrador"
### }

###Si no llega ser un admin registrado se dezplegara el mensaje de "unauthorized" lo que quiere decir que no esta autorizado para ingresar y no podra administrar la agencia de viajes
###SI es un administrador autorizado le aparecera un token con sus datos y ya podra administrar la agencia 

## **2) CLIENTE**
###Este aparatado es para crear, obtener, editar y eliminar clientes

### **2)A.- CREAR CLIENTE**
###En este aparatado podra crear un cliente primero entrara a la carpeta del postman llamada cliente y seleccionara **Crear Cliente**
###Ya estando en crear cliente ingresara un json como en el ejemplo se muestra y para crearlo le dara en send

### {
###    "Nombre" : "Su nombre",
###    "Direccion" : "Su direccion",
###    "Correo" : "Su correo",
###    "Telefono" : "Su telefono"
### }

### **2)B.- OBTENER CLIENTES**
###En este apartado podra ver todos los clientes que se tienen registrados en el sistema
###Primero entramos en la carpeta de cliente y ya despues seleccionamos el apartado **Obtener Clientes** y presiona el boton de send y se desplegaran todos los clientes registrados

### **2)C.- EDITAR CLIENTE**
###En este aparatado podra editar un cliente ya creado primero entrara a la carpeta del postman llamada cliente y seleccionara **Editar Cliente**
###Ya estando en editar cliente pondra en el url el id del cliente que quiere editar y depues ingresara un json como en el ejemplo se muestra, para guardar los datos le da al send

### {
###    "Nombre" : "Su nombre",
###    "Direccion" : "Su direccion",
###    "Correo" : "Su correo",
###    "Telefono" : "Su telefono"
### } 

### **2)D.- ELIMINAR CLIENTE**
###En este apartado podra eliminar un cliente por su id primero entrara a la carpeta del postman llamada cliente y seleccionara **Eliminar Cliente**
###Ya estando en eliminar cliente pondra en el url el id del cliente que quiere eliminar, para eliminar ese cliente nomas oprima el boton de send y el cliente sera borrado
###Si el cliente ya habia realizado reservaciones tambien se borraran 

## **3) VIAJE**
###Este aparatado es para crear, obtener, editar y eliminar viajes

### **3)A.- CREAR VIAJE**
###En este aparatado podra crear un viaje primero entrara a la carpeta del postman llamada viaje y seleccionara **Crear Viaje**
###Ya estando en crear viaje ingresara un json como en el ejemplo se muestra y para crearlo le dara en send

### {
###    "Origen" : "Ciudad De Partida",
###    "Destino" : "Ciudad De LLegada",
###    "Aerolinea" : "Compañia del Vuelo",
###    "Hotel" : "Nombre de hotel reservado",
###    "NumPersonas" : cantidad de personas que pueden reservar este viaje, (ingresar un numero)
###    "Ida": "Fecha de Salida", (poner fecha completa)
###    "Regreso":"Fecha del regreso" (poner fecha completa)
### }

### **3)B.- OBTENER VIAJES**
###En este apartado podra ver todos los viajes que se tienen registrados en el sistema
###Primero entramos en la carpeta de viaje y ya despues seleccionamos el apartado **Obtener Viajes** y presiona el boton de send y se desplegaran todos los viajes registrados

### **3)C.- EDITAR VIAJE**
###En este aparatado podra editar un viaje ya creado primero entrara a la carpeta del postman llamada viaje y seleccionara **Editar Viaje**
###Ya estando en editar viaje pondra en el url el id del viaje que quiere editar y depues ingresara un json como en el ejemplo se muestra, para guardar los datos le da al send

### {
###    "Origen" : "Ciudad De Partida",
###    "Destino" : "Ciudad De LLegada",
###    "Aerolinea" : "Compañia del Vuelo",
###    "Hotel" : "Nombre de hotel reservado",
###    "NumPersonas" : cantidad de personas que pueden reservar este viaje, (ingresar un numero)
###    "Ida": "Fecha de Salida", (poner fecha completa)
###    "Regreso":"Fecha del regreso" (poner fecha completa)
### }

### **3)D.- ELIMINAR VIAJE**
###En este apartado podra eliminar un viaje por su id primero entrara a la carpeta del postman llamada viaje y seleccionara **Eliminar Viaje**
###Ya estando en eliminar viaje pondra en el url el id del viaje que quiere eliminar, para eliminar ese viaje nomas oprima el boton de send y el viaje sera borrado
###Si el viaje ya tenia reservaciones tambien se borraran 

## **4) RESERVACION**
###Este aparatado es para crear, obtener y eliminar las reservaciones

### **4)A.- CREAR RESERVACION**
###En este aparatado podra crear una reservacion primero entrara a la carpeta del postman llamada reservacion y seleccionara **Crear Reservacion**
###Ya estando en crear reservacion ingresara un json como en el ejemplo se muestra y para crearlo le dara en send
###Si el Cliente no esta registrado no podra reservar el viaje 
###Si el viaje no esta registrado no podra reservar el viaje 
###Si ya no existen espacions en el viaje no podra reservar el viaje

### {
###    "Personas" : Cantidad de personas que reservaran el viaje, (ingresar numero)
###    "IDCliente": Id del cliente, (el cliente tendra que estar registrado)
###    "IDViaje": id del viaje (el viaje tendra que estar registrado)
### }

### **4)B.- OBTENER RESERVACION**
###En este apartado podra ver su reservacion con su numero de id de la reservacion 
###Primero entramos en la carpeta de reservacion y ya despues seleccionamos el apartado **Obtener Reservacion** y pondra su id en el url, presionar el boton de send para ver su reservacion

### **4)D.- ELIMINAR RESERVACION**
###En este apartado podra eliminar una reservacion por su id, primero entrara a la carpeta del postman llamada reservacion y seleccionara **Eliminar Reservacion**
###Ya estando en eliminar reservacion pondra en el url el id de la reservacion que quiere eliminar, para eliminar la reservacion oprima el boton de send y la reservacion sera borrada


