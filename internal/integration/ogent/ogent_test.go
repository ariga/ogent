package ogent

import (
	"context"
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
	t.client = enttest.Open(t.T(), dialect.SQLite, "file:ogent?mode=memory&cache=shared&_fk=1")
	t.handler = ogent.NewOgentHandler(t.client)
}

func (t *testSuite) TestCreate() {
	// R400
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
	t.Require().IsType(new(ogent.PetCreate), got)
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
	case http.StatusConflict:
		t.Require().IsType(new(ogent.R409), err)
		err := err.(*ogent.R409)
		ac, at = err.Code, err.Status
	default:
		panic("unimplemented")
	}
	t.Require().Equal(c, ac, err)
	t.Require().Equal(http.StatusText(c), at, err)
}
