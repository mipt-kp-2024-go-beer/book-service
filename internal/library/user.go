package library

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type UserServiceClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewUserServiceClient(baseURL string) *UserServiceClient {
	return &UserServiceClient{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

// Check if the token has the required permissions
func (client *UserServiceClient) CheckPermissions(token string, mask uint) (bool, error) {
	// Prepare the request body
	data := struct {
		Token string `json:"token"`
	}{Token: token}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return false, fmt.Errorf("error marshalling token data: %v", err)
	}

	// Make the POST request to check permissions
	resp, err := client.HTTPClient.Post("http://"+client.BaseURL+"/user/permissions", "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return false, fmt.Errorf("error sending request to user service: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to check permissions, status: %d", resp.StatusCode)
	}

	// Decode the response
	var permissionResp struct {
		Permissions string `json:"permissios"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&permissionResp); err != nil {
		return false, fmt.Errorf("error decoding permissions response: %v", err)
	}

	// Convert permissions string to an integer
	// permissions, err := strconv.Atoi(permissionResp.Permissions)
	permissions, err := strconv.ParseUint(permissionResp.Permissions, 10, 64)
	if err != nil {
		return false, fmt.Errorf("error converting permission value: %v", err)
	}

	// Check if the permission mask includes the required permission
	return (uint(permissions)&mask != 0), nil
}

const (
	// PermManageBooks allows the user to add, edit and delete books from the library
	PermManageBooks uint = 1 << 0
	// PermQueryTotalStock allows the user to get the total stored book count
	PermQueryTotalStock uint = 1 << 1
	// PermChangeTotalStock allows the user to register updates to the total stored book count.
	// Requires PermGetTotalStock as a prerequisite.
	PermChangeTotalStock uint = 1 << 2
	// PermQueryUsers allows the user to get information about other users, including their permissions.
	// Not required to get information about oneself, other rules apply.
	PermQueryUsers uint = 1 << 3
	// PermManageUsers allows the user to add, edit and delete other users.
	// Not required to manage oneself, other rules apply.
	// Requires PermQueryUsers as a prerequisite.
	PermManageUsers uint = 1 << 4
	// PermGrantPermissions allows the user to grant permissions to other users.
	// Only a subset of own permissions may be granted.
	// Requires PermQueryUsers as a prerequisite.
	PermGrantPermissions uint = 1 << 5
	// PermLoanBooks allows the user to register book takeouts and returns.
	PermLoanBooks uint = 1 << 6
	// PermQueryAvailableStock allows the user to get the number of available (not lent out) copies of a book.
	PermQueryAvailableStock uint = 1 << 7
	// PermQueryReservations allows the user to get information related to book reservations.
	PermQueryReservations uint = 1 << 8
)
