// developed by Dariel Vega Bernal

This is not the final document, it is a work in progress.

# Introduction

# API explanation

Esta aplicacion tiene la funcionalidad de automatizar las descargar la informacion de los test que se requieran, en conjunto con la los resultados de los candidatos y sus datos.

se requiere:

- Instalar golang
- Ejecutar pwd (en mac) para saber el path donde se encuentra el proyecto
- exportar las variables de ambiente (en PATH, es donde se instalo go)
export GOPATH=(valor de pwd, pero sin la carpeta del proyecto) /
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin 
- obtener la cookie 
- obtener el id de los test a descargar

ejecucion en mac:

1. Descargar el proyecto localmente

2. Abrir terminal dentro del proyect 

3. Ejecutar pwd

4. ejecutar en terminal (o ponerlo en las variables de ambiente para no tener que hacer este paso siempre que se vaya a ejecutar):
export GOPATH=(valor de pwd, pero sin la carpeta del proyecto) /
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin 

5. agregar la cookie en src/env.go en la parte de TOKEN, donde esta el string vacio.

6. agregar el id en TESTS_TO_SEARCH en src/env.go, ejemplo: []string{"123","456","789"}

7. ejecutar en terminal: gon run main.go

8. Esto generara el scripts y una carpeta con la fecha del dia, con los las respuestas obtenidas de los test como jsons.