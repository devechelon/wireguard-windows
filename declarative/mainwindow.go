// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package declarative

import (
	"github.com/lxn/walk"
)

type MainWindow struct {
	AssignTo       **walk.MainWindow
	Name           string
	MinSize        Size
	MaxSize        Size
	ContextMenu    Menu
	Font           Font
	Title          string
	Size           Size
	Layout         Layout
	Children       []Widget
	MenuActions    []*walk.Action
	ToolBarActions []*walk.Action
}

func (mw MainWindow) Create(parent walk.Container) error {
	w, err := walk.NewMainWindow()
	if err != nil {
		return err
	}

	return InitWidget(mw, w, func() error {
		if err := w.SetTitle(mw.Title); err != nil {
			return err
		}

		if err := w.SetSize(mw.Size.toW()); err != nil {
			return err
		}

		if err := addToActionList(w.Menu().Actions(), mw.MenuActions); err != nil {
			return err
		}

		imageList, err := walk.NewImageList(walk.Size{16, 16}, 0)
		if err != nil {
			return err
		}
		w.ToolBar().SetImageList(imageList)

		if err := addToActionList(w.ToolBar().Actions(), mw.ToolBarActions); err != nil {
			return err
		}

		if mw.AssignTo != nil {
			*mw.AssignTo = w
		}

		return nil
	})
}

func (mw MainWindow) WidgetInfo() (name string, minSize, maxSize Size, stretchFactor, row, rowSpan, column, columnSpan int, contextMenu *Menu) {
	return mw.Name, mw.MinSize, mw.MaxSize, 0, 0, 0, 0, 0, &mw.ContextMenu
}

func (mw MainWindow) Font_() *Font {
	return &mw.Font
}

func (mw MainWindow) ContainerInfo() (Layout, []Widget) {
	return mw.Layout, mw.Children
}
