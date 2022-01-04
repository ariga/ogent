package ogent

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"

	"github.com/ariga/ogent/internal/integration/ogent/ent"
	"github.com/ariga/ogent/internal/integration/ogent/ent/enttest"
	"github.com/ariga/ogent/internal/integration/ogent/ent/ogent"
)

func TestOgent(t *testing.T) {
	suite.Run(t, new(testSuite))
}

type testSuite struct {
	suite.Suite
	client  *ent.Client
	handler ogent.Handler
}

func (t *testSuite) SetupTest() {
	t.client = enttest.Open(t.T(), dialect.SQLite, fmt.Sprintf("file:ogent_%s?mode=memory&cache=shared&_fk=1", time.Now()))
	t.handler = ogent.NewOgentHandler(t.client)
}

func (t *testSuite) TestCreate() {
	// R409
	got, err := t.handler.CreatePet(context.Background(), ogent.CreatePetReq{})
	t.Require().NoError(err)
	t.reqErr(http.StatusConflict, got)

	// OK
	owner := t.client.User.Create().SetName("Ariel").SetAge(33).SaveX(context.Background())
	got, err = t.handler.CreatePet(context.Background(), ogent.CreatePetReq{
		Name:       "Ariels most loved Leopard",
		Weight:     ogent.NewOptInt(10),
		Birthday:   ogent.NewOptTime(time.Now()),
		Categories: nil,
		Owner:      owner.ID,
		Friends:    nil,
	})
	t.Require().NoError(err)
	t.Require().Equal(ogent.NewPetCreate(t.client.Pet.Query().WithOwner().FirstX(context.Background())), got)
}

func (t *testSuite) TestRead() {
	// R400
	got, err := t.handler.ReadPet(context.Background(), ogent.ReadPetParams{ID: 2000})
	t.Require().NoError(err)
	t.reqErr(http.StatusNotFound, got)

	// OK
	owner := t.client.User.Create().SetName("Ariel").SetAge(33).SaveX(context.Background())
	pet := t.client.Pet.Create().SetName("First Pet").SetOwner(owner).SaveX(context.Background())
	pet.Edges.Owner = owner
	got, err = t.handler.ReadPet(context.Background(), ogent.ReadPetParams{ID: 1})
	t.Require().NoError(err)
	t.Require().Equal(ogent.NewPetRead(pet), got)
}

func (t *testSuite) TestUpdate() {
	// R404
	got, err := t.handler.UpdatePet(context.Background(), ogent.UpdatePetReq{}, ogent.UpdatePetParams{ID: 2000})
	t.Require().NoError(err)
	t.reqErr(http.StatusNotFound, got)

	// OK
	owner := t.client.User.Create().SetName("Ariel").SetAge(33).SaveX(context.Background())
	pet := t.client.Pet.Create().SetName("First Pet").SetOwner(owner).SaveX(context.Background())
	pet.Edges.Owner = owner
	got, err = t.handler.UpdatePet(context.Background(), ogent.UpdatePetReq{Name: ogent.NewOptString("The changed name")}, ogent.UpdatePetParams{ID: pet.ID})
	pet.Name = "The changed name"
	t.Require().NoError(err)
	t.Require().Equal(ogent.NewPetUpdate(pet), got)
}

func (t *testSuite) reqErr(c int, err interface{}) {
	var (
		ac int
		at string
	)
	switch c {
	case http.StatusBadRequest:
		t.Require().IsType(new(ogent.R400), err)
		err := err.(*ogent.R400)
		ac, at = err.Code, err.Status
	case http.StatusNotFound:
		t.Require().IsType(new(ogent.R404), err)
		err := err.(*ogent.R404)
		ac, at = err.Code, err.Status
	case http.StatusConflict:
		t.Require().IsType(new(ogent.R409), err)
		err := err.(*ogent.R409)
		ac, at = err.Code, err.Status
	default:
		t.Failf("panic reqErr", "unknown status code: %d", c)
	}
	t.Require().Equal(c, ac, err)
	t.Require().Equal(http.StatusText(c), at, err)
}
