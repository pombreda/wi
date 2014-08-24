// Copyright 2014 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/maruel/wi/wi-plugin"
	"log"
	"time"
)

type CommandHandler func(c *command, cd wi.CommandDispatcherFull, w wi.Window, args ...string)

// command is the boilerplate wi.Command implementation.
type command struct {
	name      string
	handler   CommandHandler
	category  wi.CommandCategory
	shortDesc langMap
	longDesc  langMap
}

func (c *command) Name() string {
	return c.name
}

func (c *command) Handle(cd wi.CommandDispatcherFull, w wi.Window, args ...string) {
	c.handler(c, cd, w, args...)
}

func (c *command) Category(cd wi.CommandDispatcherFull, w wi.Window) wi.CommandCategory {
	return c.category
}

func (c *command) ShortDesc(cd wi.CommandDispatcherFull, w wi.Window) string {
	return getStr(cd.CurrentLanguage(), c.shortDesc)
}

func (c *command) LongDesc(cd wi.CommandDispatcherFull, w wi.Window) string {
	return getStr(cd.CurrentLanguage(), c.longDesc)
}

// commandAlias references another command by its name. It's important to not
// bind directly to the wi.Command reference, so that if a command is replaced
// by a plugin, that the replacement command is properly called by the alias.
type commandAlias struct {
	name    string
	command string
}

func (c *commandAlias) Name() string {
	return c.name
}

func (c *commandAlias) Handle(cd wi.CommandDispatcherFull, w wi.Window, args ...string) {
	// The alias is executed inline. This is important for command queue
	// ordering.
	cmd := wi.GetCommand(cd, w, c.command)
	if cmd != nil {
		cmd.Handle(cd, w, args...)
	} else {
		cmd = wi.GetCommand(cd, w, "alert")
		txt := fmt.Sprintf(getStr(cd.CurrentLanguage(), aliasNotFound), c.name, c.command)
		cmd.Handle(cd, w, txt)
	}
}

func (c *commandAlias) Category(cd wi.CommandDispatcherFull, w wi.Window) wi.CommandCategory {
	cmd := wi.GetCommand(cd, w, c.command)
	if cmd != nil {
		return c.Category(cd, w)
	}
	return wi.UnknownCategory
}

func (c *commandAlias) ShortDesc(cd wi.CommandDispatcherFull, w wi.Window) string {
	return fmt.Sprintf(getStr(cd.CurrentLanguage(), aliasFor), c.command)
}

func (c *commandAlias) LongDesc(cd wi.CommandDispatcherFull, w wi.Window) string {
	return fmt.Sprintf(getStr(cd.CurrentLanguage(), aliasFor), c.command)
}

type commands struct {
	commands map[string]wi.Command
}

func (c *commands) Register(cmd wi.Command) bool {
	name := cmd.Name()
	_, ok := c.commands[name]
	c.commands[name] = cmd
	return !ok
}

func (c *commands) Get(cmd string) wi.Command {
	return c.commands[cmd]
}

func makeCommands() wi.Commands {
	return &commands{make(map[string]wi.Command)}
}

// Default commands

func cmdAlert(c *command, cd wi.CommandDispatcherFull, w wi.Window, args ...string) {
	// TODO(maruel): Create an infobar that automatically dismiss itself after 5s.
	if len(args) != 1 {
		cd.ExecuteCommand(w, "alert", c.LongDesc(cd, w))
		return
	}
	// TODO(maruel): Use a 5 seconds infobar.
	wi.RootWindow(w).NewChildWindow(makeAlertView(args[0]), wi.DockingFloating)
	log.Printf("Tree:\n%s", wi.RootWindow(w).Tree())
	//w2.Activate()
	go func() {
		<-time.After(5 * time.Second)
		// TODO(maruel): Dismiss.
	}()
}

func cmdAddStatusBar(c *command, cd wi.CommandDispatcherFull, w wi.Window, args ...string) {
	// Create a tree of views that is used for alignment.
	if len(args) != 0 {
		cd.ExecuteCommand(w, "alert", c.LongDesc(cd, w))
		return
	}
	statusWindowRoot := w.NewChildWindow(makeStatusViewRoot(), wi.DockingBottom)
	statusWindowRoot.NewChildWindow(makeStatusViewName(), wi.DockingLeft)
	statusWindowRoot.NewChildWindow(makeStatusViewPosition(), wi.DockingRight)
}

func cmdOpen(c *command, cd wi.CommandDispatcherFull, w wi.Window, args ...string) {
	// The Window and View are created synchronously. The View is populated
	// asynchronously.
	log.Printf("Faking opening a file: %s", args)
}

func cmdNew(c *command, cd wi.CommandDispatcherFull, w wi.Window, args ...string) {
	if len(args) != 0 {
		cmd := wi.GetCommand(cd, w, "alias")
		cd.ExecuteCommand(w, "alert", cmd.LongDesc(cd, w))
	} else {
		w.NewChildWindow(makeView("New doc", -1, -1), wi.DockingFill)
	}
}

func cmdShell(c *command, cd wi.CommandDispatcherFull, w wi.Window, args ...string) {
	log.Printf("Faking opening a new shell: %s", args)
}

func cmdDoc(c *command, cd wi.CommandDispatcherFull, w wi.Window, args ...string) {
	// TODO(maruel): Grab the current word under selection if no args is
	// provided. Pass this token to shell.
	docArgs := make([]string, len(args)+1)
	docArgs[0] = "doc"
	copy(docArgs[1:], args)
	//dispatcher.Execute(w, "shell", docArgs...)
}

func isDirtyRecurse(cd wi.CommandDispatcherFull, w wi.Window) bool {
	for _, child := range w.ChildrenWindows() {
		if isDirtyRecurse(cd, child) {
			return true
		}
		v := child.View()
		if v.IsDirty() {
			cd.ExecuteCommand(w, "alert", fmt.Sprintf(getStr(cd.CurrentLanguage(), viewDirty), v.Title()))
			return true
		}
	}
	return false
}

func cmdQuit(c *command, cd wi.CommandDispatcherFull, w wi.Window, args ...string) {
	// TODO(maruel): For all the View, question if fine to quit via
	// view.IsDirty(). If not fine, "prompt" y/n to force quit. If n, stop there.
	// - Send a signal to each plugin.
	// - Send a signal back to the main loop.
	root := wi.RootWindow(w)
	if !isDirtyRecurse(cd, root) {
		quitFlag = true
		// PostDraw wakes up the command event loop so it detects it's time to
		// quit.
		cd.PostDraw()
	}
}

func cmdHelp(c *command, cd wi.CommandDispatcherFull, w wi.Window, args ...string) {
	// TODO(maruel): Creates a new Window with a ViewHelp.
	log.Printf("Faking help: %s", args)
}

func cmdAlias(c *command, cd wi.CommandDispatcherFull, w wi.Window, args ...string) {
	if len(args) != 3 {
		cmd := wi.GetCommand(cd, w, "alias")
		cd.ExecuteCommand(w, "alert", cmd.LongDesc(cd, w))
		return
	}
	if args[0] == "window" {
	} else if args[0] == "global" {
		w = wi.RootWindow(w)
	} else {
		cmd := wi.GetCommand(cd, w, "alias")
		cd.ExecuteCommand(w, "alert", cmd.LongDesc(cd, w))
		return
	}
	alias := &commandAlias{args[1], args[2]}
	w.View().Commands().Register(alias)
}

func cmdKeyBind(c *command, cd wi.CommandDispatcherFull, w wi.Window, args ...string) {
	if len(args) != 4 {
		cmd := wi.GetCommand(cd, w, "keybind")
		cd.ExecuteCommand(w, "alert", cmd.LongDesc(cd, w))
		return
	}
	location := args[0]
	modeName := args[1]
	keyName := args[2]
	cmdName := args[3]

	if location == "global" {
		w = wi.RootWindow(w)
	} else if location != "window" {
		cmd := wi.GetCommand(cd, w, "keybind")
		cd.ExecuteCommand(w, "alert", cmd.LongDesc(cd, w))
		return
	}

	var mode wi.KeyboardMode
	if modeName == "command" {
		mode = wi.CommandMode
	} else if modeName == "edit" {
		mode = wi.CommandMode
	} else if modeName == "all" {
		mode = wi.AllMode
	} else {
		cmd := wi.GetCommand(cd, w, "keybind")
		cd.ExecuteCommand(w, "alert", cmd.LongDesc(cd, w))
		return
	}
	w.View().KeyBindings().Set(mode, keyName, cmdName)
}

func cmdShowCommandWindow(c *command, cd wi.CommandDispatcherFull, w wi.Window, args ...string) {
	if len(args) != 0 {
		cmd := wi.GetCommand(cd, w, "show_command_window")
		cd.ExecuteCommand(w, "alert", cmd.LongDesc(cd, w))
		return
	}

	// Create the Window with the command view and attach it to the currently
	// focused Window.
	cmdWindow := makeCommandView()
	w.NewChildWindow(cmdWindow, wi.DockingFloating)
}

func cmdLogWindowTree(c *command, cd wi.CommandDispatcherFull, w wi.Window, args ...string) {
	if len(args) != 0 {
		cmd := wi.GetCommand(cd, w, "log_window_tree")
		cd.ExecuteCommand(w, "alert", cmd.LongDesc(cd, w))
		return
	}
	root := wi.RootWindow(w)
	log.Printf("Window tree:\n%s", root.Tree())
}

// Native commands.
var defaultCommands = []wi.Command{
	&command{
		"alert",
		cmdAlert,
		wi.WindowCategory,
		langMap{
			wi.LangEn: "Shows a modal message",
		},
		langMap{
			wi.LangEn: "Prints a message in a modal dialog box.",
		},
	},
	&command{
		"add_status_bar",
		cmdAddStatusBar,
		wi.WindowCategory,
		langMap{
			wi.LangEn: "Adds the standard status bar",
		},
		langMap{
			wi.LangEn: "Adds the standard status bar to the active window. This command exists so it can be overriden by a plugin, so it can create its own status bar.",
		},
	},
	&command{
		"help",
		cmdHelp,
		wi.WindowCategory,
		langMap{
			wi.LangEn: "Prints help",
		},
		langMap{
			wi.LangEn: "Prints general help or help for a particular command.",
		},
	},
	&command{
		"new",
		cmdNew,
		wi.WindowCategory,
		langMap{
			wi.LangEn: "Create a new buffer",
		},
		langMap{
			wi.LangEn: "Create a new buffer.",
		},
	},
	&command{
		"open",
		cmdOpen,
		wi.WindowCategory,
		langMap{
			wi.LangEn: "Opens a file in a new buffer",
		},
		langMap{
			wi.LangEn: "Opens a file in a new buffer.",
		},
	},
	&command{
		"quit",
		cmdQuit,
		wi.WindowCategory,
		langMap{
			wi.LangEn: "Quits",
		},
		langMap{
			wi.LangEn: "Quits the editor. Optionally bypasses writing the files to disk.",
		},
	},
	&command{
		"shell",
		cmdShell,
		wi.WindowCategory,
		langMap{
			wi.LangEn: "Opens a shell process",
		},
		langMap{
			wi.LangEn: "Opens a shell process in a new buffer.",
		},
	},
	&command{
		"doc",
		cmdDoc,
		wi.WindowCategory,
		langMap{
			wi.LangEn: "Search godoc documentation",
		},
		langMap{
			wi.LangEn: "Uses the 'doc' tool to get documentation about the text under the cursor.",
		},
	},

	&command{
		"alias",
		cmdAlias,
		wi.CommandsCategory,
		langMap{
			wi.LangEn: "Binds an alias to another command",
		},
		langMap{
			// TODO(maruel): For complex aliasing, use macro?
			wi.LangEn: "Usage: alias [window|global] <alias> <name>\nBinds an alias to another command. The alias can either be local to the window or global",
		},
	},
	&command{
		"keybind",
		cmdKeyBind,
		wi.CommandsCategory,
		langMap{
			wi.LangEn: "Binds a keyboard mapping to a command",
		},
		langMap{
			wi.LangEn: "Usage: keybind [window|global] [command|edit|all] <key> <command>\nBinds a keyboard mapping to a command. The binding can be to the active view for view-specific key binding or to the root view for global key bindings.",
		},
	},
	&command{
		"show_command_window",
		cmdShowCommandWindow,
		wi.CommandsCategory,
		langMap{
			wi.LangEn: "Shows the interactive command window",
		},
		langMap{
			wi.LangEn: "This commands exists so it can be bound to a key to pop up the interactive command window.",
		},
	},

	&command{
		"log_window_tree",
		cmdLogWindowTree,
		wi.DebugCategory,
		langMap{
			wi.LangEn: "Logs the tree in the log file",
		},
		langMap{
			wi.LangEn: "Logs the tree in the log file, this is only relevant if -verbose is used.",
		},
	},

	&commandAlias{"q", "quit"},

	// DIRECTION = up/down/left/right
	// window_DIRECTION
	// window_close
	// cursor_move_DIRECTION
	// add_text/insert/delete
	// undo/redo
	// verb/movement/multiplier
	// Modes, select (both column and normal), command.
	// 'screenshot', mainly for unit test; open a new buffer with the screenshot, so it can be saved with 'w'.
	// ...
}

// RegisterDefaultCommands registers the top-level native commands. This
// includes the window management commands, opening a new file buffer (it's a
// text editor after all) and help, quitting, etc. It doesn't includes handling
// a file buffer itself, it's up to the relevant view to add the corresponding
// commands. For example, "open" is implemented but "write" is not!
func RegisterDefaultCommands(dispatcher wi.Commands) {
	for _, cmd := range defaultCommands {
		dispatcher.Register(cmd)
	}
}
