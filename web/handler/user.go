package handler

import (
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"vpub/model"
	"vpub/web/handler/form"
)

func (h *Handler) showUserPostsView(w http.ResponseWriter, r *http.Request) {
	logged, _ := h.session.Get(r)

	var page int64 = 1
	if val, ok := r.URL.Query()["page"]; ok && len(val[0]) == 1 {
		page, _ = strconv.ParseInt(val[0], 10, 64)
	}

	user, err := h.storage.UserByName(mux.Vars(r)["userId"])
	if err != nil {
		notFound(w)
		return
	}

	//posts, showMore, err := h.storage.PostsByUsernameWithReplyCount(user.Name, h.perPage, page)
	if err != nil {
		serverError(w, err)
		return
	}

	h.renderLayout(w, "user_posts", map[string]interface{}{
		"user":     user,
		"posts":    "",
		"page":     page,
		"showMore": "",
		"nextPage": page + 1,
	}, logged)
}

func (h *Handler) showAccountView(w http.ResponseWriter, r *http.Request, user model.User) {
	h.renderLayout(w, "account", map[string]interface{}{
		"user":           user,
		csrf.TemplateTag: csrf.TemplateField(r),
	}, user)
}

func (h *Handler) saveAccount(w http.ResponseWriter, r *http.Request, user model.User) {
	accountForm := form.NewAccountForm(r)
	user.About = accountForm.About
	user.Picture = accountForm.Picture
	if err := h.storage.UpdateUser(user); err != nil {
		serverError(w, err)
		return
	}
	http.Redirect(w, r, "/~"+user.Name, http.StatusFound)
}
