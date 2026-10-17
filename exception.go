package gest

import http_exception "github.com/martinpetrdev/Gest/internal/http-exception"

type IHTTPException = http_exception.IHTTPException

func NewHTTPBadRequestException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPBadRequestStatus, msg)
}

func NewHTTPUnauthorizedException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPUnauthorizedStatus, msg)
}

func NewHTTPPaymentRequiredException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPPaymentRequiredStatus, msg)
}

func NewHTTPForbiddenException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPForbiddenStatus, msg)
}

func NewHTTPNotFoundException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPNotFoundStatus, msg)
}

func NewHTTPMethodNotAllowedException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPMethodNotAllowedStatus, msg)
}

func NewHTTPNotAcceptableException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPNotAcceptableStatus, msg)
}

func NewHTTPProxyAuthRequiredException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPProxyAuthRequiredStatus, msg)
}

func NewHTTPRequestTimeoutException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPRequestTimeoutStatus, msg)
}

func NewHTTPConflictException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPConflictStatus, msg)
}

func NewHTTPGoneException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPGoneStatus, msg)
}

func NewHTTPLengthRequiredException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPLengthRequiredStatus, msg)
}

func NewHTTPPreconditionFailedException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPPreconditionFailedStatus, msg)
}

func NewHTTPPayloadTooLargeException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPPayloadTooLargeStatus, msg)
}

func NewHTTPURITooLongException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPURITooLongStatus, msg)
}

func NewHTTPUnsupportedMediaTypeException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPUnsupportedMediaTypeStatus, msg)
}

func NewHTTPRangeNotSatisfiableException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPRangeNotSatisfiableStatus, msg)
}

func NewHTTPExpectationFailedException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPExpectationFailedStatus, msg)
}

func NewHTTPImATeapotException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPImATeapotStatus, msg)
}

func NewHTTPMisdirectedRequestException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPMisdirectedRequestStatus, msg)
}

func NewHTTPUnprocessableEntityException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPUnprocessableEntityStatus, msg)
}

func NewHTTPLockedException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPLockedStatus, msg)
}

func NewHTTPFailedDependencyException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPFailedDependencyStatus, msg)
}

func NewHTTPTooEarlyException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPTooEarlyStatus, msg)
}

func NewHTTPUpgradeRequiredException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPUpgradeRequiredStatus, msg)
}

func NewHTTPPreconditionRequiredException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPPreconditionRequiredStatus, msg)
}

func NewHTTPTooManyRequestsException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPTooManyRequestsStatus, msg)
}

func NewHTTPRequestHeaderFieldsTooLargeException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPRequestHeaderFieldsTooLargeStatus, msg)
}

func NewHTTPUnavailableForLegalReasonsException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPUnavailableForLegalReasonsStatus, msg)
}

func NewHTTPInternalServerErrorException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPInternalServerErrorStatus, msg)
}

func NewHTTPNotImplementedException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPNotImplementedStatus, msg)
}

func NewHTTPBadGatewayException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPBadGatewayStatus, msg)
}

func NewHTTPServiceUnavailableException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPServiceUnavailableStatus, msg)
}

func NewHTTPGatewayTimeoutException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPGatewayTimeoutStatus, msg)
}

func NewHTTPHTTPVersionNotSupportedException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPHTTPVersionNotSupportedStatus, msg)
}

func NewHTTPVariantAlsoNegotiatesException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPVariantAlsoNegotiatesStatus, msg)
}

func NewHTTPInsufficientStorageException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPInsufficientStorageStatus, msg)
}

func NewHTTPLoopDetectedException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPLoopDetectedStatus, msg)
}

func NewHTTPNotExtendedException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPNotExtendedStatus, msg)
}

func NewHTTPNetworkAuthenticationRequiredException(msg string) http_exception.HTTPException {
	return http_exception.New(http_exception.HTTPNetworkAuthenticationRequiredStatus, msg)
}
