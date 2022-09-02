package controller

import "github.com/beto-ouverney/nikiti-books/customerror"

// Delete is a function to delete a book
func (c *BookController) Delete(name string) *customerror.CustomError {
	err := c.Service.Delete(name)
	if err != nil {
		return err
	}

	return nil
}
