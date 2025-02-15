package main

const (
	logDeleted  = "user deleted"
	logNotFound = "user not found"
	logAdmin    = "admin deleted"
)

type user struct {
	name   string
	number int
	admin  bool
}

func logAndDelete(users map[string]user, name string) (log string) {
	user, ok := users[name]

	delete(users, name)
	if ok {
		if user.admin {
			return logAdmin
		}
	}
	if !ok {
		return logNotFound
	}
	return logDeleted
}
