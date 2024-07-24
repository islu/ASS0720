package router

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/islu/HW0720/internal/domain/common"
	"github.com/islu/HW0720/internal/usecase"
)

// @description	User status response
type UserStatusResponse struct {
	TaskName        string    `json:"taskName" example:"Onboarding Task"`
	TaskDescription string    `json:"taskDescription" example:"The user needs to swap at least 1000u"`
	TaskStartTime   time.Time `json:"taskStartTime" example:"2024-07-01T00:00:00Z"`
	TaskEndTime     time.Time `json:"taskEndTime" example:"2024-07-29T00:00:00Z"`
	WalletAddress   string    `json:"walletAddress" example:"0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce"`
	Points          int       `json:"points" example:"100"`
	TotalAmount     int64     `json:"totalAmount" example:"100"`
	Status          string    `json:"status" example:"completed"`
	CreateTime      time.Time `json:"createTime" example:"2024-07-02T00:00:00Z"`
	UpdateTime      time.Time `json:"updateTime" example:"2024-07-02T00:00:00Z"`
} //	@name	UserStatusResponse

// Get user tasks status by address
//
//	@summary		Get user tasks status by address
//	@description	Get user tasks status by address
//	@tags			user
//	@accept			json
//	@produce		json
//	@router			/user/tasks/{address}   [get]
//	@param			address	path		string	true	"Wallet address"
//	@success		200		{object}	[]router.UserStatusResponse
//	@failure		400
func GetUserTaskStatus(app *usecase.Application) gin.HandlerFunc {

	return func(c *gin.Context) {

		address := c.Param("address")
		if address == "" {
			respondWithError(c, common.NewError(common.ErrorCodeParameterInvalid, errors.New("address is empty")))
			return
		}

		// TODO: Need to check address signature?

		tasks, err := app.UserService.GetUserTaskStatus(c, address)
		if err != nil {
			err = errors.Join(errors.New("[UserHandler][GetUserTaskStatus] Get user tasks failed"), err)
			respondWithError(c, common.NewError(common.ErrorCodeInternalProcess, err))
			return
		}

		var response []UserStatusResponse
		for _, task := range tasks {
			response = append(response, UserStatusResponse{
				TaskName:        task.TaskName,
				TaskDescription: task.TaskDescription,
				TaskStartTime:   task.TaskStartTime,
				TaskEndTime:     task.TaskEndTime,
				WalletAddress:   task.WalletAddress,
				Points:          task.Points,
				TotalAmount:     task.TotalAmount,
				Status:          task.Status,
				CreateTime:      task.CreateTime,
				UpdateTime:      task.UpdateTime,
			})
		}
		respondWithJSON(c, http.StatusOK, response)
	}
}

// @description	User tasks request body
type UserTasksBody struct {
	WalletAddress string `json:"walletAddress" binding:"required" example:"0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce"`
} //	@name	UserTasksBody

// @description	User tasks response
type UserTasksResponse struct {
	TaskName        string `json:"taskName" example:"Onboarding Task"`
	TaskDescription string `json:"taskDescription" example:"The user needs to swap at least 1000u"`

	// TaskStartTime   time.Time `json:"taskStartTime" example:"2024-07-01T00:00:00Z"`
	// TaskEndTime     time.Time `json:"taskEndTime" example:"2024-07-29T00:00:00Z"`
	// WalletAddress   string    `json:"walletAddress" example:"0x1234.."`

	Points int `json:"points" example:"100"`

	// TotalAmount     int64     `json:"totalAmount" example:"100"`
	// Status     string    `json:"status" example:"completed"`
	// CreateTime time.Time `json:"createTime" example:"2024-07-02T00:00:00Z"`

	UpdateTime time.Time `json:"updateTime" example:"2024-07-02T00:00:00Z"`
} //	@name	UserTasksResponse

// Get user points history for distributed tasks
//
//	@summary		Get user points history for distributed tasks
//	@description	Get user points history for distributed tasks
//	@tags			user
//	@accept			json
//	@produce		json
//	@router			/user/points   [post]
//	@param			body	body		router.UserTasksBody	true	"User tasks request body"
//	@success		200		{object}	[]router.UserTasksResponse
//	@failure		400
func GetUserPointsHistory(app *usecase.Application) gin.HandlerFunc {

	return func(c *gin.Context) {

		var body UserTasksBody
		if err := c.BindJSON(&body); err != nil {
			err = errors.Join(errors.New("[UserHandler][GetUserPointsHistory] Bind JSON"), err)
			respondWithError(c, common.NewError(common.ErrorCodeParameterInvalid, err))
			return
		}

		// TODO: Need to check address signature?

		tasks, err := app.UserService.GetUserPointsHistory(c, body.WalletAddress)
		if err != nil {
			err = errors.Join(errors.New("[UserHandler][GetUserPointsHistory] Get user tasks failed"), err)
			respondWithError(c, common.NewError(common.ErrorCodeInternalProcess, err))
			return
		}

		var response []UserTasksResponse
		for _, task := range tasks {
			response = append(response, UserTasksResponse{
				TaskName:        task.TaskName,
				TaskDescription: task.TaskDescription,
				Points:          task.Points,
				UpdateTime:      task.UpdateTime,
			})
		}
		respondWithJSON(c, http.StatusOK, response)
	}
}

// Distribute tasks for user
//
// Assumption: Users must accept the task first
//
//	@summary		Distribute tasks for user
//	@description	Distribute tasks for user
//	@tags			user
//	@accept			json
//	@produce		json
//	@router			/user/distribute   [post]
//	@param			body	body		router.UserTasksBody	true	"User tasks request body"
//	@success		200		{object}	router.SuccessMessage
//	@failure		400
func DistributeTasks(app *usecase.Application) gin.HandlerFunc {

	return func(c *gin.Context) {

		var body UserTasksBody
		if err := c.BindJSON(&body); err != nil {
			err = errors.Join(errors.New("[UserHandler][DistributeTasks] Bind JSON"), err)
			respondWithError(c, common.NewError(common.ErrorCodeParameterInvalid, err))
			return
		}

		// TODO: Need to check address signature?

		err := app.UserService.DistributeTasks(c, body.WalletAddress)
		if err != nil {
			err = errors.Join(errors.New("[UserHandler][DistributeTasks] Get user tasks failed"), err)
			respondWithError(c, common.NewError(common.ErrorCodeInternalProcess, err))
			return
		}

		respondWithSuccess(c)
	}
}

/*
	Dashboard
*/

// @description Uniswap USDC/ETH pair request body
type UniswapUSDCETHPairBody struct {
	StartBlockNumber int64 `json:"startBlockNumber" binding:"required" example:"20358617"`
	EndBlockNumber   int64 `json:"endBlockNumber" binding:"required" example:"20358638"`
} // @name UniswapUSDCETHPairBody

// Update Uniswap USDC/ETH pair swap log
//
//	@summary		Update Uniswap USDC/ETH pair swap log
//	@description	Update Uniswap USDC/ETH pair swap log
//	@tags			dashboard
//	@accept			json
//	@produce		json
//	@router			/dashboard/uniswap-log   [post]
//	@param			body	body		router.UniswapUSDCETHPairBody	true	"Uniswap USDC/ETH pair request body"
//	@success		200		{object}	router.SuccessMessage
//	@failure		400
func UpdateUniswapUSDCETHPairSwapLog(app *usecase.Application) gin.HandlerFunc {

	return func(c *gin.Context) {

		var body UniswapUSDCETHPairBody
		if err := c.BindJSON(&body); err != nil {
			err = errors.Join(errors.New("[DashboardHandler][UpdateUniswapUSDCETHPairSwapLog] Bind JSON"), err)
			respondWithError(c, common.NewError(common.ErrorCodeParameterInvalid, err))
			return
		}

		err := app.UserService.UpdateUniswapUSDCETHPairSwapLog(c, body.StartBlockNumber, body.EndBlockNumber)
		if err != nil {
			err = errors.Join(errors.New("[DashboardHandler][UpdateUniswapUSDCETHPairSwapLog] Update Uniswap USDC/ETH pair swap log failed"), err)
			respondWithError(c, common.NewError(common.ErrorCodeInternalProcess, err))
			return
		}

		respondWithSuccess(c)
	}
}
