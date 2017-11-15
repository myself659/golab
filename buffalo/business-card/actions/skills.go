package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/mycodesmells/golang-examples/buffalo/business-card/models"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Skill)
// DB Table: Plural (skills)
// Resource: Plural (Skills)
// Path: Plural (/skills)
// View Template Folder: Plural (/templates/skills/)

// SkillsResource is the resource for the skill model
type SkillsResource struct {
	buffalo.Resource
}

// List gets all Skills. This function is mapped to the path
// GET /skills
func (v SkillsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	skills := &models.Skills{}
	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())
	// You can order your list here. Just change
	err := q.All(skills)
	// to:
	// err := q.Order("created_at desc").All(skills)
	if err != nil {
		return errors.WithStack(err)
	}
	// Make Skills available inside the html template
	c.Set("skills", skills)
	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)
	return c.Render(200, r.HTML("skills/index.html", "old_application.html"))
}

// Show gets the data for one Skill. This function is mapped to
// the path GET /skills/{skill_id}
func (v SkillsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Allocate an empty Skill
	skill := &models.Skill{}
	// To find the Skill the parameter skill_id is used.
	err := tx.Find(skill, c.Param("skill_id"))
	if err != nil {
		return errors.WithStack(err)
	}
	// Make skill available inside the html template
	c.Set("skill", skill)
	return c.Render(200, r.HTML("skills/show.html", "old_application.html"))
}

// New renders the formular for creating a new Skill.
// This function is mapped to the path GET /skills/new
func (v SkillsResource) New(c buffalo.Context) error {
	// Make skill available inside the html template
	c.Set("skill", &models.Skill{})
	return c.Render(200, r.HTML("skills/new.html", "old_application.html"))
}

// Create adds a Skill to the DB. This function is mapped to the
// path POST /skills
func (v SkillsResource) Create(c buffalo.Context) error {
	// Allocate an empty Skill
	skill := &models.Skill{}
	// Bind skill to the html form elements
	err := c.Bind(skill)
	if err != nil {
		return errors.WithStack(err)
	}
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(skill)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		// Make skill available inside the html template
		c.Set("skill", skill)
		// Make the errors available inside the html template
		c.Set("errors", verrs)
		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("skills/new.html", "old_application.html"))
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "Skill was created successfully")
	// and redirect to the skills index page
	return c.Redirect(302, "/skills/%s", skill.ID)
}

// Edit renders a edit formular for a skill. This function is
// mapped to the path GET /skills/{skill_id}/edit
func (v SkillsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Allocate an empty Skill
	skill := &models.Skill{}
	err := tx.Find(skill, c.Param("skill_id"))
	if err != nil {
		return errors.WithStack(err)
	}
	// Make skill available inside the html template
	c.Set("skill", skill)
	return c.Render(200, r.HTML("skills/edit.html", "old_application.html"))
}

// Update changes a skill in the DB. This function is mapped to
// the path PUT /skills/{skill_id}
func (v SkillsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Allocate an empty Skill
	skill := &models.Skill{}
	err := tx.Find(skill, c.Param("skill_id"))
	if err != nil {
		return errors.WithStack(err)
	}
	// Bind Skill to the html form elements
	err = c.Bind(skill)
	if err != nil {
		return errors.WithStack(err)
	}
	verrs, err := tx.ValidateAndUpdate(skill)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		// Make skill available inside the html template
		c.Set("skill", skill)
		// Make the errors available inside the html template
		c.Set("errors", verrs)
		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("skills/edit.html", "old_application.html"))
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "Skill was updated successfully")
	// and redirect to the skills index page
	return c.Redirect(302, "/skills/%s", skill.ID)
}

// Destroy deletes a skill from the DB. This function is mapped
// to the path DELETE /skills/{skill_id}
func (v SkillsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Allocate an empty Skill
	skill := &models.Skill{}
	// To find the Skill the parameter skill_id is used.
	err := tx.Find(skill, c.Param("skill_id"))
	if err != nil {
		return errors.WithStack(err)
	}
	err = tx.Destroy(skill)
	if err != nil {
		return errors.WithStack(err)
	}
	// If there are no errors set a flash message
	c.Flash().Add("success", "Skill was destroyed successfully")
	// Redirect to the skills index page
	return c.Redirect(302, "/skills")
}
