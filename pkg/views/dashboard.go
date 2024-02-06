package views

import (
	"github.com/Abdirahman04/novatrust/pkg/model"
	"github.com/Abdirahman04/novatrust/pkg/styles"
)

func Dashboard(customer model.Customer) {
  styles.Head("dashboard")
  styles.P(model.FullName(customer))
}
