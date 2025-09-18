package controller

// Controller นี้ทำหน้าที่รับ Request, ตรวจสอบและจัดรูปแบบข้อมูลเบื้องต้น, ส่งต่องานให้ Service, และสุดท้ายคือการส่งผลลัพธ์กลับไปยังผู้ใช้งาน
import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	_cardManagingModel "github.com/tehdev/summoner-rift-api/pkg/cardShopMange/model"
	_cardManagingService "github.com/tehdev/summoner-rift-api/pkg/cardShopMange/service"
	"github.com/tehdev/summoner-rift-api/pkg/custom"
)

type cardManagingControllermpl struct{
	cardManagingService _cardManagingService.CardManagingService
}

func NewCardManagingControllermpl (
	cardManagingService  _cardManagingService.CardManagingService,
	
) CardManagingController { //interface
	return &cardManagingControllermpl{cardManagingService}
}

//Call Createcard
func (c *cardManagingControllermpl) Creating(pctx echo.Context) error{
	cardCreatingReq := new(_cardManagingModel.CardCreatingReq)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(cardCreatingReq); err != nil{
		return  custom.Error(pctx,http.StatusBadRequest,err.Error())

	}

	card, err := c.cardManagingService.Creating(cardCreatingReq)
	if err != nil{
		return  custom.Error(pctx,http.StatusInternalServerError,err.Error())
	}

	return pctx.JSON(http.StatusCreated, card)
}

//call Editcard
func (c *cardManagingControllermpl) Editing(pctx echo.Context) error {
	cardID, err := c.getCardID(pctx)
	if err != nil{
		return  custom.Error(pctx,http.StatusBadRequest, err.Error())
	}

	cardEditingReq := new(_cardManagingModel.CardEditingReq)

	custromEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := custromEchoRequest.Bind(cardEditingReq); err != nil {
		return  custom.Error(pctx, http.StatusBadRequest, err.Error())
	}

	card, err := c.cardManagingService.Editing(cardID,cardEditingReq)
	if err != nil{
		return custom.Error(pctx,http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK,card)
}

//Call Archiving card
func (c *cardManagingControllermpl) Archiving(pctx echo.Context) error {
	cardID, err := c.getCardID(pctx)
	if err != nil{
		return  custom.Error(pctx,http.StatusBadRequest, err.Error())
	}

	if err := c.cardManagingService.Archiving(cardID); err != nil{
		return custom.Error(pctx,http.StatusInternalServerError, err.Error())
	}
	return pctx.NoContent(http.StatusNoContent)
}


//แปลง Params เป้นเลขงับ
func (c *cardManagingControllermpl) getCardID(pctx echo.Context) (uint64, error) {
	cardID := pctx.Param("cardID")
	cardIDUint64 ,err := strconv.ParseUint(cardID,10,64) 
	if err != nil{
		return  0, err
	}

	return cardIDUint64, nil
}