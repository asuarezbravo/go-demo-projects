
# Go Quiz Program

Este es un programa de cuestionario de línea de comandos escrito en Go. Lee preguntas y respuestas desde un archivo CSV, presenta cada pregunta al usuario y evalúa las respuestas basándose en su corrección dentro de un límite de tiempo especificado.

## Características

- **Entrada desde CSV**: Las preguntas y respuestas se leen desde un archivo CSV especificado por el usuario.
- **Cuestionario con Temporizador**: El usuario puede establecer un límite de tiempo usando una bandera.
- **Concurrencia con Goroutines y Canales**: El programa maneja de manera eficiente la entrada de datos dentro del límite de tiempo usando goroutines y canales.
- **Calificación en Tiempo Real**: Muestra la puntuación del usuario al completar el cuestionario o cuando el tiempo se agota.

## Uso

1. Compila o ejecuta el programa directamente usando Go:
   ```bash
   go run quiz.go -csv=yourfile.csv -limit=30
   ```
   Reemplaza `yourfile.csv` con la ruta a tu archivo CSV, y `30` con el límite de tiempo deseado en segundos.

2. **Banderas de Línea de Comandos**:
   - `-csv` (opcional): Ruta al archivo CSV que contiene preguntas y respuestas (por defecto es `problems.csv`).
   - `-limit` (opcional): Límite de tiempo para el cuestionario en segundos (por defecto es `30` segundos).

3. **Formato del Archivo CSV**: El archivo CSV debe tener las preguntas en la primera columna y las respuestas en la segunda, por ejemplo:
   ```csv
   question,answer
   5+5,10
   7-3,4
   ```

## Paquetes Utilizados

- **`flag`**: Permite gestionar argumentos de línea de comandos de forma sencilla y efectiva, lo que facilita el control sobre el archivo CSV de entrada y el límite de tiempo.
- **`csv`**: Proporciona una forma sencilla de leer archivos CSV y estructurarlos en una lista de preguntas y respuestas, mediante métodos como `ReadAll`, que simplifica la conversión de datos estructurados.
- **`time`**: Permite configurar un temporizador para establecer el límite de tiempo y gestionar las respuestas del usuario dentro de ese tiempo mediante el uso de `NewTimer` y `Timer.C`.

## Decisiones de Diseño

El programa incorpora varias decisiones de diseño para hacerlo modular, eficiente y fácil de usar:

### 1. Manejo de Archivos con `defer file.Close()`
   - El uso de `defer` asegura que el archivo CSV se cierra después de que se termine de leer, liberando recursos del sistema y previniendo fugas de memoria. Colocar `defer` inmediatamente después de abrir el archivo es una práctica común en Go para manejar recursos de manera segura.

### 2. Concurrencia con Goroutines y Canales
   - Utilizamos una goroutine para manejar la entrada del usuario de manera concurrente mientras el temporizador corre. Esto permite que el cuestionario finalice de inmediato cuando se alcanza el límite de tiempo, sin esperar a que el usuario responda.
   - Un canal (`answerCh`) permite que la goroutine envíe la respuesta del usuario al proceso principal. El programa selecciona entre recibir la respuesta del canal o la señal del temporizador (`timer.C`), controlando de manera eficiente el flujo del cuestionario y la respuesta dentro del límite de tiempo.

### 3. Modularidad y Función `askQuestions`
   - La lógica de preguntas y respuestas se encapsula en una función `askQuestions`, mejorando la organización del código y facilitando futuras modificaciones, como agregar nuevas reglas de preguntas o configuraciones de puntaje. `main` ahora se enfoca únicamente en la configuración inicial, lectura de archivos, y llamada a funciones principales.

### 4. Validación y Saneamiento de Datos
   - En la función `parseLines`, cada respuesta se procesa con `strings.TrimSpace` para asegurar que no haya espacios en blanco innecesarios que puedan afectar la comparación de respuestas.

### 5. Función `exit` para Manejo de Errores
   - La función `exit` centraliza los mensajes de error y la salida del programa, permitiendo una terminación limpia y consistente del programa en caso de error.

## Estructura del Código

- **`main`**: Inicializa las banderas, lee el archivo CSV y lanza el cuestionario.
- **`askQuestions`**: Función principal que presenta cada pregunta al usuario y maneja la respuesta dentro del límite de tiempo.
- **`parseLines`**: Convierte las líneas del archivo CSV en una estructura `problem`, simplificando la manipulación de preguntas y respuestas.
- **`exit`**: Imprime un mensaje de error y termina el programa para asegurar un cierre seguro y controlado en caso de errores.

## Ejemplo de Ejecución

```bash
go run quiz.go -csv=problems.csv -limit=20
```

Esto inicia un cuestionario utilizando `problems.csv` con un límite de tiempo de 20 segundos. El programa mostrará cada pregunta y pedirá una respuesta al usuario. Cuando se termine el tiempo o se respondan todas las preguntas, el programa muestra la puntuación final.

## Posibles Mejoras

- **Sugerencias o Intentos Múltiples**: Para mejorar la interactividad educativa, se podrían agregar sugerencias o permitir múltiples intentos.
- **Carga Dinámica de Preguntas**: Integrar una API o base de datos para cargar preguntas dinámicamente.
- **Desglose de Puntuación**: Mostrar las respuestas correctas e incorrectas al final ayudaría al usuario a aprender de sus errores.

## Licencia

Este proyecto es de código abierto y está disponible bajo la licencia MIT.
