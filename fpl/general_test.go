package fpl

import (
	"testing"
)

func TestGetGeneral(t *testing.T) {

	c := NewClient(nil)
	got, err := c.GetGeneral()
	if err != nil {
		t.Fatal(err)
	}
	if got == nil {
		t.Errorf("Could do not be succeed, Got: %+v", got)
	}

}

func TestListTeamInfo(t *testing.T) {

	c := NewClient(nil)

	got, err := c.ListClubsInfo()
	if err != nil {
		t.Fatal(err)
	}

	if got == nil {
		t.Errorf("Could do not be succeed, Got: %+v", got)
	}

}

func TestListEventInfo(t *testing.T) {

	c := NewClient(nil)

	got, err := c.ListEventInfo()
	if err != nil {
		t.Fatal(err)
	}

	if got == nil {
		t.Errorf("Could do not be succeed, Got: %+v", got)
	}
}

func TestListElementStatsInfo(t *testing.T) {

	c := NewClient(nil)

	got, err := c.ListElementStatsInfo()
	if err != nil {
		t.Fatal(err)
	}

	if got == nil {
		t.Errorf("Could do not be succeed, Got: %+v", got)
	}
}

func TestListElementTypesInfo(t *testing.T) {

	c := NewClient(nil)

	got, err := c.ListElementTypesInfo()
	if err != nil {
		t.Fatal(err)
	}

	if got == nil {
		t.Errorf("Could do not be succeed, Got: %+v", got)
	}
}

func TestListSettings(t *testing.T) {

	c := NewClient(nil)

	got, err := c.ListSettings()
	if err != nil {
		t.Fatal(err)
	}

	if got == nil {
		t.Errorf("Could do not be succeed, Got: %+v", got)
	}
}
