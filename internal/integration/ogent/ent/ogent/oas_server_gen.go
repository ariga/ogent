// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CreateCategory implements createCategory operation.
	//
	// POST /categories
	CreateCategory(ctx context.Context, req CreateCategoryReq) (CreateCategoryRes, error)
	// CreateCategoryPets implements createCategoryPets operation.
	//
	// POST /categories/{id}/pets
	CreateCategoryPets(ctx context.Context, req CreateCategoryPetsReq, params CreateCategoryPetsParams) (CreateCategoryPetsRes, error)
	// CreatePet implements createPet operation.
	//
	// POST /pets
	CreatePet(ctx context.Context, req CreatePetReq) (CreatePetRes, error)
	// CreatePetCategories implements createPetCategories operation.
	//
	// POST /pets/{id}/categories
	CreatePetCategories(ctx context.Context, req CreatePetCategoriesReq, params CreatePetCategoriesParams) (CreatePetCategoriesRes, error)
	// CreatePetFriends implements createPetFriends operation.
	//
	// POST /pets/{id}/friends
	CreatePetFriends(ctx context.Context, req CreatePetFriendsReq, params CreatePetFriendsParams) (CreatePetFriendsRes, error)
	// CreatePetOwner implements createPetOwner operation.
	//
	// POST /pets/{id}/owner
	CreatePetOwner(ctx context.Context, req CreatePetOwnerReq, params CreatePetOwnerParams) (CreatePetOwnerRes, error)
	// CreateUser implements createUser operation.
	//
	// POST /users
	CreateUser(ctx context.Context, req CreateUserReq) (CreateUserRes, error)
	// CreateUserPets implements createUserPets operation.
	//
	// POST /users/{id}/pets
	CreateUserPets(ctx context.Context, req CreateUserPetsReq, params CreateUserPetsParams) (CreateUserPetsRes, error)
	// DeleteCategory implements deleteCategory operation.
	//
	// DELETE /categories/{id}
	DeleteCategory(ctx context.Context, params DeleteCategoryParams) (DeleteCategoryRes, error)
	// DeletePet implements deletePet operation.
	//
	// DELETE /pets/{id}
	DeletePet(ctx context.Context, params DeletePetParams) (DeletePetRes, error)
	// DeletePetOwner implements deletePetOwner operation.
	//
	// DELETE /pets/{id}/owner
	DeletePetOwner(ctx context.Context, params DeletePetOwnerParams) (DeletePetOwnerRes, error)
	// DeleteUser implements deleteUser operation.
	//
	// DELETE /users/{id}
	DeleteUser(ctx context.Context, params DeleteUserParams) (DeleteUserRes, error)
	// ListCategory implements listCategory operation.
	//
	// GET /categories
	ListCategory(ctx context.Context, params ListCategoryParams) (ListCategoryRes, error)
	// ListCategoryPets implements listCategoryPets operation.
	//
	// GET /categories/{id}/pets
	ListCategoryPets(ctx context.Context, params ListCategoryPetsParams) (ListCategoryPetsRes, error)
	// ListPet implements listPet operation.
	//
	// GET /pets
	ListPet(ctx context.Context, params ListPetParams) (ListPetRes, error)
	// ListPetCategories implements listPetCategories operation.
	//
	// GET /pets/{id}/categories
	ListPetCategories(ctx context.Context, params ListPetCategoriesParams) (ListPetCategoriesRes, error)
	// ListPetFriends implements listPetFriends operation.
	//
	// GET /pets/{id}/friends
	ListPetFriends(ctx context.Context, params ListPetFriendsParams) (ListPetFriendsRes, error)
	// ListUser implements listUser operation.
	//
	// GET /users
	ListUser(ctx context.Context, params ListUserParams) (ListUserRes, error)
	// ListUserPets implements listUserPets operation.
	//
	// GET /users/{id}/pets
	ListUserPets(ctx context.Context, params ListUserPetsParams) (ListUserPetsRes, error)
	// ReadCategory implements readCategory operation.
	//
	// GET /categories/{id}
	ReadCategory(ctx context.Context, params ReadCategoryParams) (ReadCategoryRes, error)
	// ReadPet implements readPet operation.
	//
	// GET /pets/{id}
	ReadPet(ctx context.Context, params ReadPetParams) (ReadPetRes, error)
	// ReadPetOwner implements readPetOwner operation.
	//
	// GET /pets/{id}/owner
	ReadPetOwner(ctx context.Context, params ReadPetOwnerParams) (ReadPetOwnerRes, error)
	// ReadUser implements readUser operation.
	//
	// GET /users/{id}
	ReadUser(ctx context.Context, params ReadUserParams) (ReadUserRes, error)
	// UpdateCategory implements updateCategory operation.
	//
	// PATCH /categories/{id}
	UpdateCategory(ctx context.Context, req UpdateCategoryReq, params UpdateCategoryParams) (UpdateCategoryRes, error)
	// UpdatePet implements updatePet operation.
	//
	// PATCH /pets/{id}
	UpdatePet(ctx context.Context, req UpdatePetReq, params UpdatePetParams) (UpdatePetRes, error)
	// UpdateUser implements updateUser operation.
	//
	// PATCH /users/{id}
	UpdateUser(ctx context.Context, req UpdateUserReq, params UpdateUserParams) (UpdateUserRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	cfg config
}

func NewServer(h Handler, opts ...Option) *Server {
	srv := &Server{
		h:   h,
		cfg: newConfig(opts...),
	}
	return srv
}
