package noteHandler

import (
	"github.com/LinhNguyen411/notes-api-fiber/database"
	"github.com/LinhNguyen411/notes-api-fiber/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes []model.Note

	db.Find(&notes)

	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Notes found", "data": notes})
}

func CreateNotes(c *fiber.Ctx) error {
	db := database.DB
	note := new(model.Note)

	err := c.BodyParser(note)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	note.ID = uuid.New()

	err = db.Create(&note).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create your note", "data": err})
	}
	return c.JSON(fiber.Map{"status": "succes", "message": "Created note", "data": note})
}

func GetNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	// Read the param noteId
	id := c.Params("noteId")

	// Find the note with the given Id
	db.Find(&note, "id = ?", id)

	// If no such note present return an error
	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	// Return the note with the Id
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": note})
}

func UpdataNote(c *fiber.Ctx) error {
	type updateNote struct {
		Title    string `json:"title"`
		Subtitle string `json:"subtitle"`
		Text     string `json:"text"`
	}

	db := database.DB
	var note model.Note

	id := c.Params("noteId")

	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	var updateNoteData updateNote
	err := c.BodyParser(&updateNoteData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	note.Title = updateNoteData.Title
	note.Subtitle = updateNoteData.Subtitle
	note.Text = updateNoteData.Text

	db.Save(&note)

	return c.JSON(fiber.Map{"status": "success", "message": "Updated note", "data": note})
}

func DeleteNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	id := c.Params("noteId")
	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	// Delete the note and return error if encountered
	err := db.Delete(&note, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete note", "data": nil})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted Note"})
}
