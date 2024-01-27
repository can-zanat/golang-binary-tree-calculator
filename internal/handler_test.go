package internal

import (
	"bytes"
	"encoding/json"
	"main/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/gofiber/fiber/v2"
)

var testCaseBody1 = &model.Tree{
	Tree: &model.BinaryTree{
		Nodes: []*model.Node{
			{
				ID:    "1",
				Left:  StringPtr("2"),
				Right: StringPtr("3"),
				Value: 1,
			},
			{
				ID:    "3",
				Left:  StringPtr("6"),
				Right: StringPtr("7"),
				Value: 3,
			},
			{
				ID:    "7",
				Left:  nil,
				Right: nil,
				Value: 7,
			},
			{
				ID:    "6",
				Left:  nil,
				Right: nil,
				Value: 6,
			},
			{
				ID:    "2",
				Left:  StringPtr("4"),
				Right: StringPtr("5"),
				Value: 2,
			},
			{
				ID:    "5",
				Left:  nil,
				Right: nil,
				Value: 5,
			},
			{
				ID:    "4",
				Left:  nil,
				Right: nil,
				Value: 4,
			},
		},
		Root: "1",
	},
}
var testCaseBody2 = &model.Tree{
	Tree: &model.BinaryTree{
		Nodes: []*model.Node{
			{
				ID:    "1",
				Left:  StringPtr("2"),
				Right: StringPtr("3"),
				Value: 1,
			},
			{
				ID:    "3",
				Left:  nil,
				Right: nil,
				Value: 3,
			},
			{
				ID:    "2",
				Left:  nil,
				Right: nil,
				Value: 2,
			},
		},
		Root: "1",
	},
}
var testCaseBody3 = &model.Tree{
	Tree: &model.BinaryTree{
		Nodes: []*model.Node{
			{
				ID:    "1",
				Left:  StringPtr("-10"),
				Right: StringPtr("-5"),
				Value: 1,
			},
			{
				ID:    "-5",
				Left:  StringPtr("-20"),
				Right: StringPtr("-21"),
				Value: -5,
			},
			{
				ID:    "-21",
				Left:  StringPtr("100-2"),
				Right: StringPtr("1-3"),
				Value: -21,
			},
			{
				ID:    "1-3",
				Left:  nil,
				Right: nil,
				Value: 1,
			},
			{
				ID:    "100-2",
				Left:  nil,
				Right: nil,
				Value: 100,
			},
			{
				ID:    "-20",
				Left:  StringPtr("100"),
				Right: StringPtr("2"),
				Value: -20,
			},
			{
				ID:    "2",
				Left:  nil,
				Right: nil,
				Value: 2,
			},
			{
				ID:    "100",
				Left:  nil,
				Right: nil,
				Value: 100,
			},
			{
				ID:    "-10",
				Left:  StringPtr("30"),
				Right: StringPtr("45"),
				Value: -10,
			},
			{
				ID:    "45",
				Left:  StringPtr("3"),
				Right: StringPtr("-3"),
				Value: 45,
			},
			{
				ID:    "-3",
				Left:  nil,
				Right: nil,
				Value: -3,
			},
			{
				ID:    "3",
				Left:  nil,
				Right: nil,
				Value: 3,
			},
			{
				ID:    "30",
				Left:  StringPtr("5"),
				Right: StringPtr("1-2"),
				Value: 30,
			},
			{
				ID:    "1-2",
				Left:  nil,
				Right: nil,
				Value: 1,
			},
			{
				ID:    "5",
				Left:  nil,
				Right: nil,
				Value: 5,
			},
		},
		Root: "1",
	},
}

func TestHandler_GetUsers(t *testing.T) {
	t.Run("should return path sum for test case 1", func(t *testing.T) {
		mockService, mockServiceController := createMockService(t)
		defer mockServiceController.Finish()

		app := createTestApp()

		expectedResponse := &model.BinaryTreeSum{Sum: 18}

		mockService.
			EXPECT().
			PostBinaryTreeSum(gomock.Any()).
			Return(expectedResponse, nil).
			Times(1)

		handler := NewHandler(mockService)
		handler.RegisterRoutes(app)

		req := NewUsersRequest(http.MethodPost, "/postBinaryTreeSum", testCaseBody1)
		res, err := app.Test(req)
		defer res.Body.Close()

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
	t.Run("should return path sum for test case 2", func(t *testing.T) {
		mockService, mockServiceController := createMockService(t)
		defer mockServiceController.Finish()

		app := createTestApp()

		expectedResponse := &model.BinaryTreeSum{Sum: 18}

		mockService.
			EXPECT().
			PostBinaryTreeSum(gomock.Any()).
			Return(expectedResponse, nil).
			Times(1)

		handler := NewHandler(mockService)
		handler.RegisterRoutes(app)

		req := NewUsersRequest(http.MethodPost, "/postBinaryTreeSum", testCaseBody2)
		res, err := app.Test(req)
		defer res.Body.Close()

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
	t.Run("should return path sum for test case 3", func(t *testing.T) {
		mockService, mockServiceController := createMockService(t)
		defer mockServiceController.Finish()

		app := createTestApp()

		expectedResponse := &model.BinaryTreeSum{Sum: 18}

		mockService.
			EXPECT().
			PostBinaryTreeSum(gomock.Any()).
			Return(expectedResponse, nil).
			Times(1)

		handler := NewHandler(mockService)
		handler.RegisterRoutes(app)

		req := NewUsersRequest(http.MethodPost, "/postBinaryTreeSum", testCaseBody3)
		res, err := app.Test(req)
		defer res.Body.Close()

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
	t.Run("should return internal server error", func(t *testing.T) {
		mockService, mockServiceController := createMockService(t)
		defer mockServiceController.Finish()

		app := createTestApp()

		mockService.
			EXPECT().
			PostBinaryTreeSum(gomock.Any()).
			Return(nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")).
			Times(1)

		handler := NewHandler(mockService)
		handler.RegisterRoutes(app)

		req := NewUsersRequest(http.MethodPost, "/postBinaryTreeSum", testCaseBody1)
		res, err := app.Test(req)
		defer res.Body.Close()

		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, res.StatusCode)
	})
}

func createMockService(t *testing.T) (*Mockactions, *gomock.Controller) {
	t.Helper()

	mockServiceController := gomock.NewController(t)
	mockService := NewMockactions(mockServiceController)

	return mockService, mockServiceController
}

func createTestApp() *fiber.App {
	return fiber.New(fiber.Config{})
}

func NewUsersRequest(method, url string, body *model.Tree) *http.Request {
	requestBody, _ := json.Marshal(body)
	req := httptest.NewRequest(method, url, bytes.NewReader(requestBody))
	req.Header.Add("Content-type", "application/json")

	return req
}

func StringPtr(s string) *string {
	if s == "" {
		return nil
	}

	return &s
}
