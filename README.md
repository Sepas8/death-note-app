# Death Note App 📝💀

Aplicación Full Stack inspirada en **Death Note**. Permite gestionar personas (estado vivo/muerto), con backend en Go y frontend en React.

---

## 🧱 Tecnologías

- Backend: Go + Gorilla Mux + GORM + PostgreSQL
- Frontend: React + Vite + TypeScript
- Contenedores: Docker & Docker Compose

---

## 🚀 Instalación

### 1. Clona el proyecto

```bash
git clone https://github.com/usuario/death-note-app.git
cd death-note-app
```

### 2. Requisitos

- ✅ [Docker Desktop](https://www.docker.com/products/docker-desktop) instalado y ejecutándose

---

## ⚙️ Levantar el proyecto (Backend + DB)

Desde la carpeta raíz del proyecto:

```bash
cd backend
docker compose up --build
```

Esto levantará:

- Servidor Go en `http://localhost:8000`
- Base de datos PostgreSQL en `localhost:5432`

---

## 📦 Endpoints del backend

| Método | Endpoint          | Descripción                |
| ------ | ----------------- | -------------------------- |
| GET    | `/api/people`     | Obtener todas las personas |
| GET    | `/api/people/:id` | Obtener persona por ID     |
| POST   | `/api/people`     | Crear nueva persona        |
| PUT    | `/api/people/:id` | Editar persona             |
| DELETE | `/api/people/:id` | Eliminar persona           |

---

## 💻 Frontend (opcional si tienes)

Si el frontend está configurado:

```bash
cd frontend
npm install
npm run dev
```

Por defecto se abre en: `http://localhost:5173`

Asegúrate de que el archivo `.env` del frontend tenga la URL correcta del backend:

```env
VITE_API_URL=http://localhost:8000
```

---

## 🧪 Probar conexión

Puedes probar usando Postman, ThunderClient o `curl`:

```bash
curl http://localhost:8000/api/people
```

---

## 🗃️ Variables de entorno

Archivo `.env` en `backend/` (ya incluido o puedes copiar el ejemplo):

```env
POSTGRES_DB=deathnote
POSTGRES_USER=admin
POSTGRES_PASSWORD=secret
POSTGRES_HOST=postgres
```

---

## 🧼 Apagar todo

```bash
cd backend
docker compose down
```

---

## 📝 Autor

Hecho con cariño por [tu nombre o alias].
