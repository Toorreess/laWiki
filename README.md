# laWiki
## Descripción
Aplicación web de hosteo de wikis y entradas, desarrollado para la asignatura *Ingeniería Web* impartida en el cuarto curso del grado de *Ingeniería del Software*, en la *Universidad de Málaga*.

## Grupo 3F
- Álvaro Acedo Espejo
- Jesús Moreno Carmona
- Nicolás Reyes Trujillo
- Jaime Rodrigo Roldán Corcelles
- Juan José Serrano España
- José Torres Postigo

## Instalación
### Prerrequisitos
- Tener instalado git
- Tener instalado Docker
- (opcional) Tener instalado Go 1.23.2

### Clonar el repositorio
En caso de clonar el repositorio en vez de usar el código fuente entregado, usar el comando:
```bash
git clone https://github.com/Toorreess/laWiki.git
```

### Componer el servicio
1. Acceder al directorio raíz del proyecto (./laWiki).
2. **Si se ha clonado el repo**, se debe copiar en los directorios wiki-service/, entry-service/ y comment-service/ el archivo creds.json.

\**Si se utiliza el fichero .zip con el código fuente entregado, no es necesario este paso*.

3. Lanzar el comando de composición de Docker:
```bash
docker compose -f docker_compose.yaml up
```
Se puede añadir el flag -d para lanzarlo en segundo plano y no bloquear la consola.
