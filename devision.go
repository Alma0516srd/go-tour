// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Using the template, declare a set of concrete types that implement the set
// of predefined interface types. Then create values of these types and use
// them to complete a set of predefined tasks.
package main

import "fmt"

// administrator represents a person or other entity capable of administering
// hardware and software infrastructure.
type administrator interface {
	administrate(system string)
}

// developer represents a person or other entity capable of writing software.
type developer interface {
	develop(system string)
}

// =============================================================================

// adminlist represents a group of administrators.
type adminlist struct {
	list []administrator
}

// Enqueue adds an administrator to the adminlist.
func (l *adminlist) Enqueue(a administrator) {
	l.list = append(l.list, a)
}

// Dequeue removes an administrator from the adminlist.
func (l *adminlist) Dequeue() administrator {
	a := l.list[0]
	l.list = l.list[1:]
	return a
}

// =============================================================================

// devlist represents a group of developers.
type devlist struct {
	list []developer
}

// Enqueue adds a developer to the devlist.
func (l *devlist) Enqueue(d developer) {
	l.list = append(l.list, d)
}

// Dequeue removes a developer from the devlist.
func (l *devlist) Dequeue() developer {
	d := l.list[0]
	l.list = l.list[1:]
	return d
}

// =============================================================================

// sysadmin представляет системного администратора с полем name.
type sysadmin struct {
	name string
}

func (s sysadmin) administrate(system string) {
	fmt.Printf("Sysadmin %s administers system %s\n", s.name, system)
}

type programmer struct {
	name string
}

func (p programmer) develop(system string) {
	fmt.Printf("Programmer %s develops system %s\n", p.name, system)
}

// company представляет компанию, которая объединяет возможности
// администратора и разработчика через встраивание интерфейсов.
type company struct {
	administrator
	developer
}

// =============================================================================

func main() {

	var admins adminlist

	var devs devlist

	admins.Enqueue(sysadmin{name: "Alice"})

	devs.Enqueue(programmer{name: "Bob"})
	devs.Enqueue(programmer{name: "Carol"})

	cmp := company{
		administrator: admins.Dequeue(),
		developer:     devs.Dequeue(),
	}

	admins.Enqueue(cmp)
	devs.Enqueue(cmp)

	tasks := []struct {
		needsAdmin bool
		system     string
	}{
		{needsAdmin: false, system: "xenia"},
		{needsAdmin: true, system: "pillar"},
		{needsAdmin: false, system: "omega"},
	}

	for _, task := range tasks {

		if task.needsAdmin {

			admin := admins.Dequeue()
			admin.administrate(task.system)

			continue
		}

		dev := devs.Dequeue()
		dev.develop(task.system)
	}
}
