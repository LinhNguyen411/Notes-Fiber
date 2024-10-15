package noteRoutes

import (
	noteHandler "github.com/LinhNguyen411/notes-api-fiber/internal/handler/note"
	"github.com/gofiber/fiber/v2"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")

	note.Post("/", noteHandler.CreateNotes)

	note.Get("/", noteHandler.GetNotes)

	note.Get("/:noteId", noteHandler.GetNote)

	note.Put("/:noteId", noteHandler.UpdataNote)

	note.Delete("/:noteId", noteHandler.DeleteNote)
}
