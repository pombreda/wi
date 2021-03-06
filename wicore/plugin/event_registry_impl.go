// generated by go run ../tools/wi-event-generator/main.go ; DO NOT EDIT

package plugin

import (
	"errors"
	"sync"

	"github.com/wi-ed/wi/internal"
	"github.com/wi-ed/wi/wicore"
	"github.com/wi-ed/wi/wicore/key"
	"github.com/wi-ed/wi/wicore/lang"
)

type eventTriggerRPC struct {
	eventRegistry
}

// makeEventRegistry returns a wicore.EventRegistry and the channel to read
// from to run the events piped in.
func makeEventRegistry() (wicore.EventRegistry, internal.EventTriggerRPC, chan func()) {
	// Reduce the odds of allocation within RegistryXXX() by using relatively
	// large buffers.
	c := make(chan func(), 2048)
	e := &eventTriggerRPC{
		eventRegistry{
			deferred:                  c,
			commands:                  make([]listenerCommands, 0, 64),
			documentCreated:           make([]listenerDocumentCreated, 0, 64),
			documentCursorMoved:       make([]listenerDocumentCursorMoved, 0, 64),
			editorKeyboardModeChanged: make([]listenerEditorKeyboardModeChanged, 0, 64),
			editorLanguage:            make([]listenerEditorLanguage, 0, 64),
			terminalKeyPressed:        make([]listenerTerminalKeyPressed, 0, 64),
			terminalMetaKeyPressed:    make([]listenerTerminalMetaKeyPressed, 0, 64),
			terminalResized:           make([]listenerTerminalResized, 0, 64),
			viewActivated:             make([]listenerViewActivated, 0, 64),
			viewCreated:               make([]listenerViewCreated, 0, 64),
			windowCreated:             make([]listenerWindowCreated, 0, 64),
			windowResized:             make([]listenerWindowResized, 0, 64),
		},
	}
	return e, e, c
}

func (er *eventTriggerRPC) TriggerCommandsRPC(packet internal.PacketCommands, ignored *int) error {
	er.triggerCommands(packet.Cmds)
	return nil
}

func (er *eventTriggerRPC) TriggerDocumentCreatedRPC(packet internal.PacketDocumentCreated, ignored *int) error {
	er.triggerDocumentCreated(packet.Doc)
	return nil
}

func (er *eventTriggerRPC) TriggerDocumentCursorMovedRPC(packet internal.PacketDocumentCursorMoved, ignored *int) error {
	er.triggerDocumentCursorMoved(packet.Doc, packet.Col, packet.Row)
	return nil
}

func (er *eventTriggerRPC) TriggerEditorKeyboardModeChangedRPC(packet internal.PacketEditorKeyboardModeChanged, ignored *int) error {
	er.triggerEditorKeyboardModeChanged(packet.Mode)
	return nil
}

func (er *eventTriggerRPC) TriggerEditorLanguageRPC(packet internal.PacketEditorLanguage, ignored *int) error {
	er.triggerEditorLanguage(packet.L)
	return nil
}

func (er *eventTriggerRPC) TriggerTerminalKeyPressedRPC(packet internal.PacketTerminalKeyPressed, ignored *int) error {
	er.triggerTerminalKeyPressed(packet.K)
	return nil
}

func (er *eventTriggerRPC) TriggerTerminalMetaKeyPressedRPC(packet internal.PacketTerminalMetaKeyPressed, ignored *int) error {
	er.triggerTerminalMetaKeyPressed(packet.K)
	return nil
}

func (er *eventTriggerRPC) TriggerTerminalResizedRPC(packet internal.PacketTerminalResized, ignored *int) error {
	er.triggerTerminalResized()
	return nil
}

func (er *eventTriggerRPC) TriggerViewActivatedRPC(packet internal.PacketViewActivated, ignored *int) error {
	er.triggerViewActivated(packet.View)
	return nil
}

func (er *eventTriggerRPC) TriggerViewCreatedRPC(packet internal.PacketViewCreated, ignored *int) error {
	er.triggerViewCreated(packet.View)
	return nil
}

func (er *eventTriggerRPC) TriggerWindowCreatedRPC(packet internal.PacketWindowCreated, ignored *int) error {
	er.triggerWindowCreated(packet.Window)
	return nil
}

func (er *eventTriggerRPC) TriggerWindowResizedRPC(packet internal.PacketWindowResized, ignored *int) error {
	er.triggerWindowResized(packet.Window)
	return nil
}

func (er *eventRegistry) TriggerCommands(cmds wicore.EnqueuedCommands) {
	// TODO(maruel): Send it upstream to the editor.
}

func (er *eventRegistry) TriggerDocumentCreated(doc wicore.Document) {
	// TODO(maruel): Send it upstream to the editor.
}

func (er *eventRegistry) TriggerDocumentCursorMoved(doc wicore.Document, col, row int) {
	// TODO(maruel): Send it upstream to the editor.
}

func (er *eventRegistry) TriggerEditorKeyboardModeChanged(mode wicore.KeyboardMode) {
	// TODO(maruel): Send it upstream to the editor.
}

func (er *eventRegistry) TriggerEditorLanguage(l lang.Language) {
	// TODO(maruel): Send it upstream to the editor.
}

func (er *eventRegistry) TriggerTerminalKeyPressed(k key.Press) {
	// TODO(maruel): Send it upstream to the editor.
}

func (er *eventRegistry) TriggerTerminalMetaKeyPressed(k key.Press) {
	// TODO(maruel): Send it upstream to the editor.
}

func (er *eventRegistry) TriggerTerminalResized() {
	// TODO(maruel): Send it upstream to the editor.
}

func (er *eventRegistry) TriggerViewActivated(view wicore.View) {
	// TODO(maruel): Send it upstream to the editor.
}

func (er *eventRegistry) TriggerViewCreated(view wicore.View) {
	// TODO(maruel): Send it upstream to the editor.
}

func (er *eventRegistry) TriggerWindowCreated(window wicore.Window) {
	// TODO(maruel): Send it upstream to the editor.
}

func (er *eventRegistry) TriggerWindowResized(window wicore.Window) {
	// TODO(maruel): Send it upstream to the editor.
}

type listenerCommands struct {
	id       int
	callback func(cmds wicore.EnqueuedCommands)
}

type listenerDocumentCreated struct {
	id       int
	callback func(doc wicore.Document)
}

type listenerDocumentCursorMoved struct {
	id       int
	callback func(doc wicore.Document, col, row int)
}

type listenerEditorKeyboardModeChanged struct {
	id       int
	callback func(mode wicore.KeyboardMode)
}

type listenerEditorLanguage struct {
	id       int
	callback func(l lang.Language)
}

type listenerTerminalKeyPressed struct {
	id       int
	callback func(k key.Press)
}

type listenerTerminalMetaKeyPressed struct {
	id       int
	callback func(k key.Press)
}

type listenerTerminalResized struct {
	id       int
	callback func()
}

type listenerViewActivated struct {
	id       int
	callback func(view wicore.View)
}

type listenerViewCreated struct {
	id       int
	callback func(view wicore.View)
}

type listenerWindowCreated struct {
	id       int
	callback func(window wicore.Window)
}

type listenerWindowResized struct {
	id       int
	callback func(window wicore.Window)
}

// eventRegistry is automatically generated via wi-event-generator from the
// interface wicore.EventRegistry. It completely implements
// wicore.EventRegistry.
type eventRegistry struct {
	lock     sync.Mutex
	nextID   int
	deferred chan<- func()

	commands                  []listenerCommands
	documentCreated           []listenerDocumentCreated
	documentCursorMoved       []listenerDocumentCursorMoved
	editorKeyboardModeChanged []listenerEditorKeyboardModeChanged
	editorLanguage            []listenerEditorLanguage
	terminalKeyPressed        []listenerTerminalKeyPressed
	terminalMetaKeyPressed    []listenerTerminalMetaKeyPressed
	terminalResized           []listenerTerminalResized
	viewActivated             []listenerViewActivated
	viewCreated               []listenerViewCreated
	windowCreated             []listenerWindowCreated
	windowResized             []listenerWindowResized
}

func (er *eventRegistry) unregister(eventID int) {
	er.lock.Lock()
	defer er.lock.Unlock()
	// TODO(maruel): The buffers are never reallocated, so it's effectively a
	// memory leak.
	switch eventID & 0xff000000 {
	case 0x1000000:
		for index, value := range er.commands {
			if value.id == eventID {
				copy(er.commands[index:], er.commands[index+1:])
				er.commands = er.commands[0 : len(er.commands)-1]
				return
			}
		}
	case 0x2000000:
		for index, value := range er.documentCreated {
			if value.id == eventID {
				copy(er.documentCreated[index:], er.documentCreated[index+1:])
				er.documentCreated = er.documentCreated[0 : len(er.documentCreated)-1]
				return
			}
		}
	case 0x3000000:
		for index, value := range er.documentCursorMoved {
			if value.id == eventID {
				copy(er.documentCursorMoved[index:], er.documentCursorMoved[index+1:])
				er.documentCursorMoved = er.documentCursorMoved[0 : len(er.documentCursorMoved)-1]
				return
			}
		}
	case 0x4000000:
		for index, value := range er.editorKeyboardModeChanged {
			if value.id == eventID {
				copy(er.editorKeyboardModeChanged[index:], er.editorKeyboardModeChanged[index+1:])
				er.editorKeyboardModeChanged = er.editorKeyboardModeChanged[0 : len(er.editorKeyboardModeChanged)-1]
				return
			}
		}
	case 0x5000000:
		for index, value := range er.editorLanguage {
			if value.id == eventID {
				copy(er.editorLanguage[index:], er.editorLanguage[index+1:])
				er.editorLanguage = er.editorLanguage[0 : len(er.editorLanguage)-1]
				return
			}
		}
	case 0x6000000:
		for index, value := range er.terminalKeyPressed {
			if value.id == eventID {
				copy(er.terminalKeyPressed[index:], er.terminalKeyPressed[index+1:])
				er.terminalKeyPressed = er.terminalKeyPressed[0 : len(er.terminalKeyPressed)-1]
				return
			}
		}
	case 0x7000000:
		for index, value := range er.terminalMetaKeyPressed {
			if value.id == eventID {
				copy(er.terminalMetaKeyPressed[index:], er.terminalMetaKeyPressed[index+1:])
				er.terminalMetaKeyPressed = er.terminalMetaKeyPressed[0 : len(er.terminalMetaKeyPressed)-1]
				return
			}
		}
	case 0x8000000:
		for index, value := range er.terminalResized {
			if value.id == eventID {
				copy(er.terminalResized[index:], er.terminalResized[index+1:])
				er.terminalResized = er.terminalResized[0 : len(er.terminalResized)-1]
				return
			}
		}
	case 0x9000000:
		for index, value := range er.viewActivated {
			if value.id == eventID {
				copy(er.viewActivated[index:], er.viewActivated[index+1:])
				er.viewActivated = er.viewActivated[0 : len(er.viewActivated)-1]
				return
			}
		}
	case 0xa000000:
		for index, value := range er.viewCreated {
			if value.id == eventID {
				copy(er.viewCreated[index:], er.viewCreated[index+1:])
				er.viewCreated = er.viewCreated[0 : len(er.viewCreated)-1]
				return
			}
		}
	case 0xb000000:
		for index, value := range er.windowCreated {
			if value.id == eventID {
				copy(er.windowCreated[index:], er.windowCreated[index+1:])
				er.windowCreated = er.windowCreated[0 : len(er.windowCreated)-1]
				return
			}
		}
	case 0xc000000:
		for index, value := range er.windowResized {
			if value.id == eventID {
				copy(er.windowResized[index:], er.windowResized[index+1:])
				er.windowResized = er.windowResized[0 : len(er.windowResized)-1]
				return
			}
		}
	}
}

func (er *eventRegistry) RegisterCommands(callback func(cmds wicore.EnqueuedCommands)) wicore.EventListener {
	er.lock.Lock()
	defer er.lock.Unlock()
	i := er.nextID
	er.nextID++
	er.commands = append(er.commands, listenerCommands{i, callback})
	return &eventListener{er, i | 0x1000000}
}

func (er *eventRegistry) RegisterDocumentCreated(callback func(doc wicore.Document)) wicore.EventListener {
	er.lock.Lock()
	defer er.lock.Unlock()
	i := er.nextID
	er.nextID++
	er.documentCreated = append(er.documentCreated, listenerDocumentCreated{i, callback})
	return &eventListener{er, i | 0x2000000}
}

func (er *eventRegistry) RegisterDocumentCursorMoved(callback func(doc wicore.Document, col, row int)) wicore.EventListener {
	er.lock.Lock()
	defer er.lock.Unlock()
	i := er.nextID
	er.nextID++
	er.documentCursorMoved = append(er.documentCursorMoved, listenerDocumentCursorMoved{i, callback})
	return &eventListener{er, i | 0x3000000}
}

func (er *eventRegistry) RegisterEditorKeyboardModeChanged(callback func(mode wicore.KeyboardMode)) wicore.EventListener {
	er.lock.Lock()
	defer er.lock.Unlock()
	i := er.nextID
	er.nextID++
	er.editorKeyboardModeChanged = append(er.editorKeyboardModeChanged, listenerEditorKeyboardModeChanged{i, callback})
	return &eventListener{er, i | 0x4000000}
}

func (er *eventRegistry) RegisterEditorLanguage(callback func(l lang.Language)) wicore.EventListener {
	er.lock.Lock()
	defer er.lock.Unlock()
	i := er.nextID
	er.nextID++
	er.editorLanguage = append(er.editorLanguage, listenerEditorLanguage{i, callback})
	return &eventListener{er, i | 0x5000000}
}

func (er *eventRegistry) RegisterTerminalKeyPressed(callback func(k key.Press)) wicore.EventListener {
	er.lock.Lock()
	defer er.lock.Unlock()
	i := er.nextID
	er.nextID++
	er.terminalKeyPressed = append(er.terminalKeyPressed, listenerTerminalKeyPressed{i, callback})
	return &eventListener{er, i | 0x6000000}
}

func (er *eventRegistry) RegisterTerminalMetaKeyPressed(callback func(k key.Press)) wicore.EventListener {
	er.lock.Lock()
	defer er.lock.Unlock()
	i := er.nextID
	er.nextID++
	er.terminalMetaKeyPressed = append(er.terminalMetaKeyPressed, listenerTerminalMetaKeyPressed{i, callback})
	return &eventListener{er, i | 0x7000000}
}

func (er *eventRegistry) RegisterTerminalResized(callback func()) wicore.EventListener {
	er.lock.Lock()
	defer er.lock.Unlock()
	i := er.nextID
	er.nextID++
	er.terminalResized = append(er.terminalResized, listenerTerminalResized{i, callback})
	return &eventListener{er, i | 0x8000000}
}

func (er *eventRegistry) RegisterViewActivated(callback func(view wicore.View)) wicore.EventListener {
	er.lock.Lock()
	defer er.lock.Unlock()
	i := er.nextID
	er.nextID++
	er.viewActivated = append(er.viewActivated, listenerViewActivated{i, callback})
	return &eventListener{er, i | 0x9000000}
}

func (er *eventRegistry) RegisterViewCreated(callback func(view wicore.View)) wicore.EventListener {
	er.lock.Lock()
	defer er.lock.Unlock()
	i := er.nextID
	er.nextID++
	er.viewCreated = append(er.viewCreated, listenerViewCreated{i, callback})
	return &eventListener{er, i | 0xa000000}
}

func (er *eventRegistry) RegisterWindowCreated(callback func(window wicore.Window)) wicore.EventListener {
	er.lock.Lock()
	defer er.lock.Unlock()
	i := er.nextID
	er.nextID++
	er.windowCreated = append(er.windowCreated, listenerWindowCreated{i, callback})
	return &eventListener{er, i | 0xb000000}
}

func (er *eventRegistry) RegisterWindowResized(callback func(window wicore.Window)) wicore.EventListener {
	er.lock.Lock()
	defer er.lock.Unlock()
	i := er.nextID
	er.nextID++
	er.windowResized = append(er.windowResized, listenerWindowResized{i, callback})
	return &eventListener{er, i | 0xc000000}
}

func (er *eventRegistry) triggerCommands(cmds wicore.EnqueuedCommands) {
	er.deferred <- func() {
		items := func() []func(cmds wicore.EnqueuedCommands) {
			er.lock.Lock()
			defer er.lock.Unlock()
			items := make([]func(cmds wicore.EnqueuedCommands), 0, len(er.commands))
			for _, item := range er.commands {
				items = append(items, item.callback)
			}
			return items
		}()
		for _, item := range items {
			item(cmds)
		}
	}
}

func (er *eventRegistry) triggerDocumentCreated(doc wicore.Document) {
	er.deferred <- func() {
		items := func() []func(doc wicore.Document) {
			er.lock.Lock()
			defer er.lock.Unlock()
			items := make([]func(doc wicore.Document), 0, len(er.documentCreated))
			for _, item := range er.documentCreated {
				items = append(items, item.callback)
			}
			return items
		}()
		for _, item := range items {
			item(doc)
		}
	}
}

func (er *eventRegistry) triggerDocumentCursorMoved(doc wicore.Document, col, row int) {
	er.deferred <- func() {
		items := func() []func(doc wicore.Document, col, row int) {
			er.lock.Lock()
			defer er.lock.Unlock()
			items := make([]func(doc wicore.Document, col, row int), 0, len(er.documentCursorMoved))
			for _, item := range er.documentCursorMoved {
				items = append(items, item.callback)
			}
			return items
		}()
		for _, item := range items {
			item(doc, col, row)
		}
	}
}

func (er *eventRegistry) triggerEditorKeyboardModeChanged(mode wicore.KeyboardMode) {
	er.deferred <- func() {
		items := func() []func(mode wicore.KeyboardMode) {
			er.lock.Lock()
			defer er.lock.Unlock()
			items := make([]func(mode wicore.KeyboardMode), 0, len(er.editorKeyboardModeChanged))
			for _, item := range er.editorKeyboardModeChanged {
				items = append(items, item.callback)
			}
			return items
		}()
		for _, item := range items {
			item(mode)
		}
	}
}

func (er *eventRegistry) triggerEditorLanguage(l lang.Language) {
	er.deferred <- func() {
		items := func() []func(l lang.Language) {
			er.lock.Lock()
			defer er.lock.Unlock()
			items := make([]func(l lang.Language), 0, len(er.editorLanguage))
			for _, item := range er.editorLanguage {
				items = append(items, item.callback)
			}
			return items
		}()
		for _, item := range items {
			item(l)
		}
	}
}

func (er *eventRegistry) triggerTerminalKeyPressed(k key.Press) {
	er.deferred <- func() {
		items := func() []func(k key.Press) {
			er.lock.Lock()
			defer er.lock.Unlock()
			items := make([]func(k key.Press), 0, len(er.terminalKeyPressed))
			for _, item := range er.terminalKeyPressed {
				items = append(items, item.callback)
			}
			return items
		}()
		for _, item := range items {
			item(k)
		}
	}
}

func (er *eventRegistry) triggerTerminalMetaKeyPressed(k key.Press) {
	er.deferred <- func() {
		items := func() []func(k key.Press) {
			er.lock.Lock()
			defer er.lock.Unlock()
			items := make([]func(k key.Press), 0, len(er.terminalMetaKeyPressed))
			for _, item := range er.terminalMetaKeyPressed {
				items = append(items, item.callback)
			}
			return items
		}()
		for _, item := range items {
			item(k)
		}
	}
}

func (er *eventRegistry) triggerTerminalResized() {
	er.deferred <- func() {
		items := func() []func() {
			er.lock.Lock()
			defer er.lock.Unlock()
			items := make([]func(), 0, len(er.terminalResized))
			for _, item := range er.terminalResized {
				items = append(items, item.callback)
			}
			return items
		}()
		for _, item := range items {
			item()
		}
	}
}

func (er *eventRegistry) triggerViewActivated(view wicore.View) {
	er.deferred <- func() {
		items := func() []func(view wicore.View) {
			er.lock.Lock()
			defer er.lock.Unlock()
			items := make([]func(view wicore.View), 0, len(er.viewActivated))
			for _, item := range er.viewActivated {
				items = append(items, item.callback)
			}
			return items
		}()
		for _, item := range items {
			item(view)
		}
	}
}

func (er *eventRegistry) triggerViewCreated(view wicore.View) {
	er.deferred <- func() {
		items := func() []func(view wicore.View) {
			er.lock.Lock()
			defer er.lock.Unlock()
			items := make([]func(view wicore.View), 0, len(er.viewCreated))
			for _, item := range er.viewCreated {
				items = append(items, item.callback)
			}
			return items
		}()
		for _, item := range items {
			item(view)
		}
	}
}

func (er *eventRegistry) triggerWindowCreated(window wicore.Window) {
	er.deferred <- func() {
		items := func() []func(window wicore.Window) {
			er.lock.Lock()
			defer er.lock.Unlock()
			items := make([]func(window wicore.Window), 0, len(er.windowCreated))
			for _, item := range er.windowCreated {
				items = append(items, item.callback)
			}
			return items
		}()
		for _, item := range items {
			item(window)
		}
	}
}

func (er *eventRegistry) triggerWindowResized(window wicore.Window) {
	er.deferred <- func() {
		items := func() []func(window wicore.Window) {
			er.lock.Lock()
			defer er.lock.Unlock()
			items := make([]func(window wicore.Window), 0, len(er.windowResized))
			for _, item := range er.windowResized {
				items = append(items, item.callback)
			}
			return items
		}()
		for _, item := range items {
			item(window)
		}
	}
}

type unregister interface {
	unregister(id int)
}

type eventListener struct {
	unregister unregister
	id         int
}

func (e *eventListener) Close() error {
	if e.id == 0 {
		return errors.New("EventListener already closed")
	}
	e.unregister.unregister(e.id)
	e.id = 0
	return nil
}
