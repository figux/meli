# Prueba Magneto

El desarrollo de la prueba se realizó con Golang y Docker.

Se uso la versión 1.15.10 de Golang, en la estructura de archivos se creó una structura para manejar el JSON de entrada llamada Person.
El algoritmo que resuelve el problema se encuentra en otro paquete llamado xmenrecluting. Como optimización salimos del ciclo si encontramos mas de una secuencia de 4 letras iguales.

El algoritmo valida 4 movimientos, si hay 4 letras hacia la derecha, hacia abajo, diagonal hacia la derecha y diagonal hacia la izquierda.

Para ejecutar el programa es necesario tener instalado docker.
1. Crear una imagen de docker con el comando: docker build -t magneto .
2. Correr la imagen docker por el puerto correcto con el comando: docker run -d -p 3009:3009 --rm magneto

El proyecto se encuentra alojado en una instancia micro de AWS con el puerto 3009 abierto a internet.
Se uso POSTMAN para las pruebas y los request POST con la información.

Ejemplos de respuestas:

# Mutante
<img src="https://figuxtests.s3-us-west-2.amazonaws.com/pruebaMELI/Mutante.png"/>

# NO Mutante
<img src="https://figuxtests.s3-us-west-2.amazonaws.com/pruebaMELI/NoMutante.png"/>

# URL de la API

IP: 44.239.48.102
Port: 3009
