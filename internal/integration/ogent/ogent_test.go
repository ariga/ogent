package ogent

import (
	"context"
	"fmt"
	"net/http"
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
	t.Require().Equal(ogent.ListCategoryOKApplicationJSON(ogent.NewCategoryLists(es[0:30])), got)

	// Custom page size.
	got, err = t.handler.ListCategory(context.Background(), ogent.ListCategoryParams{ItemsPerPage: ogent.NewOptInt(10)})
	t.Require().NoError(err)
	t.Require().Equal(ogent.ListCategoryOKApplicationJSON(ogent.NewCategoryLists(es[0:10])), got)

	// Custom page.
	got, err = t.handler.ListCategory(context.Background(), ogent.ListCategoryParams{Page: ogent.NewOptInt(2)})
	t.Require().NoError(err)
	t.Require().Equal(ogent.ListCategoryOKApplicationJSON(ogent.NewCategoryLists(es[30:50])), got)

	// Custom page and page size.
	got, err = t.handler.ListCategory(context.Background(), ogent.ListCategoryParams{Page: ogent.NewOptInt(2), ItemsPerPage: ogent.NewOptInt(10)})
	t.Require().NoError(err)
	t.Require().Equal(ogent.ListCategoryOKApplicationJSON(ogent.NewCategoryLists(es[30:40])), got)
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
	t.Require().Equal(ogent.ListPetFriendsOKApplicationJSON(nil), got)

	owner := t.client.User.Create().SetName("Ariel").SetAge(33).SetFavoriteCatBreed(user.FavoriteCatBreedLeopard).SaveX(context.Background())

	// OK - no attached resource
	loner := t.client.Pet.Create().SetName("Lonely Wolf").SetOwner(owner).SaveX(context.Background())
	got, err = t.handler.ListPetFriends(context.Background(), ogent.ListPetFriendsParams{ID: loner.ID})
	t.Require().NoError(err)
	t.Require().Equal(ogent.ListPetFriendsOKApplicationJSON(nil), got)

	// Get the Lonely Wolf some friends.
	b := make([]*ent.PetCreate, 50)
	for i := range b {
		b[i] = t.client.Pet.Create().SetName("Pet " + strconv.Itoa(i+1)).SetOwner(owner).AddFriends(loner)
	}
	es := t.client.Pet.CreateBulk(b...).SaveX(context.Background())

	// Default page size.
	got, err = t.handler.ListPetFriends(context.Background(), ogent.ListPetFriendsParams{ID: loner.ID})
	t.Require().NoError(err)
	t.Require().Equal(ogent.ListPetFriendsOKApplicationJSON(ogent.NewPetFriendsLists(es[0:30])), got)

	// Custom page size.
	got, err = t.handler.ListPetFriends(context.Background(), ogent.ListPetFriendsParams{ID: loner.ID, ItemsPerPage: ogent.NewOptInt(10)})
	t.Require().NoError(err)
	t.Require().Equal(ogent.ListPetFriendsOKApplicationJSON(ogent.NewPetFriendsLists(es[0:10])), got)

	// Custom page.
	got, err = t.handler.ListPetFriends(context.Background(), ogent.ListPetFriendsParams{ID: loner.ID, Page: ogent.NewOptInt(2)})
	t.Require().NoError(err)
	t.Require().Equal(ogent.ListPetFriendsOKApplicationJSON(ogent.NewPetFriendsLists(es[30:50])), got)

	// Custom page and page size.
	got, err = t.handler.ListPetFriends(context.Background(), ogent.ListPetFriendsParams{ID: loner.ID, Page: ogent.NewOptInt(2), ItemsPerPage: ogent.NewOptInt(10)})
	t.Require().NoError(err)
	t.Require().Equal(ogent.ListPetFriendsOKApplicationJSON(ogent.NewPetFriendsLists(es[30:40])), got)
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
