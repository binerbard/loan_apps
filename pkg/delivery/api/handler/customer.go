package handler

import (
	"loan_apps/datasource/domain/schema"
	"loan_apps/datasource/domain/usecase"
	"loan_apps/scripts/utils"

	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	customerUsecase usecase.CustomerUsecase
}

func NewCustomerHandler(uc usecase.CustomerUsecase) *CustomerHandler {
	return &CustomerHandler{
		customerUsecase: uc,
	}
}


func (h *CustomerHandler) Store(ctx *fiber.Ctx) error {	
	var req schema.CustomerReq

	// Image Check
	imageKTP, err := ctx.FormFile("img_ktp")
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	ktpPath, err := utils.UploadImage(imageKTP)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	imageSelfie , err := ctx.FormFile("selfie")
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	selfiePath, err := utils.UploadImage(imageSelfie)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	if err := ctx.BodyParser(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	if err := schema.ValidatorCustomerReq(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	req.ImgSelfie = selfiePath
	req.ImgID = ktpPath

	if err := h.customerUsecase.Create(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success store customer",
	})
}

func (h *CustomerHandler) Update(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	if ID == "" {
		return utils.BadRequest(ctx, fiber.Map{
			"message": "id is empty",
		})
	}

	var req schema.CustomerReq
	if err := ctx.BodyParser(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	if err := schema.ValidatorCustomerReq(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}


	imageSelfie , err := ctx.FormFile("selfie")
	if err != nil {
		req.ImgSelfie = ""
	}else {
		selfiePath, err := utils.UploadImage(imageSelfie)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest)
			return utils.BadRequest(ctx, fiber.Map{
				"message": err.Error(),
			})
		}
		req.ImgSelfie = selfiePath
	}

	imageKTP, err := ctx.FormFile("img_ktp")
	if err != nil {
		req.ImgID = ""
	}else {
		ktpPath, err := utils.UploadImage(imageKTP)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest)
			return utils.BadRequest(ctx, fiber.Map{
				"message": err.Error(),
			})
		}
		req.ImgID = ktpPath
	}

	if err := h.customerUsecase.Update(ID, &req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success updated customer",
	})
}

func (h *CustomerHandler) Delete(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	if ID == "" {
		return utils.BadRequest(ctx, fiber.Map{
			"message": "id is empty",
		})
	}

	if err := h.customerUsecase.Delete(ID); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success deleted customer",
	})
}

func (h *CustomerHandler) Find(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	if ID == "" {
		return utils.BadRequest(ctx, fiber.Map{
			"message": "id is empty",
		})
	}

	customer, err := h.customerUsecase.FindByID(ID)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success",
		"data": customer,
	})
}

func (h *CustomerHandler) List(ctx *fiber.Ctx) error {

	customers, err := h.customerUsecase.List()
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success",
		"data": customers,
	})

}

func (h *CustomerHandler) ShowImage(ctx *fiber.Ctx) error {
	objectName := ctx.Params("filename")
	if objectName == "" {
		return utils.BadRequest(ctx, fiber.Map{
			"message": "object name is empty",
		})
	}

	objectRender, err := utils.ShowFile(objectName)
	if err != nil {
		return utils.BadRequest(ctx, fiber.Map{
			"message": "object not found",
		})
	}

	return ctx.Status(fiber.StatusOK).SendStream(objectRender)
}

