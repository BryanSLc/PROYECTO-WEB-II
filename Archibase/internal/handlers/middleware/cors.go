package middleware

import "net/http"

// Cors permite que las peticiones desde el navegador o herramientas externas
// se comuniquen con la API de ArquiDraft sin ser bloqueadas por seguridad.
func Cors(siguiente http.Handler) http.Handler {
	return http.HandlerFunc(func(respuesta http.ResponseWriter, peticion *http.Request) {
		// Permitir el acceso desde cualquier origen en entorno de desarrollo
		respuesta.Header().Set("Access-Control-Allow-Origin", "*")

		// Indicar qué métodos HTTP están permitidos en ArquiDraft
		respuesta.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Permitir cabeceras personalizadas (indispensable para recibir el token de Authorization)
		respuesta.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Si el navegador realiza la petición previa de diagnóstico (OPTIONS), respondemos 200 OK de inmediato
		if peticion.Method == "OPTIONS" {
			respuesta.WriteHeader(http.StatusOK)
			return
		}

		// Continuar el flujo normal de la petición hacia el siguiente interceptor o controlador
		siguiente.ServeHTTP(respuesta, peticion)
	})
}
