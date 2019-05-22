# seguridad_social_crud

seguridad_social_crud, API de entidades que utiliza el sistema de Seguridad Social

Este proyecto está escrito en el lenguaje Go, utilizando el framework beego (https://beego.me/).

Se utiliza para administrar los modelos, transacciones y efectos de cambios directamente en la base de datos, nada de lógica. 


# Ejecución 

Clonar el repositorio

```go
  go get github.com/udistrital/seguridad_social_crud`
```

Moverse a la carpeta del proyecto

```sh
  cd $GOPATH/src/github.com/udistrital/seguridad_social_crud
```

Actualizar la rama develop

```sh
  git pull origin develop
```


Cambiar a la rama develop

```git
  git checkout develop
```

Asignar variables de entorno (ver fichero /conf/app.conf) 

Después de asignar las variables de entorno ejecutar

`bee run`

Para crear/actualizar la documentación y archivo swagger de forma automática ejecutar dentro del proyecto

`bee run -downdoc=true -gendoc=true`


## Derechos de Autor

This program is free software: you can redistribute it 
and/or modify it under the terms of the GNU General Public 
License as published by the Free Software Foundation, either
version 3 of the License, or (at your option) any later
version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

### UNIVERSIDAD DISTRITAL FRANCISCO JOSÉ DE CALDAS

### OFICINA ASESORA DE SISTEMAS

### 2019
