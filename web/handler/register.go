package handler

import (
	"github.com/gorilla/csrf"
	"net/http"
	"vpub/model"
	"vpub/web/handler/form"
)

func (h *Handler) showRegisterView(w http.ResponseWriter, r *http.Request) {
	h.renderLayout(w, "register", map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(r),
	}, model.User{})
}

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	userForm := form.NewUserForm(r)
	user := model.User{
		Name:     userForm.Username,
		Password: userForm.Password,
	}
	showError := func(err error) {
		h.renderLayout(w, "register", map[string]interface{}{"form": *userForm, "error": err.Error(), csrf.TemplateTag: csrf.TemplateField(r)}, model.User{})
		return
	}
	if err := userForm.Validate(); err != nil {
		showError(err)
		return
	}
	//if h.storage.UserExists(user.Name) {
	//	serverError(w, errors.New("username already exists"))
	//	return
	//}
	//if ok := key.Unlock(r.FormValue("key")); !ok {
	//	forbidden(w)
	//	return
	//}
	id, err := h.storage.CreateUser(user)
	if err != nil {
		serverError(w, err)
		return
	}
	if err := h.session.Save(r, w, id); err != nil {
		serverError(w, err)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
