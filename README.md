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
```bash
git clone https://github.com/Toorreess/laWiki.git
```

### Componer el servicio
1. Acceder al directorio raíz del proyecto (./laWiki).
2. Lanzar el comando de composición de Docker:
```bash
docker compose -f docker_compose.yaml up
```
Se puede añadir el flag -d para lanzarlo en segundo plano y no bloquear la consola.
