package ogent

import (
	"fmt"
	"net/http/httptest"
	"testing"

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
	client *ent.Client
	srv    *ogent.Server
}

func (t *testSuite) SetupTest() {
	t.client = enttest.Open(t.T(), dialect.SQLite, "file:ogent?mode=memory&cache=shared&_fk=1")
	t.srv = ogent.NewServer(ogent.NewOgentHandler(t.client))
}

func (t *testSuite) TestCreate() {
	s := httptest.NewServer(t.srv)
	defer s.Close()

	c := s.Client()

	got, err := c.Post(s.URL+"/pets", "application/json", nil)
	t.Require().NoError(err)
	fmt.Println(got)
}
