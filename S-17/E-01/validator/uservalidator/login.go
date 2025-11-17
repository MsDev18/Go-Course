package uservalidator

import (
	"E-01/dto"
	"E-01/pkg/errmsg"
	"E-01/pkg/richerror"
	"fmt"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateLoginRequest(req dto.LoginRequest) (map[string]string, error) {
	const op = "uservalidator.ValidateLoginRequest"

	if err := validation.ValidateStruct(&req,
		// TODO - add 3 config
		validation.Field(&req.PhoneNumber,
			validation.Required,
			validation.Match(regexp.MustCompile(PhoneNumberRegex)).Error(errmsg.ErrMsgPhoneNumberIsNotValid),
			validation.By(v.doesPhoneNumberUnique),
		),
		validation.Field(&req.Password,validation.Required),
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

func (v Validator) doesPhoneNumberUnique(value interface{}) error {
	phoneNumber := value.(string)

	_, err := v.repo.GetUserByPhoneNumber(phoneNumber)
	if err != nil {
		return fmt.Errorf(errmsg.ErrMsgNotFound)
	}	

	return nil
}
