package router

import (
	"github.com/gin-gonic/gin"
	"github.com/islu/ASS0720/internal/usecase"
)

// @description	User status response
type UserStatusResponse struct {
	WalletAddress string `json:"walletAddress" example:"0x1234.."`
	Points        int    `json:"points" example:"100"`
	Amount        int64  `json:"amount" example:"100"`
} //	@name	UserStatusResponse

// Get user tasks status by address
//
//	@summary		Get user tasks status by address
//	@description	Get user tasks status by address
//	@tags			user
//	@accept			json
//	@produce		json
//	@router			/user/{address}   [get]
//	@param			address	path		string	true	"Wallet address"
//	@success		200		{object}	router.UserStatusResponse
//	@failure		400
func GetUserTaskStatus(app *usecase.Application) gin.HandlerFunc {

	return func(c *gin.Context) {

		// TODO: Implement logic

		respondWithSuccess(c)
	}
}

// @description	User tasks request body
type UserTasksBody struct {
	WalletAddress string `json:"walletAddress" example:"0x1234.."`
} //	@name	UserTasksBody

// @description	User tasks response
type UserTasksResponse struct {
	WalletAddress string `json:"walletAddress" example:"0x1234.."`
	Points        int    `json:"points" example:"100"`
	Amount        int64  `json:"amount" example:"100"`
} //	@name	UserTasksResponse

// Get user points history for distributed tasks
//
//	@summary		Get user points history for distributed tasks
//	@description	Get user points history for distributed tasks
//	@tags			user
//	@accept			json
//	@produce		json
//	@router			/user/tasks   [post]
//	@param			body	body		router.UserTasksBody	true	"User tasks request body"
//	@success		200		{object}	[]router.UserTasksResponse
//	@failure		400
func GetUserTaskList(app *usecase.Application) gin.HandlerFunc {

	return func(c *gin.Context) {

		// TODO: Implement logic

		respondWithSuccess(c)
	}
}
