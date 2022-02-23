// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/bits"
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
	"go.opentelemetry.io/otel/codes"
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
	_ = bits.LeadingZeros64
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
	_ = codes.Unset
)

// HandleCreateCategoryRequest handles createCategory operation.
//
// POST /categories
func (s *Server) handleCreateCategoryRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreateCategory",
		trace.WithAttributes(otelogen.OperationID("createCategory")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodeCreateCategoryRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.CreateCategory(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreateCategoryResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleCreatePetRequest handles createPet operation.
//
// POST /pets
func (s *Server) handleCreatePetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreatePet",
		trace.WithAttributes(otelogen.OperationID("createPet")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodeCreatePetRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.CreatePet(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreatePetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleCreateUserRequest handles createUser operation.
//
// POST /users
func (s *Server) handleCreateUserRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreateUser",
		trace.WithAttributes(otelogen.OperationID("createUser")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodeCreateUserRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.CreateUser(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreateUserResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleDBHealthRequest handles DBHealth operation.
//
// GET /db-health
func (s *Server) handleDBHealthRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DBHealth",
		trace.WithAttributes(otelogen.OperationID("DBHealth")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	response, err := s.h.DBHealth(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeDBHealthResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleDeleteCategoryRequest handles deleteCategory operation.
//
// DELETE /categories/{id}
func (s *Server) handleDeleteCategoryRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DeleteCategory",
		trace.WithAttributes(otelogen.OperationID("deleteCategory")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeDeleteCategoryParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.DeleteCategory(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeDeleteCategoryResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleDeletePetRequest handles deletePet operation.
//
// DELETE /pets/{id}
func (s *Server) handleDeletePetRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DeletePet",
		trace.WithAttributes(otelogen.OperationID("deletePet")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeDeletePetParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.DeletePet(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeDeletePetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleDeleteUserRequest handles deleteUser operation.
//
// DELETE /users/{id}
func (s *Server) handleDeleteUserRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DeleteUser",
		trace.WithAttributes(otelogen.OperationID("deleteUser")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeDeleteUserParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.DeleteUser(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeDeleteUserResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleListCategoryRequest handles listCategory operation.
//
// GET /categories
func (s *Server) handleListCategoryRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListCategory",
		trace.WithAttributes(otelogen.OperationID("listCategory")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeListCategoryParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ListCategory(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListCategoryResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleListCategoryPetsRequest handles listCategoryPets operation.
//
// GET /categories/{id}/pets
func (s *Server) handleListCategoryPetsRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListCategoryPets",
		trace.WithAttributes(otelogen.OperationID("listCategoryPets")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeListCategoryPetsParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ListCategoryPets(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListCategoryPetsResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleListPetRequest handles listPet operation.
//
// GET /pets
func (s *Server) handleListPetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListPet",
		trace.WithAttributes(otelogen.OperationID("listPet")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeListPetParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ListPet(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListPetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleListPetCategoriesRequest handles listPetCategories operation.
//
// GET /pets/{id}/categories
func (s *Server) handleListPetCategoriesRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListPetCategories",
		trace.WithAttributes(otelogen.OperationID("listPetCategories")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeListPetCategoriesParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ListPetCategories(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListPetCategoriesResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleListPetFriendsRequest handles listPetFriends operation.
//
// GET /pets/{id}/friends
func (s *Server) handleListPetFriendsRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListPetFriends",
		trace.WithAttributes(otelogen.OperationID("listPetFriends")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeListPetFriendsParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ListPetFriends(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListPetFriendsResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleListUserRequest handles listUser operation.
//
// GET /users
func (s *Server) handleListUserRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListUser",
		trace.WithAttributes(otelogen.OperationID("listUser")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeListUserParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ListUser(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListUserResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleListUserPetsRequest handles listUserPets operation.
//
// GET /users/{id}/pets
func (s *Server) handleListUserPetsRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListUserPets",
		trace.WithAttributes(otelogen.OperationID("listUserPets")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeListUserPetsParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ListUserPets(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListUserPetsResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleReadCategoryRequest handles readCategory operation.
//
// GET /categories/{id}
func (s *Server) handleReadCategoryRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ReadCategory",
		trace.WithAttributes(otelogen.OperationID("readCategory")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeReadCategoryParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ReadCategory(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeReadCategoryResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleReadPetRequest handles readPet operation.
//
// GET /pets/{id}
func (s *Server) handleReadPetRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ReadPet",
		trace.WithAttributes(otelogen.OperationID("readPet")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeReadPetParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ReadPet(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeReadPetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleReadPetOwnerRequest handles readPetOwner operation.
//
// GET /pets/{id}/owner
func (s *Server) handleReadPetOwnerRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ReadPetOwner",
		trace.WithAttributes(otelogen.OperationID("readPetOwner")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeReadPetOwnerParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ReadPetOwner(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeReadPetOwnerResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleReadUserRequest handles readUser operation.
//
// GET /users/{id}
func (s *Server) handleReadUserRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ReadUser",
		trace.WithAttributes(otelogen.OperationID("readUser")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeReadUserParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ReadUser(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeReadUserResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleUpdateCategoryRequest handles updateCategory operation.
//
// PATCH /categories/{id}
func (s *Server) handleUpdateCategoryRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "UpdateCategory",
		trace.WithAttributes(otelogen.OperationID("updateCategory")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeUpdateCategoryParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodeUpdateCategoryRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.UpdateCategory(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeUpdateCategoryResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleUpdatePetRequest handles updatePet operation.
//
// PATCH /pets/{id}
func (s *Server) handleUpdatePetRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "UpdatePet",
		trace.WithAttributes(otelogen.OperationID("updatePet")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeUpdatePetParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodeUpdatePetRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.UpdatePet(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeUpdatePetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

// HandleUpdateUserRequest handles updateUser operation.
//
// PATCH /users/{id}
func (s *Server) handleUpdateUserRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "UpdateUser",
		trace.WithAttributes(otelogen.OperationID("updateUser")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeUpdateUserParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodeUpdateUserRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.UpdateUser(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeUpdateUserResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		return
	}
	span.SetStatus(codes.Ok, "Ok")
}

func respondError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	data, writeErr := json.Marshal(struct {
		ErrorMessage string `json:"error_message"`
	}{
		ErrorMessage: err.Error(),
	})
	if writeErr == nil {
		w.Write(data)
	}
}