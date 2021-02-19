# Buscador de Mutantes

Con este proyecto se pretende agilizar la búsqueda de Magneto para reclutar mutantes y darle unas nalgaditas a los *X-Men*

## Requerimientos

* Golang 1.15 o sup.
* Generar un proyecto en Firebase.
* Postman.

## Configuración

Es necesario tener configurado la varibale de go: _GO111MODULE="on"_.

Si no tienes configurada una variable de entrono _PORT_ el puerto será el _3000_, en caso contrario, será el puerto que hayas configurado.

Debes descargar de Firebase el json de las credenciales de tu proyecto y después setear las variables de entorno _GOOGLE_APPLICATION_CREDENTIALS_ y _FBASE_ID_ en tu _.bashrc_, _.zshrc_, _profile_ o lo que uses, las instrucciones detalladas están [aquí](https://firebase.google.com/docs/admin/setup?authuser=0#initialize-sdk), la varibale _FBASE_ID_ no es más que el ID de tu proyecto de Firebase, lo puedes ver en el json que descargaste:

```
export GOOGLE_APPLICATION_CREDENTIALS="/path/al/firebase/credentials.json"
export FBASE_ID="id-firebase"
```

Ahora entra a tu proyecto Firebase y en Cloud Firestore crea la colección **stats**, el documento **counter** y agrega las propiedades **count_mutant_dna** y **count_human_dna**, ambas son numéricas, inicialózalas en **0**

## A Divertirnos!

Clona el repo y entra a la carpeta "mutantfinder/", una vez dentro, puedes correr el programa:

```
$ go gun main.go

// Después de instalar las dependencias te aparecerá algo así:

2021/02/18 23:18:19 Environment is ok!

 ┌───────────────────────────────────────────────────┐ 
 │                    Fiber v2.5.0                   │ 
 │               http://127.0.0.1:3000               │ 
 │                                                   │ 
 │ Handlers ............ 11  Processes ........... 1 │ 
 │ Prefork ....... Disabled  PID ............. 55783 │ 
 └───────────────────────────────────────────────────┘ 
```

El proyecto tiene un validador de las variables de entorno antes mencionadas, si no las tienes configuradas, te dará error y si están mal configuradas, el proyecto tirará errores.

### Postman

Puedes hacer pruebas con postman, [aquí](https://www.getpostman.com/collections/edc6ab5939f89c531e68) te dejo la colección, sólo tendrás que agregar una bariable global {{baseUrl}} con el valor de tu local host (lo puedes ver al iniciar el programa, te da la url)

### Descripción del API

Puedes hacer un GET a  http://127.0.0.1:3000/api.yaml y te dará la documentación en OpenAPI 3

## Test

Puedes hacer el testing del algorimo de análisi de las cadenas de ADN corriendo:

```
$ go test ./structs
```

Correrá los test de los diferentes casos del algoritmo.



**Diviérte!**
