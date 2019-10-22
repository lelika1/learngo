package main

import (
	"testing"
)

var emptyMapStr = "Map is empty"

var filledMapStr = `(k: 100, v: 100)
(k: 50, v: 50)
(k: 40, v: 40)
(k: 20, v: 20)
(k: 30, v: 30)
(k: 60, v: 60)
(k: 80, v: 80)
(k: 200, v: 200)
(k: 150, v: 150)
(k: 300, v: 300)
(k: 250, v: 250)
(k: 255, v: 255)
(k: 400, v: 400)
(k: 350, v: 350)
(k: 320, v: 320)
(k: 450, v: 450)
`

var mapLeafDelStr = `(k: 100, v: 100)
(k: 50, v: 50)
(k: 40, v: 40)
(k: 20, v: 20)
(k: 60, v: 60)
(k: 80, v: 80)
(k: 200, v: 200)
(k: 150, v: 150)
(k: 300, v: 300)
(k: 250, v: 250)
(k: 255, v: 255)
(k: 400, v: 400)
(k: 350, v: 350)
(k: 320, v: 320)
(k: 450, v: 450)
`

var mapHalfNodesDelStr = `(k: 100, v: 100)
(k: 50, v: 50)
(k: 30, v: 30)
(k: 60, v: 60)
(k: 80, v: 80)
(k: 200, v: 200)
(k: 150, v: 150)
(k: 300, v: 300)
(k: 250, v: 250)
(k: 255, v: 255)
(k: 400, v: 400)
(k: 350, v: 350)
(k: 320, v: 320)
(k: 450, v: 450)
`

var mapFullNodesDelStr = `(k: 100, v: 100)
(k: 60, v: 60)
(k: 40, v: 40)
(k: 20, v: 20)
(k: 30, v: 30)
(k: 80, v: 80)
(k: 250, v: 250)
(k: 150, v: 150)
(k: 320, v: 320)
(k: 255, v: 255)
(k: 400, v: 400)
(k: 350, v: 350)
(k: 450, v: 450)
`

var mapAfterChangesStr = `(k: 100, v: 100)
(k: 50, v: 50)
(k: 40, v: 40)
(k: 20, v: 20)
(k: 30, v: 33333)
(k: 60, v: 60)
(k: 80, v: 80)
(k: 200, v: 22222)
(k: 150, v: 150)
(k: 300, v: 300)
(k: 250, v: 250)
(k: 255, v: 255)
(k: 400, v: 400)
(k: 350, v: 350)
(k: 320, v: 320)
(k: 450, v: 450)
`

func initMap(m *Map) {
	m.Insert(100, 100)
	m.Insert(50, 50)
	m.Insert(40, 40)
	m.Insert(60, 60)
	m.Insert(20, 20)
	m.Insert(30, 30)
	m.Insert(80, 80)
	m.Insert(200, 200)
	m.Insert(150, 150)
	m.Insert(300, 300)
	m.Insert(250, 250)
	m.Insert(255, 255)
	m.Insert(400, 400)
	m.Insert(450, 450)
	m.Insert(350, 350)
	m.Insert(320, 320)
}

func TestFillMap(t *testing.T) {
	var m Map
	if m.String() != emptyMapStr {
		t.Errorf("Error. Map must be empty.")
	}

	initMap(&m)
	if m.String() != filledMapStr {
		t.Errorf("Error. Results are not the same.")
	}
}

func TestRemoveLeaf(t *testing.T) {
	var m Map
	initMap(&m)
	m.Rm(30)

	if m.String() != mapLeafDelStr {
		t.Errorf("Error. Results are not the same.")
	}
}

func TestRemoveHalfNode(t *testing.T) {
	var m Map
	initMap(&m)
	m.Rm(20)
	m.Rm(40)

	if m.String() != mapHalfNodesDelStr {
		t.Errorf("Error. Results are not the same.")
	}
}

func TestRemoveFullNode(t *testing.T) {
	var m Map
	initMap(&m)
	m.Rm(50)
	m.Rm(300)
	m.Rm(200)

	if m.String() != mapFullNodesDelStr {
		t.Errorf("Error. Results are not the same.")
	}
}

func TestRemoveRootNode(t *testing.T) {
	var m Map
	m.Insert(5, 5)
	m.Insert(7, 7)
	m.Insert(6, 6)
	m.Insert(8, 8)

	m.Rm(5)
	if m.String() != "(k: 7, v: 7)\n(k: 6, v: 6)\n(k: 8, v: 8)\n" {
		t.Errorf("Error. Results are not the same.")
	}

	m.Rm(7)
	if m.String() != "(k: 8, v: 8)\n(k: 6, v: 6)\n" {
		t.Errorf("Error. Results are not the same.")
	}

	m.Rm(8)
	if m.String() != "(k: 6, v: 6)\n" {
		t.Errorf("Error. Results are not the same.")
	}

	m.Rm(6)
	if m.String() != emptyMapStr {
		t.Errorf("Error. Map must be empty.")
	}

}

func TestFindAndChangedElemInMap(t *testing.T) {
	var m Map
	initMap(&m)

	if _, ok := m.Find(1); ok {
		t.Errorf("Error. There is no key=1 in map")
	}

	v, ok := m.Find(200)
	if !ok {
		t.Errorf("Error. Key=200 must be in map")
	}
	if *v != 200 {
		t.Errorf("Error. Value for key=200 must be equal 200.")
	}
	*v = 22222

	v, ok = m.Find(30)
	if !ok {
		t.Errorf("Error. Key=30 must be in map")
	}
	if *v != 30 {
		t.Errorf("Error. Value for key=30 must be equal 30.")
	}
	*v = 33333

	if m.String() != mapAfterChangesStr {
		t.Errorf("Error. Results are not the same.")
	}
}
