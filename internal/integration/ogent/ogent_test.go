package ogent

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"testing"
	"time"

	"ariga.io/ogent/internal/integration/ogent/ent/user"
	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"

	"ariga.io/ogent/internal/integration/ogent/ent"
	"ariga.io/ogent/internal/integration/ogent/ent/enttest"
	"ariga.io/ogent/internal/integration/ogent/ent/ogent"
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
	t.client = enttest.Open(t.T(), dialect.SQLite, fmt.Sprintf("file:ogent_%d?mode=memory&cache=shared&_fk=1", time.Now().UnixNano()))
	t.handler = ogent.NewOgentHandler(t.client)
}

func (t *testSuite) TestCreate() {
	// R409
	got, err := t.handler.CreatePet(context.Background(), ogent.CreatePetReq{})
	t.Require().NoError(err)
	t.reqErr(http.StatusConflict, got)

	// OK
	owner := t.client.User.Create().SetName("Ariel").SetAge(33).SetFavoriteCatBreed(user.FavoriteCatBreedLeopard).SaveX(context.Background())
	got, err = t.handler.CreatePet(context.Background(), ogent.CreatePetReq{
		Name:       "Ariels most loved Leopard",
		Weight:     ogent.NewOptInt(10),
		Birthday:   ogent.NewOptDateTime(time.Now()),
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
	owner := t.client.User.Create().SetName("Ariel").SetAge(33).SetFavoriteCatBreed(user.FavoriteCatBreedLeopard).SaveX(context.Background())
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
	owner := t.client.User.Create().SetName("Ariel").SetAge(33).SetFavoriteCatBreed(user.FavoriteCatBreedLeopard).SaveX(context.Background())
	pet := t.client.Pet.Create().SetName("First Pet").SetOwner(owner).SaveX(context.Background())
	pet.Edges.Owner = owner
	got, err = t.handler.UpdatePet(context.Background(), ogent.UpdatePetReq{Name: ogent.NewOptString("The changed name")}, ogent.UpdatePetParams{ID: pet.ID})
	pet.Name = "The changed name"
	t.Require().NoError(err)
	t.Require().Equal(ogent.NewPetUpdate(pet), got)
}

func (t *testSuite) TestDelete() {
	// R404
	got, err := t.handler.DeleteUser(context.Background(), ogent.DeleteUserParams{ID: 2000})
	t.Require().NoError(err)
	t.reqErr(http.StatusNotFound, got)

	// OK TODO(masseelch): This should fail with a foreign key exception once https://github.com/ent/ent/pull/1703 gets merged.
	owner := t.client.User.Create().SetName("Ariel").SetAge(33).SetFavoriteCatBreed(user.FavoriteCatBreedLeopard).SaveX(context.Background())
	got, err = t.handler.DeleteUser(context.Background(), ogent.DeleteUserParams{ID: owner.ID})
	t.Require().NoError(err)
	t.Require().Equal(new(ogent.DeleteUserNoContent), got)
}

func (t *testSuite) TestList() {
	// Add some entities.
	b := make([]*ent.CategoryCreate, 50)
	for i := range b {
		b[i] = t.client.Category.Create().SetName("Category " + strconv.Itoa(i+1))
	}
	es := t.client.Category.CreateBulk(b...).SaveX(context.Background())

	// Default page size.
	got, err := t.handler.ListCategory(context.Background(), ogent.ListCategoryParams{})
	t.Require().NoError(err)
	r := ogent.NewCategoryLists(es[0:30])
	t.Require().Equal((*ogent.ListCategoryOKApplicationJSON)(&r), got)

	// Custom page size.
	got, err = t.handler.ListCategory(context.Background(), ogent.ListCategoryParams{ItemsPerPage: ogent.NewOptInt(10)})
	t.Require().NoError(err)
	r = ogent.NewCategoryLists(es[0:10])
	t.Require().Equal((*ogent.ListCategoryOKApplicationJSON)(&r), got)

	// Custom page.
	got, err = t.handler.ListCategory(context.Background(), ogent.ListCategoryParams{Page: ogent.NewOptInt(2)})
	t.Require().NoError(err)
	r = ogent.NewCategoryLists(es[30:50])
	t.Require().Equal((*ogent.ListCategoryOKApplicationJSON)(&r), got)

	// Custom page and page size.
	got, err = t.handler.ListCategory(context.Background(), ogent.ListCategoryParams{Page: ogent.NewOptInt(2), ItemsPerPage: ogent.NewOptInt(10)})
	t.Require().NoError(err)
	r = ogent.NewCategoryLists(es[10:20])
	t.Require().Equal((*ogent.ListCategoryOKApplicationJSON)(&r), got)

	//Order By
	got, err = t.handler.ListCategory(context.Background(), ogent.ListCategoryParams{OrderBy: ogent.NewOptString("id desc")})
	t.Require().NoError(err)
	r = ogent.NewCategoryLists(es[20:50])
	sort.SliceStable(r, func(i, j int) bool {
		return r[i].ID > r[j].ID
	})
	t.Require().Equal((*ogent.ListCategoryOKApplicationJSON)(&r), got)
}

func (t *testSuite) TestReadSub() {
	// R404 - parent not found
	got, err := t.handler.ReadUserBestFriend(context.Background(), ogent.ReadUserBestFriendParams{})
	t.Require().NoError(err)
	t.reqErr(http.StatusNotFound, got)

	// R404 - no attached resource
	ariel := t.client.User.Create().SetName("Ariel").SetAge(33).SetFavoriteCatBreed(user.FavoriteCatBreedLeopard).SaveX(context.Background())
	got, err = t.handler.ReadUserBestFriend(context.Background(), ogent.ReadUserBestFriendParams{ID: ariel.ID})
	t.Require().NoError(err)
	t.reqErr(http.StatusNotFound, got)

	// OK
	elch := t.client.User.Create().SetName("MasseElch").SetAge(31).SetFavoriteCatBreed(user.FavoriteCatBreedLeopard).SetBestFriend(ariel).SaveX(context.Background())
	got, err = t.handler.ReadUserBestFriend(context.Background(), ogent.ReadUserBestFriendParams{ID: ariel.ID})
	t.Require().NoError(err)
	t.Require().Equal(ogent.NewUserBestFriendRead(elch), got)
}

func (t *testSuite) TestListSub() {
	// OK - parent not found
	got, err := t.handler.ListPetFriends(context.Background(), ogent.ListPetFriendsParams{})
	t.Require().NoError(err)
	var pf []*ent.Pet
	r := ogent.NewPetFriendsLists(pf)
	t.Require().Equal((*ogent.ListPetFriendsOKApplicationJSON)(&r), got)

	owner := t.client.User.Create().SetName("Ariel").SetAge(33).SetFavoriteCatBreed(user.FavoriteCatBreedLeopard).SaveX(context.Background())

	// OK - no attached resource
	loner := t.client.Pet.Create().SetName("Lonely Wolf").SetOwner(owner).SaveX(context.Background())
	got, err = t.handler.ListPetFriends(context.Background(), ogent.ListPetFriendsParams{ID: loner.ID})
	t.Require().NoError(err)
	t.Require().Equal((*ogent.ListPetFriendsOKApplicationJSON)(&r), got)

	// Get the Lonely Wolf some friends.
	b := make([]*ent.PetCreate, 50)
	for i := range b {
		b[i] = t.client.Pet.Create().SetName("Pet " + strconv.Itoa(i+1)).SetOwner(owner).AddFriends(loner)
	}
	es := t.client.Pet.CreateBulk(b...).SaveX(context.Background())

	// Default page size.
	got, err = t.handler.ListPetFriends(context.Background(), ogent.ListPetFriendsParams{ID: loner.ID})
	t.Require().NoError(err)
	r = ogent.NewPetFriendsLists(es[0:30])
	t.Require().Equal((*ogent.ListPetFriendsOKApplicationJSON)(&r), got)

	// Custom page size.
	got, err = t.handler.ListPetFriends(context.Background(), ogent.ListPetFriendsParams{ID: loner.ID, ItemsPerPage: ogent.NewOptInt(10)})
	t.Require().NoError(err)
	r = ogent.NewPetFriendsLists(es[0:10])
	t.Require().Equal((*ogent.ListPetFriendsOKApplicationJSON)(&r), got)

	// Custom page.
	got, err = t.handler.ListPetFriends(context.Background(), ogent.ListPetFriendsParams{ID: loner.ID, Page: ogent.NewOptInt(2)})
	t.Require().NoError(err)
	r = ogent.NewPetFriendsLists(es[30:50])
	t.Require().Equal((*ogent.ListPetFriendsOKApplicationJSON)(&r), got)

	// Custom page and page size.
	got, err = t.handler.ListPetFriends(context.Background(), ogent.ListPetFriendsParams{ID: loner.ID, Page: ogent.NewOptInt(2), ItemsPerPage: ogent.NewOptInt(10)})
	t.Require().NoError(err)
	r = ogent.NewPetFriendsLists(es[10:20])
	t.Require().Equal((*ogent.ListPetFriendsOKApplicationJSON)(&r), got)
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

func (t *testSuite) TestParseFilterToSQL() {
	f, e := preparePetFilterToSQL()
	t.Equal(len(f), len(e), "filter and expected list must have the same length")
	i := 0
	n := len(f)
	for i < n {
		q := t.client.Pet.Query()
		exp := &ogent.Expression{}
		err := ogent.SqlParser.ParseString(f[i], exp)
		t.Require().NoError(err, "fail to parse the filter: "+f[i])
		q.Where(func(s *sql.Selector) {
			ogent.ParseExpression(exp, s)
			query, _ := s.Query()
			t.Equal(e[i], query, "actual query created is not the expected one")
		})
		_, err = q.All(context.Background())
		t.Require().NoError(err, "fail to execute the query")
		i++
	}
}

func preparePetFilterToSQL() ([]string, []string) {
	var f []string //filter
	var e []string //expected
	//EQ
	f = append(f, "name eq 'Ariel'")
	e = append(e, "SELECT `pets`.`id`, `pets`.`name`, `pets`.`weight`, `pets`.`birthday` FROM `pets` WHERE `pets`.`name` = ?")
	//NE
	f = append(f, "name ne 'Milk'")
	e = append(e, "SELECT `pets`.`id`, `pets`.`name`, `pets`.`weight`, `pets`.`birthday` FROM `pets` WHERE `pets`.`name` <> ?")
	//GT
	f = append(f, "weight gt 2")
	e = append(e, "SELECT `pets`.`id`, `pets`.`name`, `pets`.`weight`, `pets`.`birthday` FROM `pets` WHERE `pets`.`weight` > ?")
	//GE
	f = append(f, "weight ge 2")
	e = append(e, "SELECT `pets`.`id`, `pets`.`name`, `pets`.`weight`, `pets`.`birthday` FROM `pets` WHERE `pets`.`weight` >= ?")
	//LT
	f = append(f, "weight lt 2")
	e = append(e, "SELECT `pets`.`id`, `pets`.`name`, `pets`.`weight`, `pets`.`birthday` FROM `pets` WHERE `pets`.`weight` < ?")
	//LE
	f = append(f, "weight le 2")
	e = append(e, "SELECT `pets`.`id`, `pets`.`name`, `pets`.`weight`, `pets`.`birthday` FROM `pets` WHERE `pets`.`weight` <= ?")
	//AND
	f = append(f, "name eq 'Ariel' and weight le 2")
	e = append(e, "SELECT `pets`.`id`, `pets`.`name`, `pets`.`weight`, `pets`.`birthday` FROM `pets` WHERE `pets`.`name` = ? AND `pets`.`weight` <= ?")
	//OR
	f = append(f, "name eq 'Ariel' or weight le 2")
	e = append(e, "SELECT `pets`.`id`, `pets`.`name`, `pets`.`weight`, `pets`.`birthday` FROM `pets` WHERE `pets`.`name` = ? OR `pets`.`weight` <= ?")
	//AND Recursive
	f = append(f, "(name eq 'Ariel' or name eq 'Milk') and weight le 2")
	e = append(e, "SELECT `pets`.`id`, `pets`.`name`, `pets`.`weight`, `pets`.`birthday` FROM `pets` WHERE (`pets`.`name` = ? OR `pets`.`name` = ?) AND `pets`.`weight` <= ?")
	//AND Recursive
	f = append(f, "weight le 2 and (name eq 'Ariel' or name eq 'Milk')")
	e = append(e, "SELECT `pets`.`id`, `pets`.`name`, `pets`.`weight`, `pets`.`birthday` FROM `pets` WHERE (`pets`.`weight` <= ? AND `pets`.`name` = ?) OR `pets`.`name` = ?")
	//OR Recursive
	f = append(f, "(name eq 'Ariel' or name eq 'Milk') or weight le 2")
	e = append(e, "SELECT `pets`.`id`, `pets`.`name`, `pets`.`weight`, `pets`.`birthday` FROM `pets` WHERE (`pets`.`name` = ? OR `pets`.`name` = ?) OR `pets`.`weight` <= ?")
	//OR Recursive
	f = append(f, "weight le 2 or (name eq 'Ariel' or name eq 'Milk')")
	e = append(e, "SELECT `pets`.`id`, `pets`.`name`, `pets`.`weight`, `pets`.`birthday` FROM `pets` WHERE (`pets`.`weight` <= ? OR `pets`.`name` = ?) OR `pets`.`name` = ?")
	//NOT
	f = append(f, "not weight le 3.5")
	e = append(e, "SELECT `pets`.`id`, `pets`.`name`, `pets`.`weight`, `pets`.`birthday` FROM `pets` WHERE NOT (`pets`.`weight` <= ?)")

	return f, e
}
