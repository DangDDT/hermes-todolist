package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"
)

const baseURL = "http://localhost:8080/api/v1"

// uniqueName generates a unique username per test run to avoid collisions.
func uniqueName(prefix string) string {
	return fmt.Sprintf("%s_%d", prefix, time.Now().UnixNano())
}

// ---- helpers ----

// apiURL joins path segments to the base URL.
func apiURL(path string) string {
	return baseURL + path
}

// doReq performs an HTTP request and returns the response.
func doReq(method, url, contentType, body string) (*http.Response, error) {
	var reqBody *bytes.Reader
	if body != "" {
		reqBody = bytes.NewReader([]byte(body))
	} else {
		reqBody = bytes.NewReader(nil)
	}
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	client := &http.Client{}
	return client.Do(req)
}

// doReqWithCookie performs an HTTP request with an access_token cookie.
func doReqWithCookie(method, url, contentType, body, token string) (*http.Response, error) {
	var reqBody *bytes.Reader
	if body != "" {
		reqBody = bytes.NewReader([]byte(body))
	} else {
		reqBody = bytes.NewReader(nil)
	}
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	req.Header.Set("Cookie", "access_token="+token)
	client := &http.Client{}
	return client.Do(req)
}

// decodeJSON decodes a JSON response body into the provided target.
func decodeJSON(resp *http.Response, target any) error {
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

// registerUser creates a new user and returns the JSON response map.
func registerUser(t *testing.T, username, password, displayName string) (map[string]any, int, error) {
	t.Helper()
	body := fmt.Sprintf(`{"username":"%s","password":"%s","display_name":"%s"}`,
		username, password, displayName)
	resp, err := doReq("POST", apiURL("/auth/register"), "application/json", body)
	if err != nil {
		return nil, 0, err
	}
	status := resp.StatusCode
	var result map[string]any
	if err := decodeJSON(resp, &result); err != nil {
		return nil, status, err
	}
	return result, status, nil
}

// loginUser authenticates and returns the JSON response map, status, and extracted token.
func loginUser(t *testing.T, username, password string) (map[string]any, int, string, error) {
	t.Helper()
	body := fmt.Sprintf(`{"username":"%s","password":"%s"}`, username, password)
	resp, err := doReq("POST", apiURL("/auth/login"), "application/json", body)
	if err != nil {
		return nil, 0, "", err
	}
	status := resp.StatusCode
	var result map[string]any
	if err := decodeJSON(resp, &result); err != nil {
		return nil, status, "", err
	}
	token := ""
	if data, ok := result["data"].(map[string]any); ok {
		if t, ok := data["token"].(string); ok {
			token = t
		}
	}
	return result, status, token, nil
}

// createTask creates a task with the given token and returns the response.
func createTask(t *testing.T, token, title, description, priority string) (map[string]any, int, error) {
	t.Helper()
	body := fmt.Sprintf(`{"title":"%s","description":"%s","priority":"%s"}`,
		title, description, priority)
	resp, err := doReqWithCookie("POST", apiURL("/tasks"), "application/json", body, token)
	if err != nil {
		return nil, 0, err
	}
	status := resp.StatusCode
	var result map[string]any
	if err := decodeJSON(resp, &result); err != nil {
		return nil, status, err
	}
	return result, status, nil
}

// getTask retrieves a task by ID with the given token.
func getTask(t *testing.T, token, taskID string) (map[string]any, int, error) {
	t.Helper()
	resp, err := doReqWithCookie("GET", apiURL("/tasks/"+taskID), "", "", token)
	if err != nil {
		return nil, 0, err
	}
	status := resp.StatusCode
	var result map[string]any
	if err := decodeJSON(resp, &result); err != nil {
		return nil, status, err
	}
	return result, status, nil
}

// listTasks lists tasks with optional query parameters.
func listTasks(t *testing.T, token string, queryParams ...string) (map[string]any, int, error) {
	t.Helper()
	url := apiURL("/tasks")
	if len(queryParams) > 0 && queryParams[0] != "" {
		url += "?" + queryParams[0]
	}
	resp, err := doReqWithCookie("GET", url, "", "", token)
	if err != nil {
		return nil, 0, err
	}
	status := resp.StatusCode
	var result map[string]any
	if err := decodeJSON(resp, &result); err != nil {
		return nil, status, err
	}
	return result, status, nil
}

// updateTask updates a task with the given fields.
func updateTask(t *testing.T, token, taskID string, fields map[string]string) (map[string]any, int, error) {
	t.Helper()
	var parts []string
	for k, v := range fields {
		parts = append(parts, fmt.Sprintf(`"%s":"%s"`, k, v))
	}
	body := "{" + strings.Join(parts, ",") + "}"
	resp, err := doReqWithCookie("PUT", apiURL("/tasks/"+taskID), "application/json", body, token)
	if err != nil {
		return nil, 0, err
	}
	status := resp.StatusCode
	var result map[string]any
	if err := decodeJSON(resp, &result); err != nil {
		return nil, status, err
	}
	return result, status, nil
}

// deleteTask deletes a task by ID.
func deleteTask(t *testing.T, token, taskID string) (int, error) {
	t.Helper()
	resp, err := doReqWithCookie("DELETE", apiURL("/tasks/"+taskID), "", "", token)
	if err != nil {
		return 0, err
	}
	status := resp.StatusCode
	resp.Body.Close()
	return status, nil
}

// getError extracts the top-level error string from a response map.
// Handles both formats: {"error":"msg","code":"..."} and {"data":{"error":"msg"}}
func getError(result map[string]any) string {
	if err, ok := result["error"].(string); ok {
		return err
	}
	if data, ok := result["data"].(map[string]any); ok {
		if err, ok := data["error"].(string); ok {
			return err
		}
	}
	return ""
}

// getData extracts the "data" sub-map from a response.
func getData(result map[string]any) map[string]any {
	if data, ok := result["data"].(map[string]any); ok {
		return data
	}
	return nil
}

// ---- tests ----

func TestHealthEndpoint(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	t.Run("HealthEndpoint", func(t *testing.T) {
		resp, err := doReq("GET", apiURL("/health"), "", "")
		if err != nil {
			t.Fatalf("health request failed: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected 200, got %d", resp.StatusCode)
		}
		var result map[string]any
		if err := decodeJSON(resp, &result); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}
		data := getData(result)
		if data == nil {
			t.Fatal("expected data in response")
		}
		if data["status"] != "ok" {
			t.Errorf("expected status 'ok', got %v", data["status"])
		}
		if data["version"] != "1.0.0" {
			t.Errorf("expected version '1.0.0', got %v", data["version"])
		}
	})
}

func TestAuthRegister(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	t.Run("RegisterNewUser", func(t *testing.T) {
		username := uniqueName("regtest")
		result, status, err := registerUser(t, username, "Password1!", "Register Test")
		if err != nil {
			t.Fatalf("register request failed: %v", err)
		}
		if status != http.StatusCreated {
			t.Fatalf("expected 201, got %d: %v", status, result)
		}
		data := getData(result)
		if data == nil {
			t.Fatal("expected data in response")
		}
		if data["username"] != username {
			t.Errorf("expected username %q, got %v", username, data["username"])
		}
		if data["id"] == "" {
			t.Errorf("expected non-empty id, got %v", data["id"])
		}
	})
}

func TestAuthRegisterDuplicate(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	t.Run("RegisterDuplicateUser", func(t *testing.T) {
		username := uniqueName("dupuser")
		// Register first time.
		_, status, err := registerUser(t, username, "Password1!", "Dup User")
		if err != nil {
			t.Fatalf("first register failed: %v", err)
		}
		if status != http.StatusCreated {
			t.Fatalf("expected 201 on first register, got %d", status)
		}
		// Register same user again.
		result, status, err := registerUser(t, username, "Password1!", "Dup User Again")
		if err != nil {
			t.Fatalf("second register request failed: %v", err)
		}
		if status != http.StatusConflict {
			t.Fatalf("expected 409, got %d", status)
		}
		errMsg := getError(result)
		if !strings.Contains(strings.ToLower(errMsg), "exists") {
			t.Errorf("expected error containing 'exists', got %q", errMsg)
		}
	})
}

func TestAuthRegisterInvalid(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	t.Run("RegisterNoPassword", func(t *testing.T) {
		body := `{"username":"nopass_user","password":"","display_name":"No Pass"}`
		resp, err := doReq("POST", apiURL("/auth/register"), "application/json", body)
		if err != nil {
			t.Fatalf("register request failed: %v", err)
		}
		if resp.StatusCode != http.StatusUnprocessableEntity {
			t.Fatalf("expected 422, got %d", resp.StatusCode)
		}
		var result map[string]any
		if err := decodeJSON(resp, &result); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}
		errMsg := getError(result)
		if !strings.Contains(strings.ToLower(errMsg), "password") {
			t.Errorf("expected error mentioning password, got %q", errMsg)
		}
	})
}

func TestAuthLogin(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	t.Run("LoginSuccess", func(t *testing.T) {
		username := uniqueName("logintest")
		_, _, err := registerUser(t, username, "Password1!", "Login Test")
		if err != nil {
			t.Fatalf("register before login failed: %v", err)
		}
		result, status, token, err := loginUser(t, username, "Password1!")
		if err != nil {
			t.Fatalf("login request failed: %v", err)
		}
		if status != http.StatusOK {
			t.Fatalf("expected 200, got %d", status)
		}
		if token == "" {
			t.Fatal("expected non-empty token in login response")
		}
		data := getData(result)
		if data == nil {
			t.Fatal("expected data in response")
		}
		if _, ok := data["expires_at"]; !ok {
			t.Errorf("expected expires_at in response data")
		}
	})
}

func TestAuthLoginWrongPassword(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	t.Run("LoginWrongPassword", func(t *testing.T) {
		username := uniqueName("wrongpw")
		_, _, err := registerUser(t, username, "Password1!", "Wrong PW")
		if err != nil {
			t.Fatalf("register before login failed: %v", err)
		}
		result, status, _, err := loginUser(t, username, "WrongPassword!")
		if err != nil {
			t.Fatalf("login request failed: %v", err)
		}
		if status != http.StatusUnauthorized {
			t.Fatalf("expected 401, got %d", status)
		}
		errMsg := getError(result)
		if errMsg == "" {
			t.Error("expected error message in response")
		}
	})
}

func TestAuthEndpointWithoutToken(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	t.Run("TasksEndpointWithoutToken", func(t *testing.T) {
		resp, err := doReq("GET", apiURL("/tasks"), "", "")
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		if resp.StatusCode != http.StatusUnauthorized {
			t.Fatalf("expected 401, got %d", resp.StatusCode)
		}
		var result map[string]any
		if err := decodeJSON(resp, &result); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}
		errMsg := getError(result)
		if !strings.Contains(strings.ToLower(errMsg), "missing") &&
			!strings.Contains(strings.ToLower(errMsg), "token") {
			t.Errorf("expected error mentioning missing/token, got %q", errMsg)
		}
	})
}

func TestCreateTask(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	t.Run("CreateTaskSuccess", func(t *testing.T) {
		username := uniqueName("createtask")
		_, _, err := registerUser(t, username, "Password1!", "Create Task")
		if err != nil {
			t.Fatalf("register failed: %v", err)
		}
		_, _, token, err := loginUser(t, username, "Password1!")
		if err != nil || token == "" {
			t.Fatalf("login failed: %v", err)
		}
		result, status, err := createTask(t, token, "My Task", "My Description", "HIGH")
		if err != nil {
			t.Fatalf("create task failed: %v", err)
		}
		if status != http.StatusCreated {
			t.Fatalf("expected 201, got %d", status)
		}
		data := getData(result)
		if data == nil {
			t.Fatal("expected data in response")
		}
		if data["id"] == "" {
			t.Error("expected non-empty task id")
		}
		if data["title"] != "My Task" {
			t.Errorf("expected title 'My Task', got %v", data["title"])
		}
		if data["status"] != "TODO" {
			t.Errorf("expected status 'TODO', got %v", data["status"])
		}
	})
}

func TestListTasks(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	t.Run("ListTasks", func(t *testing.T) {
		username := uniqueName("listtask")
		_, _, err := registerUser(t, username, "Password1!", "List Task")
		if err != nil {
			t.Fatalf("register failed: %v", err)
		}
		_, _, token, err := loginUser(t, username, "Password1!")
		if err != nil || token == "" {
			t.Fatalf("login failed: %v", err)
		}
		// Create a couple of tasks.
		createTask(t, token, "Task A", "Desc A", "LOW")
		createTask(t, token, "Task B", "Desc B", "HIGH")
		result, status, err := listTasks(t, token, "")
		if err != nil {
			t.Fatalf("list tasks failed: %v", err)
		}
		if status != http.StatusOK {
			t.Fatalf("expected 200, got %d", status)
		}
		data := getData(result)
		if data == nil {
			t.Fatal("expected data in response")
		}
		tasks, ok := data["tasks"].([]any)
		if !ok {
			t.Fatal("expected tasks array in response")
		}
		if len(tasks) < 2 {
			t.Errorf("expected at least 2 tasks, got %d", len(tasks))
		}
		if _, ok := result["meta"]; !ok {
			t.Error("expected meta in paginated response")
		}
	})
}

func TestGetTask(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	t.Run("GetTask", func(t *testing.T) {
		username := uniqueName("gettask")
		_, _, err := registerUser(t, username, "Password1!", "Get Task")
		if err != nil {
			t.Fatalf("register failed: %v", err)
		}
		_, _, token, err := loginUser(t, username, "Password1!")
		if err != nil || token == "" {
			t.Fatalf("login failed: %v", err)
		}
		createResult, _, err := createTask(t, token, "Get Me", "Get desc", "MEDIUM")
		if err != nil {
			t.Fatalf("create task failed: %v", err)
		}
		taskID := getData(createResult)["id"].(string)
		result, status, err := getTask(t, token, taskID)
		if err != nil {
			t.Fatalf("get task failed: %v", err)
		}
		if status != http.StatusOK {
			t.Fatalf("expected 200, got %d", status)
		}
		data := getData(result)
		if data == nil {
			t.Fatal("expected data in response")
		}
		if data["id"] != taskID {
			t.Errorf("expected task id %q, got %v", taskID, data["id"])
		}
		if data["title"] != "Get Me" {
			t.Errorf("expected title 'Get Me', got %v", data["title"])
		}
		if data["status"] != "TODO" {
			t.Errorf("expected status 'TODO', got %v", data["status"])
		}
		if data["priority"] != "MEDIUM" {
			t.Errorf("expected priority 'MEDIUM', got %v", data["priority"])
		}
	})
}

func TestUpdateTask(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	t.Run("UpdateTask", func(t *testing.T) {
		username := uniqueName("updatetask")
		_, _, err := registerUser(t, username, "Password1!", "Update Task")
		if err != nil {
			t.Fatalf("register failed: %v", err)
		}
		_, _, token, err := loginUser(t, username, "Password1!")
		if err != nil || token == "" {
			t.Fatalf("login failed: %v", err)
		}
		createResult, _, err := createTask(t, token, "Original", "Original desc", "LOW")
		if err != nil {
			t.Fatalf("create task failed: %v", err)
		}
		taskID := getData(createResult)["id"].(string)
		// Update status and title.
		result, status, err := updateTask(t, token, taskID, map[string]string{
			"status": "IN_PROGRESS",
			"title":  "Updated Title",
		})
		if err != nil {
			t.Fatalf("update task failed: %v", err)
		}
		if status != http.StatusOK {
			t.Fatalf("expected 200, got %d", status)
		}
		data := getData(result)
		if data == nil {
			t.Fatal("expected data in response")
		}
		if data["status"] != "IN_PROGRESS" {
			t.Errorf("expected status 'IN_PROGRESS', got %v", data["status"])
		}
		if data["title"] != "Updated Title" {
			t.Errorf("expected title 'Updated Title', got %v", data["title"])
		}
	})
}

func TestDeleteTask(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	t.Run("DeleteTask", func(t *testing.T) {
		username := uniqueName("deltask")
		_, _, err := registerUser(t, username, "Password1!", "Delete Task")
		if err != nil {
			t.Fatalf("register failed: %v", err)
		}
		_, _, token, err := loginUser(t, username, "Password1!")
		if err != nil || token == "" {
			t.Fatalf("login failed: %v", err)
		}
		createResult, _, err := createTask(t, token, "Delete Me", "To be deleted", "HIGH")
		if err != nil {
			t.Fatalf("create task failed: %v", err)
		}
		taskID := getData(createResult)["id"].(string)
		// Delete the task.
		status, err := deleteTask(t, token, taskID)
		if err != nil {
			t.Fatalf("delete request failed: %v", err)
		}
		if status != http.StatusNoContent {
			t.Fatalf("expected 204, got %d", status)
		}
		// Verify task is gone.
		_, status, err = getTask(t, token, taskID)
		if err != nil {
			t.Fatalf("get after delete failed: %v", err)
		}
		if status != http.StatusNotFound {
			t.Fatalf("expected 404 after delete, got %d", status)
		}
	})
}

func TestListTasksFilteredByStatus(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	t.Run("FilterByStatus", func(t *testing.T) {
		username := uniqueName("filtertask")
		_, _, err := registerUser(t, username, "Password1!", "Filter Task")
		if err != nil {
			t.Fatalf("register failed: %v", err)
		}
		_, _, token, err := loginUser(t, username, "Password1!")
		if err != nil || token == "" {
			t.Fatalf("login failed: %v", err)
		}
		// Create a task with TODO status.
		createTask(t, token, "Todo Item", "Todo desc", "LOW")
		// Filter by TODO.
		result, status, err := listTasks(t, token, "status=TODO")
		if err != nil {
			t.Fatalf("filtered list failed: %v", err)
		}
		if status != http.StatusOK {
			t.Fatalf("expected 200, got %d", status)
		}
		data := getData(result)
		if data == nil {
			t.Fatal("expected data in response")
		}
		tasks, ok := data["tasks"].([]any)
		if !ok {
			t.Fatal("expected tasks array in filtered response")
		}
		if len(tasks) == 0 {
			t.Error("expected at least one TODO task")
		}
		// Verify all returned tasks have TODO status.
		for i, task := range tasks {
			taskMap, ok := task.(map[string]any)
			if !ok {
				continue
			}
			if taskMap["status"] != "TODO" {
				t.Errorf("task %d: expected status TODO, got %v", i, taskMap["status"])
			}
		}
	})
}

func TestAuthorization(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	t.Run("AuthorizationUserBCannotAccessUserATask", func(t *testing.T) {
		userA := uniqueName("userauth")
		userB := uniqueName("userbauth")
		// Register both users.
		_, _, err := registerUser(t, userA, "Password1!", "User A")
		if err != nil {
			t.Fatalf("register user A failed: %v", err)
		}
		_, _, err = registerUser(t, userB, "Password1!", "User B")
		if err != nil {
			t.Fatalf("register user B failed: %v", err)
		}
		// Login both.
		_, _, tokenA, err := loginUser(t, userA, "Password1!")
		if err != nil || tokenA == "" {
			t.Fatalf("login user A failed: %v", err)
		}
		_, _, tokenB, err := loginUser(t, userB, "Password1!")
		if err != nil || tokenB == "" {
			t.Fatalf("login user B failed: %v", err)
		}
		// User A creates a task.
		createResult, _, err := createTask(t, tokenA, "User A Task", "Secret", "HIGH")
		if err != nil {
			t.Fatalf("user A create task failed: %v", err)
		}
		taskID := getData(createResult)["id"].(string)
		// User B tries to GET the task — should get 404 (not found due to ownership check).
		_, status, err := getTask(t, tokenB, taskID)
		if err != nil {
			t.Fatalf("user B get task failed: %v", err)
		}
		if status != http.StatusNotFound {
			t.Errorf("expected 404 when user B gets user A's task, got %d", status)
		}
		// User B tries to DELETE the task — should get 404.
		status, err = deleteTask(t, tokenB, taskID)
		if err != nil {
			t.Fatalf("user B delete task failed: %v", err)
		}
		if status != http.StatusNotFound {
			t.Errorf("expected 404 when user B deletes user A's task, got %d", status)
		}
		// User A can still GET their own task.
		_, status, err = getTask(t, tokenA, taskID)
		if err != nil {
			t.Fatalf("user A get task after attempts failed: %v", err)
		}
		if status != http.StatusOK {
			t.Errorf("expected 200 when user A gets own task, got %d", status)
		}
	})
}
