package matchingvalidator

import (
	"E-01/entity"
	"E-01/param"
	"E-01/pkg/errmsg"
	"E-01/pkg/richerror"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateAddToWaitingListRequest(req param.AddToWaitingListRequest) (map[string]string, error) {
	const op = "matchingvalidator.AddToWaitingList"

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Category,
			validation.Required,
			validation.By(v.isCategoryValid),
		),
	); err != nil {

		fieldErrors := make(map[string]string)

		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}

		return fieldErrors, richerror.New(op).
			WithMessage(errmsg.ErrMsgInValidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).
			WithErr(err)
	}
	return nil, nil
}
func (v Validator) isCategoryValid(value interface{}) error {
	category := value.(entity.Category)
	fmt.Println(category)
	fmt.Println(category.IsValid())
	if !category.IsValid() {
		return fmt.Errorf(errmsg.ErrMsgCategoryIsNotValid)
	}
	return nil
}
